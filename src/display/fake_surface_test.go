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

	t.Run("SetFillColor", func(t *testing.T) {
		instance.SetFillColor(0x00ffccff)
		commands := instance.GetCommands()

		command := commands[0]
		assert.Equal(t, command.Name, "SetFillColor")
		assert.Equal(t, command.Args[0], uint(0x00ffccff))
	})
}
