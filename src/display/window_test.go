package display

import (
	"assert"
	"testing"
)

func TestWindow(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		surface := &FakeSurface{}
		renderer := CreateRenderer(surface, func(s Surface) {
			Window(s)
		})

		win := renderer.GetRoot()
		assert.NotNil(win)
	})
}
