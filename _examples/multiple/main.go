package main

import (
	"fmt"
	"github.com/wyy-go/wview"
	"html/template"
	"net/http"
	"time"
)

func main() {
	fm := make(template.FuncMap)
	fm["copy"] = func() string {
		return time.Now().Format("2006")
	}

	gvFront := wview.New(wview.Config{
		Root:         "views/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{"partials/ad"},
		Funcs:        fm,
		DisableCache: true,
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := gvFront.Render(w, http.StatusOK, "index", wview.M{
			"title": "Frontend title!",
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}
	})

	// =========== Backend =========== //

	gvBackend := wview.New(wview.Config{
		Root:         "views/backend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        fm,
		DisableCache: true,
	})

	http.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		err := gvBackend.Render(w, http.StatusOK, "index", wview.M{
			"title": "Backend title!",
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}
	})

	fmt.Printf("Server start on :9090")
	http.ListenAndServe(":9090", nil)
}
