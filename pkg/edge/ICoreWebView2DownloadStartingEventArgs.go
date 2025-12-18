package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2DownloadStartingEventArgsVtbl struct {
	_IUnknownVtbl
	GetDownloadOperation ComProc
	GetCancel            ComProc
	PutCancel            ComProc
	GetResultFilePath    ComProc
	PutResultFilePath    ComProc
	GetHandled           ComProc
	PutHandled           ComProc
	GetDeferral          ComProc
}

type ICoreWebView2DownloadStartingEventArgs struct {
	vtbl *_ICoreWebView2DownloadStartingEventArgsVtbl
}

func (i *ICoreWebView2DownloadStartingEventArgs) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetDownloadOperation() (*ICoreWebView2DownloadOperation, error) {
	var op *ICoreWebView2DownloadOperation
	_, _, err := i.vtbl.GetDownloadOperation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&op)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return op, nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetCancel() (bool, error) {
	var cancel bool
	_, _, err := i.vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return cancel, nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) PutCancel(cancel bool) error {
	_, _, err := i.vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetResultFilePath() (string, error) {
	var path *uint16
	_, _, err := i.vtbl.GetResultFilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&path)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	if path == nil {
		return "", nil
	}
	res := windows.UTF16PtrToString(path)
	windows.CoTaskMemFree(unsafe.Pointer(path))
	return res, nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) PutResultFilePath(resultFilePath string) error {
	_path, err := windows.UTF16PtrFromString(resultFilePath)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.PutResultFilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_path)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetHandled() (bool, error) {
	var handled bool
	_, _, err := i.vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return handled, nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) PutHandled(handled bool) error {
	_, _, err := i.vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// GetDeferral returns the underlying ICoreWebView2Deferral*.
// The concrete type is not wrapped here; use COM directly if needed.
func (i *ICoreWebView2DownloadStartingEventArgs) GetDeferral() (uintptr, error) {
	var deferral uintptr
	_, _, err := i.vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return deferral, nil
}
