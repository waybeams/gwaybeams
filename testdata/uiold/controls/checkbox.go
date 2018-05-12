package controls

import (
	"ui"
	"ui/control"
	"uiold/opts"
)

// Checkbox is a stub control pending implementation.
var Checkbox = control.Define("Checkbox", control.New,
	opts.IsFocusable(true),
	opts.LayoutType(ui.HorizontalFlowLayoutType))
