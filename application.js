window.addEventListener('DOMContentLoaded', event => {
    document.getElementsByTagName('body')[0].innerHTML = `
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
  <button id=clear>Clear</button>
</div>`

    var style = document.createElement('style')
    style.innerHTML = `
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
}`
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

    document.getElementById('quit').addEventListener('click', event => {
	var getById = id => Array.from(document.querySelectorAll('#' + id +' > .tasks > div')).map(task => {
	    return {
		checked: task.getElementsByTagName("input")[0].checked,
		label: task.getElementsByTagName("label")[0].innerHTML
	    }
	})
        terminate({
            window: {
                w: window.innerWidth,
                h: window.innerHeight
            },
            tasks: {
                do: getById('do'),
                schedule: getById('schedule'),
                delegate: getById('delegate'),
                cancel: getById('cancel'),
            }
        })
    })

    load().then(configuration => {
	var setById = id => {
	    var tasks = document.querySelector('#' + id + ' > .tasks')
	    if (configuration.tasks[id]) {
		configuration.tasks[id].forEach(task => {
		    var div = document.createElement('div')
		    var box = document.createElement('input')
		    box.type = 'checkbox'
		    box.checked = task.checked
		    div.appendChild(box)
		    var label = document.createElement('label')
		    label.innerHTML = task.label
		    div.appendChild(label)
		    tasks.appendChild(div)
		})
	    }
	}
	if (configuration.tasks) {
	    setById('do')
	    setById('schedule')
	    setById('delegate')
	    setById('cancel')
	}
    })
})
