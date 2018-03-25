package display

import (
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

		if styles.GetBgColor() != 0xfc0 {
			t.Errorf("Expected BgColor to be assigned, but was %d", styles.GetBgColor())
		}
	})
}
