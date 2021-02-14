package main

import (
	"fmt"

	"github.com/webview/webview"
)

const (
	CSS = `
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
  height: 40px;
}`
)

// Version Program version
var Version = "development"

func main() {
	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle(fmt.Sprintf("dde %s", Version))
	w.SetSize(800, 600, webview.HintNone)

	w.Bind("css", func() string {
		return CSS
	})

	w.Bind("quit", func() {
		w.Terminate()
	})

	w.Navigate(`data:text/html,
<!doctype html>
<html>
  <body>
    <div class="matrix">
      <div id="do">
        <h1>Do</h1>
        <h5>Important/Urgent</h5>
        <div>
          <input type="checkbox" id="do-0" checked>
          <label for="do-0">Do something in Golang</label>
        </div>
      </div>
      <div id="schedule">
        <h1>Schedule</h1>
        <h5>Important/Not Urgent</h5>
      </div>
      <div id="delegate">
        <h1>Delegate</h1>
        <h5>Not Important/Urgent</h5>
      </div>
      <div id="cancel">
        <h1>Cancel</h1>
        <h5>Not Important/Not Urgent</h5>
      </div>
    </div>
    <div class="bar">
      <button onclick="quit();">Quit</button>
    </div>
    <script>
      window.onload = function () {
        var ref = document.querySelector('script');
        css().then(function (code) {
          var el = document.createElement('style');
          el.innerHTML = code;
          ref.parentNode.insertBefore(el, ref);
        });
      };
    </script>
  </body>
</html>`)

	w.Run()
}
