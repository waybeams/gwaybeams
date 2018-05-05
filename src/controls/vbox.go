package controls

import (
	"component"
	"opts"
	"ui"
)

// VBox is a base component with a vertical flow layout.
var VBox = component.Define("VBox", component.New, opts.LayoutType(ui.VerticalFlowLayoutType))
