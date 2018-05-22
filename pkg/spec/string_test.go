package spec_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("Callable", func(t *testing.T) {
		str := spec.String(ctrl.HBox())
		assert.Equal(str, "HBox(Width: 0.00, Height: 0.00)")
	})

	t.Run("Handles nil spec", func(t *testing.T) {
		assert.Equal(spec.String(nil), "")
	})

	t.Run("Handles configured attrs", func(t *testing.T) {
		str := spec.String(ctrl.HBox(
			opts.Width(300.12345),
			opts.Height(200.00),
		))
		assert.Equal(str, "HBox(Width: 300.12, Height: 200.00)")
	})

	t.Run("Handles Children", func(t *testing.T) {

		tree := ctrl.VBox(
			opts.Child(ctrl.Label(opts.Text("Header"))),
			opts.Child(ctrl.Box(
				opts.Child(ctrl.Button(opts.Text("One"))),
				opts.Child(ctrl.Button(opts.Text("Two"))),
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
		assert.Equal("\n"+str, "\n"+result)
	})
}
