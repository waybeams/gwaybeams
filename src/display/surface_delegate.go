package display

import "fmt"

type SurfaceDelegate struct {
	delegateTo Surface
	offsetX    float64
	offsetY    float64
}

func (s *SurfaceDelegate) MoveTo(x float64, y float64) {
	s.delegateTo.MoveTo(x+s.offsetX, y+s.offsetY)
}

func (s *SurfaceDelegate) SetRgba(r, g, b, a float64) {
	s.delegateTo.SetRgba(r, g, b, a)
}

func (s *SurfaceDelegate) SetLineWidth(width float64) {
	s.delegateTo.SetLineWidth(width)
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

func (s *SurfaceDelegate) FillPreserve() {
	s.delegateTo.FillPreserve()
}

func (s *SurfaceDelegate) getYOffsetFor(d Displayable) float64 {
	current := d
	offset := 0.0
	for current != nil {
		offset += current.GetY()
		current = d.GetParent()
	}
	return offset
}

func (s *SurfaceDelegate) getXOffsetFor(d Displayable) float64 {
	current := d
	offset := 0.0
	for current != nil {
		offset += current.GetX()
		current = d.GetParent()
	}
	return offset
}

func (s *SurfaceDelegate) GetOffsetSurfaceFor(d Displayable) Surface {
	fmt.Println("Surface Delegate GetOffsetSurface")
	x := s.getXOffsetFor(d)
	y := s.getYOffsetFor(d)
	return &SurfaceDelegate{
		delegateTo: s.delegateTo,
		offsetX:    x,
		offsetY:    y,
	}
}

func NewSurfaceDelegate(delegateTo Surface) *SurfaceDelegate {
	return &SurfaceDelegate{delegateTo: delegateTo}
}
