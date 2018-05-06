package controls

import (
	"events"
	"fmt"
	"opts"
	. "ui"
)

func avgFloats(values []float64) float64 {
	count := float64(len(values))
	if count == 0.0 {
		return 0.0
	}

	sum := float64(0.0)
	for _, value := range values {
		sum += value
	}
	return sum / count
}

// FPS will render a collection of float32 readings by displaying a
func FPS(c Context, instanceOpts ...Option) Displayable {
	readings := []float64{}
	avgFps := avgFloats(readings)
	message := fmt.Sprintf("%v fps", avgFps)

	enterFrameHandler := func(e events.Event) {}

	return Label(c,
		opts.Text(message),
		opts.OnFrameEntered(enterFrameHandler),
		opts.LayoutType(StackLayoutType),
		opts.Children(func(c Context) {
			Sparkline(c, &SparklineModel{Values: readings})
		}))
}
