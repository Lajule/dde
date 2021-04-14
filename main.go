package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/webview/webview"
)

const (
	// MinWidth is the minimal window width.
	MinWidth = 800

	// MinHeight is the minimal window height.
	MinHeight = 600
)

// Input is used to load and save tasks.
type Input struct {
	Width  int   `json:"width"`
	Height int   `json:"height"`
	Tasks  Tasks `json:"tasks"`
}

// Tasks gives access to all tasks.
type Tasks struct {
	Do       []*Task `json:"do"`
	Schedule []*Task `json:"schedule"`
	Delegate []*Task `json:"delegate"`
	Cancel   []*Task `json:"cancel"`
}

// Task is a simple task.
type Task struct {
	Checked bool   `json:"checked"`
	Label   string `json:"label"`
}

var (
	// Version contains the program version.
	Version = "development"

	// Debug activates the debug mode.
	Debug = flag.Bool("d", false, "Debug mode")

	// Filename contains the input filename.
	Filename = flag.String("f", "dde.json", "Input file")
)

//go:embed app.js
var app string

func load() (int, int, Tasks) {
	input := Input{}

	data, err := os.ReadFile(*Filename)
	if err != nil {
		return 0, 0, Tasks{}
	}

	if err = json.Unmarshal(data, &input); err != nil {
		return 0, 0, Tasks{}
	}

	return input.Width, input.Height, input.Tasks
}

func dump(width, height int, tasks Tasks) {
	input := Input{
		Width:  width,
		Height: height,
		Tasks:  tasks,
	}

	data, err := json.Marshal(input)
	if err != nil {
		return
	}

	os.WriteFile(*Filename, data, 0644)
}

func main() {
	flag.Parse()

	width, height, tasks := load()

	window := webview.New(*Debug)
	defer window.Destroy()

	window.SetTitle(fmt.Sprintf("DDE %s", Version))
	window.SetSize(MinWidth, MinHeight, webview.HintMin)
	window.SetSize(width, height, webview.HintNone)

	window.Bind("load", func() Tasks {
		return tasks
	})

	window.Bind("update", dump)

	window.Bind("terminate", window.Terminate)

	window.Init(app)

	window.Navigate("data:text/html,<!doctype html><html></html>")

	window.Run()
}
