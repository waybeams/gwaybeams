package controls_test

import (
	"assert"
	. "controls"
	"ctx"
	"events"
	. "opts"
	"testing"
)

func TestButton(t *testing.T) {
	t.Run("Callable", func(t *testing.T) {
		btn := Button(ctx.New(), Text("Submit"))
		assert.Equal(t, btn.Text(), "Submit")
	})

	t.Run("Focusable", func(t *testing.T) {
		var calledWith events.Event
		var clickHandler = func(e events.Event) {
			calledWith = e
		}
		button := Box(ctx.New(), Text("Submit"), OnClick(clickHandler))
		button.Emit(events.New(events.Clicked, button, nil))
		assert.Equal(t, calledWith.Target(), button)
	})

	t.Run("Uses label dimensions", func(t *testing.T) {
		btn := Button(ctx.NewTestContext(), TestOptions(), Text("Submit Something"))
		assert.Equal(t, btn.Width(), 102)
		assert.Equal(t, btn.Height(), 31)
	})
}
