package wview

import (
	"net/http"
)

var view *ViewEngine

// Use setting default instance engine
func Use(engine *ViewEngine) {
	view = engine
}

// Render render view template with default instance
func Render(w http.ResponseWriter, status int, name string, data interface{}) error {
	if view == nil {
		view = Default()
		// return fmt.Errorf("instance not yet initialized, please call Init() first before Render()")
	}
	return view.Render(w, status, name, data)
}
