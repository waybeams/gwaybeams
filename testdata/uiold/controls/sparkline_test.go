package controls

import (
	"surfaces"
	"testing"
	"uiold/context"
)

func TestSparkline(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		values := []float64{2.3, 11.3, 18.4, 13.5, 16.2, 63.1}
		sparkline := Sparkline(context.New(), &SparklineModel{Values: values})

		surface := surface.NewFake()
		sparkline.Draw(surface)
	})
}
