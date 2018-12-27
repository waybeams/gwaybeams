package ctrl_test

import (
	"github.com/waybeams/assert"
	ctrl2 "github.com/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"testing"
)

func TestStyles(t *testing.T) {
	t.Run("Box style", func(t *testing.T) {
		styles := ctrl2.CreateStyles()
		instance := ctrl.VBox(styles.Box)
		assert.Equal(instance.PaddingTop(), 10)
		assert.Equal(instance.Gutter(), 10)
	})

	t.Run("Button style", func(t *testing.T) {
		styles := ctrl2.CreateStyles()
		instance := ctrl.VBox(styles.Button)
		assert.Equal(instance.BgColor(), 0xf8f8f8ff)
	})

	t.Run("Header style", func(t *testing.T) {
		styles := ctrl2.CreateStyles()
		instance := ctrl.VBox(styles.Header)
		assert.Equal(instance.FontFace(), "Roboto Light")
		assert.Equal(instance.FontSize(), 100)
	})

	t.Run("Main style", func(t *testing.T) {
		styles := ctrl2.CreateStyles()
		instance := ctrl.VBox(styles.Main)
		assert.Equal(instance.FontFace(), "Roboto")
		assert.Equal(instance.FontSize(), 24)
	})
}
