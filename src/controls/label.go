package controls

import (
	"component"
	"opts"
	"ui"
	"views"
)

// Label is a component with a text title that is rendered over the background.
var Label = component.Define("Label", component.New,
	opts.LayoutType(ui.NoLayoutType),
	opts.IsFocusable(true),
	opts.IsText(true),
	opts.View(views.LabelView))
