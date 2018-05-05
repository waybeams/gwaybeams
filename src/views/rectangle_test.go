package views_test

import (
	"assert"
	"component"
	"controls"
	"ctx"
	. "opts"
	"surface"
	"testing"
	. "views"
)

func TestRectangleView(t *testing.T) {

	t.Run("Sends some commands to surface", func(t *testing.T) {
		surface := &surface.Fake{}
		instance := component.New()
		RectangleView(surface, instance)

		commands := surface.GetCommands()
		assert.NotNil(t, commands)
	})

	t.Run("Uses zero x and y", func(t *testing.T) {
		surface := &surface.Fake{}
		instance := controls.TestComponent(ctx.New(), BgColor(0xff0000ff), Width(100), Height(120))
		RectangleView(surface, instance)

		commands := surface.GetCommands()
		assert.Equal(t, commands[0].Name, "BeginPath", "Command Name")
		assert.Equal(t, commands[1].Name, "Rect", "Command Name")
	})
}
