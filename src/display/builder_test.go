package display

import (
	"testing"
)

func TestFactory(t *testing.T) {
	surface := &FakeSurface{}

	t.Run("CreateBuilder", func(t *testing.T) {
		builder := CreateBuilder(surface, func(s Surface) {
			Box(s, &Opts{Width: 100, Height: 100, StyleName: "abcd"})
		})

		builder.Build()

		box := builder.GetRoot()
		if box == nil {
			t.Error("Expected builder.GetRoot() to return a valid box")
		}
	})
}
