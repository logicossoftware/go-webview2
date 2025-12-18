# Example

This folder contains a small runnable example application using `github.com/logicossoftware/go-webview2`.

## Run

From the repo root:

- `go run ./example`

It opens a window that demonstrates:

- `Bind()` Go functions callable from JavaScript
- Updating the native window title from JS via a bound Go function
- Logging from JS to Go

Note: This project is Windows-only because it wraps WebView2.
