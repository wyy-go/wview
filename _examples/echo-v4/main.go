package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wyy-go/wview/plugin/echoview-v4"
	"net/http"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set Renderer
	e.Renderer = echoview.Default()

	// Routes
	e.GET("/", func(c echo.Context) error {
		// render with master
		return c.Render(http.StatusOK, "index", echo.Map{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	e.GET("/page", func(c echo.Context) error {
		// render only file, must full name with extension
		return c.Render(http.StatusOK, "page.html", echo.Map{"title": "Page file title!!"})
	})

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}
