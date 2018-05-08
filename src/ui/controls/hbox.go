package controls

import (
	"ui"
	"ui/control"
	"ui/opts"
)

// HBox is a base control with a horizontal flow layout.
var HBox = control.Define("HBox", control.New, opts.LayoutType(ui.HorizontalFlowLayoutType))
