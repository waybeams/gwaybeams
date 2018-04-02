package display

import ()

type SurfaceDelegate struct {
	delegateTo Surface
	offsetX    float64
	offsetY    float64
}

func (s *SurfaceDelegate) MoveTo(x float64, y float64) {
	s.delegateTo.MoveTo(x+s.offsetX, y+s.offsetY)
}

func (s *SurfaceDelegate) SetFillColor(color uint) {
	s.delegateTo.SetFillColor(color)
}

func (s *SurfaceDelegate) SetStrokeColor(color uint) {
	s.delegateTo.SetStrokeColor(color)
}

func (s *SurfaceDelegate) SetStrokeWidth(width float64) {
	s.delegateTo.SetStrokeWidth(width)
}

func (s *SurfaceDelegate) Stroke() {
	s.delegateTo.Stroke()
}

func (s *SurfaceDelegate) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	xc += s.offsetX
	yc += s.offsetY
	s.delegateTo.Arc(xc, yc, radius, angle1, angle2)
}

func (s *SurfaceDelegate) DrawRectangle(x float64, y float64, width float64, height float64) {
	x += s.offsetX
	y += s.offsetY
	s.delegateTo.DrawRectangle(x, y, width, height)
}

func (s *SurfaceDelegate) Fill() {
	s.delegateTo.Fill()
}

func (s *SurfaceDelegate) GetOffsetSurfaceFor(d Displayable) Surface {
	return NewSurfaceDelegateFor(d, s)
}

func NewSurfaceDelegateFor(d Displayable, delegateTo Surface) Surface {
	return &SurfaceDelegate{
		delegateTo: delegateTo,
		offsetX:    d.GetXOffset(),
		offsetY:    d.GetYOffset(),
	}
}
