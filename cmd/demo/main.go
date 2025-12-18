package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/logicossoftware/go-webview2"
	"github.com/logicossoftware/go-webview2/pkg/edge"
)

func main() {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		DownloadStartingCallback: func(_ *edge.ICoreWebView2, args *edge.ICoreWebView2DownloadStartingEventArgs) {
			op, err := args.GetDownloadOperation()
			if err != nil {
				log.Printf("DownloadStarting: GetDownloadOperation failed: %v", err)
				return
			}
			uri, _ := op.GetUri()
			defaultPath, _ := args.GetResultFilePath()
			total, _ := op.GetTotalBytesToReceive()
			log.Printf("DownloadStarting: uri=%q totalBytes=%d defaultPath=%q", uri, total, defaultPath)

			// Hide the default download UI (so you can build your own).
			_ = args.PutHandled(true)

			// Example: redirect downloads to a temp folder, preserving the suggested filename.
			if defaultPath != "" {
				newPath := filepath.Join(os.TempDir(), filepath.Base(defaultPath))
				_ = args.PutResultFilePath(newPath)
				log.Printf("DownloadStarting: redirected to %q", newPath)
			}
		},
		WindowOptions: webview2.WindowOptions{
			Title:  "Minimal webview example",
			Width:  800,
			Height: 600,
			IconId: 2, // icon resource id
			Center: true,
		},
	})
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetSize(800, 600, webview2.HintFixed)
	w.Navigate("https://demo.smartscreen.msft.net/")
	w.Run()
}
