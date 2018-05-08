package controls

import (
	"ui"
	"ui/control"
	"ui/opts"
)

// Checkbox is a stub control pending implementation.
var Checkbox = control.Define("Checkbox", control.New,
	opts.IsFocusable(true),
	opts.LayoutType(ui.HorizontalFlowLayoutType))
