package controls

import (
	. "ui/opts"
	. "ui"
)

type SparklineModel struct {
	DisplayPeriod int // Number of entries shown
	Values        []float64
}

// Sparkline is a graphical display of data that can be treated like a word.
// https://en.wikipedia.org/wiki/Edward_Tufte#Sparkline
// Simply provide a set of readings, followed by standard component options to
// see a tiny chart rendered.
var Sparkline = func(c Context, model *SparklineModel, options ...[]Option) Displayable {
	return Box(c, TraitNames("Sparkline"), View(func(s Surface, d Displayable) error {
		// log.Println("DRAW SPARKLINE with:", model.Values)
		return nil
	}))
}
