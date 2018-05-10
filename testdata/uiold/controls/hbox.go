package controls

import (
	"ui"
	"ui/control"
	"uiold/opts"
)

// HBox is a base control with a horizontal flow layout.
var HBox = control.Define("HBox", control.New, opts.LayoutType(ui.HorizontalFlowLayoutType))
