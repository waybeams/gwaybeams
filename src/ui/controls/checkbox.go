package controls

import (
	"ui/comp"
	"ui/opts"
	"ui"
)

// Checkbox is a stub component pending implementation.
var Checkbox = comp.Define("Checkbox", comp.New,
	opts.IsFocusable(true),
	opts.LayoutType(ui.HorizontalFlowLayoutType))
