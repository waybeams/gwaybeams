package display

var DefaultRectangleRadius = 3.0

func RectangleView(s Surface, d Displayable) error {
	s.BeginPath()
	s.Rect(d.X(), d.Y(), d.Width(), d.Height())
	s.SetFillColor(uint(d.BgColor()))
	s.Fill()

	s.BeginPath()
	s.Rect(d.X()-0.5, d.Y()-0.5, d.Width()+1, d.Height()+1)
	s.SetStrokeWidth(float64(d.StrokeSize()))
	s.SetStrokeColor(uint(d.StrokeColor()))
	s.Stroke()
	return nil
}

func RoundedRectView(s Surface, d Displayable) error {
	// TODO(lbayes): Get the radius from component values.
	s.BeginPath()
	s.RoundedRect(d.X(), d.Y(), d.Width(), d.Height(), DefaultRectangleRadius)
	s.SetFillColor(uint(d.BgColor()))
	s.Fill()

	s.BeginPath()
	s.RoundedRect(d.X()-0.5, d.Y()-0.5, d.Width()+1, d.Height()+1, DefaultRectangleRadius)
	s.SetStrokeWidth(float64(d.StrokeSize()))
	s.SetStrokeColor(uint(d.StrokeColor()))
	s.Stroke()
	return nil
}

func LabelView(s Surface, d Displayable) error {
	err := RectangleView(s, d)
	fontSize := d.FontSize()
	s.SetFontSize(float64(d.FontSize()))
	s.SetFontFace(d.FontFace())
	s.SetFillColor(uint(d.FontColor()))
	s.Text(d.X()+d.PaddingLeft(), d.Y()+d.PaddingTop()+float64(fontSize), d.Text())
	return err
}
