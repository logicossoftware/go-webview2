package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2_2Vtbl struct {
	iCoreWebView2Vtbl
	AddWebResourceResponseReceived    ComProc
	RemoveWebResourceResponseReceived ComProc
	NavigateWithWebResourceRequest    ComProc
	AddDomContentLoaded               ComProc
	RemoveDomContentLoaded            ComProc
	GetCookieManager                  ComProc
	GetEnvironment                    ComProc
}

type ICoreWebView2_2 struct {
	vtbl *iCoreWebView2_2Vtbl
}

func (i *ICoreWebView2_2) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_2) AddWebResourceResponseReceivedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddWebResourceResponseReceived.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_2) RemoveWebResourceResponseReceived(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveWebResourceResponseReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_2) NavigateWithWebResourceRequest(request *ICoreWebView2WebResourceRequest) error {
	_, _, err := i.vtbl.NavigateWithWebResourceRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(request)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_2) AddDomContentLoadedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddDomContentLoaded.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_2) RemoveDomContentLoaded(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveDomContentLoaded.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// GetCookieManager returns the underlying ICoreWebView2CookieManager*.
// The concrete type is not wrapped here; use COM directly if needed.
func (i *ICoreWebView2_2) GetCookieManager() (uintptr, error) {
	var mgr uintptr
	_, _, err := i.vtbl.GetCookieManager.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&mgr)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return mgr, nil
}

func (i *ICoreWebView2_2) GetEnvironment() (*ICoreWebView2Environment, error) {
	var env *ICoreWebView2Environment
	_, _, err := i.vtbl.GetEnvironment.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&env)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return env, nil
}
