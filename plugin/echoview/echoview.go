package echoview

import (
	"io"

	"github.com/labstack/echo"
	"github.com/wyy-go/wview"
)

const templateEngineKey = "view-engine-echo"

// ViewEngine view engine for echo
type ViewEngine struct {
	*wview.ViewEngine
}

// New new view engine
func New(config wview.Config) *ViewEngine {
	return Wrap(wview.New(config))
}

// Wrap wrap view engine for goview.ViewEngine
func Wrap(engine *wview.ViewEngine) *ViewEngine {
	return &ViewEngine{
		ViewEngine: engine,
	}
}

// Default new default config view engine
func Default() *ViewEngine {
	return New(wview.DefaultConfig)
}

// Render render template for echo interface
func (e *ViewEngine) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return e.RenderWriter(w, name, data)
}

// Render html render for template
// You should use helper func `Middleware()` to set the supplied
// TemplateEngine and make `Render()` work validly.
func Render(ctx echo.Context, code int, name string, data interface{}) error {
	if val := ctx.Get(templateEngineKey); val != nil {
		if e, ok := val.(*ViewEngine); ok {
			return e.Render(ctx.Response().Writer, name, data, ctx)
		}
	}
	return ctx.Render(code, name, data)
}

// NewMiddleware echo middleware for func `echoview.Render()`
func NewMiddleware(config wview.Config) echo.MiddlewareFunc {
	return Middleware(New(config))
}

// Middleware echo middleware wrapper
func Middleware(e *ViewEngine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(templateEngineKey, e)
			return next(c)
		}
	}
}
