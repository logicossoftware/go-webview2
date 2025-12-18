package winloader

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"syscall"

	"golang.org/x/sys/windows"
)

type Module struct {
	handle windows.Handle
	path   string
}

type Proc struct {
	addr uintptr
}

var fileMu sync.Mutex

// LoadFromMemory persists the DLL bytes to disk and loads it with LoadLibrary.
//
// Note: this is not a true in-memory PE loader; Windows generally requires a
// file-backed mapping for LoadLibrary-style APIs. This implementation exists to
// avoid external dependencies while keeping the embedded DLL workflow.
func LoadFromMemory(dll []byte) (Module, error) {
	if len(dll) == 0 {
		return Module{}, errors.New("winloader: empty DLL")
	}

	sum := sha256.Sum256(dll)
	name := "go-winloader-" + hex.EncodeToString(sum[:]) + ".dll"
	path := filepath.Join(os.TempDir(), name)

	fileMu.Lock()
	if _, err := os.Stat(path); err != nil {
		// Only write if missing. If multiple processes race, last writer wins.
		_ = os.MkdirAll(filepath.Dir(path), 0o755)
		if writeErr := os.WriteFile(path, dll, 0o644); writeErr != nil {
			fileMu.Unlock()
			return Module{}, writeErr
		}
	}
	fileMu.Unlock()

	h, err := windows.LoadLibrary(path)
	if err != nil {
		return Module{}, err
	}
	return Module{handle: h, path: path}, nil
}

func (m Module) Proc(name string) Proc {
	addr, err := windows.GetProcAddress(m.handle, name)
	if err != nil {
		return Proc{}
	}
	return Proc{addr: addr}
}

func (p Proc) Call(args ...uint64) (uint64, uint64, error) {
	if p.addr == 0 {
		return 0, 0, syscall.EINVAL
	}

	argv := make([]uintptr, 0, len(args))
	for _, a := range args {
		argv = append(argv, uintptr(a))
	}

	r1, r2, lastErr := syscall.SyscallN(p.addr, argv...)
	return uint64(r1), uint64(r2), lastErr
}
