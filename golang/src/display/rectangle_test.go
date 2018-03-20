package display

import (
	"assert"
	"testing"
)

func TestDrawRectangle(t *testing.T) {
	surface := &FakeSurface{}

	t.Run("Draw with Sprite", func(t *testing.T) {
		sprite := NewSprite()
		DrawRectangle(surface, sprite)

		commands := surface.GetCommands()
		assert.NotNil(commands)
		// assert.Equal(len(commands), 0)
	})
}
