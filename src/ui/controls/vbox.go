package controls

import (
	"ui/comp"
	"ui/opts"
	"ui"
)

// VBox is a base component with a vertical flow layout.
var VBox = comp.Define("VBox", comp.New, opts.LayoutType(ui.VerticalFlowLayoutType))
