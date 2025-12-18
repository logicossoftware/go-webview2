package edge

import (
	"unsafe"

	"github.com/logicossoftware/go-webview2/internal/w32"
	"golang.org/x/sys/windows"
)

type _ICoreWebView2ControllerVtbl struct {
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
}

type ICoreWebView2Controller struct {
	vtbl *_ICoreWebView2ControllerVtbl
}

func (i *ICoreWebView2Controller) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Controller) GetBounds() (*w32.Rect, error) {
	var err error
	var bounds w32.Rect
	_, _, err = i.vtbl.GetBounds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&bounds)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return &bounds, nil
}

func (i *ICoreWebView2Controller) PutBounds(bounds w32.Rect) error {
	var err error

	_, _, err = i.vtbl.PutBounds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&bounds)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller) AddAcceleratorKeyPressed(eventHandler *ICoreWebView2AcceleratorKeyPressedEventHandler, token *_EventRegistrationToken) error {
	var err error
	_, _, err = i.vtbl.AddAcceleratorKeyPressed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller) PutIsVisible(isVisible bool) error {
	var err error

	_, _, err = i.vtbl.PutIsVisible.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(boolToInt(isVisible)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller) GetIsVisible() (bool, error) {
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

func (i *ICoreWebView2Controller) GetZoomFactor() (float64, error) {
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

func (i *ICoreWebView2Controller) PutZoomFactor(zoomFactor float64) error {
	_, _, err := i.vtbl.PutZoomFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&zoomFactor)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// AddZoomFactorChangedRaw registers a zoom-factor-changed handler.
// handler must be a COM object pointer implementing the expected WebView2 handler interface.
func (i *ICoreWebView2Controller) AddZoomFactorChangedRaw(handler uintptr, token *_EventRegistrationToken) error {
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

func (i *ICoreWebView2Controller) RemoveZoomFactorChanged(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveZoomFactorChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller) SetBoundsAndZoomFactor(bounds w32.Rect, zoomFactor float64) error {
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

// AddMoveFocusRequestedRaw registers a move-focus-requested handler.
func (i *ICoreWebView2Controller) AddMoveFocusRequestedRaw(handler uintptr, token *_EventRegistrationToken) error {
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

func (i *ICoreWebView2Controller) RemoveMoveFocusRequested(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveMoveFocusRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// AddGotFocusRaw registers a got-focus handler.
func (i *ICoreWebView2Controller) AddGotFocusRaw(handler uintptr, token *_EventRegistrationToken) error {
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

func (i *ICoreWebView2Controller) RemoveGotFocus(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveGotFocus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// AddLostFocusRaw registers a lost-focus handler.
func (i *ICoreWebView2Controller) AddLostFocusRaw(handler uintptr, token *_EventRegistrationToken) error {
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

func (i *ICoreWebView2Controller) RemoveLostFocus(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveLostFocus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller) RemoveAcceleratorKeyPressed(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveAcceleratorKeyPressed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller) GetParentWindow() (uintptr, error) {
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

func (i *ICoreWebView2Controller) PutParentWindow(hwnd uintptr) error {
	_, _, err := i.vtbl.PutParentWindow.Call(
		uintptr(unsafe.Pointer(i)),
		hwnd,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller) Close() error {
	_, _, err := i.vtbl.Close.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller) GetCoreWebView2() (*ICoreWebView2, error) {
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

func (i *ICoreWebView2Controller) GetICoreWebView2Controller2() *ICoreWebView2Controller2 {

	var result *ICoreWebView2Controller2

	iidICoreWebView2Controller2 := NewGUID("{c979903e-d4ca-4228-92eb-47ee3fa96eab}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Controller2)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Controller) NotifyParentWindowPositionChanged() error {
	var err error
	_, _, err = i.vtbl.NotifyParentWindowPositionChanged.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller) MoveFocus(reason uintptr) error {
	var err error
	_, _, err = i.vtbl.MoveFocus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(reason),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
