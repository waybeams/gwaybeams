package display

import (
	"assert"
	"testing"
)

func TestBuilder(t *testing.T) {
	surface := &FakeSurface{}

	t.Run("CreateBuilder", func(t *testing.T) {
		builder := NewBuilder2()
		builder.WithHint(Floating, true)
		builder.WithHint(Resizable, true)
		builder.WithHint(Visible, true)
		builder.WithSurfaceType(CairoSurfaceType)
		builder.WithFrameRate(12)
		builder.WithSize(640, 480)
		builder.WithTitle("Hello World")

		err := builder.Build(func(s Surface) {
			// components here
		})

		assert.Nil(err)
		assert.NotNil(builder)
	})

	t.Run("CreateLegacyBuilder", func(t *testing.T) {
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
