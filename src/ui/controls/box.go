package controls

import (
	"ui"
	"ui/control"
	"ui/opts"
)

// Box is a base control with a Stack layout.
var Box = control.Define("Box", control.New, opts.LayoutType(ui.StackLayoutType))
