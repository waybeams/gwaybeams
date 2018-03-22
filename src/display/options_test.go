package display

import (
	"testing"
)

type FakeComponent struct {
	SpriteComponent
}

func NewFake() Displayable {
	return &FakeComponent{}
}

// Create a new factory using our component creation function reference.
var Fake = NewComponentFactory(NewFake)

func TestComponentFactory(t *testing.T) {

	t.Run("Custom type", func(t *testing.T) {
		b := NewBuilder()
		Fake(b)
	})

	t.Run("Padding", func(t *testing.T) {
		sprite, _ := NewBuilder().Build(func(b Builder) {
			Sprite(b, Padding(10))
		})

		if sprite.GetPadding() != 10 {
			t.Error("Expected option to set padding")
		}
		if sprite.GetHorizontalPadding() != 20 {
			t.Error("Expected Padding to update HorizontalPadding")
		}
		if sprite.GetVerticalPadding() != 20 {
			t.Error("Expected Padding to update VerticalPadding")
		}
		if sprite.GetPaddingBottom() != 10 {
			t.Error("Expected Padding to update PaddingBottom")
		}
		if sprite.GetPaddingLeft() != 10 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if sprite.GetPaddingRight() != 10 {
			t.Error("Expected Padding to update PaddingRight")
		}
		if sprite.GetPaddingTop() != 10 {
			t.Error("Expected Padding to update PaddingTop")
		}
	})

	t.Run("Padding with specifics", func(t *testing.T) {
		sprite, _ := NewBuilder().Build(func(b Builder) {
			Sprite(b, Padding(10), PaddingLeft(15))
		})
		if sprite.GetVerticalPadding() != 20 {
			t.Error("Expected additive HorizontalPadding")
		}
		if sprite.GetHorizontalPadding() != 25 {
			t.Error("Expected additive HorizontalPadding")
		}
		if sprite.GetPaddingLeft() != 15 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if sprite.GetPaddingRight() != 10 {
			t.Error("Expected Padding to update PaddingRight")
		}
	})

	t.Run("Padding with specifics is NOT order dependent", func(t *testing.T) {
		sprite, _ := NewBuilder().Build(func(b Builder) {
			Sprite(b, PaddingLeft(15), Padding(10))
		})
		if sprite.GetHorizontalPadding() != 25 {
			t.Error("Expected additive HorizontalPadding")
		}
	})

	t.Run("Padding with specifics will clobber a ZERO setting", func(t *testing.T) {
		sprite, _ := NewBuilder().Build(func(b Builder) {
			Sprite(b, PaddingLeft(0), Padding(10))
		})
		// We only look for the "ZERO VALUE" when trying to figure out if we should
		// clobber. But users can set this, so we're a little jammed up here, unless
		// we flag on any/all interrelated value options. :-(
		if sprite.GetHorizontalPadding() != 20 {
			t.Error("Expected additive HorizontalPadding")
		}
	})

	t.Run("Specific Paddings", func(t *testing.T) {
		sprite, _ := NewBuilder().Build(func(b Builder) {
			Sprite(b, PaddingBottom(1), PaddingRight(2), PaddingLeft(3), PaddingTop(4))
		})
		if sprite.GetVerticalPadding() != 5 {
			t.Error("Expected additive HorizontalPadding")
		}
		if sprite.GetHorizontalPadding() != 5 {
			t.Error("Expected additive HorizontalPadding")
		}
		if sprite.GetPaddingLeft() != 3 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if sprite.GetPaddingRight() != 2 {
			t.Error("Expected Padding to update PaddingRight")
		}
		if sprite.GetPaddingTop() != 4 {
			t.Error("Expected Padding to update PaddingTop")
		}
		if sprite.GetPaddingBottom() != 1 {
			t.Error("Expected Padding to update PaddingBottom")
		}
	})
}
