package spec_test

import (
	"controls"
	"github.com/waybeams/assert"
	"opts"
	"spec"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("Callable", func(t *testing.T) {
		str := spec.String(controls.HBox())
		assert.Equal(t, str, "HBox(Width: 0.00, Height: 0.00)")
	})

	t.Run("Handles nil spec", func(t *testing.T) {
		assert.Equal(t, spec.String(nil), "")
	})

	t.Run("Handles configured attrs", func(t *testing.T) {
		str := spec.String(controls.HBox(
			opts.Width(300.12345),
			opts.Height(200.00),
		))
		assert.Equal(t, str, "HBox(Width: 300.12, Height: 200.00)")
	})

	t.Run("Handles Children", func(t *testing.T) {

		tree := controls.VBox(
			opts.Child(controls.Label(opts.Text("Header"))),
			opts.Child(controls.Box(
				opts.Child(controls.Button(opts.Text("One"))),
				opts.Child(controls.Button(opts.Text("Two"))),
			)),
		)
		result := `VBox(Width: 10.00, Height: 10.00
	Label(Width: 0.00, Height: 0.00, Text: Header)
	Box(Width: 10.00, Height: 10.00
		Button(Width: 10.00, Height: 10.00, Text: One)
		Button(Width: 10.00, Height: 10.00, Text: Two)
	)
)`
		str := spec.String(tree)
		assert.Equal(t, "\n"+str, "\n"+result)
	})
}
