package display

import (
	"testing"
)

func TestWindow(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		surface := &FakeSurface{}
		renderer := CreateBuilder(surface, func(s Surface) {
			Window(s)
		})

		renderer.Build()

		win := renderer.GetRoot()
		if win == nil {
			t.Error("Expected renderer.GetRoot() to return a valid window")
		}
	})
}
