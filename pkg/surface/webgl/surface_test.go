package webgl_test

import (
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/surface/webgl"
)

func TestWebglSurface(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		w := webgl.NewSurface(
			webgl.Font("Roboto", "../../testdata/Roboto-Regular.ttf"),
		)
		assert.NotNil(w)
	})
}
