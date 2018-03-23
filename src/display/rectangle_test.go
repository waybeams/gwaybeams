package display

import (
	"assert"
	"testing"
)

func TestDrawRectangle(t *testing.T) {
	surface := &FakeSurface{}

	t.Run("Draw with Box", func(t *testing.T) {
		sprite := NewComponent()
		DrawRectangle(surface, sprite)

		commands := surface.GetCommands()
		assert.NotNil(commands)
		// assert.Equal(len(commands), 0)
	})
}
