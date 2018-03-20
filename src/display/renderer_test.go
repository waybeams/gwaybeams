package display

import (
	"assert"
	"testing"
)

func TestFactory(t *testing.T) {
	surface := &FakeSurface{}

	t.Run("CreateRender", func(t *testing.T) {
		renderer := CreateRenderer(surface, func(s Surface) {
			Box(s, &Opts{Width: 100, Height: 100, BackgroundColor: 0xfc0}, func() {})
			// Box(s, &Opts{X: 10, Y: 10, Width: 80, Height: 80, BackgroundColor: 0xcf0})
			// })
		})

		box := renderer.GetRoot()
		assert.NotNil(box)
	})
}
