package controls_test

import (
	"assert"
	"controls"
	"ctx"
	"opts"
	"testing"
)

func TestLabel(t *testing.T) {
	t.Run("Label", func(t *testing.T) {
		label := controls.Label(ctx.New(), opts.Title("Hello World"))
		assert.Equal(t, label.Title(), "Hello World")
		// assert.Equal(t, label.Width(), 213, "Width detected")
	})
}
