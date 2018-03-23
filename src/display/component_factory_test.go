package display

import (
	"testing"
)

type FakeComponent struct {
	Component
}

func NewFake() Displayable {
	return &FakeComponent{}
}

// Create a new factory using our component creation function reference.
var Fake = NewComponentFactory(NewFake)

func TestComponentFactory(t *testing.T) {
	t.Run("No Builder", func(t *testing.T) {
		box, _ := Box(nil, Id("root"), Children(func(b Builder) {
			Box(b, Id("one"))
			Box(b, Id("two"))
		}))
		if box.GetId() != "root" {
			t.Error("Expected a configured Box component")
		}
	})

	t.Run("Custom type", func(t *testing.T) {
		fake, _ := Fake(nil)
		if fake == nil {
			t.Error("Expected builder to return new component")
		}
	})

	t.Run("Padding", func(t *testing.T) {
		sprite, _ := Box(nil, Padding(10))

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
		box, _ := Box(nil, Padding(10), PaddingLeft(15))
		if box.GetVerticalPadding() != 20 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.GetHorizontalPadding() != 25 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.GetPaddingLeft() != 15 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if box.GetPaddingRight() != 10 {
			t.Error("Expected Padding to update PaddingRight")
		}
	})

	t.Run("Padding with specifics is NOT order dependent", func(t *testing.T) {
		box, _ := Box(nil, PaddingLeft(15), Padding(10))

		if box.GetHorizontalPadding() != 25 {
			t.Error("Expected additive HorizontalPadding")
		}
	})

	t.Run("Padding with specifics will clobber a ZERO setting", func(t *testing.T) {
		box, _ := Box(nil, PaddingLeft(0), Padding(10))

		// We only look for the "ZERO VALUE" when trying to figure out if we should
		// clobber. But users can set this, so we're a little jammed up here, unless
		// we flag on any/all interrelated value options. :-(
		if box.GetHorizontalPadding() != 20 {
			t.Error("Expected additive HorizontalPadding")
		}
	})

	t.Run("Specific Paddings", func(t *testing.T) {
		box, _ := Box(nil, PaddingBottom(1), PaddingRight(2), PaddingLeft(3), PaddingTop(4))

		if box.GetVerticalPadding() != 5 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.GetHorizontalPadding() != 5 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.GetPaddingLeft() != 3 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if box.GetPaddingRight() != 2 {
			t.Error("Expected Padding to update PaddingRight")
		}
		if box.GetPaddingTop() != 4 {
			t.Error("Expected Padding to update PaddingTop")
		}
		if box.GetPaddingBottom() != 1 {
			t.Error("Expected Padding to update PaddingBottom")
		}
	})
}
