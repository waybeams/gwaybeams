package nano_test

import (
	"github.com/waybeams/assert"
	"fmt"
	"github.com/shibukawa/nanovgo"
	"github.com/waybeams/waybeams/pkg/surface/nano"
	"testing"
)

func TestNanoSurface(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		instance := nano.New()
		assert.NotNil(instance)
	})

	t.Run("Debug Flag", func(t *testing.T) {
		instance := nano.New(nano.Debug())
		flags := instance.Flags()
		fmt.Println("YO:", (0 | 4), "vs", nanovgo.Debug)
		assert.Equal(flags, nanovgo.Debug)
	})

	t.Run("All Flags", func(t *testing.T) {
		instance := nano.New(nano.AntiAlias(), nano.StencilStrokes(), nano.Debug())
		flags := instance.Flags()
		assert.Equal(flags, nanovgo.AntiAlias|nanovgo.StencilStrokes|nanovgo.Debug)
	})
}
