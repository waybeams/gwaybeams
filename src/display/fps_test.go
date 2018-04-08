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
		win, _ := TestWindow(NewBuilder(), Children(func(b Builder) {
			FPS(b)
		}))
		assert.NotNil(t, win)
		// Exploring ideas around spawning a window while developing components.
		// Currently failing b/c OS X wants all UI calls on the main thread and
		// Go tests don't seem to be playing nice with that.
		// win.(Window).Init()
	})
}
