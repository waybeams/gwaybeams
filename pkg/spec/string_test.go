package spec_test

import (
	"path/filepath"
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/env/fake"
	"github.com/waybeams/waybeams/pkg/layout"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

func TestString(t *testing.T) {
	t.Skip("Checking how much fmt imports hurt browser binary size")

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
			opts.Child(ctrl.Box(
				opts.Child(ctrl.Button(
					opts.Width(10),
					opts.Height(10),
					opts.Text("One"),
				)),
			)),
		)
		layout.Layout(tree, fakes.NewSurfaceFrom(filepath.Join("..", "..")))
		result := `VBox(Width: 46.00, Height: 34.00
	Box(Width: 46.00, Height: 34.00
		Button(Width: 46.00, Height: 34.00, Text: One)
	)
)`
		str := spec.String(tree)
		assert.Equal("\n"+str, "\n"+result)
	})
}
