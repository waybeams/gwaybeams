package controls

import (
	"ui"
	"ui/control"
	"uiold/opts"
)

// VBox is a base control with a vertical flow layout.
var VBox = control.Define("VBox", control.New, opts.LayoutType(ui.VerticalFlowLayoutType))
