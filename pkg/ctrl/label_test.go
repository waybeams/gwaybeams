package ctrl_test

import (
	"path/filepath"
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/layout"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/surface"
	"github.com/waybeams/waybeams/pkg/surface/nano"
)

func TestLabel(t *testing.T) {
	t.Run("Measured", func(t *testing.T) {
		label := ctrl.Label(
			opts.FontSize(24),
			opts.FontFace("Roboto"),
			opts.Text("Hello World"))

		assert.NotNil(label)
		s := nano.NewWithRoboto()
		layout.Layout(label, s)

		assert.Equal(label.Width(), 107)
		assert.Equal(label.Height(), 24)
	})

	t.Run("Component measure includes ascenders and descenders", func(t *testing.T) {
		var create = func() spec.ReadWriter {
			return ctrl.HBox(
				opts.BgColor(0x333333ff),
				opts.StrokeColor(0x000000ff),
				opts.StrokeSize(1),
				opts.Child(ctrl.Label(opts.Text("a"))), // vertically thin
				opts.Child(ctrl.Label(opts.Text("b"))), // ascender
				opts.Child(ctrl.Label(opts.Text("y"))), // lowercase descender
				opts.Child(ctrl.Label(opts.Text("A"))), // uppercase full size
				opts.Child(ctrl.Label(opts.Text("Q"))), // uppercase descender
			)
		}

		root := create()

		s := nano.NewWithRoboto()
		layout.Layout(root, s)

		kids := root.Children()
		assert.Equal(kids[0].Width(), 11, "one.w")
		assert.Equal(kids[0].Height(), 24, "one.h")

		assert.Equal(kids[1].Width(), 11, "two.w")
		assert.Equal(kids[1].Height(), 24, "two.h")

		assert.Equal(kids[2].Width(), 10, "three.w")
		assert.Equal(kids[2].Height(), 24, "three.h")

		assert.Equal(kids[3].Width(), 13, "four.w")
		assert.Equal(kids[3].Height(), 24, "four.h")

		assert.Equal(kids[4].Width(), 14, "five.w")
		assert.Equal(kids[4].Height(), 24, "five.h")

		// Draw a single letter and verify asc/desc/lineH offsets:
		fakeSurface := surface.NewFakeFrom(filepath.Join("..", ".."))
		layout.Draw(root.ChildAt(0), fakeSurface)
		cmds := fakeSurface.GetCommands()

		assert.Equal(len(cmds), 4)
		assert.Equal(cmds[1].Name, "SetFontFace")
		// NOTE(lbayes): The following will fail if AddFont is not called in the
		// fake surface.
		args := cmds[3].Args
		assert.Equal(args[0], 0)
		assert.Equal(args[1], 13)
		assert.Equal(args[2], "a")
	})
}
