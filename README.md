![GitHub Repo stars](https://img.shields.io/github/stars/wyy-go/yview?style=social)
![GitHub](https://img.shields.io/github/license/wyy-go/yview)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wyy-go/yview)
![GitHub all releases](https://img.shields.io/github/downloads/wyy-go/yview/total)
![GitHub CI Status](https://img.shields.io/github/workflow/status/wyy-go/yview/ci?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/wyy-go/yview)](https://goreportcard.com/report/github.com/wyy-go/yview)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/wyy-go/yview?tab=doc)
[![codecov](https://codecov.io/gh/wyy-go/yview/branch/main/graph/badge.svg)](https://codecov.io/gh/wyy-go/yview)

# yview
yview is a lightweight, minimalist and idiomatic template library based on golang [html/template](https://golang.org/pkg/html/template/) for building Go web application.

## Contents

- [Install](#install)
- [Features](#features)
- [Supports](#supports)
  - [Gin Framework](https://github.com/wyy-go/yview/tree/main/plugin/ginview)
  - [Iris Framework](https://github.com/wyy-go/yview/tree/main/plugin/irisview)
  - [Echo Framework](https://github.com/wyy-go/yview/tree/main/plugin/echoview)
  - [Go.Rice](https://github.com/wyy-go/yview/tree/main/plugin/gorice)
- [Usage](#usage)
  - [Overview](#overview)
  - [Config](#config)
  - [Include syntax](#include-syntax)
  - [Render name](#render-name)
  - [Custom template functions](#custom-template-functions)
- [Examples](#examples)
  - [Basic example](#basic-example)
  - [Gin example](#gin-example)
  - [Iris example](#iris-example)
  - [Iris multiple example](#iris-multiple-example)
  - [Echo example](#echo-example)
  - [Go-chi example](#go-chi-example)
  - [Advance example](#advance-example)
  - [Multiple example](#multiple-example)
  - [go.rice example](#gorice-example)
  - [more examples](#more-examples)


## Install

```bash
go get github.com/wyy-go/yview
```


## Features

* **Lightweight** - use golang html/template syntax.
* **Easy** - easy use for your web application.
* **Fast** - Support configure cache template.
* **Include syntax** - Support include file.
* **Master layout** - Support configure master layout file.
* **Extension** - Support configure template file extension.
* **Easy** - Support configure templates directory.
* **Auto reload** - Support dynamic reload template(disable cache mode).
* **Multiple Engine** - Support multiple templates for frontend and backend.
* **No external dependencies** - plain ol' Go html/template.
* **Gorice** - Support gorice for package resources.
* **Gin/Iris/Echo/Chi** - Support gin framework, Iris framework, echo framework, go-chi framework.


## Supports

- **[ginview](https://github.com/wyy-go/yview/tree/main/plugin/ginview)** yview for gin framework
- **[irisview](https://github.com/wyy-go/yview/tree/main/plugin/irisview)** yview for Iris framework
- **[echoview](https://github.com/wyy-go/yview/tree/main/plugin/echoview)** yview for echo framework
- **[gorice](https://github.com/wyy-go/yview/tree/main/plugin/gorice)** yview for go.rice
- go embed


## Usage

### Overview

Project structure:

```go
|-- app/views/
    |--- index.html          
    |--- page.html
    |-- layouts/
        |--- footer.html
        |--- master.html
```

Use default instance:

```go
//write http.ResponseWriter
//"index" -> index.html
yview.Render(writer, http.StatusOK, "index", yview.M{})
```

Use new instance with config:

```go
gv := yview.New(yview.Config{
    Root:      "views",
    Extension: ".tpl",
    Master:    "layouts/master",
    Partials:  []string{"partials/ad"},
    Funcs: template.FuncMap{
        "sub": func(a, b int) int {
            return a - b
        },
        "copy": func() string {
            return time.Now().Format("2006")
        },
    },
    DisableCache: true,
Delims:    Delims{Left: "{{", Right: "}}"},
})

//Set new instance
yview.Use(gv)

//write http.ResponseWriter
yview.Render(writer, http.StatusOK, "index", yview.M{})
```


Use multiple instance with config:

```go
//============== Frontend ============== //
yvFrontend := yview.New(yview.Config{
    Root:      "views/frontend",
    Extension: ".tpl",
    Master:    "layouts/master",
    Partials:  []string{"partials/ad"},
    Funcs: template.FuncMap{
        "sub": func(a, b int) int {
            return a - b
        },
        "copy": func() string {
            return time.Now().Format("2006")
        },
    },
    DisableCache: true,
Delims:       Delims{Left: "{{", Right: "}}"},
})

//write http.ResponseWriter
yvFrontend.Render(writer, http.StatusOK, "index", yview.M{})

//============== Backend ============== //
yvBackend := yview.New(yview.Config{
    Root:      "views/backend",
    Extension: ".tpl",
    Master:    "layouts/master",
    Partials:  []string{"partials/ad"},
    Funcs: template.FuncMap{
        "sub": func(a, b int) int {
            return a - b
        },
        "copy": func() string {
            return time.Now().Format("2006")
        },
    },
    DisableCache: true,
Delims:       Delims{Left: "{{", Right: "}}"},
})

//write http.ResponseWriter
yvBackend.Render(writer, http.StatusOK, "index", yview.M{})

```

### Config

```go
yview.Config{
    Root:      "views", //template root path
    Extension: ".tpl", //file extension
    Master:    "layouts/master", //master layout file
    Partials:  []string{"partials/head"}, //partial files
    Funcs: template.FuncMap{
        "sub": func(a, b int) int {
            return a - b
        },
        // more funcs
    },
    DisableCache: false, //if disable cache, auto reload template file for debug.
    Delims:       Delims{Left: "{{", Right: "}}"},
}
```

### Include syntax

```go
//template file
{{include "layouts/footer"}}
```

### Render name: 

Render name use `index` without `.html` extension, that will render with master layout.

- **"index"** - Render with master layout.
- **"index.html"** - Not render with master layout.

```
Notice: `.html` is default template extension, you can change with config
```


Render with master

```go
//use name without extension `.html`
yview.Render(w, http.StatusOK, "index", yview.M{})
```

The `w` is instance of  `http.ResponseWriter`

Render only file(not use master layout)

```go
//use full name with extension `.html`
yview.Render(w, http.StatusOK, "page.html", yview.M{})
```

### Custom template functions

We have two type of functions `global functions`, and `temporary functions`.

`Global functions` are set within the `config`.

```go
yview.Config{
	Funcs: template.FuncMap{
		"reverse": e.Reverse,
	},
}
```

```go
//template file
{{ reverse "route-name" }}
```

`Temporary functions` are set inside the handler.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	err := yview.Render(w, http.StatusOK, "index", yview.M{
		"reverse": e.Reverse,
	})
	if err != nil {
		fmt.Fprintf(w, "Render index error: %v!", err)
	}
})
```

```go
//template file
{{ call $.reverse "route-name" }}
```



## Examples

See [_examples/](https://github.com/wyy-go/yview/tree/main/_examples/) for a variety of examples.


### Basic example

```go
package main

import (
	"fmt"
	"github.com/wyy-go/yview"
	"net/http"
)

func main() {

	//render index use `index` without `.html` extension, that will render with master layout.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := yview.Render(w, http.StatusOK, "index", yview.M{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}

	})

	//render page use `page.tpl` with '.html' will only file template without master layout.
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		err := yview.Render(w, http.StatusOK, "page.html", yview.M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", nil)

}

```

Project structure:

```go
|-- app/views/
    |--- index.html          
    |--- page.html
    |-- layouts/
        |--- footer.html
        |--- master.html
    

See in "examples/basic" folder
```

[Basic example](https://github.com/wyy-go/yview/tree/main/_examples/basic)


### Gin example

```bash
go get github.com/wyy-go/yview/plugin/ginview
```

```go
package main

import (
	"github.com/wyy-go/yview/plugin/ginview"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = yview.Default()

	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	router.Run(":9090")
}

```

Project structure:

```go
|-- app/views/
    |--- index.html          
    |--- page.html
    |-- layouts/
        |--- footer.html
        |--- master.html
    

See in "examples/basic" folder
```

[Gin example](https://github.com/wyy-go/yview/tree/main/_examples/gin)

### Iris example

```bash
$ go get github.com/wyy-go/yview/plugin/irisview
```

```go
package main

import (
	"github.com/wyy-go/yview/main/irisview"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// Register the yview template engine.
	app.RegisterView(irisview.Default())

	app.Get("/", func(ctx iris.Context) {
		// Render with master.
		ctx.View("index", iris.Map{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	app.Get("/page", func(ctx iris.Context) {
		// Render only file, must full name with extension.
		ctx.View("page.html", iris.Map{"title": "Page file title!!"})
	})

	app.Listen(":9090")
}
```

Project structure:

```go
|-- app/views/
    |--- index.html          
    |--- page.html
    |-- layouts/
        |--- footer.html
        |--- master.html
    

See in "examples/iris" folder
```

[Iris example](https://github.com/wyy-go/yview/tree/main/_examples/iris)


### Iris multiple example

```go
package main

import (
	"html/template"
	"time"

	"github.com/wyy-go/yview"
	"github.com/wyy-go/yview/plugin/irisview"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// Register a new template engine.
	app.RegisterView(irisview.New(yview.Config{
		Root:      "views/frontend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	}))

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index", iris.Map{
			"title": "Frontend title!",
		})
	})

	//=========== Backend ===========//

	// Assign a new template middleware.
	mw := irisview.NewMiddleware(yview.Config{
		Root:      "views/backend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	backendGroup := app.Party("/admin", mw)

	backendGroup.Get("/", func(ctx iris.Context) {
		// Use the ctx.View as you used to. Zero changes to your codebase,
		// even if you use multiple templates.
		ctx.View("index", iris.Map{
			"title": "Backend title!",
		})
	})

	app.Listen(":9090")
}
```

Project structure:

```go
|-- app/views/
    |-- fontend/
        |--- index.html
        |-- layouts/
            |--- footer.html
            |--- head.html
            |--- master.html
        |-- partials/
     	   |--- ad.html
    |-- backend/
        |--- index.html
        |-- layouts/
            |--- footer.html
            |--- head.html
            |--- master.html
        
See in "examples/iris-multiple" folder
```

[Iris multiple example](https://github.com/wyy-go/yview/tree/main/_examples/iris-multiple)

### Echo example

Echo <=v3 version:

```bash
go get github.com/wyy-go/yview/plugin/echoview
```

Echo v4 version:

```bash
go get github.com/wyy-go/yview/plugin/echoview-v4
```


```go
package main

import (
	"github.com/wyy-go/yview/plugin/echoview"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Set Renderer
	e.Renderer = echoview.Default()

	// Routes
	e.GET("/", func(c echo.Context) error {
		//render with master
		return c.Render(http.StatusOK, "index", echo.Map{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	e.GET("/page", func(c echo.Context) error {
		//render only file, must full name with extension
		return c.Render(http.StatusOK, "page.html", echo.Map{"title": "Page file title!!"})
	})

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}

```

Project structure:

```go
|-- app/views/
    |--- index.html          
    |--- page.html
    |-- layouts/
        |--- footer.html
        |--- master.html
    

See in "examples/basic" folder
```

[Echo example](https://github.com/wyy-go/yview/tree/main/_examples/echo)
[Echo v4 example](https://github.com/wyy-go/yview/tree/main/_examples/echo-v4)


### Go-chi example

```go
package main

import (
	"fmt"
	"github.com/wyy-go/yview"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {

	r := chi.NewRouter()

	//render index use `index` without `.html` extension, that will render with master layout.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := yview.Render(w, http.StatusOK, "index", yview.M{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}
	})

	//render page use `page.tpl` with '.html' will only file template without master layout.
	r.Get("/page", func(w http.ResponseWriter, r *http.Request) {
		err := yview.Render(w, http.StatusOK, "page.html", yview.M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", r)

}

```

Project structure:

```go
|-- app/views/
    |--- index.html          
    |--- page.html
    |-- layouts/
        |--- footer.html
        |--- master.html
    

See in "examples/basic" folder
```

[Chi example](https://github.com/wyy-go/yview/tree/main/_examples/go-chi)



### Advance example

```go
package main

import (
	"fmt"
	"github.com/wyy-go/yview"
	"html/template"
	"net/http"
	"time"
)

func main() {

	gv := yview.New(yview.Config{
		Root:      "views",
		Extension: ".tpl",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	//Set new instance
	yview.Use(gv)

	//render index use `index` without `.html` extension, that will render with master layout.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := yview.Render(w, http.StatusOK, "index", yview.M{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}

	})

	//render page use `page.tpl` with '.html' will only file template without master layout.
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		err := yview.Render(w, http.StatusOK, "page.tpl", yview.M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", nil)
}

```

Project structure:

```go
|-- app/views/
    |--- index.tpl          
    |--- page.tpl
    |-- layouts/
        |--- footer.tpl
        |--- head.tpl
        |--- master.tpl
    |-- partials/
        |--- ad.tpl
    

See in "examples/advance" folder
```

[Advance example](https://github.com/wyy-go/yview/tree/main/_examples/advance)

### Multiple example

```go
package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/wyy-go/yview"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views/fontend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	router.GET("/", func(ctx *gin.Context) {
		// `HTML()` is a helper func to deal with multiple TemplateEngine's.
		// It detects the suitable TemplateEngine for each path automatically.
		gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Fontend title!",
		})
	})

	//=========== Backend ===========//

	//new middleware
	mw := gintemplate.NewMiddleware(gintemplate.TemplateConfig{
		Root:      "views/backend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	// You should use helper func `Middleware()` to set the supplied
	// TemplateEngine and make `HTML()` work validly.
	backendGroup := router.Group("/admin", mw)

	backendGroup.GET("/", func(ctx *gin.Context) {
		// With the middleware, `HTML()` can detect the valid TemplateEngine.
		gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Backend title!",
		})
	})

	router.Run(":9090")
}


```

Project structure:

```go
|-- app/views/
    |-- fontend/
        |--- index.html
        |-- layouts/
            |--- footer.html
            |--- head.html
            |--- master.html
        |-- partials/
     	   |--- ad.html
    |-- backend/
        |--- index.html
        |-- layouts/
            |--- footer.html
            |--- head.html
            |--- master.html
        
See in "examples/multiple" folder
```

[Multiple example](https://github.com/wyy-go/yview/tree/main/_examples/multiple)


### go.rice example

```bash
go get github.com/wyy-go/yview/plugin/gorice
```

```go
package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/wyy-go/yview"
	"github.com/wyy-go/yview/plugin/gorice"
	"net/http"
)

func main() {

	//static
	staticBox := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(staticBox.HTTPBox()))
	http.Handle("/static/", staticFileServer)

	//new view engine
	gv := gorice.New(rice.MustFindBox("views"))
	//set engine for default instance
	yview.Use(gv)

	//render index use `index` without `.html` extension, that will render with master layout.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := yview.Render(w, http.StatusOK, "index", yview.M{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}

	})

	//render page use `page.tpl` with '.html' will only file template without master layout.
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		err := yview.Render(w, http.StatusOK, "page.html", yview.M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", nil)
}

```

Project structure:

```go
|-- app/views/
    |--- index.html          
    |--- page.html
    |-- layouts/
        |--- footer.html
        |--- master.html
|-- app/static/  
    |-- css/
        |--- bootstrap.css   	
    |-- img/
        |--- gopher.png

See in "examples/gorice" folder
```

[gorice example](https://github.com/wyy-go/yview/tree/main/_examples/gorice)

### More examples

See [_examples/](https://github.com/wyy-go/yview/tree/main/_examples/) for a variety of examples.


## Todo

 [ ] Add Partials support directory or glob

 [ ] Add functions support.

