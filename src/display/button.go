package display

import (
	"events"
)

// ApplyOptions will apply the provided options to the received Event target.
func OptionsHandler(options ...ComponentOption) EventHandler {
	return func(e Event) {
		target := e.DisplayTarget()
		for _, option := range options {
			err := option(target)
			if err != nil {
				panic(err)
			}
		}
	}
}

// Button is a stub component pending implementation.
var Button = NewComponentFactory("Button", NewComponent,
	IsFocusable(true),
	Padding(5),
	OnState("active", BgColor(0xce3262ff)),
	OnState("hovered", BgColor(0x00acd7ff)),
	OnState("pressed", BgColor(0x5dc9e2ff)),
	OnState("disabled", BgColor(0xdbd9d6ff)),

	On(events.Entered, OptionsHandler(SetState("hovered"))),
	On(events.Exited, OptionsHandler(SetState("active"))),
	On(events.Pressed, OptionsHandler(SetState("pressed"))),
	On(events.Released, OptionsHandler(SetState("hovered"))),
	Children(func(b Builder, btn Displayable) {
		Label(b, IsFocusable(false), IsText(false), StrokeSize(0), FlexWidth(1), FlexHeight(1), Text(btn.Text()))
	}))
