package controls

import (
	"component"
	"opts"
	"ui"
	"views"
)

type LabelComponent struct {
	component.Component
}

func NewLabel() *LabelComponent {
	return &LabelComponent{}
}

// Label is a component with a text title that is rendered over the background.
var Label = component.Define("Label",
	func() ui.Displayable { return NewLabel() },
	opts.LayoutType(ui.NoLayoutType),
	opts.IsFocusable(true),
	opts.IsText(true),
	opts.View(views.LabelView))
