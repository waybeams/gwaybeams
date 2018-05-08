package controls

import (
	"ui/comp"
	"ui/opts"
	"ui"
)

// HBox is a base component with a horizontal flow layout.
var HBox = comp.Define("HBox", comp.New, opts.LayoutType(ui.HorizontalFlowLayoutType))
