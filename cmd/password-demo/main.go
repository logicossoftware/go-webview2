package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/logicossoftware/go-webview2"
)

type credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type store struct {
	mu    sync.RWMutex
	creds map[string]map[string]credential
	last  map[string]string
}

func newStore() *store {
	return &store{creds: make(map[string]map[string]credential), last: make(map[string]string)}
}

func (s *store) save(siteKey, username, password string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.creds[siteKey] == nil {
		s.creds[siteKey] = make(map[string]credential)
	}
	s.creds[siteKey][username] = credential{Username: username, Password: password}
	s.last[siteKey] = username
}

func (s *store) get(siteKey string) (credential, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	byUser := s.creds[siteKey]
	if len(byUser) == 0 {
		return credential{}, false
	}
	if u := s.last[siteKey]; u != "" {
		if c, ok := byUser[u]; ok {
			return c, true
		}
	}
	for _, c := range byUser {
		return c, true
	}
	return credential{}, false
}

type storedCredential struct {
	SiteKey        string `json:"siteKey"`
	Username       string `json:"username"`
	PasswordLength int    `json:"passwordLength"`
}

func (s *store) list(siteKey string) []storedCredential {
	s.mu.RLock()
	defer s.mu.RUnlock()
	byUser := s.creds[siteKey]
	if len(byUser) == 0 {
		return nil
	}
	items := make([]storedCredential, 0, len(byUser))
	for _, c := range byUser {
		items = append(items, storedCredential{SiteKey: siteKey, Username: c.Username, PasswordLength: len(c.Password)})
	}
	return items
}

func (s *store) clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	clear(s.creds)
	clear(s.last)
}

type passkeyCredential struct {
	IDB64URL string    `json:"idB64Url"`
	Created  time.Time `json:"created"`
}

type passkeyStore struct {
	mu    sync.RWMutex
	creds map[string]map[string][]passkeyCredential // siteKey -> username -> creds
}

func newPasskeyStore() *passkeyStore {
	return &passkeyStore{creds: make(map[string]map[string][]passkeyCredential)}
}

func (p *passkeyStore) add(siteKey, username, credIDB64URL string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.creds[siteKey] == nil {
		p.creds[siteKey] = make(map[string][]passkeyCredential)
	}
	p.creds[siteKey][username] = append(p.creds[siteKey][username], passkeyCredential{IDB64URL: credIDB64URL, Created: time.Now()})
}

func (p *passkeyStore) list(siteKey, username string) []passkeyCredential {
	p.mu.RLock()
	defer p.mu.RUnlock()
	byUser := p.creds[siteKey]
	if byUser == nil {
		return nil
	}
	items := byUser[username]
	out := make([]passkeyCredential, len(items))
	copy(out, items)
	return out
}

type session struct {
	SiteKey  string
	Username string
	AuthType string
}

type sessionStore struct {
	mu       sync.RWMutex
	sessions map[string]session
}

func newSessionStore() *sessionStore {
	return &sessionStore{sessions: make(map[string]session)}
}

func randomB64URL(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

func (s *sessionStore) newSession(siteKey, username, authType string) string {
	tok := randomB64URL(32)
	s.mu.Lock()
	s.sessions[tok] = session{SiteKey: siteKey, Username: username, AuthType: authType}
	s.mu.Unlock()
	return tok
}

func (s *sessionStore) get(tok string) (session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ss, ok := s.sessions[tok]
	return ss, ok
}

func (s *sessionStore) clear(tok string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, tok)
}

type demoServer struct {
	baseURL   string
	listener  net.Listener
	httpSrv   *http.Server
	passwords *store
	passkeys  *passkeyStore
	sessions  *sessionStore
}

func (d *demoServer) shutdown(ctx context.Context) {
	if d.httpSrv != nil {
		_ = d.httpSrv.Shutdown(ctx)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func readJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(dst)
}

func getSessionToken(r *http.Request) (string, bool) {
	c, err := r.Cookie("demo_session")
	if err != nil {
		return "", false
	}
	if c.Value == "" {
		return "", false
	}
	return c.Value, true
}

func setSessionCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "demo_session",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int((24 * time.Hour).Seconds()),
	})
}

func clearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "demo_session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})
}

func (d *demoServer) serve() (*demoServer, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	d.listener = ln
	d.baseURL = "http://" + ln.Addr().String()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		_, _ = w.Write([]byte(loginHTML))
	})

	mux.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		if tok, ok := getSessionToken(r); ok {
			if _, ok := d.sessions.get(tok); ok {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.Header().Set("Cache-Control", "no-store")
				_, _ = w.Write([]byte(appHTML))
				return
			}
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	mux.HandleFunc("/vault", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		_, _ = w.Write([]byte(vaultHTML))
	})

	mux.HandleFunc("/api/me", func(w http.ResponseWriter, r *http.Request) {
		siteKey := d.baseURL
		if tok, ok := getSessionToken(r); ok {
			if s, ok := d.sessions.get(tok); ok {
				writeJSON(w, http.StatusOK, map[string]any{"ok": true, "siteKey": siteKey, "username": s.Username, "authType": s.AuthType})
				return
			}
		}
		writeJSON(w, http.StatusOK, map[string]any{"ok": false, "siteKey": siteKey})
	})

	mux.HandleFunc("/api/logout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if tok, ok := getSessionToken(r); ok {
			d.sessions.clear(tok)
		}
		clearSessionCookie(w)
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	})

	mux.HandleFunc("/api/password/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := readJSON(r, &req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": err.Error()})
			return
		}
		req.Username = strings.TrimSpace(req.Username)
		if req.Username == "" || req.Password == "" {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": "username and password required"})
			return
		}

		d.passwords.save(d.baseURL, req.Username, req.Password)
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	})

	mux.HandleFunc("/api/password/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := readJSON(r, &req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": err.Error()})
			return
		}
		req.Username = strings.TrimSpace(req.Username)
		if req.Username == "" || req.Password == "" {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": "username and password required"})
			return
		}

		// Demo validation: just compare against the password vault.
		stored, ok := func() (credential, bool) {
			d.passwords.mu.RLock()
			defer d.passwords.mu.RUnlock()
			byUser := d.passwords.creds[d.baseURL]
			if byUser == nil {
				return credential{}, false
			}
			c, ok := byUser[req.Username]
			return c, ok
		}()
		if !ok || stored.Password != req.Password {
			writeJSON(w, http.StatusUnauthorized, map[string]any{"ok": false, "error": "invalid username or password"})
			return
		}

		tok := d.sessions.newSession(d.baseURL, req.Username, "password")
		setSessionCookie(w, tok)
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	})

	mux.HandleFunc("/api/passkey/register/options", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Username string `json:"username"`
		}
		if err := readJSON(r, &req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": err.Error()})
			return
		}
		req.Username = strings.TrimSpace(req.Username)
		if req.Username == "" {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": "username required"})
			return
		}

		challenge := randomB64URL(32)
		userID := randomB64URL(16)
		existing := d.passkeys.list(d.baseURL, req.Username)
		exclude := make([]map[string]any, 0, len(existing))
		for _, c := range existing {
			exclude = append(exclude, map[string]any{"type": "public-key", "id": c.IDB64URL})
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"ok": true,
			"publicKey": map[string]any{
				"challenge": challenge,
				"rp": map[string]any{
					"name": "go-webview2 demo",
					"id":   "127.0.0.1",
				},
				"user": map[string]any{
					"id":          userID,
					"name":        req.Username,
					"displayName": req.Username,
				},
				"pubKeyCredParams": []map[string]any{
					{"type": "public-key", "alg": -7},   // ES256
					{"type": "public-key", "alg": -257}, // RS256
				},
				"timeout":     60000,
				"attestation": "none",
				"authenticatorSelection": map[string]any{
					"userVerification": "preferred",
				},
				"excludeCredentials": exclude,
			},
		})
	})

	mux.HandleFunc("/api/passkey/register/finish", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Username          string `json:"username"`
			CredentialIDB64   string `json:"credentialId"`
			RawIDB64          string `json:"rawId"`
			ClientDataJSONB64 string `json:"clientDataJSON"`
			AttObjB64         string `json:"attestationObject"`
		}
		if err := readJSON(r, &req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": err.Error()})
			return
		}
		req.Username = strings.TrimSpace(req.Username)
		if req.Username == "" || req.RawIDB64 == "" {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": "username and rawId required"})
			return
		}

		// Demo note: we do not verify attestation or clientDataJSON here.
		_ = req.CredentialIDB64
		_ = req.ClientDataJSONB64
		_ = req.AttObjB64

		d.passkeys.add(d.baseURL, req.Username, req.RawIDB64)
		tok := d.sessions.newSession(d.baseURL, req.Username, "passkey")
		setSessionCookie(w, tok)
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	})

	mux.HandleFunc("/api/passkey/login/options", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Username string `json:"username"`
		}
		if err := readJSON(r, &req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": err.Error()})
			return
		}
		req.Username = strings.TrimSpace(req.Username)
		if req.Username == "" {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": "username required"})
			return
		}

		creds := d.passkeys.list(d.baseURL, req.Username)
		if len(creds) == 0 {
			writeJSON(w, http.StatusNotFound, map[string]any{"ok": false, "error": "no passkeys registered for this user"})
			return
		}
		allow := make([]map[string]any, 0, len(creds))
		for _, c := range creds {
			allow = append(allow, map[string]any{"type": "public-key", "id": c.IDB64URL})
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"ok": true,
			"publicKey": map[string]any{
				"challenge":        randomB64URL(32),
				"timeout":          60000,
				"rpId":             "127.0.0.1",
				"userVerification": "preferred",
				"allowCredentials": allow,
			},
		})
	})

	mux.HandleFunc("/api/passkey/login/finish", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Username             string `json:"username"`
			RawIDB64             string `json:"rawId"`
			ClientDataJSONB64    string `json:"clientDataJSON"`
			AuthenticatorDataB64 string `json:"authenticatorData"`
			SignatureB64         string `json:"signature"`
			UserHandleB64        string `json:"userHandle"`
		}
		if err := readJSON(r, &req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": err.Error()})
			return
		}
		req.Username = strings.TrimSpace(req.Username)
		if req.Username == "" || req.RawIDB64 == "" {
			writeJSON(w, http.StatusBadRequest, map[string]any{"ok": false, "error": "username and rawId required"})
			return
		}
		_ = req.ClientDataJSONB64
		_ = req.AuthenticatorDataB64
		_ = req.SignatureB64
		_ = req.UserHandleB64

		creds := d.passkeys.list(d.baseURL, req.Username)
		found := false
		for _, c := range creds {
			if c.IDB64URL == req.RawIDB64 {
				found = true
				break
			}
		}
		if !found {
			writeJSON(w, http.StatusUnauthorized, map[string]any{"ok": false, "error": "unknown credential"})
			return
		}

		// Demo note: we do not verify the assertion signature server-side.
		tok := d.sessions.newSession(d.baseURL, req.Username, "passkey")
		setSessionCookie(w, tok)
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	})

	d.httpSrv = &http.Server{Handler: mux}
	go func() {
		_ = d.httpSrv.Serve(ln)
	}()
	return d, nil
}

const loginHTML = `<!doctype html>
<html>
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Demo Site: Login</title>
  <style>
    :root { color-scheme: light dark; }
    body { font-family: Segoe UI, Arial, sans-serif; margin: 24px; }
    .row { display:flex; gap: 18px; align-items:flex-start; flex-wrap:wrap; }
    .card { border: 1px solid rgba(127,127,127,0.35); border-radius: 10px; padding: 16px; width: 420px; }
    label { display:block; margin-top: 10px; font-size: 13px; opacity: 0.85; }
    input { width: 100%; padding: 10px; font-size: 14px; box-sizing:border-box; }
    button { margin-top: 12px; padding: 10px 14px; font-size: 14px; }
    .muted { opacity: 0.75; }
    .ok { color: #1a7f37; }
    .bad { color: #b42318; }
    code { background: rgba(127,127,127,0.18); padding: 2px 6px; border-radius: 4px; }
    .actions { display:flex; gap: 10px; flex-wrap:wrap; }
  </style>
</head>
<body>
  <h1>Demo site login</h1>
  <p class="muted">
    This is a tiny local website (localhost), embedded in WebView2. It supports <b>password</b> and <b>passkey (WebAuthn)</b> sign-in.
    The "password manager" part is a Go in-memory vault exposed via <code>Bind()</code>.
  </p>

  <div class="row">
    <div class="card">
      <h2>Password</h2>
      <label>Username</label>
      <input id="u" type="text" autocomplete="username" placeholder="alice" />
      <label>Password</label>
      <input id="p" type="password" autocomplete="current-password" placeholder="••••••••" />

      <div class="actions">
        <button id="btnRegister">Create account</button>
        <button id="btnLogin">Login</button>
        <button id="btnVault" onclick="location.href='/vault'">Open vault</button>
      </div>

      <p id="pwStatus" class="muted"></p>
      <p class="muted">Tip: after creating an account, reload this page (Ctrl+R) to see autofill.</p>
    </div>

    <div class="card">
      <h2>Passkey (WebAuthn)</h2>
      <p class="muted">Works best on Windows with Windows Hello enabled.</p>
      <div class="actions">
        <button id="btnPkRegister">Register passkey</button>
        <button id="btnPkLogin">Login with passkey</button>
      </div>
      <p id="pkStatus" class="muted"></p>
      <p class="muted">If this fails, check WebView2 runtime version and that a platform authenticator is available.</p>
    </div>
  </div>

  <script>
    const $ = (id) => document.getElementById(id);
    const siteKey = () => (location.origin && location.origin !== 'null') ? location.origin : 'unknown';

    function setStatus(el, ok, msg) {
      el.className = ok ? 'ok' : 'bad';
      el.textContent = msg;
    }

    async function postJson(url, body) {
      const res = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body || {}),
      });
      const text = await res.text();
      let json;
      try { json = JSON.parse(text); } catch { json = { ok: false, error: text || ('HTTP ' + res.status) }; }
      if (!res.ok) {
        const err = new Error(json && json.error ? json.error : ('HTTP ' + res.status));
        err.httpStatus = res.status;
        throw err;
      }
      return json;
    }

    function b64urlToBytes(b64url) {
      const pad = '='.repeat((4 - (b64url.length % 4)) % 4);
      const b64 = (b64url + pad).replace(/-/g, '+').replace(/_/g, '/');
      const bin = atob(b64);
      const bytes = new Uint8Array(bin.length);
      for (let i = 0; i < bin.length; i++) bytes[i] = bin.charCodeAt(i);
      return bytes;
    }

    function bytesToB64url(bytes) {
      let bin = '';
      for (let i = 0; i < bytes.length; i++) bin += String.fromCharCode(bytes[i]);
      const b64 = btoa(bin);
      return b64.replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/g, '');
    }

    function abToB64url(ab) { return bytesToB64url(new Uint8Array(ab)); }
    function b64urlToAB(b64url) { return b64urlToBytes(b64url).buffer; }

    async function maybeAutofill() {
      if (typeof window.pw_get !== 'function') return;
      try {
        const cred = await window.pw_get(siteKey());
        if (cred && cred.username) $('u').value = cred.username;
        if (cred && cred.password) $('p').value = cred.password;
      } catch {}
    }

    async function onRegisterPassword() {
      const username = $('u').value.trim();
      const password = $('p').value;
      try {
        await postJson('/api/password/register', { username, password });
        if (typeof window.pw_save === 'function') {
          await window.pw_save(siteKey(), username, password);
        }
        setStatus($('pwStatus'), true, 'Account created. (Saved to demo password vault too.)');
      } catch (e) {
        setStatus($('pwStatus'), false, 'Register failed: ' + (e.message || e));
      }
    }

    async function onLoginPassword() {
      const username = $('u').value.trim();
      const password = $('p').value;
      try {
        await postJson('/api/password/login', { username, password });
        if (typeof window.pw_save === 'function') {
          await window.pw_save(siteKey(), username, password);
        }
        location.href = '/app';
      } catch (e) {
        setStatus($('pwStatus'), false, 'Login failed: ' + (e.message || e));
      }
    }

    async function onRegisterPasskey() {
      const username = $('u').value.trim();
      if (!username) return setStatus($('pkStatus'), false, 'Enter a username first.');
      try {
        if (!window.PublicKeyCredential || !navigator.credentials) {
          throw new Error('WebAuthn not available in this runtime.');
        }
        const opts = await postJson('/api/passkey/register/options', { username });
        const pk = opts.publicKey;
        pk.challenge = b64urlToAB(pk.challenge);
        pk.user.id = b64urlToBytes(pk.user.id);
        if (Array.isArray(pk.excludeCredentials)) {
          pk.excludeCredentials = pk.excludeCredentials.map((c) => ({ ...c, id: b64urlToAB(c.id) }));
        }
        const cred = await navigator.credentials.create({ publicKey: pk });
        const rawId = abToB64url(cred.rawId);
        const attObj = abToB64url(cred.response.attestationObject);
        const clientDataJSON = abToB64url(cred.response.clientDataJSON);
        await postJson('/api/passkey/register/finish', {
          username,
          credentialId: cred.id,
          rawId,
          attestationObject: attObj,
          clientDataJSON,
        });
        setStatus($('pkStatus'), true, 'Passkey registered. Redirecting to app…');
        location.href = '/app';
      } catch (e) {
        setStatus($('pkStatus'), false, 'Passkey register failed: ' + (e.message || e));
      }
    }

    async function onLoginPasskey() {
      const username = $('u').value.trim();
      if (!username) return setStatus($('pkStatus'), false, 'Enter a username first.');
      try {
        const opts = await postJson('/api/passkey/login/options', { username });
        const pk = opts.publicKey;
        pk.challenge = b64urlToAB(pk.challenge);
        if (Array.isArray(pk.allowCredentials)) {
          pk.allowCredentials = pk.allowCredentials.map((c) => ({ ...c, id: b64urlToAB(c.id) }));
        }
        const assertion = await navigator.credentials.get({ publicKey: pk });
        const rawId = abToB64url(assertion.rawId);
        const clientDataJSON = abToB64url(assertion.response.clientDataJSON);
        const authData = abToB64url(assertion.response.authenticatorData);
        const sig = abToB64url(assertion.response.signature);
        const userHandle = assertion.response.userHandle ? abToB64url(assertion.response.userHandle) : '';
        await postJson('/api/passkey/login/finish', {
          username,
          rawId,
          clientDataJSON,
          authenticatorData: authData,
          signature: sig,
          userHandle,
        });
        location.href = '/app';
      } catch (e) {
        setStatus($('pkStatus'), false, 'Passkey login failed: ' + (e.message || e));
      }
    }

    async function boot() {
      await maybeAutofill();
      try {
        const me = await fetch('/api/me').then(r => r.json());
        if (me && me.ok) location.href = '/app';
      } catch {}
      $('btnRegister').addEventListener('click', (e) => { e.preventDefault(); onRegisterPassword(); });
      $('btnLogin').addEventListener('click', (e) => { e.preventDefault(); onLoginPassword(); });
      $('btnPkRegister').addEventListener('click', (e) => { e.preventDefault(); onRegisterPasskey(); });
      $('btnPkLogin').addEventListener('click', (e) => { e.preventDefault(); onLoginPasskey(); });
    }

    boot();
  </script>
</body>
</html>`

const appHTML = `<!doctype html>
<html>
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Demo Site: App</title>
  <style>
    :root { color-scheme: light dark; }
    body { font-family: Segoe UI, Arial, sans-serif; margin: 24px; }
    .muted { opacity: 0.75; }
    .card { border: 1px solid rgba(127,127,127,0.35); border-radius: 10px; padding: 16px; max-width: 720px; }
    button { margin-top: 12px; padding: 10px 14px; font-size: 14px; }
    code { background: rgba(127,127,127,0.18); padding: 2px 6px; border-radius: 4px; }
    .actions { display:flex; gap: 10px; flex-wrap:wrap; }
  </style>
</head>
<body>
  <h1>Welcome</h1>
  <div class="card">
    <p id="me" class="muted">Loading session…</p>
    <div class="actions">
      <button onclick="location.href='/vault'">Open vault</button>
      <button id="logout">Logout</button>
      <button onclick="location.href='/login'">Back to login</button>
    </div>
    <p class="muted">
      Notes: passkey flows are demo-only and <b>do not verify signatures</b> server-side.
      Password vault is stored in a Go map via <code>Bind()</code>.
    </p>
  </div>
  <script>
    async function post(url, body) {
      return fetch(url, { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(body || {}) });
    }
    async function boot() {
      const me = await fetch('/api/me').then(r => r.json());
      document.getElementById('me').textContent = me && me.ok
        ? ('Signed in as ' + me.username + ' via ' + me.authType + ' on ' + me.siteKey)
        : 'Not signed in.';
      document.getElementById('logout').addEventListener('click', async () => {
        await post('/api/logout');
        location.href = '/login';
      });
    }
    boot();
  </script>
</body>
</html>`

const vaultHTML = `<!doctype html>
<html>
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Demo Site: Vault</title>
  <style>
    :root { color-scheme: light dark; }
    body { font-family: Segoe UI, Arial, sans-serif; margin: 24px; }
    .muted { opacity: 0.75; }
    .card { border: 1px solid rgba(127,127,127,0.35); border-radius: 10px; padding: 16px; max-width: 820px; }
    button { margin-top: 12px; padding: 10px 14px; font-size: 14px; }
    table { width: 100%; border-collapse: collapse; margin-top: 10px; }
    th, td { text-align: left; padding: 8px 10px; border-bottom: 1px solid rgba(127,127,127,0.25); }
    code { background: rgba(127,127,127,0.18); padding: 2px 6px; border-radius: 4px; }
    .actions { display:flex; gap: 10px; flex-wrap:wrap; }
  </style>
</head>
<body>
  <h1>Demo password vault</h1>
  <p class="muted">This page shows what the Go in-memory vault has captured for this site.</p>

  <div class="card">
    <div class="actions">
      <button onclick="location.href='/login'">Back to login</button>
      <button onclick="location.href='/app'">Go to app</button>
      <button id="refresh">Refresh</button>
      <button id="clear">Clear vault</button>
    </div>
    <p id="status" class="muted"></p>
    <table>
      <thead><tr><th>Username</th><th>Password</th></tr></thead>
      <tbody id="rows"></tbody>
    </table>
    <p class="muted">API: <code>pw_list(siteKey)</code>, <code>pw_clear()</code></p>
  </div>

  <script>
    const siteKey = () => (location.origin && location.origin !== 'null') ? location.origin : 'unknown';
    const rows = document.getElementById('rows');
    const status = document.getElementById('status');

    function mask(n) { return n > 0 ? '•'.repeat(Math.min(n, 12)) : ''; }

    async function refresh() {
      rows.innerHTML = '';
      if (typeof window.pw_list !== 'function') {
        status.textContent = 'pw_list is not available.';
        return;
      }
      const list = await window.pw_list(siteKey());
      if (!list || list.length === 0) {
        status.textContent = 'No stored passwords for ' + siteKey();
        return;
      }
      status.textContent = 'Stored passwords for ' + siteKey();
      for (const item of list) {
        const tr = document.createElement('tr');
        const tdU = document.createElement('td');
        const tdP = document.createElement('td');
        tdU.textContent = item.username;
        tdP.textContent = mask(item.passwordLength) + ' (' + item.passwordLength + ')';
        tr.appendChild(tdU);
        tr.appendChild(tdP);
        rows.appendChild(tr);
      }
    }

    document.getElementById('refresh').addEventListener('click', (e) => { e.preventDefault(); refresh(); });
    document.getElementById('clear').addEventListener('click', async (e) => {
      e.preventDefault();
      if (typeof window.pw_clear === 'function') await window.pw_clear();
      await refresh();
    });
    refresh();
  </script>
</body>
</html>`

func main() {
	st := newStore()
	pk := newPasskeyStore()
	sessions := newSessionStore()

	ds := &demoServer{passwords: st, passkeys: pk, sessions: sessions}
	_, err := ds.serve()
	if err != nil {
		log.Fatalf("failed to start demo server: %v", err)
	}
	defer ds.shutdown(context.Background())

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Password manager demo (localhost + passkeys)",
			Width:  900,
			Height: 650,
			Center: true,
		},
	})
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()

	// Password vault ("password manager") APIs.
	_ = w.Bind("pw_save", func(siteKey, username, password string) error {
		log.Printf("pw_save: site=%q user=%q (password length=%d)", siteKey, username, len(password))
		st.save(siteKey, username, password)
		return nil
	})
	_ = w.Bind("pw_get", func(siteKey string) (*credential, error) {
		if c, ok := st.get(siteKey); ok {
			return &c, nil
		}
		return nil, nil
	})
	_ = w.Bind("pw_list", func(siteKey string) ([]storedCredential, error) {
		return st.list(siteKey), nil
	})
	_ = w.Bind("pw_clear", func() error {
		log.Printf("pw_clear")
		st.clear()
		return nil
	})

	log.Printf("demo site running at %s", ds.baseURL)

	w.Navigate(ds.baseURL + "/login")

	w.Run()
}
