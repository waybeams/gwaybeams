package display

import (
	"assert"
	"testing"
)

func setUpBuilderAndRoot() (Builder, Displayable) {
	b := NewBuilder()
	root := NewComponent()
	b.Push(root)
	return b, root
}

func TestStyles(t *testing.T) {

	t.Run("Declaration", func(t *testing.T) {
		box, _ := VBox(NewBuilder(), AttrStyles(BgColor(0xfc0), FontFace("sans"), FontSize(12)))
		styles := box.GetStyles()

		assert.Equal(t, styles.GetBgColor(), 0xfc0, "BgColor")
		assert.Equal(t, styles.GetFontFace(), "sans", "FontFace")
		assert.Equal(t, styles.GetFontSize(), 12, "FontSize")
	})
}
