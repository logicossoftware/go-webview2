package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2WebResourceResponseVtbl struct {
	_IUnknownVtbl
	GetContent      ComProc
	PutContent      ComProc
	GetHeaders      ComProc
	GetStatusCode   ComProc
	PutStatusCode   ComProc
	GetReasonPhrase ComProc
	PutReasonPhrase ComProc
}

type ICoreWebView2WebResourceResponse struct {
	vtbl *_ICoreWebView2WebResourceResponseVtbl
}

func (i *ICoreWebView2WebResourceResponse) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

// GetContent returns the underlying IStream* for the response body.
func (i *ICoreWebView2WebResourceResponse) GetContent() (uintptr, error) {
	var stream uintptr
	_, _, err := i.vtbl.GetContent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&stream)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return stream, nil
}

// PutContent sets the underlying IStream* for the response body.
func (i *ICoreWebView2WebResourceResponse) PutContent(stream uintptr) error {
	_, _, err := i.vtbl.PutContent.Call(
		uintptr(unsafe.Pointer(i)),
		stream,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// GetHeaders returns the underlying ICoreWebView2HttpResponseHeaders*.
// The concrete type is not wrapped here; use COM directly if needed.
func (i *ICoreWebView2WebResourceResponse) GetHeaders() (uintptr, error) {
	var headers uintptr
	_, _, err := i.vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&headers)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return headers, nil
}

func (i *ICoreWebView2WebResourceResponse) GetStatusCode() (int32, error) {
	var code int32
	_, _, err := i.vtbl.GetStatusCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&code)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return code, nil
}

func (i *ICoreWebView2WebResourceResponse) PutStatusCode(code int32) error {
	_, _, err := i.vtbl.PutStatusCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(code),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2WebResourceResponse) GetReasonPhrase() (string, error) {
	var phrase *uint16
	_, _, err := i.vtbl.GetReasonPhrase.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&phrase)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	res := windows.UTF16PtrToString(phrase)
	windows.CoTaskMemFree(unsafe.Pointer(phrase))
	return res, nil
}

func (i *ICoreWebView2WebResourceResponse) PutReasonPhrase(phrase string) error {
	_phrase, err := windows.UTF16PtrFromString(phrase)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.PutReasonPhrase.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_phrase)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
