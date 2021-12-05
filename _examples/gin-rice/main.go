package main

import (
	"html/template"
	"net/http"
	"time"

	rice "github.com/GeertJohan/go.rice"

	"github.com/gin-gonic/gin"
	"github.com/wyy-go/wview"
	"github.com/wyy-go/wview/plugin/ginview"
	"github.com/wyy-go/wview/plugin/gorice"
)

func main() {

	router := gin.Default()

	fm := make(template.FuncMap)
	fm["copy"] = func() string {
		return time.Now().Format("2006")
	}

	// new template engine
	basic := gorice.NewWithConfig(rice.MustFindBox("views/frontend"), wview.Config{
		Root:         "views/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{"partials/ad"},
		Funcs:        fm,
		DisableCache: true,
	})
	router.HTMLRender = ginview.Wrap(basic)

	router.GET("/", func(ctx *gin.Context) {
		// `HTML()` is a helper func to deal with multiple TemplateEngine's.
		// It detects the suitable TemplateEngine for each path automatically.
		ginview.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Frontend title!",
		})
	})

	// =========== Backend =========== //

	adminView := gorice.NewWithConfig(rice.MustFindBox("views/backend"), wview.Config{
		Root:         "views/backend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        fm,
		DisableCache: true,
	})

	//new middleware
	mw := ginview.Middleware(ginview.Wrap(adminView))

	// You should use helper func `Middleware()` to set the supplied
	// TemplateEngine and make `HTML()` work validly.
	backendGroup := router.Group("/admin", mw)

	backendGroup.GET("/", func(ctx *gin.Context) {
		// With the middleware, `HTML()` can detect the valid TemplateEngine.
		ginview.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Backend title!",
		})
	})

	router.Run(":9090")
}
