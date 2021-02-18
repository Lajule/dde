package main

import (
	"flag"
	"fmt"

	"github.com/webview/webview"
)

//go:generate go run generator.go

// Version Program version
var Version = "development"

var debug = flag.Bool("d", false, "Debug mode")
var config = flag.String("c", ".dde.json", "Config filename")

func main() {
	flag.Parse()

	w := webview.New(*debug)
	defer w.Destroy()

	w.SetTitle(fmt.Sprintf("DDE %s", Version))
	w.SetSize(800, 600, webview.HintNone)

	w.Bind("load", func() Tasks {
		return Load(*config)
	})

	w.Bind("update", func(tasks Tasks) {
		tasks.Dump(*config)
	})

	w.Bind("terminate", func() {
		w.Terminate()
	})

	w.Init(string(Application))

	w.Navigate(`data:text/html,<!doctype html><html></html>`)
	w.Run()
}
