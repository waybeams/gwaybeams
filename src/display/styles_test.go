package display

import (
	"testing"
)

func setUpBuilderAndRoot() (Builder, Displayable) {
	b := NewBuilder()
	root := NewSprite()
	b.Push(root)
	return b, root
}

func TestStyles(t *testing.T) {

	t.Run("Declaration", func(t *testing.T) {
		t.Skip()
		b, root := setUpBuilderAndRoot()

		StyleFor(b, Selector("Sprite"), BgColor(0xfc0))

		styles := root.GetStyles()

		if styles.GetBgColor() != 0xfc0 {
			t.Errorf("Expected BgColor to be assigned, but was %d", styles.GetBgColor())
		}
		// b, root := setUpBuilderAndRoot()
		// Style(b, BgColor(0xffcc00))
	})
}
