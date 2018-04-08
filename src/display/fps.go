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

var FPS = NewComponentFactory(
	"FPS",
	NewComponent,
	FontColor(0x333333ff),
	FontSize(18),
	Padding(5),
	Children(func(b Builder, d Displayable) {
		Box(b,
			BgColor(0x33ff33ff),
			Height(60),
			Width(150),
			View(NewFpsView()))

		// Set timeout and invalidate the outer component?
	}))
