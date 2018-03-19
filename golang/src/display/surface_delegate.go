package display

type SurfaceDelegate struct {
	DelegateTo Surface
}

func (s *SurfaceDelegate) MoveTo(x float64, y float64) {
	s.DelegateTo.MoveTo(x, y)
}

func (s *SurfaceDelegate) SetRgba(r, g, b, a float64) {
	s.DelegateTo.SetRgba(r, g, b, a)
}

func (s *SurfaceDelegate) SetLineWidth(width float64) {
	s.DelegateTo.SetLineWidth(width)
}

func (s *SurfaceDelegate) Stroke() {
	s.DelegateTo.Stroke()
}

func (s *SurfaceDelegate) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	s.DelegateTo.Arc(xc, yc, radius, angle1, angle2)
}

func (s *SurfaceDelegate) DrawRectangle(x float64, y float64, width float64, height float64) {
	s.DelegateTo.DrawRectangle(x, y, width, height)
}

func (s *SurfaceDelegate) Fill() {
	s.DelegateTo.Fill()
}

func (s *SurfaceDelegate) FillPreserve() {
	s.DelegateTo.FillPreserve()
}

func NewSurfaceDelegate(delegateTo Surface) *SurfaceDelegate {
	return &SurfaceDelegate{DelegateTo: delegateTo}
}
