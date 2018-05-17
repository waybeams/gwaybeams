package glfw_test

import (
	"github.com/waybeams/assert"
	g "github.com/go-gl/glfw/v3.2/glfw"
	"spec"
	"testing"
	"win/glfw"
)

func TestGlfwWindow(t *testing.T) {
	t.Run("Instantiable as spec.Window", func(t *testing.T) {
		var win spec.Window
		win = glfw.New()
		assert.NotNil(t, win)
	})

	t.Run("Size", func(t *testing.T) {
		win := glfw.New(glfw.Width(10), glfw.Height(20))
		w, h := win.Width(), win.Height()
		assert.Equal(t, w, 10)
		assert.Equal(t, h, 20)
	})

	t.Run("Title", func(t *testing.T) {
		win := glfw.New(glfw.Title("Hello World"))
		assert.Equal(t, win.Title(), "Hello World")
	})

	t.Run("Hint", func(t *testing.T) {
		win := glfw.New(glfw.Hint(g.Focused, 0))

		// There are a number of default hints
		hints := win.Hints()
		h := hints[len(hints)-1]
		assert.Equal(t, h.Key, g.Focused)
		assert.Equal(t, h.Value, 0)
	})
}
