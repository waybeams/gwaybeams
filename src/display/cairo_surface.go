package display

import (
	"github.com/golang-ui/cairo"
)

type cairoSurfaceAdapter struct {
	context *cairo.Cairo
}

func (c *cairoSurfaceAdapter) MoveTo(x float64, y float64) {
	cairo.MoveTo(c.context, x, y)
}

func (c *cairoSurfaceAdapter) SetFillColor(color uint) {
	// NOTE: SetFillColor and SetStrokeColor convert to cairo.SetSourceRgba, which is order-dependent!
	// This means that these calls can't just happen anywhere in a list of calls without unexpected
	// behavior.
	r, g, b, a := HexIntToRgbaFloat64(color)
	cairo.SetSourceRgba(c.context, r, g, b, a)
}

func (c *cairoSurfaceAdapter) SetStrokeColor(color uint) {
	// NOTE: SetFillColor and SetStrokeColor convert to cairo.SetSourceRgba, which is order-dependent!
	// This means that these calls can't just happen anywhere in a list of calls without unexpected
	// behavior.
	r, g, b, a := HexIntToRgbaFloat64(color)
	cairo.SetSourceRgba(c.context, r, g, b, a)
}

func (c *cairoSurfaceAdapter) SetStrokeWidth(width float64) {
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
	// NOTE: Cairo has Fill() and FillPreserve(), but the preserve version allows us to stroke the rectangle.
	// This may not be the right thing to do here, but FWIW, it feels more consistent with Nanovg.
	cairo.FillPreserve(c.context)
}

func (c *cairoSurfaceAdapter) GetOffsetSurfaceFor(d Displayable) Surface {
	return NewSurfaceDelegateFor(d, c)
}

func (c *cairoSurfaceAdapter) SetFontSize(size float64) {
	// c.context.SetFontSize(float32(size))
}

func (c *cairoSurfaceAdapter) SetFontFace(face string) {
	// c.context.SetFontFace(face)
}

func (c *cairoSurfaceAdapter) Text(x float64, y float64, text string) {
	// c.context.Text(float32(x), float32(y), text)
}

func NewCairoSurfaceAdapter(cairo *cairo.Cairo) Surface {
	return &cairoSurfaceAdapter{context: cairo}
}
