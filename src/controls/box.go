package controls

import (
	"component"
	"opts"
	"ui"
)

// Box is a base component with a Stack layout.
var Box = component.Define("Box", component.New, opts.LayoutType(ui.StackLayoutType))
