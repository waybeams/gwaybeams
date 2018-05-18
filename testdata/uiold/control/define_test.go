package control

import (
	"github.com/waybeams/assert"
	"testing"
	. "ui"
	"uiold/context"
	"ui/control"
	"ui/controls"
	"uiold/opts"
)

// fooStruct is a control definition.
type fooStruct struct {
	control.Control
}

// newFoo creates and returns an instance of fooStruct.
func newFoo() *fooStruct {
	return &fooStruct{}
}

// foo is a control definition. Sadly, we must wrap concrete constructors
// with a function that returns a Displayable :barf:
var foo = control.Define("foo",
	func() Displayable { return newFoo() })

func TestControlFactory(t *testing.T) {
	t.Run("Default State", func(t *testing.T) {
		box := controls.Box(context.New())
		// These two assertions don't appear to be passing my custom equality check. :barf:
		if box.HAlign() != AlignLeft {
			t.Error("Expected AlignLeft, but got: %v", box.HAlign())
		}
		// These two assertions don't appear to be passing my custom equality check. :barf:
		if box.LayoutType() != StackLayoutType {
			t.Error("Expected StackLayout")
		}
		// Width and Height are inferred to zero on request. Clients can ask for StaticWidth and Height
		// for the explicitly configured value.
		assert.Equal(box.Height(), 0.0, "GetHeight is derived to zero")
		assert.Equal(box.Width(), 0.0, "GetWidth is derived to zero")

		assert.Equal(box.ActualHeight(), -1.0, "ActualHeight")
		assert.Equal(box.ActualWidth(), -1.0, "ActualWidth")
		assert.Equal(box.FlexHeight(), -1.0, "GetFlexHeight")
		assert.Equal(box.FlexWidth(), -1.0, "GetFlexWidth")
		assert.Equal(box.MaxHeight(), -1.0, "GetMaxHeight")
		assert.Equal(box.MaxWidth(), -1.0, "GetMaxWidth")
		assert.Equal(box.MinHeight(), -1.0, "GetMinHeight")
		assert.Equal(box.MinWidth(), -1.0, "GetMinWidth")
		assert.Equal(box.Padding(), -1.0, "GetPadding")
		assert.Equal(box.PaddingBottom(), -1.0)
		/*
			assert.Equal(box.GetPaddingLeft(), -1.0)
			assert.Equal(box.GetPaddingRight(), -1.0)
			assert.Equal(box.GetPaddingTop(), -1.0)
			assert.Equal(box.GetPrefHeight(), -1.0)
			assert.Equal(box.GetPrefWidth(), -1.0)
			assert.Equal(box.GetVAlign(), AlignTop)
			assert.Equal(box.GetX(), -1.0)
			assert.Equal(box.GetY(), -1.0)
			assert.Equal(box.GetZ(), -1.0)
			assert.Equal(box.GetWidth(), -1.0)
		*/
	})

	t.Run("No Builder", func(t *testing.T) {
		box := controls.Box(context.New(), opts.ID("root"), opts.Children(func(c Context) {
			controls.Box(c, opts.ID("one"))
			controls.Box(c, opts.ID("two"))
		}))
		if box.ID() != "root" {
			t.Error("Expected a configured Box control")
		}
	})

	t.Run("Custom type", func(t *testing.T) {
		instance := foo(context.New())
		if instance == nil {
			t.Error("Expected builder to return new control")
		}
	})

	t.Run("Padding", func(t *testing.T) {
		sprite := controls.Box(context.New(), opts.Padding(10))

		if sprite.Padding() != 10 {
			t.Error("Expected option to set padding")
		}
		if sprite.HorizontalPadding() != 20 {
			t.Error("Expected Padding to update HorizontalPadding")
		}
		if sprite.VerticalPadding() != 20 {
			t.Error("Expected Padding to update VerticalPadding")
		}
		if sprite.PaddingBottom() != 10 {
			t.Error("Expected Padding to update PaddingBottom")
		}
		if sprite.PaddingLeft() != 10 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if sprite.PaddingRight() != 10 {
			t.Error("Expected Padding to update PaddingRight")
		}
		if sprite.PaddingTop() != 10 {
			t.Error("Expected Padding to update PaddingTop")
		}
	})

	t.Run("Padding with specifics", func(t *testing.T) {
		box := controls.Box(context.New(), opts.Padding(10), opts.PaddingLeft(15))
		if box.VerticalPadding() != 20 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.HorizontalPadding() != 25 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.PaddingLeft() != 15 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if box.PaddingRight() != 10 {
			t.Error("Expected Padding to update PaddingRight")
		}
	})

	t.Run("Padding with specifics is NOT order dependent", func(t *testing.T) {
		box := controls.Box(context.New(), opts.PaddingLeft(15), opts.Padding(10))

		if box.HorizontalPadding() != 25 {
			t.Error("Expected additive HorizontalPadding")
		}
	})

	t.Run("Padding with specifics will NOT clobber a ZERO setting", func(t *testing.T) {
		box := controls.Box(context.New(), opts.PaddingLeft(0), opts.Padding(10))

		if box.PaddingLeft() != 0 {
			t.Error("Padding option should not clobber a previously set value of Zero")
		}

		if box.HorizontalPadding() != 10 {
			t.Error("Expected zero value padding left to be respected")
		}

		if box.VerticalPadding() != 20 {
			t.Error("Padding should apply to both axis")
		}
	})

	t.Run("Specific Paddings", func(t *testing.T) {
		box := controls.Box(context.New(), opts.PaddingBottom(1), opts.PaddingRight(2), opts.PaddingLeft(3), opts.PaddingTop(4))

		if box.VerticalPadding() != 5 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.HorizontalPadding() != 5 {
			t.Error("Expected additive HorizontalPadding")
		}
		if box.PaddingLeft() != 3 {
			t.Error("Expected Padding to update PaddingLeft")
		}
		if box.PaddingRight() != 2 {
			t.Error("Expected Padding to update PaddingRight")
		}
		if box.PaddingTop() != 4 {
			t.Error("Expected Padding to update PaddingTop")
		}
		if box.PaddingBottom() != 1 {
			t.Error("Expected Padding to update PaddingBottom")
		}
	})
}
