package glfw

import (
	"github.com/waybeams/waybeams/pkg/controls"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/layout"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/surface"
	"testing"
	"github.com/waybeams/waybeams/pkg/win"
)

func TestGlfwInput(t *testing.T) {
	t.Run("Emits entered and exited events", func(t *testing.T) {
		root := controls.VBox(
			opts.Key("Root"),
			opts.BgColor(0xffcc00ff),
			opts.Width(100),
			opts.Height(100),
			opts.Child(controls.Button(
				opts.Key("Button"),
				opts.FlexWidth(1),
				opts.FlexHeight(1),
			)),
			// opts.Child(controls.TextInput(
			opts.Child(controls.Label(
				opts.Key("TextInput"),
				opts.IsFocusable(true),
				opts.IsTextInput(true),
				opts.FlexWidth(1),
				opts.FlexHeight(1),
			)),
			opts.Child(controls.Label(
				opts.Key("Label"),
				opts.FlexWidth(1),
				opts.IsFocusable(true),
				opts.IsText(true),
				opts.FlexHeight(1),
			)),
		)
		layout.Layout(root, surface.NewFake())

		received := []events.Event{}
		var handler = func(e events.Event) {
			received = append(received, e)
		}
		root.On(events.Exited, handler)
		root.On(events.Entered, handler)

		fakeSource := win.NewFakeGestureSource()
		input := NewGlfwInput(fakeSource)

		fakeSource.SetCursorPos(10, 10)
		input.Update(root)
		assert.Equal(received[0].Name(), events.Entered)
		assert.Equal(spec.Path(received[0].Target().(spec.Reader)), spec.Path(root.ChildAt(0)), "entered 1")
		assert.Equal(len(received), 1)

		fakeSource.SetCursorPos(10, 40)
		input.Update(root)

		assert.Equal(len(received), 3)
		assert.Equal(received[1].Name(), events.Exited)
		assert.Equal(spec.Path(received[1].Target().(spec.Reader)), spec.Path(root.ChildAt(0)), "exited 1")

		assert.Equal(received[2].Name(), events.Entered)
		assert.Equal(spec.Path(received[2].Target().(spec.Reader)), spec.Path(root.ChildAt(1)), "entered 2")
		assert.Equal(fakeSource.CursorName, glfw.IBeamCursor)

		fakeSource.SetCursorPos(10, 70)
		input.Update(root)

		assert.Equal(len(received), 5, "received should be five")

		assert.Equal(received[3].Name(), events.Exited)
		assert.Equal(spec.Path(received[3].Target().(spec.Reader)), spec.Path(root.ChildAt(1)), "exited 2")

		assert.Equal(received[4].Name(), events.Entered)
		assert.Equal(spec.Path(received[4].Target().(spec.Reader)), spec.Path(root.ChildAt(2)), "entered 2")
	})
}
