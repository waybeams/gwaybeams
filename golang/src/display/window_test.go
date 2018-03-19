package display

import (
	"assert"
	"testing"
)

func TestWindow(t *testing.T) {
	f := NewFactory()

	t.Run("Creation", func(t *testing.T) {
		Window(f)
		assert.NotNil(f.GetRoot())
	})
}
