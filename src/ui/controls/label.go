package controls

import (
	"events"
	"math"
	"ui"
	"ui/control"
	"ui/opts"
	"views"
)

const DefaultTextForHeight = "Q"

func CreateLabelMeasureHandler(propValue func(d ui.Displayable) string) events.EventHandler {
	// Cache the last reading for text size
	var measuredText string
	var measuredSize int
	var minHeight float64

	return func(e events.Event) {
		label := e.Target().(ui.Displayable)
		face := label.FontFace()
		currentText := propValue(label)
		currentSize := label.FontSize()

		if currentText == "" {
			// TODO(lbayes): Instead of doing this ridiculous hackery, go get FontMetrics and
			// add ascender and descender sizes to whatever dimensions we get?
			currentText = DefaultTextForHeight
		}

		shouldUpdate := face != "" &&
			(measuredText != currentText || measuredSize != currentSize)

		if shouldUpdate {
			font := label.Context().Font(face)
			if font != nil {
				measuredText = currentText
				measuredSize = label.FontSize()

				// Update the Font Atlas with the current/updated size.
				font.SetSize(float32(measuredSize))
				w, bounds := font.Bounds(measuredText)
				h := float64(bounds[3]-bounds[1]) + label.VerticalPadding()

				if currentText == DefaultTextForHeight && minHeight == 0.0 {
					// First pass, we're using "Q" for measure, make the
					// at least that tall.
					minHeight = h
				}

				h = math.Max(h, float64(minHeight))
				label.SetTextX(float64(bounds[0]))
				label.SetTextY(float64(bounds[1]))

				label.SetMinHeight(h)
				label.SetMinWidth(float64(w) + label.HorizontalPadding())
			}
		}
	}
}

// Label is a control with a text title that is rendered over the background.
var Label = control.Define("Label",
	control.New,
	opts.OnConfigured(CreateLabelMeasureHandler(func(d ui.Displayable) string {
		return d.Text()
	})),
	opts.LayoutType(ui.NoLayoutType),
	opts.IsFocusable(true),
	opts.IsText(true),
	opts.View(views.LabelView))
