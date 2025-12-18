package edge

import (
	"unsafe"

	"github.com/logicossoftware/go-webview2/internal/w32"
	"golang.org/x/sys/windows"
)

type _ICoreWebView2Controller2Vtbl struct {
	_IUnknownVtbl
	GetIsVisible                      ComProc
	PutIsVisible                      ComProc
	GetBounds                         ComProc
	PutBounds                         ComProc
	GetZoomFactor                     ComProc
	PutZoomFactor                     ComProc
	AddZoomFactorChanged              ComProc
	RemoveZoomFactorChanged           ComProc
	SetBoundsAndZoomFactor            ComProc
	MoveFocus                         ComProc
	AddMoveFocusRequested             ComProc
	RemoveMoveFocusRequested          ComProc
	AddGotFocus                       ComProc
	RemoveGotFocus                    ComProc
	AddLostFocus                      ComProc
	RemoveLostFocus                   ComProc
	AddAcceleratorKeyPressed          ComProc
	RemoveAcceleratorKeyPressed       ComProc
	GetParentWindow                   ComProc
	PutParentWindow                   ComProc
	NotifyParentWindowPositionChanged ComProc
	Close                             ComProc
	GetCoreWebView2                   ComProc
	GetDefaultBackgroundColor         ComProc
	PutDefaultBackgroundColor         ComProc
}

type ICoreWebView2Controller2 struct {
	vtbl *_ICoreWebView2Controller2Vtbl
}

func (i *ICoreWebView2Controller2) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Controller2) GetIsVisible() (bool, error) {
	var visible bool
	_, _, err := i.vtbl.GetIsVisible.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&visible)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return visible, nil
}

func (i *ICoreWebView2Controller2) PutIsVisible(isVisible bool) error {
	_, _, err := i.vtbl.PutIsVisible.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(boolToInt(isVisible)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) GetBounds() (*w32.Rect, error) {
	var bounds w32.Rect
	_, _, err := i.vtbl.GetBounds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&bounds)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return &bounds, nil
}

func (i *ICoreWebView2Controller2) PutBounds(bounds w32.Rect) error {
	_, _, err := i.vtbl.PutBounds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&bounds)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) GetZoomFactor() (float64, error) {
	var zoom float64
	_, _, err := i.vtbl.GetZoomFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&zoom)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return zoom, nil
}

func (i *ICoreWebView2Controller2) PutZoomFactor(zoomFactor float64) error {
	_, _, err := i.vtbl.PutZoomFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&zoomFactor)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) AddZoomFactorChangedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddZoomFactorChanged.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) RemoveZoomFactorChanged(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveZoomFactorChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) SetBoundsAndZoomFactor(bounds w32.Rect, zoomFactor float64) error {
	_, _, err := i.vtbl.SetBoundsAndZoomFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&bounds)),
		uintptr(unsafe.Pointer(&zoomFactor)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) MoveFocus(reason uintptr) error {
	_, _, err := i.vtbl.MoveFocus.Call(
		uintptr(unsafe.Pointer(i)),
		reason,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) AddMoveFocusRequestedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddMoveFocusRequested.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) RemoveMoveFocusRequested(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveMoveFocusRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) AddGotFocusRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddGotFocus.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) RemoveGotFocus(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveGotFocus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) AddLostFocusRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddLostFocus.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) RemoveLostFocus(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveLostFocus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) AddAcceleratorKeyPressed(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddAcceleratorKeyPressed.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) RemoveAcceleratorKeyPressed(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveAcceleratorKeyPressed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) GetParentWindow() (uintptr, error) {
	var hwnd uintptr
	_, _, err := i.vtbl.GetParentWindow.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hwnd)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return hwnd, nil
}

func (i *ICoreWebView2Controller2) PutParentWindow(hwnd uintptr) error {
	_, _, err := i.vtbl.PutParentWindow.Call(
		uintptr(unsafe.Pointer(i)),
		hwnd,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) NotifyParentWindowPositionChanged() error {
	_, _, err := i.vtbl.NotifyParentWindowPositionChanged.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) Close() error {
	_, _, err := i.vtbl.Close.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller2) GetCoreWebView2() (*ICoreWebView2, error) {
	var wv *ICoreWebView2
	_, _, err := i.vtbl.GetCoreWebView2.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&wv)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return wv, nil
}

func (i *ICoreWebView2Controller2) GetDefaultBackgroundColor() (*COREWEBVIEW2_COLOR, error) {
	var err error
	var backgroundColor *COREWEBVIEW2_COLOR
	_, _, err = i.vtbl.GetDefaultBackgroundColor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&backgroundColor)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return backgroundColor, nil
}

func (i *ICoreWebView2Controller2) PutDefaultBackgroundColor(backgroundColor COREWEBVIEW2_COLOR) error {
	var err error

	// Cast to a uint32 as that's what the call is expecting
	col := *(*uint32)(unsafe.Pointer(&backgroundColor))

	_, _, err = i.vtbl.PutDefaultBackgroundColor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(col),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
