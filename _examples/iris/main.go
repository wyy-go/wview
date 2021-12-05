package main

import (
	"github.com/kataras/iris/v12"
	"github.com/wyy-go/wview/plugin/irisview"
)

func main() {
	app := iris.New()

	// Register the goview template engine.
	app.RegisterView(irisview.Default())

	app.Get("/", func(ctx iris.Context) {
		// Render with master.
		ctx.View("index", iris.Map{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	app.Get("/page", func(ctx iris.Context) {
		// Render only file, must full name with extension.
		ctx.View("page.html", iris.Map{"title": "Page file title!!"})
	})

	app.Listen(":9090")
}
