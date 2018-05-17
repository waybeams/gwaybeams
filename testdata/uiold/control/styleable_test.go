package control

import (
	"github.com/waybeams/assert"
	"testing"
	"uiold/context"
	. "ui/controls"
)

func TestStyleable(t *testing.T) {

	t.Run("GetFontFace", func(t *testing.T) {
		root := Box(context.New())
		assert.Equal(t, root.FontFace(), "Roboto")
		assert.Equal(t, root.FontSize(), 24)
		assert.Equal(t, root.BgColor(), 0xce3262ff, "BgColor")
		assert.Equal(t, root.StrokeColor(), 0xffffffff, "StrokeColor")
	})
}
