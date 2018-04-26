package display

import (
	"assert"
	"testing"
)

func TestStyleable(t *testing.T) {

	t.Run("GetFontFace", func(t *testing.T) {
		root, _ := Box(NewBuilder())
		assert.Equal(t, root.FontFace(), "Roboto")
		assert.Equal(t, root.FontSize(), 12)
		assert.Equal(t, root.BgColor(), 0x999999ff, "BgColor")
		assert.Equal(t, root.StrokeColor(), 0x333333ff, "StrokeColor")
	})
}
