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

    const sections = ['do', 'schedule', 'delegate', 'cancel']

    const loadTasks = (configuration) => {
	const setById = id => {
	    const tasks = document.querySelector('#' + id + ' > .tasks')
	    if (configuration[id]) {
		configuration[id].forEach(task => {
		    tasks.appendChild(newTask(task.checked, task.label))
		})
	    }
	}

	sections.forEach(id => {
	    setById(id)
	})
    }

    const updateTasks = () => {
	const getById = id => Array.from(document.querySelectorAll('#' + id +' > .tasks > div')).map(task => {
	    return {
		checked: task.getElementsByTagName("input")[0].checked,
		label: task.getElementsByTagName("label")[0].innerHTML
	    }
	})

	update(sections.reduce((result, id) => {
	    result[id] = getById(id)
	    return result
	}, {}))
    }

    const newTask = (checked, text) => {
        const div = document.createElement('div')
	// Generate unique identifier
	div.id = '_' + Math.random().toString(36).substr(2, 9)
	// Allow drag and drop
	div.draggable = true
	div.addEventListener('dragstart', event => {
	    event.dataTransfer.setData('text/plain', event.target.id)
	})

        const box = document.createElement('input')
        box.type = 'checkbox'
	box.checked = checked
        div.appendChild(box)

        const label = document.createElement('label')
        label.innerHTML = text
	div.appendChild(label)

	return div
    }

    // Handle "Add" buttons click
    document.querySelectorAll('.add').forEach(btn => {
        const tasks = btn.parentNode.parentNode.getElementsByClassName('tasks')[0]
        const input = btn.parentNode.getElementsByTagName('input')[0]
        const addTask = () => {
	    tasks.appendChild(newTask(false, input.value))
	    // Move to the end list and reset input
            tasks.scrollTop = tasks.scrollHeight
            input.value = ''
        }

        input.addEventListener('keydown', event => {
	    // Add a task if 'enter' key is pressed
            if (event.keyCode == 13 && event.target.value != '') {
                addTask()
            }
        })

        btn.addEventListener('click', event => {
            if (input.value != '') {
                addTask()
            }
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

    // Handle drop, consider ".task" as dropzone
    document.querySelectorAll('.tasks').forEach(tasks => {
	tasks.addEventListener('dragover', event => {
	    event.preventDefault()
	})

	tasks.addEventListener('drop', event => {
	    event.preventDefault()
	    // Drop element only if target is a dropzone
	    if (event.target.classList.contains('tasks')) {
		const id = event.dataTransfer.getData('text/plain')
		event.target.appendChild(document.getElementById(id))
	    }
	})
    })

    // Persist changes
    document.addEventListener('change', event => {
	updateTasks()
    })

    // Load tasks from config
    load().then(result => {
	loadTasks(result)
    })
})
