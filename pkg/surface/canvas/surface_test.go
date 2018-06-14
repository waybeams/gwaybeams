package canvas_test

import (
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/surface/canvas"
)

func TestWebglSurface(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		w := canvas.NewSurface(
			canvas.Font("Roboto", "../../testdata/Roboto-Regular.ttf"),
		)
		assert.NotNil(w)
	})
}
