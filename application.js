window.addEventListener('DOMContentLoaded', event => {
    // Add DOM elements first
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

    // Then inject stylesheet
    const style = document.createElement('style')
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
  display: flex;
  flex-direction: column;
  flex: 1 1 calc((100% / 2) - 40px);
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
  flex: 1 1 auto;
  border: 0;
}
.matrix > div > .tasks {
  flex: 1 1 auto;
  padding: 1px;
  height: 100px;
  overflow-y: auto;
}
.bar {
  display: flex;
  flex-direction: row-reverse;
  height: 40px;
}`
    document.getElementsByTagName('head')[0].appendChild(style)

    // Persist changes into file
    const updateTasks = () => {
	const getById = id => Array.from(document.querySelectorAll('#' + id +' > .tasks > div')).map(task => {
	    return {
		checked: task.getElementsByTagName("input")[0].checked,
		label: task.getElementsByTagName("label")[0].innerHTML
	    }
	})
	update({
	    do: getById('do'),
	    schedule: getById('schedule'),
	    delegate: getById('delegate'),
	    cancel: getById('cancel')
	})
    }

    const newTask = (checked, text) => {
        const div = document.createElement('div')
	div.id = '_' + Math.random().toString(36).substr(2, 9)
	// Allow drag
	div.draggable = true
	div.addEventListener('dragstart', event => {
	    event.dataTransfer.setData('text/plain', event.target.id)
	})
        const box = document.createElement('input')
	box.id = '_' + Math.random().toString(36).substr(2, 9)
        box.type = 'checkbox'
	box.checked = checked
	box.addEventListener('change', event => {
	    updateTasks()
	})
        div.appendChild(box)
        const label = document.createElement('label')
	label.htmlFor = box.id
        label.innerHTML = text
	div.appendChild(label)
	return div
    }

    document.querySelectorAll('.tasks').forEach(tasks => {
	const button = tasks.parentNode.getElementsByTagName('button')[0]
        const input = tasks.parentNode.getElementsByTagName('input')[0]
	// Handle "Add" action
        const addTask = () => {
	    tasks.appendChild(newTask(false, input.value))
	    // Move to the end list and reset input
            tasks.scrollTop = tasks.scrollHeight
            input.value = ''
	    updateTasks()
        }
        input.addEventListener('keydown', event => {
	    // Add a task if 'enter' key is pressed
            if (event.keyCode == 13 && event.target.value != '') {
                addTask()
            }
        })
        button.addEventListener('click', event => {
            if (input.value != '') {
                addTask()
            }
        })
	// Handle drop
	tasks.parentNode.addEventListener('dragover', event => {
	    event.preventDefault()
	})
	tasks.parentNode.addEventListener('drop', event => {
	    event.preventDefault()
	    const id = event.dataTransfer.getData('text/plain')
	    tasks.appendChild(document.getElementById(id))
	    updateTasks()
	})
    })

    // Remove completed tasks from the DOM
    document.getElementById('clear').addEventListener('click', event => {
        document.querySelectorAll('input[type=checkbox]').forEach(box => {
            if (box.checked) {
                box.parentNode.remove()
            }
        })
	updateTasks()
    })

    // Kill the app
    document.getElementById('quit').addEventListener('click', event => {
        terminate()
    })

    // Load tasks from file
    load().then(result => {
	const setById = id => {
	    const tasks = document.querySelector('#' + id + ' > .tasks')
	    if (result[id]) {
		result[id].forEach(task => {
		    tasks.appendChild(newTask(task.checked, task.label))
		})
	    }
	}
	setById('do')
	setById('schedule')
	setById('delegate')
	setById('cancel')
    })
})
