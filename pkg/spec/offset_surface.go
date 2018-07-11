package spec

// OffsetSurface provides a Surface interface to a concrete Surface
// implementation, but will offset any global coordinates to the local
// coordinate space.
type OffsetSurface struct {
	delegateTo Surface
	offsetX    float64
	offsetY    float64
}

func (s *OffsetSurface) BeginFrame() {
	s.delegateTo.BeginFrame()
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
func (s *OffsetSurface) GetOffsetSurfaceFor(r Reader) Surface {
	return NewOffsetSurface(r, s)
}

func (s *OffsetSurface) AddFont(name string, path string) {
	s.delegateTo.AddFont(name, path)
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

func (s *OffsetSurface) TextBounds(face string, size float64, text string) (x, y, w, h float64) {
	return s.delegateTo.TextBounds(face, size, text)
}

func (s *OffsetSurface) SetWidth(w float64) {
	s.delegateTo.SetWidth(w)
}

func (s *OffsetSurface) SetHeight(h float64) {
	s.delegateTo.SetHeight(h)
}

func (s *OffsetSurface) Width() float64 {
	return s.delegateTo.Width()
}

func (s *OffsetSurface) Height() float64 {
	return s.delegateTo.Height()
}

// NewOffsetSurface creates a new surface delegate.
func NewOffsetSurface(r Reader, delegateTo Surface) Surface {
	parent := r.Parent()
	var x, y float64
	if parent != nil {
		x, y = parent.X(), parent.Y()
	}

	return &OffsetSurface{
		delegateTo: delegateTo,
		offsetX:    x,
		offsetY:    y,
	}
}
