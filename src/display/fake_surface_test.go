package display

import (
	"assert"
	"testing"
)

func TestFakeSurface(t *testing.T) {
	instance := FakeSurface{}

	t.Run("Instantiable", func(t *testing.T) {
		assert.NotNil(instance)
	})

	t.Run("SetRgba", func(t *testing.T) {
		instance.SetRgba(0.8, 0.7, 0.6, 0.5)
		commands := instance.GetCommands()

		command := commands[0]
		assert.Equal(t, command.Name, "SetRgba")
		// Args are turned into strings for easier test comparisons
		assert.Equal(t, command.Args[0], 0.8)
		assert.Equal(t, command.Args[1], 0.7)
		assert.Equal(t, command.Args[2], 0.6)
		assert.Equal(t, command.Args[3], 0.5)
	})
}
