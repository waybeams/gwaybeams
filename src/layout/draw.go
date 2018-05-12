package layout

import (
	"spec"
	"views"
)

var DefaultView spec.RenderHandler = views.RectangleView

// Draw the provided spec tree onto the provided Surface
func Draw(r spec.Reader, s spec.Surface) {
	view := r.View()
	if view == nil {
		view = DefaultView
	}
	view(s, r)

	for _, child := range r.Children() {
		Draw(child, s)
	}
}
