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

	t.Run("Try it", func(t *testing.T) {
		fps, _ := FPS(NewBuilder())
		assert.NotNil(t, fps)
	})
}
