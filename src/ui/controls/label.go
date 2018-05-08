package controls

import (
	"ui"
	"ui/control"
	"ui/opts"
	"views"
)

type LabelControl struct {
	control.Control

	measuredText string
	measuredSize int
}

func (l *LabelControl) SetFontSize(size int) {
	l.measuredSize = 0
	l.Control.SetFontSize(size)
	l.Measure()
}

// Layout the Label by first measuring Text and configuring our min dimensions.
func (l *LabelControl) SetText(currentText string) {
	l.Control.SetText(currentText)
	l.Measure()
}

func (l *LabelControl) Measure() {
	face := l.FontFace()
	currentText := l.Text()
	currentSize := l.FontSize()

	shouldUpdate := face != "" &&
		(l.measuredText != currentText || l.measuredSize != currentSize)

	if shouldUpdate {
		font := l.Context().Font(face)
		if font != nil {
			l.measuredText = currentText
			l.measuredSize = l.FontSize()

			// Update the Font Atlas with the current/updated size.
			font.SetSize(float32(l.measuredSize))
			w, bounds := font.Bounds(l.measuredText)
			h := bounds[3] - bounds[1]

			l.SetTextX(float64(bounds[0]))
			l.SetTextY(float64(bounds[1]))

			l.SetMinHeight(float64(h) + l.VerticalPadding())
			l.SetMinWidth(float64(w) + l.HorizontalPadding())
		}
	}
}

func NewLabel() *LabelControl {
	return &LabelControl{}
}

// Label is a control with a text title that is rendered over the background.
var Label = control.Define("Label",
	func() ui.Displayable { return NewLabel() },
	opts.LayoutType(ui.NoLayoutType),
	opts.IsFocusable(true),
	opts.IsText(true),
	opts.View(views.LabelView))
