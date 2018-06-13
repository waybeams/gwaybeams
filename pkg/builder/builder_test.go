package builder_test

import (
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/builder"
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/surface"
	"github.com/waybeams/waybeams/pkg/win"
)

func TestBuilder(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		var b spec.Builder
		b = builder.New()
		assert.NotNil(b)
	})

	t.Run("Surface", func(t *testing.T) {
		fakeSurface := surface.NewFakeFrom("../../")
		b := builder.New(builder.Surface(fakeSurface))
		assert.Equal(b.Surface(), fakeSurface)
	})

	t.Run("Window", func(t *testing.T) {
		win := win.NewFake()
		assert.NotNil(win)
	})
}
