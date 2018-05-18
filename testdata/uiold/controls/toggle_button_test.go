package controls

import (
	"github.com/waybeams/assert"
	"testing"
	"uiold/context"
)

func TestToggleButton(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		btn := ToggleButton(context.New())
		assert.Equal(btn.ChildCount(), 1)
		assert.Equal(btn.State(), ToggleUnselected)
	})

	t.Run("Selected", func(t *testing.T) {
		btn := ToggleButton(context.New())
		btn.SetState(ToggleSelected)
		btn.Context().Builder().Update(btn)
		assert.Equal(btn.State(), ToggleSelected)
	})
}
