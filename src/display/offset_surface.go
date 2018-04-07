package display

import ()

// OffsetSurface provides a Surface interface to a concrete Surface
// implementation, but will offset any global coordinates to the local
// coordinate space.
type OffsetSurface struct {
	delegateTo Surface
	offsetX    float64
	offsetY    float64
}

// Arc draws an arc from the x,y point along angle 1 and 2 at the provided radius.
func (s *OffsetSurface) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	xc += s.offsetX
	yc += s.offsetY
	s.delegateTo.Arc(xc, yc, radius, angle1, angle2)
}

// DrawRectangle draws a rectangle from x and y to width and height.
func (s *OffsetSurface) DrawRectangle(x float64, y float64, width float64, height float64) {
	x += s.offsetX
	y += s.offsetY
	s.delegateTo.DrawRectangle(x, y, width, height)
}

// Fill will fill the previously drawn shape.
func (s *OffsetSurface) Fill() {
	s.delegateTo.Fill()
}

// SetStrokeWidth configures the width in pixels of the next shape.
func (s *OffsetSurface) SetStrokeWidth(width float64) {
	s.delegateTo.SetStrokeWidth(width)
}

// SetFillColor configures the fill color as an RGBA hex value (0xffcc00ff)
func (s *OffsetSurface) SetFillColor(color uint) {
	s.delegateTo.SetFillColor(color)
}

// SetStrokeColor configures the stroke color as an RGBA hex value (0xffcc00ff)
func (s *OffsetSurface) SetStrokeColor(color uint) {
	s.delegateTo.SetStrokeColor(color)
}

// Stroke draws a stroke around the previous shape.
func (s *OffsetSurface) Stroke() {
	s.delegateTo.Stroke()
}

// GetOffsetSurfaceFor provides offset surface for nested components so that
// they can use local coordinates for positioning.
func (s *OffsetSurface) GetOffsetSurfaceFor(d Displayable) Surface {
	return NewOffsetSurface(d, s)
}

func (s *OffsetSurface) SetFontSize(size float64) {
	s.delegateTo.SetFontSize(size)
}

func (s *OffsetSurface) SetFontFace(face string) {
	s.delegateTo.SetFontFace(face)
}

func (s *OffsetSurface) Text(x float64, y float64, text string) {
	x += s.offsetX
	y += s.offsetY
	s.delegateTo.Text(x, y, text)
}

// NewSurfaceDelegateFor creates a new surface delegate.
func NewOffsetSurface(d Displayable, delegateTo Surface) Surface {
	parent := d.GetParent()
	var x, y float64
	if parent != nil {
		x, y = parent.GetXOffset(), parent.GetYOffset()
	} else {
		x, y = d.GetXOffset(), d.GetYOffset()
	}

	return &OffsetSurface{
		delegateTo: delegateTo,
		offsetX:    x,
		offsetY:    y,
	}
}
