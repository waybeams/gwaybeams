package display

import (
	"assert"
	"testing"
)

func TestFakeSurface(t *testing.T) {
	instance := FakeSurface{}

	t.Run("Instantiable", func(t *testing.T) {
		assert.NotNil(t, instance)
	})

	t.Run("SetRgba", func(t *testing.T) {
		instance.SetRgba(0x00, 0xff, 0xcc, 0xff)
		commands := instance.GetCommands()

		command := commands[0]
		assert.Equal(t, command.Name, "SetRgba")
		// Args are turned into strings for easier test comparisons
		assert.Equal(t, command.Args[0], uint(0x00))
		assert.Equal(t, command.Args[1], uint(0xff))
		assert.Equal(t, command.Args[2], uint(0xcc))
		assert.Equal(t, command.Args[3], uint(0xff))
	})
}
