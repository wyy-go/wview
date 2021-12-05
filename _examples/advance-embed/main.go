package main

import (
	"embed"
	"fmt"
	"github.com/wyy-go/wview"
	"html/template"
	"net/http"
	"time"
)

//go:embed views
var views embed.FS

func main() {
	fm := make(template.FuncMap)
	fm["safeHTML"] = func(v string) template.HTML {
		return template.HTML(v)
	}

	fm["copy"] = func() string {
		return time.Now().Format("2006")
	}

	yv := wview.New(wview.Config{
		Root:         "views",
		Extension:    ".tpl",
		Master:       "layouts/master",
		Partials:     []string{"partials/ad"},
		Funcs:        fm,
		DisableCache: true,
		EnableEmbed:  true,
		Views:        views,
	})

	// Set new instance
	wview.Use(yv)

	rawContent := `This is <b>HTML</b> content! Posted on <time datetime="2021-12-05 15:02:03">May 16</time> by wyy-go.`

	// render index use `index` without `.tpl` extension, that will render with master layout.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := wview.Render(w, http.StatusOK, "index", wview.M{
			"Title":       "Index title!",
			"HtmlContent": template.HTML(rawContent),
			"RawContent":  rawContent,
			"tempConvertHTML": func(v string) template.HTML {
				return template.HTML(v)
			},
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}

	})

	// render page use `page.tpl` with '.tpl' will only file template without master layout.
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		err := wview.Render(w, http.StatusOK, "page.tpl", wview.M{
			"Title": "Page file title!!",
		})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", nil)
}
