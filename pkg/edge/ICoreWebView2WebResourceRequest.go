package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2WebResourceRequestVtbl struct {
	_IUnknownVtbl
	GetUri     ComProc
	PutUri     ComProc
	GetMethod  ComProc
	PutMethod  ComProc
	GetContent ComProc
	PutContent ComProc
	GetHeaders ComProc
}

type ICoreWebView2WebResourceRequest struct {
	vtbl *_ICoreWebView2WebResourceRequestVtbl
}

func (i *ICoreWebView2WebResourceRequest) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2WebResourceRequest) GetUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _uri *uint16
	_, _, err = i.vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_uri)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	uri := windows.UTF16PtrToString(_uri)
	windows.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

func (i *ICoreWebView2WebResourceRequest) PutUri(uri string) error {
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.PutUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2WebResourceRequest) GetMethod() (string, error) {
	var _method *uint16
	_, _, err := i.vtbl.GetMethod.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_method)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	method := windows.UTF16PtrToString(_method)
	windows.CoTaskMemFree(unsafe.Pointer(_method))
	return method, nil
}

func (i *ICoreWebView2WebResourceRequest) PutMethod(method string) error {
	_method, err := windows.UTF16PtrFromString(method)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.PutMethod.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_method)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// GetContent returns the underlying IStream* for the request body.
func (i *ICoreWebView2WebResourceRequest) GetContent() (uintptr, error) {
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

// PutContent sets the underlying IStream* for the request body.
func (i *ICoreWebView2WebResourceRequest) PutContent(stream uintptr) error {
	_, _, err := i.vtbl.PutContent.Call(
		uintptr(unsafe.Pointer(i)),
		stream,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// GetHeaders returns the underlying ICoreWebView2HttpRequestHeaders*.
// The concrete type is not wrapped here; use COM directly if needed.
func (i *ICoreWebView2WebResourceRequest) GetHeaders() (uintptr, error) {
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
