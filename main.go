package main

import (
	"fmt"

	"github.com/webview/webview"
)

// Version Program version
var Version = "development"

func main() {
	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle(fmt.Sprintf("dde %s", Version))
	w.SetSize(800, 600, webview.HintMin)

	w.Bind("save", func() {
	})

	w.Bind("terminate", func() {
		w.Terminate()
	})

	w.Init(`
    window.addEventListener('DOMContentLoaded', event => {
      var style = document.createElement('style')
      style.innerHTML = '\
* {\
  margin: 0;\
  padding: 0;\
  box-sizing: border-box;\
}\
html,\
body {\
  height: 100%;\
  font-family: "Exo", sans-serif;\
  font-size: 12px;\
}\
.matrix {\
  display: flex;\
  flex-wrap: wrap;\
  height: calc(100% - 40px);\
}\
.matrix > div {\
  flex-grow: 1;\
  flex-basis: calc(100% / 2);\
  height: calc((100% / 2) - 40px);\
  overflow-y: scroll;\
}\
.bar {\
  display: flex;\
  flex-direction: row-reverse;\
  height: 40px;\
}'
      var head = document.getElementsByTagName('head')[0];
      head.appendChild(style)
      var btns = document.querySelectorAll('.add')
      btns.forEach(btn => {
        btn.addEventListener('click', event => {
          var section = btn.parentNode
          var tasks = section.getElementsByClassName('tasks')[0]
          var input = section.getElementsByTagName('input')[0]
          if (input.value != '') {
            var div = document.createElement('div')
            var box = document.createElement('input')
            box.type = 'checkbox'
            div.appendChild(box)
            var label = document.createElement('label')
            label.innerHTML = input.value
            div.appendChild(label)
            tasks.appendChild(div)
            input.value = ''
          }
        })
      })
      var clear = document.getElementById('clear')
      clear.addEventListener('click', event => {
        var boxes = document.querySelectorAll('input[type=checkbox]')
        boxes.forEach(box => {
          if (box.checked) {
            var node = box.parentNode
            node.remove()
          }
        })
      })
      var quit = document.getElementById('quit')
      quit.addEventListener('click', event => { terminate() })
    })`)

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
