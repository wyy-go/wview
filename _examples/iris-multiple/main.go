package main

import (
	"html/template"
	"time"

	"github.com/wyy-go/wview"
	"github.com/wyy-go/wview/plugin/irisview"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	fm := make(template.FuncMap)
	fm["copy"] = func() string {
		return time.Now().Format("2006")
	}

	// Register a new template engine.
	app.RegisterView(irisview.New(wview.Config{
		Root:         "views/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{"partials/ad"},
		Funcs:        fm,
		DisableCache: true,
	}))

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index", iris.Map{
			"title": "Frontend title!",
		})
	})

	// =========== Backend =========== //

	// Assign a new template middleware.
	mw := irisview.NewMiddleware(wview.Config{
		Root:         "views/backend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        fm,
		DisableCache: true,
	})

	backendGroup := app.Party("/admin", mw)

	backendGroup.Get("/", func(ctx iris.Context) {
		// Use the ctx.View as you used to. Zero changes to your codebase,
		// even if you use multiple templates.
		ctx.View("index", iris.Map{
			"title": "Backend title!",
		})
	})

	app.Listen(":9090")
}
