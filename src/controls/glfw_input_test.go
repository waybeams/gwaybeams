package controls_test

import (
	"assert"
	. "controls"
	"ctx"
	"events"
	"github.com/go-gl/glfw/v3.2/glfw"
	. "opts"
	"testing"
	. "ui"
)

type FakeGestureSource struct {
	xpos          float64
	ypos          float64
	CursorName    glfw.StandardCursor
	CharCallback  CharCallback
	MouseCallback MouseButtonCallback
}

func (f *FakeGestureSource) SetCursorPos(xpos, ypos float64) {
	f.xpos = xpos
	f.ypos = ypos
}

func (f *FakeGestureSource) GetCursorPos() (xpos, ypos float64) {
	return f.xpos, f.ypos
}

func (f *FakeGestureSource) SetCursorByName(name glfw.StandardCursor) {
	f.CursorName = name
}

func (f *FakeGestureSource) SetCharCallback(callback CharCallback) events.Unsubscriber {
	f.CharCallback = callback
	return func() bool {
		f.CharCallback = nil
		return true
	}
}

func (f *FakeGestureSource) SetMouseButtonCallback(callback MouseButtonCallback) events.Unsubscriber {
	f.MouseCallback = callback
	return func() bool {
		f.MouseCallback = nil
		return true
	}
}

func TestGlfwInput(t *testing.T) {
	t.Run("Emits entered and exited events", func(t *testing.T) {
		root := VBox(ctx.New(), BgColor(0xffcc00ff), Width(100), Height(100), Children(func(c Context) {
			Button(c, FlexWidth(1), FlexHeight(1))
			TextInput(c, FlexWidth(1), FlexHeight(1))
			Label(c, FlexWidth(1), FlexHeight(1))
		}))
		received := []events.Event{}
		var handler = func(e events.Event) {
			received = append(received, e)
		}
		root.On(events.Exited, handler)
		root.On(events.Entered, handler)

		fakeSource := &FakeGestureSource{}
		input := NewGlfwInput(root, fakeSource)

		fakeSource.SetCursorPos(10, 10)
		input.Update()
		assert.Equal(t, received[0].Name(), events.Entered)
		assert.Equal(t, received[0].Target().(Composable).Path(), root.ChildAt(0).Path(), "entered 1")
		assert.Equal(t, len(received), 1)

		fakeSource.SetCursorPos(10, 40)
		input.Update()

		assert.Equal(t, len(received), 3)
		assert.Equal(t, received[1].Name(), events.Exited)
		assert.Equal(t, received[1].Target().(Composable).Path(), root.ChildAt(0).Path(), "exited 1")

		assert.Equal(t, received[2].Name(), events.Entered)
		assert.Equal(t, received[2].Target().(Composable).Path(), root.ChildAt(1).Path(), "entered 2")
		assert.Equal(t, fakeSource.CursorName, glfw.IBeamCursor)

		fakeSource.SetCursorPos(10, 70)
		input.Update()

		assert.Equal(t, len(received), 5, "received should be five")

		assert.Equal(t, received[3].Name(), events.Exited)
		assert.Equal(t, received[3].Target().(Composable).Path(), root.ChildAt(1).Path(), "exited 2")

		assert.Equal(t, received[4].Name(), events.Entered)
		assert.Equal(t, received[4].Target().(Composable).Path(), root.ChildAt(2).Path(), "entered 2")
	})
}
