package spec_test

import (
	"testing"
)

func TestOffsetSurface(t *testing.T) {
	/*
		t.Run("Receives offset for padding", func(t *testing.T) {
			surface := &Fake{}
			var root, child spec.ReadWrite
			root = controls.Box(context.New(), opts.Padding(10), opts.Width(100), opts.Height(100), opts.Children(func(b ui.Context) {
				child = controls.Box(b, opts.FlexWidth(1), opts.FlexHeight(1))
			}))
			root.Draw(surface)
			commands := surface.GetCommands()

			assert.Equal(root.XOffset(), 0)
			assert.Equal(root.YOffset(), 0)
			assert.Equal(child.XOffset(), 10)
			assert.Equal(child.YOffset(), 10)

			// TODO(lbayes): Extract this mess out to the Fake,
			// and implement some custom validations/assertions on that entity.

			rectangles := []Command{}
			for _, command := range commands {
				if command.Name == "Rect" {
					rectangles = append(rectangles, command)
				}
			}

			assert.Equal(len(rectangles), 4, "Rectangle count")

			rootArgs := rectangles[0].Args
			assert.Equal(rootArgs[0], 0, "root x")
			assert.Equal(rootArgs[1], 0, "root y")
			assert.Equal(rootArgs[2], 100, "root width")
			assert.Equal(rootArgs[3], 100, "root height")

			childArgs := rectangles[1].Args
			assert.Equal(childArgs[0], -0.5, "child x")
			assert.Equal(childArgs[1], -0.5, "child y")
			assert.Equal(childArgs[2], 101, "child width")
			assert.Equal(childArgs[3], 101, "child height")
		})
	*/
}
