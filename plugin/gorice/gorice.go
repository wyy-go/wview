package gorice

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/wyy-go/wview"
)

// New new gorice template engine, default views root.
func New(viewsRootBox *rice.Box) *wview.ViewEngine {
	return NewWithConfig(viewsRootBox, wview.DefaultConfig)
}

// NewWithConfig gorice template engine
// Important!!! The viewsRootBox's name and config.Root must be consistent.
func NewWithConfig(viewsRootBox *rice.Box, config wview.Config) *wview.ViewEngine {
	config.Root = viewsRootBox.Name()
	engine := wview.New(config)
	engine.SetFileHandler(FileHandler(viewsRootBox))
	return engine
}

// FileHandler Support go.rice file handler
func FileHandler(viewsRootBox *rice.Box) wview.FileHandler {
	return func(config wview.Config, tplFile string) (content string, err error) {
		// get file contents as string
		return viewsRootBox.String(tplFile + config.Extension)
	}
}
