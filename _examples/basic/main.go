package main

import (
	"fmt"
	"github.com/wyy-go/wview"
	"net/http"
)

func main() {

	//render index use `index` without `.html` extension, that will render with master layout.
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

	//render page use `page.html` with '.html' will only file template without master layout.
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		err := wview.Render(w, http.StatusOK, "page.html", wview.M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", nil)

}
