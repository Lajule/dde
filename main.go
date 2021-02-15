package main

import (
	"fmt"

	"github.com/webview/webview"
)

/*
html,
body {
  height: 100%;
  margin: 0;
  font-family: "Exo", sans-serif;
  font-size: 16px;
}
.matrix {
  display: flex;
  flex-wrap: wrap;
  height: calc(100% - 40px);
}
.matrix > div {
  flex-grow: 1;
  flex-basis: calc((100% / 2) - 40px);
  padding: 20px;
}
#do {
  background: lightgreen;
}
#schedule {
  background: lightblue;
}
#delegate {
  background: orange;
}
#cancel {
  background: red;
}
.bar {
  display: flex;
  flex-direction: row-reverse;
  height: 40px;
}
*/

// Version Program version
var Version = "development"

func main() {
	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle(fmt.Sprintf("dde %s", Version))
	w.SetSize(800, 600, webview.HintNone)

	w.Bind("save", func() {
	})

	w.Bind("quit", func() {
		w.Terminate()
	})

	w.Eval(`
    window.addEventListener('DOMContentLoaded', event => {
      var btns = document.querySelectorAll('.add')
      btns.forEach(btn => {
        btn.addEventListener('click', event => {
          section = btn.parentNode
          var task = section.getElementsByTagName('input')[0]
          if (task.value != '') {
            var div = document.createElement('div')
            var input = document.createElement('input')
            input.type = 'checkbox'
            div.appendChild(input)
            var label = document.createElement('label')
            label.innerHTML = task.value
            div.appendChild(label)
            section.appendChild(div)
            task.value = ''
          }
        })
      })
      clear = document.getElementById('clear')
      clear.addEventListener('click', event => {
        boxes = document.querySelectorAll('input[type=checkbox]')
      })
    })
`)

	w.Navigate(`data:text/html,
<!doctype html>
<html>
  <div class=matrix>
    <div id=do>
      <h1>Do</h1>
      <h5>Important/Urgent</h5>
      <button class=add>Add</button>
      <input type=text>
      <div class=tasks></div>
    </div>
    <div id=schedule>
      <h1>Schedule</h1>
      <h5>Important/Not Urgent</h5>
      <button class=add>Add</button>
      <input type=text>
      <div class=tasks></div>
    </div>
    <div id=delegate>
      <h1>Delegate</h1>
      <h5>Not Important/Urgent</h5>
      <button class=add>Add</button>
      <input type=text>
      <div class=tasks></div>
    </div>
    <div id=cancel>
      <h1>Cancel</h1>
      <h5>Not Important/Not Urgent</h5>
      <button class=add>Add</button>
      <input type=text>
      <div class=tasks></div>
    </div>
  </div>
  <div class=bar>
    <button id=quit>Quit</button>
    <button id=save>Save</button>
    <button id=clear>Clear</button>
  </div>
</html>`)

	w.Run()
}
