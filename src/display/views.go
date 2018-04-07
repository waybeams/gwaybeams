package display

func RectangleView(s Surface, d Displayable) error {
	s.SetFillColor(uint(d.GetBgColor()))
	s.SetStrokeWidth(float64(d.GetStrokeSize()))
	s.SetStrokeColor(uint(d.GetStrokeColor()))
	s.DrawRectangle(d.GetX(), d.GetY(), d.GetWidth(), d.GetHeight())
	s.Fill()
	s.Stroke()
	return nil
}

func LabelView(s Surface, d Displayable) error {
	RectangleView(s, d)
	fontSize := d.GetFontSize()
	s.SetFontSize(float64(d.GetFontSize()))
	s.SetFontFace(d.GetFontFace())
	s.SetFillColor(uint(d.GetFontColor()))
	s.Text(d.GetX()+d.GetPaddingLeft(), d.GetY()+d.GetPaddingTop()+float64(fontSize), d.GetText())
	return nil
}
