package controls_test

import (
	"assert"
	"controls"
	"opts"
	"testing"
)

func TestLabel(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		label := controls.Label(
			opts.FontSize(24),
			opts.FontFace("Roboto"),
			opts.Text("Hello World"))

		assert.NotNil(t, label)
		// b.Layout()
		// label := b.Root()
		// assert.Equal(t, label.Width(), 107)
		// assert.Equal(t, label.Height(), 19)
	})
}
