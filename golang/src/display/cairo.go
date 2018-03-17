package display

import (
	"github.com/golang-ui/cairo"
)

type cairoAdapter struct {
	context *cairo.Cairo
}

func (c *cairoAdapter) SetRgba(r, g, b, a float64) {
	cairo.SetSourceRgba(c.context, r, g, b, a)
}

func (c *cairoAdapter) SetLineWidth(width float64) {
	cairo.SetLineWidth(c.context, width)
}

func (c *cairoAdapter) Stroke() {
	cairo.Stroke(c.context)
}

func (c *cairoAdapter) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	cairo.Arc(c.context, xc, yc, radius, angle1, angle2)
}

func (c *cairoAdapter) MakeRectangle(x float64, y float64, width float64, height float64) {
	cairo.MakeRectangle(c.context, x, y, width, height)
}

func (c *cairoAdapter) Fill() {
	cairo.Fill(c.context)
}

func (c *cairoAdapter) FillPreserve() {
	cairo.FillPreserve(c.context)
}

func NewCairoAdapter(cairoSurface *cairo.Cairo) Surface {
	return &cairoAdapter{context: cairoSurface}
}
