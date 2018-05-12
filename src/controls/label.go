package controls

import (
	"math"
	"opts"
	"spec"
)

const DefaultTextForHeight = "Q"

type LabelSpec struct {
	spec.Spec

	measuredText string
	measuredSize float64
}

func (l *LabelSpec) Measure(s spec.Surface) {
	face := l.FontFace()
	currentText := l.Text()
	currentSize := l.FontSize()

	if currentText == "" {
		// TODO(lbayes): Instead of doing this ridiculous hackery, go get FontMetrics and
		// add ascender and descender sizes to whatever dimensions we get?
		currentText = DefaultTextForHeight
	}

	shouldUpdate := face != "" &&
		(l.measuredText != currentText || l.measuredSize != currentSize)

	// Don't do work if it's not necessary.
	if shouldUpdate {
		minHeight := l.MinHeight()
		font := s.Font(face)
		if font != nil {
			l.measuredText = currentText
			l.measuredSize = l.FontSize()

			// Update the Font Atlas with the current/updated size.
			font.SetSize(float32(l.measuredSize))
			w, bounds := font.Bounds(l.measuredText)
			h := float64(bounds[3]-bounds[1]) + l.VerticalPadding()

			// if currentText == DefaultTextForHeight && minHeight == 0 {
			// First pass, we're using "Q" for measure, make the label
			// at least that tall.
			// l.SetMinHeight(h)
			// }

			h = math.Max(h, minHeight)
			l.SetTextX(float64(bounds[0]))
			l.SetTextY(float64(bounds[1]))

			l.SetMinHeight(h)
			l.SetMinWidth(float64(w) + l.HorizontalPadding())
		}
	}
}

func Label(options ...spec.Option) *LabelSpec {
	defaults := []spec.Option{
		opts.IsMeasured(true),
	}
	label := &LabelSpec{}
	spec.ApplyAll(label, defaults, options)
	return label
}
