package controls_test

import (
	"controls"
	"github.com/waybeams/assert"
	"layout"
	"opts"
	"spec"
	"surface"
	"surface/nano"
	"testing"
)

func TestLabel(t *testing.T) {
	t.Run("Measured", func(t *testing.T) {
		label := controls.Label(
			opts.FontSize(24),
			opts.FontFace("Roboto"),
			opts.Text("Hello World"))

		assert.NotNil(t, label)
		s := nano.NewWithRoboto()
		layout.Layout(label, s)

		assert.Equal(t, label.Width(), 107)
		assert.Equal(t, label.Height(), 24)
	})

	t.Run("Component measure includes ascenders and descenders", func(t *testing.T) {
		var create = func() spec.ReadWriter {
			return controls.HBox(
				opts.BgColor(0x333333ff),
				opts.StrokeColor(0x000000ff),
				opts.StrokeSize(1),
				opts.Child(controls.Label(opts.Text("a"))), // vertically thin
				opts.Child(controls.Label(opts.Text("b"))), // ascender
				opts.Child(controls.Label(opts.Text("y"))), // lowercase descender
				opts.Child(controls.Label(opts.Text("A"))), // uppercase full size
				opts.Child(controls.Label(opts.Text("Q"))), // uppercase descender
			)
		}

		root := create()

		s := nano.NewWithRoboto()
		layout.Layout(root, s)

		kids := root.Children()
		assert.Equal(t, kids[0].Width(), 11, "one.w")
		assert.Equal(t, kids[0].Height(), 24, "one.h")

		assert.Equal(t, kids[1].Width(), 11, "two.w")
		assert.Equal(t, kids[1].Height(), 24, "two.h")

		assert.Equal(t, kids[2].Width(), 10, "three.w")
		assert.Equal(t, kids[2].Height(), 24, "three.h")

		assert.Equal(t, kids[3].Width(), 13, "four.w")
		assert.Equal(t, kids[3].Height(), 24, "four.h")

		assert.Equal(t, kids[4].Width(), 14, "five.w")
		assert.Equal(t, kids[4].Height(), 24, "five.h")

		// Draw a single letter and verify asc/desc/lineH offsets:
		fakeSurface := surface.NewFake()
		layout.Draw(root.ChildAt(0), fakeSurface)
		cmds := fakeSurface.GetCommands()

		assert.Equal(t, len(cmds), 4)
		assert.Equal(t, cmds[3].Name, "Text")
		args := cmds[3].Args
		assert.Equal(t, args[0], 0)
		assert.Equal(t, args[1], 13)
		assert.Equal(t, args[2], "a")
	})
}