package display

import (
	"events"
	"fmt"
)

// Button is a stub component pending implementation.
var Button = NewComponentFactory("Button", NewComponent, IsFocusable(true), Children(func(b Builder) {
	var enterHandler = func(e Event) {
		fmt.Println("ENTERED:", e.Target().(Displayable).Path())
		e.Target().(Displayable).SetState("hovered")
	}
	var exitedHandler = func(e Event) {
		fmt.Println("EXITED:", e.Target().(Displayable).Path())
		e.Target().(Displayable).SetState("active")
	}
	Box(b, FlexWidth(1), FlexHeight(1),
		On(events.Entered, enterHandler),
		On(events.Exited, exitedHandler),
		AddState("active", BgColor(0xcececeff)),
		AddState("hovered", BgColor(0xffcc00ff)),
		AddState("pressed", BgColor(0xff0000ff)),
		AddState("disabled", BgColor(0xccccccff)),
	)
}))
