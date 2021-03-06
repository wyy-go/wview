package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wyy-go/wview/plugin/ginview"
	"net/http"
)

func main() {
	router := gin.Default()

	// new template engine
	router.HTMLRender = ginview.Default()

	router.GET("/", func(ctx *gin.Context) {
		// render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		// render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	router.Run(":9090")
}
