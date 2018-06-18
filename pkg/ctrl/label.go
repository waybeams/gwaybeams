package ctrl

import (
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/views"
)

type LabelSpec struct {
	spec.Spec

	measuredText     string
	measuredFontSize float64
}

func (l *LabelSpec) Measure(s spec.Surface) {
	x, y, w, h := s.TextBounds(l.FontFace(), l.FontSize(), l.Text())
	l.SetTextX(x)
	l.SetTextY(y)
	l.SetContentWidth(w)
	l.SetContentHeight(h)
}

func Label(options ...spec.Option) *LabelSpec {
	label := &LabelSpec{}
	label.SetSpecName("Label")
	label.SetIsMeasured(true)
	label.SetView(views.LabelView)

	spec.Apply(label, options...)
	return label
}
