# GinView

[![GoDoc Widget]][GoDoc] 

goview support for gin template

## Install
```bash

go get -u github.com/wyy-go/yview

go get -u github.com/wyy-go/yview/plugin/ginview

```

### Example

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
	router.HTMLRender = ginview.Default()

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
           
## More examples

See [_examples/](https://github.com/wyy-go/yview/tree/main/_examples/) for a variety of examples.

[GoDoc]: https://godoc.org/github.com/wyy-go/yview/plugin/ginview
[GoDoc Widget]: https://godoc.org/github.com/wyy-go/yview/plugin/ginview?status.svg
