package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2_4Vtbl struct {
	iCoreWebView2_3Vtbl
	AddFrameCreated        ComProc
	RemoveFrameCreated     ComProc
	AddDownloadStarting    ComProc
	RemoveDownloadStarting ComProc
}

type ICoreWebView2_4 struct {
	vtbl *iCoreWebView2_4Vtbl
}

func (i *ICoreWebView2_4) AddFrameCreatedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddFrameCreated.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_4) RemoveFrameCreated(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveFrameCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_4) AddDownloadStartingRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddDownloadStarting.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_4) RemoveDownloadStarting(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveDownloadStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GetICoreWebView2_4() *ICoreWebView2_4 {
	var result *ICoreWebView2_4

	iidICoreWebView2_4 := NewGUID("{20d02d59-6df2-42dc-bd06-f98a694b1302}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_4)),
		uintptr(unsafe.Pointer(&result)),
	)

	return result
}

func (e *Chromium) GetICoreWebView2_4() *ICoreWebView2_4 {
	return e.webview.GetICoreWebView2_4()
}
