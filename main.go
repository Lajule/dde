package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/webview/webview"
)

// Version Program version
var Version = "development"

//go:embed app.js
var app string

var debug = flag.Bool("d", false, "Debug mode")
var file = flag.String("c", "config.json", "Configuration filename")

func main() {
	flag.Parse()

	w, h, t := Load(*file)

	wv := webview.New(*debug)
	defer wv.Destroy()

	wv.SetTitle(fmt.Sprintf("DDE %s", Version))

	wv.SetSize(w, h, webview.HintNone)

	wv.Bind("load", func() Tasks {
		return t
	})

	wv.Bind("update", func(width, height int, tasks Tasks) {
		Dump(*file, width, height, tasks)
	})

	wv.Bind("terminate", func() {
		wv.Terminate()
	})

	wv.Init(app)

	wv.Navigate("data:text/html,<!doctype html><html></html>")
	wv.Run()
}
