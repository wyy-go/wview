package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/wyy-go/wview"
	"net/http"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
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

	r.Get("/page", func(w http.ResponseWriter, r *http.Request) {
		err := wview.Render(w, http.StatusOK, "page.html", wview.M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", r)

}
