package controls

import (
	"component"
	"opts"
	"ui"
)

// HBox is a base component with a horizontal flow layout.
var HBox = component.Define("HBox", component.New, opts.LayoutType(ui.HorizontalFlowLayoutType))
