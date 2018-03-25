package display

import (
	"assert"
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
	t.Run("Default State", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		// These two assertions don't appear to be passing my custom equality check. :barf:
		if box.GetHAlign() != LeftAlign {
			t.Error("Expected LeftAlign, but got: %v", box.GetHAlign())
		}
		// These two assertions don't appear to be passing my custom equality check. :barf:
		if box.GetLayoutType() != StackLayoutType {
			t.Error("Expected StackLayout")
		}
		// Width and Height are inferred to zero on request. Clients can ask for StaticWidth and Height
		// for the explicitly configured value.
		assert.Equal(t, box.GetHeight(), 0.0, "GetHeight is derived to zero")
		assert.Equal(t, box.GetWidth(), 0.0, "GetWidth is derived to zero")

		assert.Equal(t, box.GetActualHeight(), -1.0, "ActualHeight")
		assert.Equal(t, box.GetActualWidth(), -1.0, "ActualWidth")
		assert.Equal(t, box.GetFlexHeight(), -1.0, "GetFlexHeight")
		assert.Equal(t, box.GetFlexWidth(), -1.0, "GetFlexWidth")
		assert.Equal(t, box.GetMaxHeight(), -1.0, "GetMaxHeight")
		assert.Equal(t, box.GetMaxWidth(), -1.0, "GetMaxWidth")
		assert.Equal(t, box.GetMinHeight(), -1.0, "GetMinHeight")
		assert.Equal(t, box.GetMinWidth(), -1.0, "GetMinWidth")
		assert.Equal(t, box.GetPadding(), -1.0, "GetPadding")
		assert.Equal(t, box.GetPaddingBottom(), -1.0)
		/*
			assert.Equal(t, box.GetPaddingLeft(), -1.0)
			assert.Equal(t, box.GetPaddingRight(), -1.0)
			assert.Equal(t, box.GetPaddingTop(), -1.0)
			assert.Equal(t, box.GetPrefHeight(), -1.0)
			assert.Equal(t, box.GetPrefWidth(), -1.0)
			assert.Equal(t, box.GetVAlign(), TopAlign)
			assert.Equal(t, box.GetX(), -1.0)
			assert.Equal(t, box.GetY(), -1.0)
			assert.Equal(t, box.GetZ(), -1.0)
			assert.Equal(t, box.GetWidth(), -1.0)
		*/
	})

	t.Run("No Builder", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Id("root"), Children(func(b Builder) {
			Box(b, Id("one"))
			Box(b, Id("two"))
		}))
		if box.GetId() != "root" {
			t.Error("Expected a configured Box component")
		}
	})

	t.Run("Child with no builder should fail", func(t *testing.T) {
		unexpectedReslt, err := Box(nil)

		if unexpectedReslt != nil {
			t.Error("Should not have returned a component with no Builder")
		}
		if err == nil {
			t.Error("Expected an error when no component was provided")
		}
	})

	t.Run("Custom type", func(t *testing.T) {
		fake, _ := Fake(NewBuilder())
		if fake == nil {
			t.Error("Expected builder to return new component")
		}
	})

	t.Run("Padding", func(t *testing.T) {
		sprite, _ := Box(NewBuilder(), Padding(10))

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
		box, _ := Box(NewBuilder(), Padding(10), PaddingLeft(15))
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
		box, _ := Box(NewBuilder(), PaddingLeft(15), Padding(10))

		if box.GetHorizontalPadding() != 25 {
			t.Error("Expected additive HorizontalPadding")
		}
	})

	t.Run("Padding with specifics will clobber a ZERO setting", func(t *testing.T) {
		box, _ := Box(NewBuilder(), PaddingLeft(0), Padding(10))

		// We only look for the "ZERO VALUE" when trying to figure out if we should
		// clobber. But users can set this, so we're a little jammed up here, unless
		// we flag on any/all interrelated value options. :-(
		if box.GetHorizontalPadding() != 20 {
			t.Error("Expected additive HorizontalPadding")
		}
	})

	t.Run("Specific Paddings", func(t *testing.T) {
		box, _ := Box(NewBuilder(), PaddingBottom(1), PaddingRight(2), PaddingLeft(3), PaddingTop(4))

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

	t.Run("NewComponentFactoryFrom", func(t *testing.T) {
		BigVBox := NewComponentFactoryFrom(VBox, MinWidth(200), MinHeight(200))

		instance, _ := BigVBox(NewBuilder())
		if instance.GetMinWidth() != 200 {
			t.Error("Expected default MinWidth")
		}
	})
}
