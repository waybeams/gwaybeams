package nano_test

import (
	"assert"
	"fmt"
	"github.com/shibukawa/nanovgo"
	"surface/nano"
	"testing"
)

func TestNanoSurface(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		instance := nano.New()
		assert.NotNil(t, instance)
	})

	t.Run("Debug Flag", func(t *testing.T) {
		instance := nano.New(nano.Debug())
		flags := instance.Flags()
		fmt.Println("YO:", (0 | 4), "vs", nanovgo.Debug)
		assert.Equal(t, flags, nanovgo.Debug)
	})

	t.Run("All Flags", func(t *testing.T) {
		instance := nano.New(nano.AntiAlias(), nano.StencilStrokes(), nano.Debug())
		flags := instance.Flags()
		assert.Equal(t, flags, nanovgo.AntiAlias|nanovgo.StencilStrokes|nanovgo.Debug)
	})
}
