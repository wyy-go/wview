package ginview

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/wyy-go/wview"
)

const templateEngineKey = "view-engine-gin"

// ViewEngine view engine for gin
type ViewEngine struct {
	*wview.ViewEngine
}

// ViewRender view render implement gin interface
type ViewRender struct {
	Engine *ViewEngine
	Name   string
	Data   interface{}
}

// New new view engine for gin
func New(config wview.Config) *ViewEngine {
	return Wrap(wview.New(config))
}

// Wrap wrap view engine for goview.ViewEngine
func Wrap(engine *wview.ViewEngine) *ViewEngine {
	return &ViewEngine{
		ViewEngine: engine,
	}
}

// Default new default engine
func Default() *ViewEngine {
	return New(wview.DefaultConfig)
}

// Instance implement gin interface
func (e *ViewEngine) Instance(name string, data interface{}) render.Render {
	return ViewRender{
		Engine: e,
		Name:   name,
		Data:   data,
	}
}

// HTML render html
func (e *ViewEngine) HTML(ctx *gin.Context, code int, name string, data interface{}) {
	instance := e.Instance(name, data)
	ctx.Render(code, instance)
}

// Render (YAML) marshals the given interface object and writes data with custom ContentType.
func (v ViewRender) Render(w http.ResponseWriter) error {
	return v.Engine.RenderWriter(w, v.Name, v.Data)
}

// WriteContentType write html content type
func (v ViewRender) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = wview.HTMLContentType
	}
}

// NewMiddleware gin middleware for func `gintemplate.HTML()`
func NewMiddleware(config wview.Config) gin.HandlerFunc {
	return Middleware(New(config))
}

// Middleware gin middleware wrapper
func Middleware(e *ViewEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(templateEngineKey, e)
	}
}

// HTML html render for template
// You should use helper func `Middleware()` to set the supplied
// TemplateEngine and make `HTML()` work validly.
func HTML(ctx *gin.Context, code int, name string, data interface{}) {
	if val, ok := ctx.Get(templateEngineKey); ok {
		if e, ok := val.(*ViewEngine); ok {
			e.HTML(ctx, code, name, data)
			return
		}
	}
	ctx.HTML(code, name, data)
}
