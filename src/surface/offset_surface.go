package surface

import (
	"font"
	"spec"
)

// OffsetSurface provides a Surface interface to a concrete Surface
// implementation, but will offset any global coordinates to the local
// coordinate space.
type OffsetSurface struct {
	delegateTo spec.Surface
	offsetX    float64
	offsetY    float64
}

func (s *OffsetSurface) BeginFrame(w, h float64) {
	s.delegateTo.BeginFrame(w, h)
}

func (s *OffsetSurface) EndFrame() {
	s.delegateTo.EndFrame()
}

func (s *OffsetSurface) Init() {
	s.delegateTo.Init()
}

func (s *OffsetSurface) Close() {
	s.delegateTo.Close()
}

func (s *OffsetSurface) Font(name string) *font.Font {
	return s.delegateTo.Font(name)
}

func (s *OffsetSurface) CreateFont(name, path string) {
	s.delegateTo.CreateFont(name, path)
}

// Arc draws an arc from the x,y point along angle 1 and 2 at the provided radius.
func (s *OffsetSurface) Arc(xc, yc, radius, angle1, angle2 float64) {
	xc += s.offsetX
	yc += s.offsetY
	s.delegateTo.Arc(xc, yc, radius, angle1, angle2)
}

// BeginPath should be called before a Stroke or Fill
func (s *OffsetSurface) BeginPath() {
	s.delegateTo.BeginPath()
}

// DebugDumpPathCache will print the current Path cache to log.
func (s *OffsetSurface) DebugDumpPathCache() {
	s.delegateTo.DebugDumpPathCache()
}

// Rect draws a rectangle from x and y to width and height.
func (s *OffsetSurface) Rect(x, y, width, height float64) {
	x += s.offsetX
	y += s.offsetY
	s.delegateTo.Rect(x, y, width, height)
}

// RoundedRect draws a rectangle from x and y to width and height with the
// provided radius.
func (s *OffsetSurface) RoundedRect(x, y, width, height, radius float64) {
	x += s.offsetX
	y += s.offsetY
	s.delegateTo.RoundedRect(x, y, width, height, radius)
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

// GetOffsetSurfaceFor provides offset surface for nested control so that
// they can use local coordinates for positioning.
func (s *OffsetSurface) GetOffsetSurfaceFor(r spec.Reader) spec.Surface {
	return NewOffsetSurface(r, s)
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

// NewOffsetSurface creates a new surface delegate.
func NewOffsetSurface(r spec.Reader, delegateTo spec.Surface) spec.Surface {
	parent := r.Parent()
	var x, y float64
	if parent != nil {
		x, y = parent.X(), parent.Y()
		// x, y = parent.XOffset(), parent.YOffset()
	} else {
		x, y = 0, 0
		// x, y = r.XOffset(), r.YOffset()
	}

	return &OffsetSurface{
		delegateTo: delegateTo,
		offsetX:    x,
		offsetY:    y,
	}
}
