package controls

import (
	"component"
	"events"
	. "opts"
	. "ui"
)

// Button is a stub component pending implementation.
var Button = component.Define("Button", component.New,
	LayoutType(StackLayoutType),
	IsFocusable(true),
	Padding(10),
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
			IsFocusable(false),
			IsText(false),
			Text(btn.Text()))
	}))
