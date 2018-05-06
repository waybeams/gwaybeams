package controls_test

import (
	"assert"
	. "controls"
	"ctx"
	"testing"
)

func TestToggleButton(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		btn := ToggleButton(ctx.New())
		assert.Equal(t, btn.ChildCount(), 1)
		assert.Equal(t, btn.State(), ToggleUnselected)
	})

	t.Run("Selected", func(t *testing.T) {
		btn := ToggleButton(ctx.New())
		btn.SetState(ToggleSelected)
		btn.Context().Builder().Update(btn)
		assert.Equal(t, btn.State(), ToggleSelected)
	})
}
