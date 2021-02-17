package main

import (
	"encoding/json"
	"io/ioutil"
)

// Window Window size
type Window struct {
	W int
	H int
}

// Tasks All tasks
type Tasks struct {
	Do       []*Task `json:"do"`
	Schedule []*Task `json:"schedule"`
	Delegate []*Task `json:"delegate"`
	Cancel   []*Task `json:"cancel"`
}

// Task A task
type Task struct {
	Checked bool   `json:"checked"`
	Label   string `json:"label"`
}

// Configuration program configuration
type Configuration struct {
	Window *Window `json:"window"`
	Tasks  *Tasks  `json:"tasks"`
}

// Dump Dump configuration
func (c Configuration) Dump(f string) {
	data, err := json.Marshal(c)
	if err != nil {
		return
	}

	ioutil.WriteFile(f, data, 0644)
}

// Load Load configuration file
func Load(f string) Configuration {
	configuration := Configuration{}

	data, err := ioutil.ReadFile(f)
	if err != nil {
		return configuration
	}

	_ = json.Unmarshal(data, &configuration)
	return configuration
}
