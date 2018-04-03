package display

type LabelComponent struct {
	Component
}

func (l *LabelComponent) Draw(surface Surface) {
	DrawRectangle(surface, l)

	fontSize := l.GetFontSize()
	surface.SetFontSize(float64(l.GetFontSize()))
	surface.SetFontFace(l.GetFontFace())
	// TODO(lbayes): Wire up font color!
	surface.SetFillColor(uint(l.GetFontColor()))
	surface.Text(l.GetX()+l.GetPaddingLeft(), l.GetY()+l.GetPaddingTop()+float64(fontSize), l.GetText())
}

func NewLabelComponent() Displayable {
	return &LabelComponent{}
}

var Label = NewComponentFactory(NewLabelComponent)
