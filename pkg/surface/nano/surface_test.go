package nano_test

import (
	"testing"

	"github.com/shibukawa/nanovgo"
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/surface/nano"
)

func TestNanoSurface(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		instance := nano.NewSurface()
		assert.NotNil(instance)
	})

	t.Run("Debug Flag", func(t *testing.T) {
		instance := nano.NewSurface(nano.Debug())
		flags := instance.Flags()
		assert.Equal(flags, nanovgo.Debug)
	})

	t.Run("All Flags", func(t *testing.T) {
		instance := nano.NewSurface(nano.AntiAlias(), nano.StencilStrokes(), nano.Debug())
		flags := instance.Flags()
		assert.Equal(flags, nanovgo.AntiAlias|nanovgo.StencilStrokes|nanovgo.Debug)
	})
}
