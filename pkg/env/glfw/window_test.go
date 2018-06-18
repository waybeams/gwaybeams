package glfw_test

import (
	"testing"

	g "github.com/go-gl/glfw/v3.2/glfw"
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/env/glfw"
	"github.com/waybeams/waybeams/pkg/spec"
)

func TestGlfwWindow(t *testing.T) {
	t.Run("Instantiable as spec.Window", func(t *testing.T) {
		var win spec.Window
		win = glfw.NewWindow()
		assert.NotNil(win)
	})

	t.Run("Size", func(t *testing.T) {
		win := glfw.NewWindow(glfw.Width(10), glfw.Height(20))
		w, h := win.Width(), win.Height()
		assert.Equal(w, 10)
		assert.Equal(h, 20)
	})

	t.Run("Title", func(t *testing.T) {
		win := glfw.NewWindow(glfw.Title("Hello World"))
		assert.Equal(win.Title(), "Hello World")
	})

	t.Run("Hint", func(t *testing.T) {
		win := glfw.NewWindow(glfw.Hint(g.Focused, 0))

		// There are a number of default hints
		hints := win.Hints()
		h := hints[len(hints)-1]
		assert.Equal(h.Key, g.Focused)
		assert.Equal(h.Value, 0)
	})
}
