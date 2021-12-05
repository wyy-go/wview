package irisview

import (
	"io"

	"github.com/wyy-go/wview"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/view"
)

const templateEngineKey = "view-engine-iris"

// ViewEngine view engine for Iris.
type ViewEngine struct {
	*wview.ViewEngine
	ext string
}

var _ view.Engine = (*ViewEngine)(nil)

// New view engine for Iris.
func New(config wview.Config) *ViewEngine {
	e := Wrap(wview.New(config))
	e.ext = config.Extension
	return e
}

// Wrap wraps a view engine for goview.ViewEngine.
func Wrap(engine *wview.ViewEngine) *ViewEngine {
	return &ViewEngine{
		ViewEngine: engine,
		ext:        wview.DefaultConfig.Extension,
	}
}

// Default returns a new default view engine.
func Default() *ViewEngine {
	return New(wview.DefaultConfig)
}

// Load does nothing here, templates are loaded through goview.
func (e *ViewEngine) Load() error {
	return nil
}

// Ext returns the file extension, it's empty on this case because
// the goviw engine supports filenames without extension.
func (e *ViewEngine) Ext() string {
	return ""
}

// ExecuteWriter executes a template by its name.
// It supports multiple templates, see `NewMiddleware` and `Middleware` too.
func (e *ViewEngine) ExecuteWriter(w io.Writer, filename string, _ string, bindingData interface{}) error {
	if ctx, ok := w.(iris.Context); ok {
		if v := ctx.Values().Get(templateEngineKey); v != nil {
			if e, ok := v.(*ViewEngine); ok {
				return e.ViewEngine.RenderWriter(w, filename, bindingData)
			}
		}
	}

	return e.ViewEngine.RenderWriter(w, filename, bindingData)
}

// NewMiddleware Iris middleware for multiple templates.
func NewMiddleware(config wview.Config) iris.Handler {
	return Middleware(New(config))
}

// Middleware Iris middleware wrapper.
func Middleware(e *ViewEngine) iris.Handler {
	return func(ctx iris.Context) {
		ctx.Values().Set(templateEngineKey, e)
		ctx.Next()
	}
}
