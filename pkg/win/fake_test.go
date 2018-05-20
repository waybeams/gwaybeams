package win_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/win"
	"testing"
)

func TestWinFake(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		w := win.NewFake()
		w.SetWidth(30)
		w.SetHeight(40)
		assert.Equal(w.Width(), 30)
		assert.Equal(w.Height(), 40)
	})

	t.Run("GestureSource", func(t *testing.T) {
		g := win.NewFakeGestureSource()
		g.SetCursorPos(20, 30)
		x, y := g.GetCursorPos()

		assert.Equal(x, 20)
		assert.Equal(y, 30)
	})
}
