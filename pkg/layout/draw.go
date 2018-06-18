package layout

import (
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/views"
)

// Draw the provided spec tree onto the provided Surface
func Draw(r spec.Reader, s spec.Surface) {
	s = spec.NewOffsetSurface(r, s)
	view := r.View()
	if view == nil {
		view = views.RectangleView
	}
	view(s, r)

	for _, child := range r.Children() {
		Draw(child, s)
	}
}
