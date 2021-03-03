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

// Size Get window size
func (c Config) Size() (int, int) {
	if c.W == 0 || c.H == 0 {
		return 800, 600
	}

	return c.W, c.H
}

// Dump Dump configuration
func Dump(f string, w, h int, tasks Tasks) {
	c := Config{
		W: w,
		H: h,
		Tasks: tasks,
	}

	data, err := json.Marshal(c)
	if err != nil {
		return
	}

	ioutil.WriteFile(f, data, 0644)
}

// Load Load configuration file
func Load(f string) (int, int, Tasks) {
	c := Config{}

	data, err := ioutil.ReadFile(f)
	if err != nil {
		return 800, 600, Tasks{}
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		return 800, 600, Tasks{}
	}

	if c.W == 0 || c.H == 0 {
		return 800, 600, c.Tasks
	}

	return c.W, c.H, c.Tasks
}
