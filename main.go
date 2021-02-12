package main

import (
	"log"

	"github.com/webview/webview"
)

// Version Program version
var Version = "development"

func main() {
	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("dde")

	w.Bind("noop", func() string {
		log.Println("hello")
		return Version
	})

	w.Bind("add", func(a, b int) int {
		return a + b
	})

	w.Bind("quit", func() {
		w.Terminate()
	})

	w.Navigate(`data:text/html,
		<!doctype html>
		<html>
			<body>
			<button onclick="quit();">Quit</button>
			<script>
				window.onload = function () {
                                        var p = document.createElement('p');
					p.innerHTML = ` + "`hello, ${navigator.userAgent}`" + `;

					document.body.appendChild(p)

					noop().then(function (res) {
						console.log('noop res', res);
						add(1, 2).then(function (res) {
							console.log('add res', res);
						});
					});
				};
			</script>
			</body>
		</html>
	`)

	w.Run()
}
