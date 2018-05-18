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
		assert.Equal(root.FontFace(), "Roboto")
		assert.Equal(root.FontSize(), 24)
		assert.Equal(root.BgColor(), 0xce3262ff, "BgColor")
		assert.Equal(root.StrokeColor(), 0xffffffff, "StrokeColor")
	})
}
