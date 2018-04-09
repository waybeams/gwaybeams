package display

import "fmt"

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
func FPS(b Builder, instanceOpts ...ComponentOption) (Displayable, error) {
	readings := []float64{}
	avgFps := avgFloats(readings)
	message := fmt.Sprintf("%v fps", avgFps)

	enterFrameHandler := func(d Displayable) {
		fmt.Println("ON ENTER FRAME!")
	}

	return Label(b,
		Text(message),
		OnEnterFrame(enterFrameHandler),
		LayoutType(StackLayoutType),
		Children(func(b Builder) {
			Sparkline(b, &SparklineModel{Values: readings})
		}))
}
