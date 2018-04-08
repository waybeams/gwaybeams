package display

func RectangleView(s Surface, d Displayable) error {
	s.SetFillColor(uint(d.BgColor()))
	s.SetStrokeWidth(float64(d.StrokeSize()))
	s.SetStrokeColor(uint(d.StrokeColor()))
	s.DrawRectangle(d.X(), d.Y(), d.Width(), d.Height())
	s.Fill()
	s.Stroke()
	return nil
}

func LabelView(s Surface, d Displayable) error {
	RectangleView(s, d)
	fontSize := d.FontSize()
	s.SetFontSize(float64(d.FontSize()))
	s.SetFontFace(d.FontFace())
	s.SetFillColor(uint(d.FontColor()))
	s.Text(d.X()+d.PaddingLeft(), d.Y()+d.PaddingTop()+float64(fontSize), d.Text())
	return nil
}
