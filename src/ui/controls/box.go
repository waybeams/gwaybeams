package controls

import (
	"ui"
	"ui/comp"
	"ui/opts"
)

// Box is a base component with a Stack layout.
var Box = comp.Define("Box", comp.New, opts.LayoutType(ui.StackLayoutType))
