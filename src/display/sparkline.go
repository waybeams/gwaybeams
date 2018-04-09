package display

type SparklineModel struct {
	DisplayPeriod int // Number of entries shown
	Values        []float64
}

// Sparline is a graphical display of data that can be treated like a word.
// https://en.wikipedia.org/wiki/Edward_Tufte#Sparkline
// Simply provide a set of readings, followed by standard component options to
// see a tiny chart rendered.
var Sparkline = func(b Builder, model *SparklineModel, options ...[]ComponentOption) (Displayable, error) {

	return Box(b, TypeName("Sparkline"), View(func(s Surface, d Displayable) error {
		// log.Println("DRAW SPARKLINE with:", model.Values)
		return nil
	}))
}
