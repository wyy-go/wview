package main

import (
	"embed"
	"github.com/wyy-go/wview"
	"github.com/wyy-go/wview/plugin/echoview"
	"html/template"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//go:embed views
var views embed.FS

//go:embed static
var static embed.FS

//go:embed static/favicon.ico
var favicon []byte

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	fm := make(template.FuncMap)
	fm["copy"] = func() string {
		return time.Now().Format("2006")
	}

	//new template engine
	e.Renderer = echoview.New(wview.Config{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        fm,
		DisableCache: true,
		EnableEmbed:  true,
		Views:        views,
	})

	// static file
	e.Any("/static/*filepath", func(context echo.Context) error {
		staticServer := http.FileServer(http.FS(static))
		staticServer.ServeHTTP(context.Response(), context.Request())
		return nil
	})

	// favicon
	e.GET("/favicon.ico", func(context echo.Context) error {
		if _, err := context.Response().Write(favicon); err != nil {
			return err
		}
		return nil
	})

	e.GET("/", func(context echo.Context) error {
		return echoview.Render(context, http.StatusOK, "index", echo.Map{
			"title": "Frontend title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	e.GET("/page", func(context echo.Context) error {
		return echoview.Render(context, http.StatusOK, "page.html", echo.Map{
			"title": "Page file title!!",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}
