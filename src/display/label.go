package display

type LabelComponent struct {
	Component
}

func (l *LabelComponent) Draw(surface Surface) {
	surface.SetFillColor(0xffcc00ff)
	surface.SetStrokeColor(0x333333ff)
	surface.DrawRectangle(l.GetX(), l.GetY(), l.GetWidth(), l.GetHeight())
	surface.Fill()
	surface.Stroke()

	fontSize := 64.0
	surface.SetFontSize(fontSize)
	surface.SetFontFace("sans")
	surface.SetFillColor(0x333333ff)
	surface.Text(l.GetX()+l.GetPaddingLeft(), l.GetY()+l.GetPaddingTop()+fontSize, l.GetText())
}

func NewLabelComponent() Displayable {
	return &LabelComponent{}
}

var Label = NewComponentFactory(NewLabelComponent)
