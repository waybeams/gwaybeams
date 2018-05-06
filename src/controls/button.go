package controls

import (
	"component"
	"events"
	. "opts"
	. "ui"
)

// ApplyOptions will apply the provided options to the received Event target.
func OptionsHandler(options ...Option) events.EventHandler {
	return func(e events.Event) {
		target := e.Target().(Displayable)
		for _, option := range options {
			err := option(target)
			if err != nil {
				panic(err)
			}
		}
	}
}

// Button is a stub component pending implementation.
var Button = component.Define("Button", component.New,
	LayoutType(StackLayoutType),
	IsFocusable(true),
	PaddingBottom(10),
	PaddingLeft(10),
	PaddingRight(10),
	PaddingTop(5),
	OnState("active", BgColor(0xce3262ff)),
	OnState("hovered", BgColor(0x00acd7ff)),
	OnState("pressed", BgColor(0x5dc9e2ff)),
	OnState("disabled", BgColor(0xdbd9d6ff)),

	On(events.Entered, OptionsHandler(SetState("hovered"))),
	On(events.Exited, OptionsHandler(SetState("active"))),
	On(events.Pressed, OptionsHandler(SetState("pressed"))),
	On(events.Released, OptionsHandler(SetState("hovered"))),
	Children(func(c Context, btn Displayable) {
		Label(c,
			X(10),
			Y(0),
			IsFocusable(false),
			IsText(false),
			StrokeSize(0),
			Text(btn.Text()))
	}))
