package builder_test

import (
	"assert"
	"builder"
	"spec"
	"surface"
	"testing"
	"win"
)

func TestBuilder(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		var b spec.Builder
		b = builder.New()
		assert.NotNil(t, b)
	})

	t.Run("Surface", func(t *testing.T) {
		fakeSurface := surface.NewFake()
		b := builder.New(builder.Surface(fakeSurface))
		assert.Equal(t, b.Surface(), fakeSurface)
	})

	t.Run("Window", func(t *testing.T) {
		win := win.NewFake()
		assert.NotNil(t, win)
	})
}
