package display

import (
	"assert"
	"testing"
)

func TestToggleButton(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		btn, err := ToggleButton(NewBuilder())
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(t, btn.ChildCount(), 1)
		assert.Equal(t, btn.State(), ToggleUnselected)
	})

	t.Run("Selected", func(t *testing.T) {
		btn, _ := ToggleButton(NewBuilder())
		btn.SetState(ToggleSelected)
		btn.Builder().Update(btn)
		assert.Equal(t, btn.State(), ToggleSelected)
	})

	t.Run("CLICKED", func(t *testing.T) {
		// btn, _ := ToggleButton(NewBuilder())
		// btn.ChildAt(0).Bubble(NewEvent(events.Clicked, ))

	})
}
