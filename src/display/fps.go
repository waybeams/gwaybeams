package display

import "fmt"

func NewFpsView() RenderHandler {
	message := fmt.Sprintf("%v fps")

	return func(s Surface, d Displayable) error {
		RectangleView(s, d)
		fontSize := d.FontSize()
		s.SetFontSize(float64(d.FontSize()))
		s.SetFontFace(d.FontFace())
		s.SetFillColor(uint(d.FontColor()))
		s.Text(d.X()+d.PaddingLeft(), d.Y()+d.PaddingTop()+float64(fontSize), message)
		return nil
	}
}

func averageValues(values []float32) float32 {
	count := float32(len(values))
	if count == 0.0 {
		return 0.0
	}

	sum := float32(0.0)
	for _, value := range values {
		sum += value
	}
	return sum / count
}

var fpsFactory = NewComponentFactory(
	"FPS",
	NewComponent,
	FontColor(0x333333ff),
	FontSize(18),
	Padding(5),
)

func FPS(b Builder, instanceOpts ...ComponentOption) (Displayable, error) {
	readings := []float32{1.2, 14.3, 3.4, 6.7}

	averageReadings := averageValues(readings)
	label := fmt.Sprintf("%v fps", averageReadings)

	var fpsView = func(s Surface, d Displayable) error {
		RectangleView(s, d)
		fontSize := d.FontSize()
		s.SetFontSize(float64(d.FontSize()))
		s.SetFontFace(d.FontFace())
		s.SetFillColor(uint(d.FontColor()))
		s.Text(d.X()+d.PaddingLeft(), d.Y()+d.PaddingTop()+float64(fontSize), label)
		return nil
	}

	instanceOpts = append(instanceOpts, Children(func(b Builder, d Displayable) {
		Box(b,
			BgColor(0x33ff33ff),
			Height(60),
			Width(150),
			View(fpsView),
		)
	}))

	instance, err := fpsFactory(b, instanceOpts...)

	if err != nil {
		return nil, err
	}
	return instance, nil
}
