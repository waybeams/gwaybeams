package display

import (
	"github.com/shibukawa/nanovgo"
)

type nanoSurface struct {
	context *nanovgo.Context
}

func (n *nanoSurface) MoveTo(x float64, y float64) {
	n.context.MoveTo(float32(x), float32(y))
}

func (n *nanoSurface) SetFillColor(color uint) {
	r, g, b, a := HexIntToRgbaFloat32(color)
	n.context.SetFillColor(nanovgo.Color{r, g, b, a})
}

func (n *nanoSurface) SetStrokeColor(color uint) {
	r, g, b, a := HexIntToRgbaFloat32(color)
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

func (n *nanoSurface) DrawRectangle(x float64, y float64, width float64, height float64) {
	n.context.BeginPath()
	n.context.Rect(float32(x), float32(y), float32(width), float32(height))
}

func (n *nanoSurface) Fill() {
	n.context.Fill()
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

func (n *nanoSurface) GetOffsetSurfaceFor(d Displayable) Surface {
	return NewSurfaceDelegateFor(d, n)
}

func NewNanoSurface(context *nanovgo.Context) *nanoSurface {
	return &nanoSurface{context: context}
}
