// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("application.js")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("application.go", []byte(fmt.Sprintf(`package main

var Application = %#v
`, data)), 0644)
	if err != nil {
		panic(err)
	}
}
