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

	c := Load(*file)

	w := webview.New(*debug)
	defer w.Destroy()

	w.SetTitle(fmt.Sprintf("DDE %s", Version))
	w.SetSize(c.W, c.H, webview.HintNone)

	w.Bind("load", func() Tasks {
		return c.Tasks
	})

	w.Bind("update", func(config Config) {
		config.Dump(*file)
	})

	w.Bind("terminate", func() {
		w.Terminate()
	})

	w.Init(app)

	w.Navigate(`data:text/html,<!doctype html><html></html>`)
	w.Run()
}
