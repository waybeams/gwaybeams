package controls

import (
	"component"
	"opts"
	"ui"
)

// Checkbox is a stub component pending implementation.
var Checkbox = component.Define("Checkbox", component.New,
	opts.IsFocusable(true),
	opts.LayoutType(ui.HorizontalFlowLayoutType))
