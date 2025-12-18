package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2DownloadOperationVtbl struct {
	_IUnknownVtbl
	AddBytesReceivedChanged       ComProc
	RemoveBytesReceivedChanged    ComProc
	AddEstimatedEndTimeChanged    ComProc
	RemoveEstimatedEndTimeChanged ComProc
	AddStateChanged               ComProc
	RemoveStateChanged            ComProc
	GetUri                        ComProc
	GetContentDisposition         ComProc
	GetMimeType                   ComProc
	GetTotalBytesToReceive        ComProc
	GetBytesReceived              ComProc
	GetEstimatedEndTime           ComProc
	GetResultFilePath             ComProc
	GetState                      ComProc
	GetInterruptReason            ComProc
	Cancel                        ComProc
	Pause                         ComProc
	Resume                        ComProc
	GetCanResume                  ComProc
}

type ICoreWebView2DownloadOperation struct {
	vtbl *_ICoreWebView2DownloadOperationVtbl
}

func (i *ICoreWebView2DownloadOperation) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DownloadOperation) AddBytesReceivedChangedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddBytesReceivedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) RemoveBytesReceivedChanged(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveBytesReceivedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) AddEstimatedEndTimeChangedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddEstimatedEndTimeChanged.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) RemoveEstimatedEndTimeChanged(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveEstimatedEndTimeChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) AddStateChangedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddStateChanged.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) RemoveStateChanged(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveStateChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) GetUri() (string, error) {
	var uri *uint16
	_, _, err := i.vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&uri)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	if uri == nil {
		return "", nil
	}
	res := windows.UTF16PtrToString(uri)
	windows.CoTaskMemFree(unsafe.Pointer(uri))
	return res, nil
}

func (i *ICoreWebView2DownloadOperation) GetContentDisposition() (string, error) {
	var cd *uint16
	_, _, err := i.vtbl.GetContentDisposition.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cd)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	if cd == nil {
		return "", nil
	}
	res := windows.UTF16PtrToString(cd)
	windows.CoTaskMemFree(unsafe.Pointer(cd))
	return res, nil
}

func (i *ICoreWebView2DownloadOperation) GetMimeType() (string, error) {
	var mt *uint16
	_, _, err := i.vtbl.GetMimeType.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&mt)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	if mt == nil {
		return "", nil
	}
	res := windows.UTF16PtrToString(mt)
	windows.CoTaskMemFree(unsafe.Pointer(mt))
	return res, nil
}

func (i *ICoreWebView2DownloadOperation) GetTotalBytesToReceive() (int64, error) {
	var total int64
	_, _, err := i.vtbl.GetTotalBytesToReceive.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&total)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return total, nil
}

func (i *ICoreWebView2DownloadOperation) GetBytesReceived() (int64, error) {
	var received int64
	_, _, err := i.vtbl.GetBytesReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&received)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return received, nil
}

func (i *ICoreWebView2DownloadOperation) GetEstimatedEndTime() (string, error) {
	var et *uint16
	_, _, err := i.vtbl.GetEstimatedEndTime.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&et)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	if et == nil {
		return "", nil
	}
	res := windows.UTF16PtrToString(et)
	windows.CoTaskMemFree(unsafe.Pointer(et))
	return res, nil
}

func (i *ICoreWebView2DownloadOperation) GetResultFilePath() (string, error) {
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

func (i *ICoreWebView2DownloadOperation) GetState() (COREWEBVIEW2_DOWNLOAD_STATE, error) {
	var state COREWEBVIEW2_DOWNLOAD_STATE
	_, _, err := i.vtbl.GetState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&state)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return state, nil
}

func (i *ICoreWebView2DownloadOperation) GetInterruptReason() (COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON, error) {
	var reason COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON
	_, _, err := i.vtbl.GetInterruptReason.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&reason)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return reason, nil
}

func (i *ICoreWebView2DownloadOperation) Cancel() error {
	_, _, err := i.vtbl.Cancel.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) Pause() error {
	_, _, err := i.vtbl.Pause.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) Resume() error {
	_, _, err := i.vtbl.Resume.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) GetCanResume() (bool, error) {
	var can bool
	_, _, err := i.vtbl.GetCanResume.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&can)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return can, nil
}
