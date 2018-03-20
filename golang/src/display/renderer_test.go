package display

import (
	"assert"
	"testing"
)

func TestFactory(t *testing.T) {
	surface := &FakeSurface{}
	render := CreateRenderer(surface)

	t.Run("Instantiable", func(t *testing.T) {
		assert.NotNil(render)
		assert.NotNil(surface)
	})

	t.Run("SetRgba", func(t *testing.T) {
		render(func(s Surface) {
			Box(s, &Opts{Width: 100, Height: 100, BackgroundColor: 0xfc0}, func() {})
			// Box(s, &Opts{X: 10, Y: 10, Width: 80, Height: 80, BackgroundColor: 0xcf0})
			// })
		})
	})
}
