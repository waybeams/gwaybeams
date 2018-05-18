package views

import (
	"github.com/waybeams/waybeams/pkg/spec"
)

var DefaultRectangleRadius = 3.0

func RectangleView(s spec.Surface, r spec.Reader) {
	// fmt.Println("Rectangle with:", spec.Path(r), "x:", r.X(), "y:", r.Y(), "w:", r.Width(), "h:", r.Height())
	s.BeginPath()
	s.Rect(r.X(), r.Y(), r.Width(), r.Height())
	s.SetFillColor(r.BgColor())
	s.Fill()

	s.BeginPath()
	s.Rect(r.X()-0.5, r.Y()-0.5, r.Width()+1, r.Height()+1)
	s.SetStrokeWidth(r.StrokeSize())
	s.SetStrokeColor(r.StrokeColor())
	s.Stroke()
}

func RoundedRectView(s spec.Surface, r spec.Reader) {
	// TODO(lbayes): Get the radius from control values.
	s.BeginPath()
	s.RoundedRect(r.X(), r.Y(), r.Width(), r.Height(), DefaultRectangleRadius)
	s.SetFillColor(r.BgColor())
	s.Fill()

	s.BeginPath()
	s.RoundedRect(r.X()-0.5, r.Y()-0.5, r.Width()+1, r.Height()+1, DefaultRectangleRadius)
	s.SetStrokeWidth(r.StrokeSize())
	s.SetStrokeColor(r.StrokeColor())
	s.Stroke()
}

func LabelView(s spec.Surface, r spec.Reader) {
	if r.BgColor() != 0 || r.StrokeColor() != 0 {
		RectangleView(s, r)
	}
	if r.Text() != "" {
		s.SetFontSize(r.FontSize())
		s.SetFontFace(r.FontFace())
		s.SetFillColor(r.FontColor())
		s.Text(r.TextX(), r.TextY(), r.Text())
	}
}
