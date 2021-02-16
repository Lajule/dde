package main

//curl -X POST -s --data-urlencode 'input=p { color : red; }' https://html-minifier.com/raw
/*
<!doctype html>
<html>
  <div class=matrix>
    <div id=do>
      <h1>Do</h1>
      <h5>Important/Urgent</h5>
      <div class=bar>
        <button class=add>Add</button>
        <input type=text>
      </div>
      <div class=tasks></div>
    </div>
    <div id=schedule>
      <h1>Schedule</h1>
      <h5>Important/Not Urgent</h5>
      <div class=bar>
        <button class=add>Add</button>
        <input type=text>
      </div>
      <div class=tasks></div>
    </div>
    <div id=delegate>
      <h1>Delegate</h1>
      <h5>Not Important/Urgent</h5>
      <div class=bar>
        <button class=add>Add</button>
        <input type=text>
      </div>
      <div class=tasks></div>
    </div>
    <div id=cancel>
      <h1>Cancel</h1>
      <h5>Not Important/Not Urgent</h5>
      <div class=bar>
        <button class=add>Add</button>
        <input type=text>
      </div>
      <div class=tasks></div>
    </div>
  </div>
  <div class=bar>
    <button id=quit>Quit</button>
    <button id=save>Save</button>
    <button id=clear>Clear</button>
  </div>
</html>
*/

//curl -X POST -s --data-urlencode 'input=p { color : red; }' https://cssminifier.com/raw
/*
html,
body {
  height: 100%;
  margin: 0;
  font-family: "Exo", sans-serif;
  font-size: 12px;
}
button {
  border: 0;
  background: none;
}
.matrix {
  display: flex;
  flex-wrap: wrap;
  height: calc(100% - 40px);
}
.matrix > div {
  flex-grow: 1;
  flex-basis: calc((100% / 2) - 40px);
  padding: 10px;
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
.matrix > div > .bar > input {
  flex-grow: 1;
  border: 0;
}
.matrix > div > .tasks {
  position: absolute;
  width: calc((100% / 2) - 20px);
  height: calc((100% / 2) - 180px);
  overflow-y: scroll;
}
.bar {
  display: flex;
  flex-direction: row-reverse;
  height: 40px;
}
*/

//curl -X POST -s --data-urlencode 'input=console.log( 1 );' https://javascript-minifier.com/raw
/*
window.addEventListener('DOMContentLoaded', event => {
    var style = document.createElement('style')
    style.innerHTML = ``
    document.getElementsByTagName('head')[0].appendChild(style)
    document.querySelectorAll('.add').forEach(btn => {
        var tasks = btn.parentNode.parentNode.getElementsByClassName('tasks')[0]
        var input = btn.parentNode.parentNode.getElementsByTagName('input')[0]
        var add = () => {
            var div = document.createElement('div')
            var box = document.createElement('input')
            box.type = 'checkbox'
            div.appendChild(box)
            var label = document.createElement('label')
            label.innerHTML = input.value
            div.appendChild(label)
            tasks.appendChild(div)
            tasks.scrollTop = tasks.scrollHeight
            input.value = ''
        }
        input.addEventListener('keydown', event => {
            if (event.keyCode == 13 && event.target.value != '') {
                add()
            }
        })
        btn.addEventListener('click', event => {
            if (input.value != '') {
                add()
            }
        })
    })
    document.getElementById('clear').addEventListener('click', event => {
        document.querySelectorAll('input[type=checkbox]').forEach(box => {
            if (box.checked) {
                box.parentNode.remove()
            }
        })
    })
    document.getElementById('save').addEventListener('click', event => {
        update({
            window: {
                w: window.innerWidth,
                h: window.innerHeight
            },
            tasks: {
                do: [],
                schedule: [],
                delegate: [],
                cancel: []
            }
        })
    })
    document.getElementById('quit').addEventListener('click', event => {
        terminate()
    })
})
*/

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

	w.Bind("update", func(data map[string]interface{}) {
		fmt.Printf("%#v\n", data)
	})

	w.Bind("terminate", func() {
		w.Terminate()
	})

	w.Init(`window.addEventListener("DOMContentLoaded",e=>{var t=document.createElement("style");t.innerHTML="body,html{height:100%;margin:0;font-family:Exo,sans-serif;font-size:12px}button{border:0;background:0 0}.matrix{display:flex;flex-wrap:wrap;height:calc(100% - 40px)}.matrix>div{flex-grow:1;flex-basis:calc((100% / 2) - 40px);padding:10px}#do{background:#90ee90}#schedule{background:#add8e6}#delegate{background:orange}#cancel{background:red}.matrix>div>.bar>input{flex-grow:1;border:0}.matrix>div>.tasks{position:absolute;width:calc((100% / 2) - 20px);height:calc((100% / 2) - 180px);overflow-y:scroll}.bar{display:flex;flex-direction:row-reverse;height:40px}",document.getElementsByTagName("head")[0].appendChild(t),document.querySelectorAll(".add").forEach(e=>{var t=e.parentNode.parentNode.getElementsByClassName("tasks")[0],d=e.parentNode.parentNode.getElementsByTagName("input")[0],a=()=>{var e=document.createElement("div"),a=document.createElement("input");a.type="checkbox",e.appendChild(a);var n=document.createElement("label");n.innerHTML=d.value,e.appendChild(n),t.appendChild(e),t.scrollTop=t.scrollHeight,d.value=""};d.addEventListener("keydown",e=>{13==e.keyCode&&""!=e.target.value&&a()}),e.addEventListener("click",e=>{""!=d.value&&a()})}),document.getElementById("clear").addEventListener("click",e=>{document.querySelectorAll("input[type=checkbox]").forEach(e=>{e.checked&&e.parentNode.remove()})}),document.getElementById("save").addEventListener("click",e=>{update({window:{w:window.innerWidth,h:window.innerHeight},tasks:{do:[],schedule:[],delegate:[],cancel:[]}})}),document.getElementById("quit").addEventListener("click",e=>{terminate()})});`)

	w.Navigate(`data:text/html,<!doctype html><html><div class=matrix><div id=do><h1>Do</h1><h5>Important/Urgent</h5><div class=bar> <button class=add>Add</button> <input type=text></div><div class=tasks></div></div><div id=schedule><h1>Schedule</h1><h5>Important/Not Urgent</h5><div class=bar> <button class=add>Add</button> <input type=text></div><div class=tasks></div></div><div id=delegate><h1>Delegate</h1><h5>Not Important/Urgent</h5><div class=bar> <button class=add>Add</button> <input type=text></div><div class=tasks></div></div><div id=cancel><h1>Cancel</h1><h5>Not Important/Not Urgent</h5><div class=bar> <button class=add>Add</button> <input type=text></div><div class=tasks></div></div></div><div class=bar> <button id=quit>Quit</button> <button id=save>Save</button> <button id=clear>Clear</button></div></html>`)

	w.Run()
}
