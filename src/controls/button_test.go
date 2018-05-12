package controls

import (
	"assert"
	"events"
	"opts"
	"testing"
)

func TestButton(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		b := Button()
		assert.NotNil(t, b)
		assert.True(t, b.IsMeasured())
		assert.True(t, b.IsFocusable())
	})

	t.Run("states", func(t *testing.T) {
		b := Button(opts.Text("abcd"))

		assert.Equal(t, b.State(), "active")
		b.Emit(events.New(events.Entered, b, nil))
		assert.Equal(t, b.State(), "hovered")
		b.Emit(events.New(events.Pressed, b, nil))
		assert.Equal(t, b.State(), "pressed")
		b.Emit(events.New(events.Released, b, nil))
		assert.Equal(t, b.State(), "hovered")
		b.Emit(events.New(events.Exited, b, nil))
		assert.Equal(t, b.State(), "active")
	})
}
