package controls

import (
	"assert"
	"testing"
	"ui/context"
)

func TestToggleButton(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		btn := ToggleButton(context.New())
		assert.Equal(t, btn.ChildCount(), 1)
		assert.Equal(t, btn.State(), ToggleUnselected)
	})

	t.Run("Selected", func(t *testing.T) {
		btn := ToggleButton(context.New())
		btn.SetState(ToggleSelected)
		btn.Context().Builder().Update(btn)
		assert.Equal(t, btn.State(), ToggleSelected)
	})
}
