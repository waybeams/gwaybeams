package controls

import (
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/views"
)

type LabelSpec struct {
	spec.Spec

	measuredText     string
	measuredFontSize float64
	// ascender         float32
	// descender        float32
	// lineHeight       float32
}

func (l *LabelSpec) Measure(s spec.Surface) {
	face := l.FontFace()
	currentText := l.Text()
	currentSize := l.FontSize()

	shouldUpdate := face != "" &&
		(l.measuredText != currentText || l.measuredFontSize != currentSize)

	// Don't do work if it's not necessary.
	if shouldUpdate {
		font := s.Font(face)
		if font != nil {
			l.measuredText = currentText
			l.measuredFontSize = l.FontSize()

			// Update the Font Atlas with the current/updated size.
			font.SetSize(float32(l.measuredFontSize))

			_, _, lineH := font.VerticalMetrics()
			w, bounds := font.Bounds(l.measuredText)
			h := float64(lineH)

			// fmt.Println("LABEL Text:", currentText, asc, desc, lineH, "TextY?", bounds[1])
			// fmt.Println("BOUNDS:", bounds)
			l.SetTextX(float64(bounds[0]))
			l.SetTextY(float64(bounds[1]))

			l.SetMinWidth(float64(w) + l.HorizontalPadding())

			minHeight := l.MinHeight()
			newMinHeight := h + l.VerticalPadding()
			if newMinHeight > minHeight {
				l.SetMinHeight(newMinHeight)
			}
		}
	}
}

func Label(options ...spec.Option) *LabelSpec {
	defaults := []spec.Option{
		opts.SpecName("Label"),
		opts.IsMeasured(true),
		opts.View(views.LabelView),
	}
	label := &LabelSpec{}
	spec.ApplyAll(label, defaults, options)
	return label
}
