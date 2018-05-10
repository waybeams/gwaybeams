package surface

import (
	"github.com/shibukawa/nanovgo"
	"helpers"
	"ui"
)

type nanoSurface struct {
	context *nanovgo.Context
}

func (n *nanoSurface) CreateFont(name, path string) {
	n.context.CreateFont(name, path)
}

func (n *nanoSurface) MoveTo(x float64, y float64) {
	n.context.MoveTo(float32(x), float32(y))
}

func (n *nanoSurface) SetFillColor(color uint) {
	r, g, b, a := helpers.HexIntToRgbaFloat32(color)
	n.context.SetFillColor(nanovgo.Color{r, g, b, a})
}

func (n *nanoSurface) SetStrokeColor(color uint) {
	r, g, b, a := helpers.HexIntToRgbaFloat32(color)
	n.context.SetStrokeColor(nanovgo.Color{r, g, b, a})
}

func (n *nanoSurface) SetStrokeWidth(width float64) {
	n.context.SetStrokeWidth(float32(width))
}

func (n *nanoSurface) Stroke() {
	n.context.Stroke()
}

func (n *nanoSurface) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	// TODO(lbayes): Update external Surface to include direction and facilitate for Cairo
	n.context.Arc(float32(xc), float32(yc), float32(radius), float32(angle1), float32(angle2), nanovgo.Clockwise)
}

func (n *nanoSurface) BeginPath() {
	n.context.BeginPath()
}

func (n *nanoSurface) DebugDumpPathCache() {
	n.context.DebugDumpPathCache()
}

func (n *nanoSurface) Fill() {
	n.context.Fill()
}

func (n *nanoSurface) Rect(x, y, width, height float64) {
	n.context.Rect(float32(x), float32(y), float32(width), float32(height))
}

func (n *nanoSurface) RoundedRect(x, y, width, height, radius float64) {
	n.context.RoundedRect(float32(x), float32(y), float32(width), float32(height), float32(radius))
}

func (n *nanoSurface) SetFontSize(size float64) {
	n.context.SetFontSize(float32(size))
}

func (n *nanoSurface) SetFontFace(face string) {
	n.context.SetFontFace(face)
}

func (n *nanoSurface) Text(x float64, y float64, text string) {
	// TODO(lbayes): Add validation that ensures required calls have been made before calling this function (e.g., SetFontFace)
	n.context.Text(float32(x), float32(y), text)
}

func (n *nanoSurface) GetOffsetSurfaceFor(d ui.Displayable) ui.Surface {
	return NewOffsetSurface(d, n)
}

func NewNano(context *nanovgo.Context) *nanoSurface {
	return &nanoSurface{context: context}
}
