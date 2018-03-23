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
		// b, root := setUpBuilderAndRoot()
		// Style(b, BgColor(0xffcc00))
	})
}
