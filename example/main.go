//go:build windows
// +build windows

package main

import (
	"log"
	"time"

	"github.com/logicossoftware/go-webview2"
)

func main() {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "go-webview2 example",
			Width:  900,
			Height: 650,
			Center: true,
		},
	})
	if w == nil {
		log.Fatalln("Failed to initialize WebView2.")
	}
	defer w.Destroy()

	w.SetSize(900, 650, webview2.HintNone)

	// Simple Go->JS bridge examples.
	_ = w.Bind("goAdd", func(a, b int) (int, error) {
		return a + b, nil
	})

	_ = w.Bind("goNow", func() (string, error) {
		return time.Now().Format(time.RFC3339), nil
	})

	_ = w.Bind("goSetTitle", func(title string) {
		w.Dispatch(func() {
			w.SetTitle(title)
		})
	})

	_ = w.Bind("goLog", func(msg string) {
		log.Printf("JS: %s", msg)
	})

	w.SetHtml(exampleHTML)
	w.Run()
}

const exampleHTML = `<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>go-webview2 example</title>
    <style>
      :root { color-scheme: light dark; }
      body { font-family: system-ui, Segoe UI, Arial, sans-serif; margin: 24px; }
      .row { display: flex; gap: 12px; flex-wrap: wrap; align-items: center; }
      input { padding: 8px 10px; min-width: 220px; }
      button { padding: 8px 12px; cursor: pointer; }
      pre { padding: 12px; border-radius: 8px; background: rgba(127,127,127,0.15); overflow: auto; }
      .card { border: 1px solid rgba(127,127,127,0.35); border-radius: 12px; padding: 16px; margin-top: 16px; }
      a { color: inherit; }
    </style>
  </head>
  <body>
    <h1>go-webview2 example</h1>
    <p>
      This page demonstrates <code>Bind</code> RPC calls and basic UI updates.
      DevTools are enabled (<code>Debug: true</code>).
    </p>

    <div class="card">
      <h2>Go bindings</h2>
      <div class="row">
        <button id="btnAdd">Call goAdd(40, 2)</button>
        <button id="btnNow">Call goNow()</button>
        <button id="btnTitle">Set window title</button>
        <input id="title" value="Title from JS" />
      </div>
      <div class="row" style="margin-top: 10px;">
        <button id="btnLog">goLog('hello from JS')</button>
        <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/" target="_blank">
          WebView2 Win32 docs
        </a>
      </div>

      <h3>Output</h3>
      <pre id="out">(click a button)</pre>
    </div>

    <script>
      const out = document.getElementById('out');
      const titleInput = document.getElementById('title');

      function write(value) {
        out.textContent = typeof value === 'string' ? value : JSON.stringify(value, null, 2);
      }

      document.getElementById('btnAdd').addEventListener('click', async () => {
        const result = await window.goAdd(40, 2);
        write({ goAdd: result });
      });

      document.getElementById('btnNow').addEventListener('click', async () => {
        const result = await window.goNow();
        write({ goNow: result });
      });

      document.getElementById('btnTitle').addEventListener('click', async () => {
        await window.goSetTitle(titleInput.value);
        write({ goSetTitle: 'ok' });
      });

      document.getElementById('btnLog').addEventListener('click', async () => {
        await window.goLog('hello from JS');
        write({ goLog: 'logged to Go stdout' });
      });

      // Basic sanity check.
      window.goLog('example page loaded');
    </script>
  </body>
</html>`
