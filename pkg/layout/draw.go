package layout

import (
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/surface"
	"github.com/waybeams/waybeams/pkg/views"
)

var DefaultView spec.RenderHandler = views.RectangleView

// Draw the provided spec tree onto the provided Surface
func Draw(r spec.Reader, s spec.Surface) {
	s = surface.NewOffsetSurface(r, s)
	view := r.View()
	if view == nil {
		view = DefaultView
	}
	view(s, r)

	for _, child := range r.Children() {
		Draw(child, s)
	}
}
