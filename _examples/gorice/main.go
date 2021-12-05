package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/wyy-go/wview"
	"github.com/wyy-go/wview/plugin/gorice"
	"net/http"
)

func main() {

	// static
	staticBox := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(staticBox.HTTPBox()))
	http.Handle("/static/", staticFileServer)

	// new view engine
	gv := gorice.New(rice.MustFindBox("views"))
	// set engine for default instance
	wview.Use(gv)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := wview.Render(w, http.StatusOK, "index", wview.M{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}

	})

	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		err := wview.Render(w, http.StatusOK, "page.html", wview.M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", nil)
}
