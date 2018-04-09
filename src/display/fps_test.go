package display

import (
	"assert"
	"testing"
)

func TestFps(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		instance, _ := FPS(NewBuilder())
		if instance == nil {
			t.Error("Expected an instance")
		}
	})

	t.Run("Renders FPS", func(t *testing.T) {
		instance, _ := FPS(NewBuilder())

		assert.NotNil(t, instance)
	})
}
