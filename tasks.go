package main

import (
	"encoding/json"
	"io/ioutil"
)

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

// Dump Dump tasks
func (t Tasks) Dump(f string) {
	data, err := json.Marshal(t)
	if err != nil {
		return
	}

	ioutil.WriteFile(f, data, 0644)
}

// Load Load tasks file
func Load(f string) Tasks {
	t := Tasks{}

	data, err := ioutil.ReadFile(f)
	if err != nil {
		return t
	}

	_ = json.Unmarshal(data, &t)
	return t
}
