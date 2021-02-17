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
	w.SetSize(800, 600, webview.HintMin)

	w.Bind("update", func(data map[string]interface{}) {
		fmt.Printf("%#v\n", data)
	})

	w.Bind("terminate", func() {
		w.Terminate()
	})

	w.Init(string(Application))

	w.Navigate(`data:text/html,<!doctype html><html></html>`)

	w.Run()
}
