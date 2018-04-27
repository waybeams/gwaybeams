package display_test

import (
	"assert"
	. "display"
	"events"
	"testing"
)

type FakeGestureWindow struct {
	xpos float64
	ypos float64
}

func (f *FakeGestureWindow) SetCursorPos(xpos, ypos float64) {
	f.xpos = xpos
	f.ypos = ypos
}

func (f *FakeGestureWindow) GetCursorPos() (xpos, ypos float64) {
	return f.xpos, f.ypos
}

func TestGlfwInput(t *testing.T) {
	t.Run("Emits entered and exited events", func(t *testing.T) {
		root, _ := HBox(NewBuilder(), Width(100), Height(100), Children(func(b Builder) {
			Button(b, FlexWidth(1), FlexHeight(1))
			Button(b, FlexWidth(1), FlexHeight(1))
		}))
		received := []Event{}
		var handler = func(e Event) {
			received = append(received, e)
		}
		root.On(events.Exited, handler)
		root.On(events.Entered, handler)

		fakeWindow := &FakeGestureWindow{}
		input := NewGlfwInput(root, fakeWindow)

		fakeWindow.SetCursorPos(10, 10)
		input.Update()
		assert.Equal(t, received[0].Name(), events.Entered)
		assert.Equal(t, received[0].Target(), root.FirstChild(), "entered 1")

		fakeWindow.SetCursorPos(60, 10)
		input.Update()
		assert.Equal(t, received[1].Name(), events.Exited)
		assert.Equal(t, received[1].Target(), root.FirstChild(), "exited 1")

		assert.Equal(t, received[2].Name(), events.Entered)
		assert.Equal(t, received[2].Target(), root.LastChild(), "entered 2")

		assert.Equal(t, len(received), 3)
	})
}
