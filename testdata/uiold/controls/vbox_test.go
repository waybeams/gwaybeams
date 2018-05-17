package controls

import (
	"github.com/waybeams/assert"
	"testing"
	"ui"
	"uiold/context"
	"uiold/opts"
)

func TestVBox(t *testing.T) {

	t.Run("Simple Children", func(t *testing.T) {
		root := VBox(context.New(), opts.Height(100), opts.Children(func(c ui.Context) {
			Box(c, opts.FlexHeight(1))
			Box(c, opts.FlexHeight(1))
		}))

		one := root.ChildAt(0)
		two := root.ChildAt(1)
		assert.Equal(t, one.Height(), 50)
		assert.Equal(t, two.Height(), 50)
	})
}
