//go:build windows
// +build windows

package edge

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/logicossoftware/go-webview2/internal/w32"

	"github.com/logicossoftware/go-webview2/webviewloader"
	"golang.org/x/sys/windows"
)

func init() {
	runtime.LockOSThread()

	r, _, _ := w32.Ole32CoInitializeEx.Call(0, 2)
	if int(r) < 0 {
		log.Printf("Warning: CoInitializeEx call failed: E=%08x", r)
	}
}

type _EventRegistrationToken struct {
	Value int64
}

type CoreWebView2PermissionKind uint32

const (
	CoreWebView2PermissionKindUnknownPermission CoreWebView2PermissionKind = iota
	CoreWebView2PermissionKindMicrophone
	CoreWebView2PermissionKindCamera
	CoreWebView2PermissionKindGeolocation
	CoreWebView2PermissionKindNotifications
	CoreWebView2PermissionKindOtherSensors
	CoreWebView2PermissionKindClipboardRead
)

type CoreWebView2PermissionState uint32

const (
	CoreWebView2PermissionStateDefault CoreWebView2PermissionState = iota
	CoreWebView2PermissionStateAllow
	CoreWebView2PermissionStateDeny
)

func createCoreWebView2EnvironmentWithOptions(browserExecutableFolder, userDataFolder *uint16, environmentOptions uintptr, environmentCompletedHandle *iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler) (uintptr, error) {
	return webviewloader.CreateCoreWebView2EnvironmentWithOptions(
		browserExecutableFolder,
		userDataFolder,
		environmentOptions,
		uintptr(unsafe.Pointer(environmentCompletedHandle)),
	)
}

// IUnknown

type _IUnknownVtbl struct {
	QueryInterface ComProc
	AddRef         ComProc
	Release        ComProc
}

type _IUnknownImpl interface {
	QueryInterface(refiid, object uintptr) uintptr
	AddRef() uintptr
	Release() uintptr
}

// ICoreWebView2

type iCoreWebView2Vtbl struct {
	_IUnknownVtbl
	GetSettings                            ComProc
	GetSource                              ComProc
	Navigate                               ComProc
	NavigateToString                       ComProc
	AddNavigationStarting                  ComProc
	RemoveNavigationStarting               ComProc
	AddContentLoading                      ComProc
	RemoveContentLoading                   ComProc
	AddSourceChanged                       ComProc
	RemoveSourceChanged                    ComProc
	AddHistoryChanged                      ComProc
	RemoveHistoryChanged                   ComProc
	AddNavigationCompleted                 ComProc
	RemoveNavigationCompleted              ComProc
	AddFrameNavigationStarting             ComProc
	RemoveFrameNavigationStarting          ComProc
	AddFrameNavigationCompleted            ComProc
	RemoveFrameNavigationCompleted         ComProc
	AddScriptDialogOpening                 ComProc
	RemoveScriptDialogOpening              ComProc
	AddPermissionRequested                 ComProc
	RemovePermissionRequested              ComProc
	AddProcessFailed                       ComProc
	RemoveProcessFailed                    ComProc
	AddScriptToExecuteOnDocumentCreated    ComProc
	RemoveScriptToExecuteOnDocumentCreated ComProc
	ExecuteScript                          ComProc
	CapturePreview                         ComProc
	Reload                                 ComProc
	PostWebMessageAsJSON                   ComProc
	PostWebMessageAsString                 ComProc
	AddWebMessageReceived                  ComProc
	RemoveWebMessageReceived               ComProc
	CallDevToolsProtocolMethod             ComProc
	GetBrowserProcessID                    ComProc
	GetCanGoBack                           ComProc
	GetCanGoForward                        ComProc
	GoBack                                 ComProc
	GoForward                              ComProc
	GetDevToolsProtocolEventReceiver       ComProc
	Stop                                   ComProc
	AddNewWindowRequested                  ComProc
	RemoveNewWindowRequested               ComProc
	AddDocumentTitleChanged                ComProc
	RemoveDocumentTitleChanged             ComProc
	GetDocumentTitle                       ComProc
	AddHostObjectToScript                  ComProc
	RemoveHostObjectFromScript             ComProc
	OpenDevToolsWindow                     ComProc
	AddContainsFullScreenElementChanged    ComProc
	RemoveContainsFullScreenElementChanged ComProc
	GetContainsFullScreenElement           ComProc
	AddWebResourceRequested                ComProc
	RemoveWebResourceRequested             ComProc
	AddWebResourceRequestedFilter          ComProc
	RemoveWebResourceRequestedFilter       ComProc
	AddWindowCloseRequested                ComProc
	RemoveWindowCloseRequested             ComProc
}

type ICoreWebView2 struct {
	vtbl *iCoreWebView2Vtbl
}

func (i *ICoreWebView2) GetSettings() (*ICoreWebViewSettings, error) {
	var err error
	var settings *ICoreWebViewSettings
	_, _, err = i.vtbl.GetSettings.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&settings)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return settings, nil
}

func (i *ICoreWebView2) Navigate(url string) error {
	_url, err := windows.UTF16PtrFromString(url)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.Navigate.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_url)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) NavigateToString(htmlContent string) error {
	_html, err := windows.UTF16PtrFromString(htmlContent)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.NavigateToString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_html)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// AddScriptToExecuteOnDocumentCreated adds a script that will run on every document.
// This helper intentionally ignores the optional completion handler.
func (i *ICoreWebView2) AddScriptToExecuteOnDocumentCreated(javaScript string) error {
	_js, err := windows.UTF16PtrFromString(javaScript)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.AddScriptToExecuteOnDocumentCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_js)),
		0,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// ExecuteScript runs the given JavaScript in the current document.
// This helper intentionally ignores the optional completion handler.
func (i *ICoreWebView2) ExecuteScript(javaScript string) error {
	_js, err := windows.UTF16PtrFromString(javaScript)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.ExecuteScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_js)),
		0,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) PostWebMessageAsString(message string) error {
	_msg, err := windows.UTF16PtrFromString(message)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.PostWebMessageAsString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_msg)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) PostWebMessageAsJSON(messageJSON string) error {
	_msg, err := windows.UTF16PtrFromString(messageJSON)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.PostWebMessageAsJSON.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_msg)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddWebMessageReceived(eventHandler *iCoreWebView2WebMessageReceivedEventHandler, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddWebMessageReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddPermissionRequested(eventHandler *iCoreWebView2PermissionRequestedEventHandler, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddPermissionRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddWebResourceRequested(eventHandler *iCoreWebView2WebResourceRequestedEventHandler, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddWebResourceRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// ICoreWebView2Environment

type iCoreWebView2EnvironmentVtbl struct {
	_IUnknownVtbl
	CreateCoreWebView2Controller     ComProc
	CreateWebResourceResponse        ComProc
	GetBrowserVersionString          ComProc
	AddNewBrowserVersionAvailable    ComProc
	RemoveNewBrowserVersionAvailable ComProc
}

type ICoreWebView2Environment struct {
	vtbl *iCoreWebView2EnvironmentVtbl
}

func (e *ICoreWebView2Environment) CreateWebResourceResponse(content []byte, statusCode int, reasonPhrase string, headers string) (*ICoreWebView2WebResourceResponse, error) {
	var err error
	var stream uintptr

	if len(content) > 0 {
		// Create stream for response
		stream, err = w32.SHCreateMemStream(content)
		if err != nil {
			return nil, err
		}
	}

	// Convert string 'uri' to *uint16
	_reason, err := windows.UTF16PtrFromString(reasonPhrase)
	if err != nil {
		return nil, err
	}
	// Convert string 'uri' to *uint16
	_headers, err := windows.UTF16PtrFromString(headers)
	if err != nil {
		return nil, err
	}
	var response *ICoreWebView2WebResourceResponse
	_, _, err = e.vtbl.CreateWebResourceResponse.Call(
		uintptr(unsafe.Pointer(e)),
		stream,
		uintptr(statusCode),
		uintptr(unsafe.Pointer(_reason)),
		uintptr(unsafe.Pointer(_headers)),
		uintptr(unsafe.Pointer(&response)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return response, nil

}

// ICoreWebView2WebMessageReceivedEventArgs

type iCoreWebView2WebMessageReceivedEventArgsVtbl struct {
	_IUnknownVtbl
	GetSource                ComProc
	GetWebMessageAsJSON      ComProc
	TryGetWebMessageAsString ComProc
}

type iCoreWebView2WebMessageReceivedEventArgs struct {
	vtbl *iCoreWebView2WebMessageReceivedEventArgsVtbl
}

func (i *iCoreWebView2WebMessageReceivedEventArgs) TryGetWebMessageAsString() (string, error) {
	var msg *uint16
	_, _, err := i.vtbl.TryGetWebMessageAsString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&msg)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	if msg == nil {
		return "", nil
	}
	res := w32.Utf16PtrToString(msg)
	windows.CoTaskMemFree(unsafe.Pointer(msg))
	return res, nil
}

// ICoreWebView2PermissionRequestedEventArgs

type iCoreWebView2PermissionRequestedEventArgsVtbl struct {
	_IUnknownVtbl
	GetURI             ComProc
	GetPermissionKind  ComProc
	GetIsUserInitiated ComProc
	GetState           ComProc
	PutState           ComProc
	GetDeferral        ComProc
}

type iCoreWebView2PermissionRequestedEventArgs struct {
	vtbl *iCoreWebView2PermissionRequestedEventArgsVtbl
}

func (i *iCoreWebView2PermissionRequestedEventArgs) GetPermissionKind() (CoreWebView2PermissionKind, error) {
	var kind CoreWebView2PermissionKind
	_, _, err := i.vtbl.GetPermissionKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return kind, nil
}

func (i *iCoreWebView2PermissionRequestedEventArgs) PutState(state CoreWebView2PermissionState) error {
	_, _, err := i.vtbl.PutState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(state),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler

type iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl interface {
	_IUnknownImpl
	EnvironmentCompleted(res uintptr, env *ICoreWebView2Environment) uintptr
}

type iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler struct {
	vtbl *iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl
	impl iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl
}

func _ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownQueryInterface(this *iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownAddRef(this *iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownRelease(this *iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInvoke(this *iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler, res uintptr, env *ICoreWebView2Environment) uintptr {
	return this.impl.EnvironmentCompleted(res, env)
}

var iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerFn = iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInvoke),
}

func newICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler(impl iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl) *iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler {
	return &iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler{
		vtbl: &iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerFn,
		impl: impl,
	}
}

// ICoreWebView2WebMessageReceivedEventHandler

type iCoreWebView2WebMessageReceivedEventHandlerImpl interface {
	_IUnknownImpl
	MessageReceived(sender *ICoreWebView2, args *iCoreWebView2WebMessageReceivedEventArgs) uintptr
}

type iCoreWebView2WebMessageReceivedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type iCoreWebView2WebMessageReceivedEventHandler struct {
	vtbl *iCoreWebView2WebMessageReceivedEventHandlerVtbl
	impl iCoreWebView2WebMessageReceivedEventHandlerImpl
}

func _ICoreWebView2WebMessageReceivedEventHandlerIUnknownQueryInterface(this *iCoreWebView2WebMessageReceivedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2WebMessageReceivedEventHandlerIUnknownAddRef(this *iCoreWebView2WebMessageReceivedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2WebMessageReceivedEventHandlerIUnknownRelease(this *iCoreWebView2WebMessageReceivedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2WebMessageReceivedEventHandlerInvoke(this *iCoreWebView2WebMessageReceivedEventHandler, sender *ICoreWebView2, args *iCoreWebView2WebMessageReceivedEventArgs) uintptr {
	return this.impl.MessageReceived(sender, args)
}

var iCoreWebView2WebMessageReceivedEventHandlerFn = iCoreWebView2WebMessageReceivedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerInvoke),
}

func newICoreWebView2WebMessageReceivedEventHandler(impl iCoreWebView2WebMessageReceivedEventHandlerImpl) *iCoreWebView2WebMessageReceivedEventHandler {
	return &iCoreWebView2WebMessageReceivedEventHandler{
		vtbl: &iCoreWebView2WebMessageReceivedEventHandlerFn,
		impl: impl,
	}
}

// ICoreWebView2PermissionRequestedEventHandler

type iCoreWebView2PermissionRequestedEventHandlerImpl interface {
	_IUnknownImpl
	PermissionRequested(sender *ICoreWebView2, args *iCoreWebView2PermissionRequestedEventArgs) uintptr
}

type iCoreWebView2PermissionRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type iCoreWebView2PermissionRequestedEventHandler struct {
	vtbl *iCoreWebView2PermissionRequestedEventHandlerVtbl
	impl iCoreWebView2PermissionRequestedEventHandlerImpl
}

func _ICoreWebView2PermissionRequestedEventHandlerIUnknownQueryInterface(this *iCoreWebView2PermissionRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2PermissionRequestedEventHandlerIUnknownAddRef(this *iCoreWebView2PermissionRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2PermissionRequestedEventHandlerIUnknownRelease(this *iCoreWebView2PermissionRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2PermissionRequestedEventHandlerInvoke(this *iCoreWebView2PermissionRequestedEventHandler, sender *ICoreWebView2, args *iCoreWebView2PermissionRequestedEventArgs) uintptr {
	return this.impl.PermissionRequested(sender, args)
}

var iCoreWebView2PermissionRequestedEventHandlerFn = iCoreWebView2PermissionRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2PermissionRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2PermissionRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2PermissionRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2PermissionRequestedEventHandlerInvoke),
}

func newICoreWebView2PermissionRequestedEventHandler(impl iCoreWebView2PermissionRequestedEventHandlerImpl) *iCoreWebView2PermissionRequestedEventHandler {
	return &iCoreWebView2PermissionRequestedEventHandler{
		vtbl: &iCoreWebView2PermissionRequestedEventHandlerFn,
		impl: impl,
	}
}

func (i *ICoreWebView2) AddWebResourceRequestedFilter(uri string, resourceContext COREWEBVIEW2_WEB_RESOURCE_CONTEXT) error {
	var err error
	// Convert string 'uri' to *uint16
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.AddWebResourceRequestedFilter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(resourceContext),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
func (i *ICoreWebView2) AddNavigationCompleted(eventHandler *ICoreWebView2NavigationCompletedEventHandler, token *_EventRegistrationToken) error {
	var err error
	_, _, err = i.vtbl.AddNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveNavigationCompleted(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GetSource() (string, error) {
	var src *uint16
	_, _, err := i.vtbl.GetSource.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&src)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	res := windows.UTF16PtrToString(src)
	windows.CoTaskMemFree(unsafe.Pointer(src))
	return res, nil
}

func (i *ICoreWebView2) Reload() error {
	_, _, err := i.vtbl.Reload.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// CapturePreviewRaw captures a preview image of the current page.
// stream must be an IStream* destination. handler must be a COM object pointer implementing ICoreWebView2CapturePreviewCompletedHandler.
func (i *ICoreWebView2) CapturePreviewRaw(imageFormat COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT, stream uintptr, handler uintptr) error {
	_, _, err := i.vtbl.CapturePreview.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(imageFormat),
		stream,
		handler,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) Stop() error {
	_, _, err := i.vtbl.Stop.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GetBrowserProcessID() (uint32, error) {
	var pid uint32
	_, _, err := i.vtbl.GetBrowserProcessID.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pid)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return pid, nil
}

func (i *ICoreWebView2) GetCanGoBack() (bool, error) {
	var can bool
	_, _, err := i.vtbl.GetCanGoBack.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&can)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return can, nil
}

func (i *ICoreWebView2) GetCanGoForward() (bool, error) {
	var can bool
	_, _, err := i.vtbl.GetCanGoForward.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&can)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return can, nil
}

func (i *ICoreWebView2) GoBack() error {
	_, _, err := i.vtbl.GoBack.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GoForward() error {
	_, _, err := i.vtbl.GoForward.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GetDocumentTitle() (string, error) {
	var title *uint16
	_, _, err := i.vtbl.GetDocumentTitle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&title)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	res := windows.UTF16PtrToString(title)
	windows.CoTaskMemFree(unsafe.Pointer(title))
	return res, nil
}

func (i *ICoreWebView2) OpenDevToolsWindow() error {
	_, _, err := i.vtbl.OpenDevToolsWindow.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GetContainsFullScreenElement() (bool, error) {
	var contains bool
	_, _, err := i.vtbl.GetContainsFullScreenElement.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&contains)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return contains, nil
}

func (i *ICoreWebView2) AddContainsFullScreenElementChangedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddContainsFullScreenElementChanged.Call(
		uintptr(unsafe.Pointer(i)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveContainsFullScreenElementChanged(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveContainsFullScreenElementChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveWebResourceRequestedFilter(uri string, resourceContext COREWEBVIEW2_WEB_RESOURCE_CONTEXT) error {
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.RemoveWebResourceRequestedFilter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(resourceContext),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveWebMessageReceived(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveWebMessageReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemovePermissionRequested(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemovePermissionRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveWebResourceRequested(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveWebResourceRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// Generic event registration helpers (Raw)

func (i *ICoreWebView2) AddNavigationStartingRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddNavigationStarting.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveNavigationStarting(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveNavigationStarting.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddContentLoadingRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddContentLoading.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveContentLoading(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveContentLoading.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddSourceChangedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddSourceChanged.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveSourceChanged(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveSourceChanged.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddHistoryChangedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddHistoryChanged.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveHistoryChanged(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveHistoryChanged.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddFrameNavigationStartingRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddFrameNavigationStarting.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveFrameNavigationStarting(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveFrameNavigationStarting.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddFrameNavigationCompletedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddFrameNavigationCompleted.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveFrameNavigationCompleted(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveFrameNavigationCompleted.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddScriptDialogOpeningRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddScriptDialogOpening.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveScriptDialogOpening(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveScriptDialogOpening.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddProcessFailedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddProcessFailed.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveProcessFailed(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveProcessFailed.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// RemoveScriptToExecuteOnDocumentCreated removes a previously added script by ID.
func (i *ICoreWebView2) RemoveScriptToExecuteOnDocumentCreated(scriptID string) error {
	_id, err := windows.UTF16PtrFromString(scriptID)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.RemoveScriptToExecuteOnDocumentCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_id)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// CallDevToolsProtocolMethodRaw calls a DevTools Protocol method.
// handler must be a COM object pointer implementing ICoreWebView2CallDevToolsProtocolMethodCompletedHandler.
func (i *ICoreWebView2) CallDevToolsProtocolMethodRaw(methodName string, parametersAsJSON string, handler uintptr) error {
	_method, err := windows.UTF16PtrFromString(methodName)
	if err != nil {
		return err
	}
	_params, err := windows.UTF16PtrFromString(parametersAsJSON)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.CallDevToolsProtocolMethod.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_method)),
		uintptr(unsafe.Pointer(_params)),
		handler,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// GetDevToolsProtocolEventReceiver returns the underlying ICoreWebView2DevToolsProtocolEventReceiver* for an event name.
func (i *ICoreWebView2) GetDevToolsProtocolEventReceiver(eventName string) (uintptr, error) {
	_name, err := windows.UTF16PtrFromString(eventName)
	if err != nil {
		return 0, err
	}
	var receiver uintptr
	_, _, err = i.vtbl.GetDevToolsProtocolEventReceiver.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&receiver)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return receiver, nil
}

func (i *ICoreWebView2) AddNewWindowRequestedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddNewWindowRequested.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveNewWindowRequested(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveNewWindowRequested.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddDocumentTitleChangedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddDocumentTitleChanged.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveDocumentTitleChanged(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveDocumentTitleChanged.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// AddHostObjectToScriptRaw adds a host object via a raw COM object pointer.
func (i *ICoreWebView2) AddHostObjectToScriptRaw(name string, rawObject uintptr) error {
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.AddHostObjectToScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		rawObject,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveHostObjectFromScript(name string) error {
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	_, _, err = i.vtbl.RemoveHostObjectFromScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddWindowCloseRequestedRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := i.vtbl.AddWindowCloseRequested.Call(uintptr(unsafe.Pointer(i)), handler, uintptr(unsafe.Pointer(token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveWindowCloseRequested(token _EventRegistrationToken) error {
	_, _, err := i.vtbl.RemoveWindowCloseRequested.Call(uintptr(unsafe.Pointer(i)), uintptr(unsafe.Pointer(&token)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

// ICoreWebView2Environment helpers

func (e *ICoreWebView2Environment) GetBrowserVersionString() (string, error) {
	var ver *uint16
	_, _, err := e.vtbl.GetBrowserVersionString.Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(&ver)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	res := windows.UTF16PtrToString(ver)
	windows.CoTaskMemFree(unsafe.Pointer(ver))
	return res, nil
}

// CreateCoreWebView2ControllerRaw starts controller creation.
// completedHandler must be a COM object pointer implementing ICoreWebView2CreateCoreWebView2ControllerCompletedHandler.
func (e *ICoreWebView2Environment) CreateCoreWebView2ControllerRaw(parentWindow uintptr, completedHandler uintptr) error {
	_, _, err := e.vtbl.CreateCoreWebView2Controller.Call(
		uintptr(unsafe.Pointer(e)),
		parentWindow,
		completedHandler,
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (e *ICoreWebView2Environment) AddNewBrowserVersionAvailableRaw(handler uintptr, token *_EventRegistrationToken) error {
	_, _, err := e.vtbl.AddNewBrowserVersionAvailable.Call(
		uintptr(unsafe.Pointer(e)),
		handler,
		uintptr(unsafe.Pointer(token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (e *ICoreWebView2Environment) RemoveNewBrowserVersionAvailable(token _EventRegistrationToken) error {
	_, _, err := e.vtbl.RemoveNewBrowserVersionAvailable.Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
