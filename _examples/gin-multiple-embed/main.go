package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/wyy-go/wview"
	"github.com/wyy-go/wview/plugin/ginview"
	"html/template"
	"net/http"
	"time"
)

//go:embed views
var views embed.FS

//go:embed static
var static embed.FS

//go:embed static/favicon.ico
var favicon []byte

func main() {
	router := gin.Default()

	fm := make(template.FuncMap)
	fm["copy"] = func() string {
		return time.Now().Format("2006")
	}

	// new template engine
	router.HTMLRender = ginview.New(wview.Config{
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
	router.Any("/static/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(static))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

	// favicon
	router.GET("/favicon.ico", func(context *gin.Context) {
		context.Writer.Write(favicon)
	})

	router.GET("/", func(ctx *gin.Context) {
		// `HTML()` is a helper func to deal with multiple TemplateEngine's.
		// It detects the suitable TemplateEngine for each path automatically.
		ginview.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		// `HTML()` is a helper func to deal with multiple TemplateEngine's.
		// It detects the suitable TemplateEngine for each path automatically.
		ginview.HTML(ctx, http.StatusOK, "page.html", gin.H{
			"title": "Page file title!!",
		})
	})

	router.Run(":9090")
}
