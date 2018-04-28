package display

import (
	"events"
)

var buttonEnteredHandler = func(e Event) {
	e.DisplayTarget().SetState("hovered")
}

var buttonExitedHandler = func(e Event) {
	e.DisplayTarget().SetState("active")
}

var buttonPressedHandler = func(e Event) {
	e.Target().(Displayable).SetState("pressed")
}

var buttonReleasedHandler = func(e Event) {
	target := e.Target().(Displayable)
	// TODO(lbayes): Only set active if mouse is still over the button
	target.SetState("active")
}

// ApplyOptions will apply the provided options to the received Event target.
func ApplyOptions(options ...ComponentOption) EventHandler {
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

	On(events.Entered, ApplyOptions(SetState("hovered"))),
	On(events.Exited, ApplyOptions(SetState("active"))),
	On(events.Pressed, ApplyOptions(SetState("pressed"))),
	On(events.Released, ApplyOptions(SetState("hovered"))),
	Children(func(b Builder, btn Displayable) {
		Label(b, IsFocusable(false), IsText(false), StrokeSize(0), FlexWidth(1), FlexHeight(1), Text(btn.Text()))
	}))
