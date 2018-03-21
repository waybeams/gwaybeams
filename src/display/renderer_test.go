package display

import (
	"testing"
)

func TestFactory(t *testing.T) {
	surface := &FakeSurface{}

	t.Run("CreateRender", func(t *testing.T) {
		renderer := CreateRenderer(surface, func(s Surface) {
			Box(s, &Opts{Width: 100, Height: 100, StyleName: "abcd"})
		})

		renderer.Render()

		box := renderer.GetRoot()
		if box == nil {
			t.Error("Expected renderer.GetRoot() to return a valid box")
		}
	})
}
