package display

import (
	"github.com/golang-ui/cairo"
)

func uintColorToFloat(color uint) float64 {
	if color == 0 {
		return 0
	} else {
		return float64(color) / 255.0
	}
}

type cairoSurfaceAdapter struct {
	context *cairo.Cairo
}

func (c *cairoSurfaceAdapter) MoveTo(x float64, y float64) {
	cairo.MoveTo(c.context, x, y)
}

func (c *cairoSurfaceAdapter) SetRgba(r, g, b, a uint) {
	cairo.SetSourceRgba(c.context, uintColorToFloat(r), uintColorToFloat(g), uintColorToFloat(b), uintColorToFloat(a))
}

func (c *cairoSurfaceAdapter) SetLineWidth(width float64) {
	cairo.SetLineWidth(c.context, width)
}

func (c *cairoSurfaceAdapter) Stroke() {
	cairo.Stroke(c.context)
}

func (c *cairoSurfaceAdapter) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	cairo.Arc(c.context, xc, yc, radius, angle1, angle2)
}

func (c *cairoSurfaceAdapter) DrawRectangle(x float64, y float64, width float64, height float64) {
	cairo.MakeRectangle(c.context, x, y, width, height)
}

func (c *cairoSurfaceAdapter) Fill() {
	cairo.Fill(c.context)
}

func (c *cairoSurfaceAdapter) FillPreserve() {
	cairo.FillPreserve(c.context)
}

func (c *cairoSurfaceAdapter) GetOffsetSurfaceFor(d Displayable) Surface {
	return NewSurfaceDelegateFor(d, c)
}

func NewCairoSurfaceAdapter(cairo *cairo.Cairo) Surface {
	return &cairoSurfaceAdapter{context: cairo}
}
