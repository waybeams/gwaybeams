package glfw_test

import (
	"github.com/waybeams/assert"
	g "github.com/go-gl/glfw/v3.2/glfw"
	"github.com/waybeams/waybeams/pkg/spec"
	"testing"
	"github.com/waybeams/waybeams/pkg/win/glfw"
)

func TestGlfwWindow(t *testing.T) {
	t.Run("Instantiable as spec.Window", func(t *testing.T) {
		var win spec.Window
		win = glfw.New()
		assert.NotNil(win)
	})

	t.Run("Size", func(t *testing.T) {
		win := glfw.New(glfw.Width(10), glfw.Height(20))
		w, h := win.Width(), win.Height()
		assert.Equal(w, 10)
		assert.Equal(h, 20)
	})

	t.Run("Title", func(t *testing.T) {
		win := glfw.New(glfw.Title("Hello World"))
		assert.Equal(win.Title(), "Hello World")
	})

	t.Run("Hint", func(t *testing.T) {
		win := glfw.New(glfw.Hint(g.Focused, 0))

		// There are a number of default hints
		hints := win.Hints()
		h := hints[len(hints)-1]
		assert.Equal(h.Key, g.Focused)
		assert.Equal(h.Value, 0)
	})
}
