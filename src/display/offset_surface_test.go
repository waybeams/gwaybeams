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

		assert.Equal(t, root.GetXOffset(), 0)
		assert.Equal(t, root.GetYOffset(), 0)
		assert.Equal(t, child.GetXOffset(), 10)
		assert.Equal(t, child.GetYOffset(), 10)

		// TODO(lbayes): Extract this mess out to the FakeSurface,
		// and implement some custom validations/assertions on that entity.

		for i := len(commands) - 1; i >= 0; i-- {
			// Find the last call to DrawRectangle and ensure the values were offset properly.
			command := commands[i]
			if command.Name == "DrawRectangle" {
				args := command.Args
				assert.Equal(t, 10, args[0], "x")
				assert.Equal(t, 10, args[1], "y")
				assert.Equal(t, 80, args[2], "width")
				assert.Equal(t, 80, args[3], "height")
				break
			}
		}
	})
}
