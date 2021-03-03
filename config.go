package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config Application configuration
type Config struct {
	W     int    `json:"w"`
	H     int    `json:"h"`
	Tasks Tasks `json:"tasks"`
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

// Dump Dump configuration
func (c Config) Dump(f string) {
	data, err := json.Marshal(c)
	if err != nil {
		return
	}

	ioutil.WriteFile(f, data, 0644)
}

// Load Load configuration file
func Load(f string) Config {
	c := Config{}

	data, err := ioutil.ReadFile(f)
	if err != nil {
		return c
	}

	json.Unmarshal(data, &c)

	return c
}

