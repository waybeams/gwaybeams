package display

import (
	"assert"
	"testing"
)

func TestLabel(t *testing.T) {
	t.Run("Simple Label", func(t *testing.T) {
		label, _ := Label(NewBuilder(), Title("Hello World"))
		assert.Equal(t, label.GetTitle(), "Hello World")
	})

	t.Run("Label draws", func(t *testing.T) {

	})
}
