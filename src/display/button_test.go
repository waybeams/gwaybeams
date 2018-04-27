package display

import (
	"assert"
	"events"
	"testing"
)

func TestButton(t *testing.T) {
	t.Run("Callable", func(t *testing.T) {
		btn, _ := Button(NewBuilder(), Text("Submit"))
		assert.Equal(t, btn.Text(), "Submit")
	})

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
