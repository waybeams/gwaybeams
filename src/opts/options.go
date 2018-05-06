package opts

import (
	"events"
	"ui"
)

// ActualWidth will set Component.ActualWidth.
func ActualWidth(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetActualWidth(value)
	}
}

// ActualHeight will set Component.ActualHeight.
func ActualHeight(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetActualHeight(value)
	}
}

func BgColor(color int) ui.Option {
	return func(d ui.Displayable) {
		d.SetBgColor(color)
	}
}

func Blurred() ui.Option {
	return func(d ui.Displayable) {
		d.Blur()
	}
}

// Children will compose child components onto the current component. The composer
// type must be a function with a signature that matches one of the following:
//   A) func()
//   B) func(b Builder)
//   C) func(d Displayable)
//   D) func(b Builder, d Displayable)
// The outermost Children function usually should receive a builder instance that
// all children will receive and isolated Component definitions generally require
// both arguments to the outer composer.
func Children(composer interface{}) ui.Option {
	return func(d ui.Displayable) {
		d.SetComposer(composer)
	}
}

func Data(data interface{}) ui.Option {
	return func(d ui.Displayable) {
		d.SetData(data)
	}
}

// ExcludeFromLayout will configure Component.ExcludeFromLayout.
func ExcludeFromLayout(value bool) ui.Option {
	return func(d ui.Displayable) {
		d.SetExcludeFromLayout(value)
	}
}

// FlexHeight will set Component.FlexHeight.
func FlexHeight(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetFlexHeight(value)
	}
}

// FlexWidth will set Component.FlexWidth.
func FlexWidth(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetFlexWidth(value)
	}
}

func Focused() ui.Option {
	return func(d ui.Displayable) {
		d.Focus()
	}
}

func FontColor(color int) ui.Option {
	return func(d ui.Displayable) {
		d.SetFontColor(color)
	}
}

func FontFace(face string) ui.Option {
	return func(d ui.Displayable) {
		d.SetFontFace(face)
	}
}

func FontSize(size int) ui.Option {
	return func(d ui.Displayable) {
		d.SetFontSize(size)
	}
}

func Gutter(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetGutter(value)
	}
}

// ID will set the Component.Id.
func ID(value string) ui.Option {
	return func(d ui.Displayable) {
		d.Model().ID = value
	}
}

func IsFocusable(value bool) ui.Option {
	return func(d ui.Displayable) {
		d.SetIsFocusable(value)
	}
}

func IsText(value bool) ui.Option {
	return func(d ui.Displayable) {
		d.SetIsText(value)
	}
}

func IsTextInput(value bool) ui.Option {
	return func(d ui.Displayable) {
		d.SetIsTextInput(value)
	}
}

// HAlign will set Component.HAlign.
func HAlign(align ui.Alignment) ui.Option {
	return func(d ui.Displayable) {
		d.SetHAlign(align)
	}
}

// Height will set Component.Height.
func Height(value float64) ui.Option {
	return func(d ui.Displayable) {
		// TODO(lbayes): Should use accessor!
		d.Model().Height = value
	}
}

func Key(value string) ui.Option {
	return func(d ui.Displayable) {
		d.Model().Key = value
	}
}

// LayoutType will set Component.LayoutType.
func LayoutType(layoutType ui.LayoutTypeValue) ui.Option {
	return func(d ui.Displayable) {
		d.SetLayoutType(layoutType)
	}
}

// MaxHeight will set Component.MaxHeight.
func MaxHeight(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetMaxHeight(value)
	}
}

// MaxWidth will set Component.MaxWidth.
func MaxWidth(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetMaxWidth(value)
	}
}

// MinHeight will set Component.MinHeight.
func MinHeight(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetMinHeight(value)
	}
}

// MinWidth will set Component.MinWidth.
func MinWidth(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetMinWidth(value)
	}
}

// Padding will set Component.Padding, which will effectively set padding for
// all four sides as well (bottom, top, left, right, horizontal and vertical).
func Padding(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetPadding(value)
	}
}

// PaddingBottom will set Component.PaddingBottom.
func PaddingBottom(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetPaddingBottom(value)
	}
}

// PaddingLeft will set Component.PaddingLeft.
func PaddingLeft(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetPaddingLeft(value)
	}
}

// PaddingRight will set Component.PaddingRight.
func PaddingRight(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetPaddingRight(value)
	}
}

// PaddingTop will set Component.PaddingTop.
func PaddingTop(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetPaddingTop(value)
	}
}

// PrefHeight will set Component.PrefHeight.
func PrefHeight(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetPrefHeight(value)
	}
}

// PrefWidth will set Component.PrefWidth.
func PrefWidth(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetPrefWidth(value)
	}
}

// Size will set Component.Width and Component.Height.
func Size(width, height float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetWidth(width)
		d.SetHeight(height)
	}
}

func StrokeColor(color int) ui.Option {
	return func(d ui.Displayable) {
		d.SetStrokeColor(color)
	}
}

func StrokeSize(size int) ui.Option {
	return func(d ui.Displayable) {
		d.SetStrokeSize(size)
	}
}

func Text(value string) ui.Option {
	return func(d ui.Displayable) {
		// TODO(lbayes): Sanitize text as user input values can be placed in here.
		// TODO(lbayes): Localize text using Localization map.
		d.SetText(value)
	}
}

// Title will set Component.Title.
func Title(value string) ui.Option {
	return func(d ui.Displayable) {
		d.SetTitle(value)
	}
}

func TraitNames(names ...string) ui.Option {
	return func(d ui.Displayable) {
		d.SetTraitNames(names...)
	}
}

// VAlign will set Component.VAlign.
func VAlign(align ui.Alignment) ui.Option {
	return func(d ui.Displayable) {
		d.SetVAlign(align)
	}
}

func View(view ui.RenderHandler) ui.Option {
	return func(d ui.Displayable) {
		d.SetView(view)
	}
}

func Visible(visible bool) ui.Option {
	return func(d ui.Displayable) {
		d.SetVisible(visible)
	}
}

// Width will set Component.Width.
func Width(value float64) ui.Option {
	return func(d ui.Displayable) {
		d.Model().Width = value
	}
}

// X will set Component.X.
func X(pos float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetX(pos)
	}
}

// Y will set Component.Y.
func Y(pos float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetY(pos)
	}
}

// Z will set Component.Z.
func Z(pos float64) ui.Option {
	return func(d ui.Displayable) {
		d.SetZ(pos)
	}
}

//-------------------------------------------
// Special Adapters
//-------------------------------------------

// A Bag is a collection of Options that is itself an Option.
func Bag(opts ...ui.Option) ui.Option {
	return func(d ui.Displayable) {
		for _, opt := range opts {
			opt(d)
		}
	}
}

// OptionsHandler will apply the provided options to the received Event target.
func OptionsHandler(options ...ui.Option) events.EventHandler {
	return func(e events.Event) {
		target := e.Target().(ui.Displayable)
		for _, option := range options {
			option(target)
		}
	}
}

func TestOptions() ui.Option {
	return Bag(
		FontFace("Roboto"),
		FontSize(12),
		BgColor(0xffcc00ff),
	)
}

//-------------------------------------------
// Event Helpers
//-------------------------------------------

// On will apply the provided handler to the provided event name.
func On(eventName string, handler events.EventHandler) ui.Option {
	return func(d ui.Displayable) {
		d.PushUnsub(d.On(eventName, handler))
	}
}

func OnClick(handler events.EventHandler) ui.Option {
	return func(d ui.Displayable) {
		d.PushUnsub(d.On(events.Clicked, handler))
	}
}

func OnEnter(handler events.EventHandler) ui.Option {
	return func(d ui.Displayable) {
		d.PushUnsub(d.On(events.Entered, handler))
	}
}

func OnFrameEntered(handler events.EventHandler) ui.Option {
	return func(d ui.Displayable) {
		d.PushUnsub(d.Context().OnFrameEntered(handler))
	}
}

//-------------------------------------------
// State Helpers
//-------------------------------------------

// TODO(lbayes): Consider introducing AppendState and ReplaceState
func OnState(name string, options ...ui.Option) ui.Option {
	return func(d ui.Displayable) {
		d.OnState(name, options...)
	}
}

func SetState(name string) ui.Option {
	return func(d ui.Displayable) {
		d.SetState(name)
	}
}
