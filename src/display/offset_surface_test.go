package display

import (
	"assert"
	"testing"
)

func TestOffsetSurface(t *testing.T) {
	t.Run("Receives offset for padding", func(t *testing.T) {
		surface := &FakeSurface{}
		var root, child Displayable
		root, _ = Box(NewBuilder(), Padding(10), Width(100), Height(100), Children(func(b Builder) {
			child, _ = Box(b, FlexWidth(1), FlexHeight(1))
		}))
		root.Layout()
		root.Draw(surface)
		commands := surface.GetCommands()

		assert.Equal(t, root.XOffset(), 0)
		assert.Equal(t, root.YOffset(), 0)
		assert.Equal(t, child.XOffset(), 10)
		assert.Equal(t, child.YOffset(), 10)

		// TODO(lbayes): Extract this mess out to the FakeSurface,
		// and implement some custom validations/assertions on that entity.

		rectangles := []SurfaceCommand{}
		for _, command := range commands {
			if command.Name == "DrawRectangle" {
				rectangles = append(rectangles, command)
			}
		}

		assert.Equal(t, len(rectangles), 2, "Rectangle count")

		rootArgs := rectangles[0].Args
		assert.Equal(t, rootArgs[0], 0, "root x")
		assert.Equal(t, rootArgs[1], 0, "root y")
		assert.Equal(t, rootArgs[2], 100, "root width")
		assert.Equal(t, rootArgs[3], 100, "root height")

		childArgs := rectangles[1].Args
		assert.Equal(t, childArgs[0], 10, "child x")
		assert.Equal(t, childArgs[1], 10, "child y")
		assert.Equal(t, childArgs[2], 80, "child width")
		assert.Equal(t, childArgs[3], 80, "child height")
	})
}
