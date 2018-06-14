package canvas_test

import (
	"testing"

	"github.com/gopherjs/gopherjs/js"
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/surface/canvas"
)

func TestWebglSurface(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		elem := &js.Object{}
		w := canvas.NewSurface(canvas.PageContext(elem))
		assert.NotNil(w)
	})
}
