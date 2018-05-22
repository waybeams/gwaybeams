package ctrl_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"testing"
)

func TestButton(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		b := ctrl.Button()
		assert.NotNil(b)
		assert.True(b.IsMeasured())
		assert.True(b.IsFocusable())
	})

	t.Run("states", func(t *testing.T) {
		b := ctrl.Button(opts.Text("abcd"))

		assert.Equal(b.State(), "active")
		b.Emit(events.New(events.Entered, b, nil))
		assert.Equal(b.State(), "hovered")
		b.Emit(events.New(events.Pressed, b, nil))
		assert.Equal(b.State(), "pressed")
		b.Emit(events.New(events.Released, b, nil))
		assert.Equal(b.State(), "hovered")
		b.Emit(events.New(events.Exited, b, nil))
		assert.Equal(b.State(), "active")
	})
}
