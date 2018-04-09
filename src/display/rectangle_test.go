package display

import (
	"assert"
	"testing"
)

func TestRectangleView(t *testing.T) {

	t.Run("Sends some commands to surface", func(t *testing.T) {
		surface := &FakeSurface{}
		instance := NewComponent()
		RectangleView(surface, instance)

		commands := surface.GetCommands()
		assert.NotNil(t, commands)
	})

	t.Run("Uses zero x and y", func(t *testing.T) {
		surface := &FakeSurface{}
		instance, _ := TestComponent(NewBuilder(), BgColor(0xff0000ff), Width(100), Height(120))
		RectangleView(surface, instance)

		commands := surface.GetCommands()
		assert.Equal(t, commands[0].Name, "BeginPath", "Command Name")
		assert.Equal(t, commands[1].Name, "DrawRectangle", "Command Name")
	})
}
