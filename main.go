package main

import (
	"fmt"

	"github.com/webview/webview"
)

//go:generate go run generator.go

// Version Program version
var Version = "development"

func main() {
	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle(fmt.Sprintf("dde %s", Version))

	configuration := Load("dde.json")

	w.SetSize(configuration.Window.W, configuration.Window.H, webview.HintNone)

	w.Bind("load", func() Configuration {
		return configuration
	})

	w.Bind("terminate", func(newConfiguration Configuration) {
		newConfiguration.Dump("dde.json")
		w.Terminate()
	})

	w.Init(string(Application))

	w.Navigate(`data:text/html,<!doctype html><html></html>`)

	w.Run()
}
