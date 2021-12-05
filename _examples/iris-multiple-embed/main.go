package main

import (
	"embed"
	"github.com/kataras/iris/v12/context"
	"html/template"
	"net/http"
	"time"

	"github.com/wyy-go/wview"
	"github.com/wyy-go/wview/plugin/irisview"

	"github.com/kataras/iris/v12"
)

//go:embed views
var views embed.FS

//go:embed static
var static embed.FS

//go:embed static/favicon.ico
var favicon []byte

func main() {
	app := iris.New()

	fm := make(template.FuncMap)
	fm["copy"] = func() string {
		return time.Now().Format("2006")
	}

	// Register a new template engine.
	app.RegisterView(irisview.New(wview.Config{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        fm,
		DisableCache: true,
		EnableEmbed:  true,
		Views:        views,
	}))

	// static file
	app.Any("/static/*filepath", func(context context.Context) {
		staticServer := http.FileServer(http.FS(static))
		staticServer.ServeHTTP(context.ResponseWriter(), context.Request())
	})

	// favicon
	app.Get("/favicon.ico", func(context context.Context) {
		context.Write(favicon)
	})

	app.Get("/", func(context context.Context) {
		context.View("index", iris.Map{
			"title": "Frontend title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	app.Get("/page", func(ctx iris.Context) {
		ctx.View("page.html", iris.Map{
			"title": "Page file title!!",
		})
	})

	app.Listen(":9090")
}
