package display

import (
	"assert"
	"testing"
)

type FakeToggle struct {
	ToggleButtonComponent
}

func TestToggleButton(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		instance, _ := ToggleButton(NewBuilder())
		assert.NotNil(t, instance)
	})

	t.Run("Child component", func(t *testing.T) {
		instance := &FakeToggle{}
		sel := Selected(true)
		sel(instance)
		assert.True(t, instance.Selected())
	})
}
