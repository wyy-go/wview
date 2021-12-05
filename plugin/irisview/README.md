# IrisView

[![GoDoc Widget]][GoDoc] 

yview support for Iris template.

## Install

```sh
$ go get -u github.com/wyy-go/yview
$ go get -u github.com/wyy-go/yview/plugin/irisview
```

### Example

```go
package main

import (
	"github.com/wyy-go/yview/plugin/irisview"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// Register the goview template engine.
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
    

See in "examples/basic" folder
```

[Iris example](https://github.com/wyy-go/yview/tree/main/_examples/iris)
           
## More examples

See [_examples/](https://github.com/wyy-go/yview/tree/main/_examples/) for a variety of examples.

[GoDoc]: https://godoc.org/github.com/wyy-go/yview/plugin/irisview
[GoDoc Widget]: https://godoc.org/github.com/wyy-go/yview/plugin/irisview?status.svg
