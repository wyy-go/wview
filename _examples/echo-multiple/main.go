package main

import (
	"github.com/wyy-go/wview"
	"github.com/wyy-go/wview/plugin/echoview"
	"html/template"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

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

	// new template engine
	e.Renderer = echoview.New(wview.Config{
		Root:         "views/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{"partials/ad"},
		Funcs:        fm,
		DisableCache: true,
	})

	e.GET("/", func(ctx echo.Context) error {
		// `HTML()` is a helper func to deal with multiple TemplateEngine's.
		// It detects the suitable TemplateEngine for each path automatically.
		return echoview.Render(ctx, http.StatusOK, "index", echo.Map{
			"title": "Frontend title!",
		})
	})

	// =========== Backend =========== //

	// new middleware
	mw := echoview.NewMiddleware(wview.Config{
		Root:         "views/backend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        fm,
		DisableCache: true,
	})

	// You should use helper func `Middleware()` to set the supplied
	// TemplateEngine and make `HTML()` work validly.
	backendGroup := e.Group("/admin", mw)

	backendGroup.GET("/", func(ctx echo.Context) error {
		// With the middleware, `HTML()` can detect the valid TemplateEngine.
		return echoview.Render(ctx, http.StatusOK, "index", echo.Map{
			"title": "Backend title!",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}
