package display

import (
	"assert"
	"testing"
)

func TestDrawRectangle(t *testing.T) {

	t.Run("Sends some commands to surface", func(t *testing.T) {
		surface := &FakeSurface{}
		instance := NewComponent()
		DrawRectangle(surface, instance)

		commands := surface.GetCommands()
		assert.NotNil(t, commands)
	})

	t.Run("Uses zero x and y", func(t *testing.T) {
		surface := &FakeSurface{}
		instance, _ := TestComponent(NewBuilder(), Width(100), Height(120))
		DrawRectangle(surface, instance)

		commands := surface.GetCommands()
		assert.Equal(t, commands[0].Name, "SetFillColor", "Command Name")
		assert.Equal(t, commands[0].Args[0], 0xccccccff, "Color Value")
	})
}
