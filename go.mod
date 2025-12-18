module github.com/logicossoftware/go-webview2

go 1.25

require (
	github.com/logicossoftware/go-winloader v0.0.0-20250406163304-c1995be93bd1
	golang.org/x/sys v0.0.0-20210218145245-beda7e5e158e
)

replace github.com/logicossoftware/go-winloader => ./third_party/go-winloader
