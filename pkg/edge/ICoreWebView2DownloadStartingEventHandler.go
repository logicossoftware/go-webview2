package edge

type _ICoreWebView2DownloadStartingEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type iCoreWebView2DownloadStartingEventHandler struct {
	vtbl *_ICoreWebView2DownloadStartingEventHandlerVtbl
	impl _ICoreWebView2DownloadStartingEventHandlerImpl
}

func _ICoreWebView2DownloadStartingEventHandlerIUnknownQueryInterface(this *iCoreWebView2DownloadStartingEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DownloadStartingEventHandlerIUnknownAddRef(this *iCoreWebView2DownloadStartingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DownloadStartingEventHandlerIUnknownRelease(this *iCoreWebView2DownloadStartingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DownloadStartingEventHandlerInvoke(this *iCoreWebView2DownloadStartingEventHandler, sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr {
	return this.impl.DownloadStarting(sender, args)
}

type _ICoreWebView2DownloadStartingEventHandlerImpl interface {
	_IUnknownImpl
	DownloadStarting(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr
}

var _ICoreWebView2DownloadStartingEventHandlerFn = _ICoreWebView2DownloadStartingEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2DownloadStartingEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2DownloadStartingEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2DownloadStartingEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2DownloadStartingEventHandlerInvoke),
}

func newICoreWebView2DownloadStartingEventHandler(impl _ICoreWebView2DownloadStartingEventHandlerImpl) *iCoreWebView2DownloadStartingEventHandler {
	return &iCoreWebView2DownloadStartingEventHandler{
		vtbl: &_ICoreWebView2DownloadStartingEventHandlerFn,
		impl: impl,
	}
}
