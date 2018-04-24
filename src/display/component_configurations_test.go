package display

import (
	"assert"
	"events"
	"testing"
)

func TestVBox(t *testing.T) {

	t.Run("Simple Children", func(t *testing.T) {
		var err error

		root, _ := VBox(NewBuilder(), Height(100), Children(func(b Builder) {
			_, err = Box(b, FlexHeight(1))
			Box(b, FlexHeight(1))
		}))

		one := root.ChildAt(0)
		two := root.ChildAt(1)
		if err != nil {
			t.Error(err)
		}
		if one.Height() != 50 {
			t.Errorf("one expected 50, but was %v", one.Height())
		}
		if two.Height() != 50 {
			t.Errorf("two expected 50, but was %v", one.Height())
		}
	})
}

func TestLabel(t *testing.T) {
	t.Run("Label", func(t *testing.T) {
		label, _ := Label(NewBuilder(), Title("Hello World"))
		assert.Equal(t, label.Title(), "Hello World")
	})
}

func TestButton(t *testing.T) {
	t.Run("Focusable", func(t *testing.T) {
		var calledWith Event
		var clickHandler = func(e Event) {
			calledWith = e
		}
		button, _ := Box(NewBuilder(), Text("Submit"), OnClick(clickHandler))
		button.Emit(NewEvent(events.Clicked, button, nil))
		assert.Equal(t, calledWith.Target(), button)
	})
}
