package display

import (
	"errors"
	"github.com/golang-ui/cairo"
	"log"
)

type cairoSurface struct {
	context *cairo.Cairo
}

func (c *cairoSurface) MoveTo(x float64, y float64) {
	log.Println("MoveTo")
	cairo.MoveTo(c.context, x, y)
}

func (c *cairoSurface) SetRgba(r, g, b, a float64) {
	log.Println("SetRgba")
	cairo.SetSourceRgba(c.context, r, g, b, a)
}

func (c *cairoSurface) SetLineWidth(width float64) {
	log.Println("SetLineWidth")
	cairo.SetLineWidth(c.context, width)
}

func (c *cairoSurface) Stroke() {
	log.Println("Stroke")
	cairo.Stroke(c.context)
}

func (c *cairoSurface) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	log.Println("Arc")
	cairo.Arc(c.context, xc, yc, radius, angle1, angle2)
}

func (c *cairoSurface) DrawRectangle(x float64, y float64, width float64, height float64) {
	log.Println("DrawRectangle")
	cairo.MakeRectangle(c.context, x, y, width, height)
}

func (c *cairoSurface) Fill() {
	cairo.Fill(c.context)
}

func (c *cairoSurface) FillPreserve() {
	cairo.FillPreserve(c.context)
}

func (c *cairoSurface) Push(d Displayable) error {
	return errors.New("Unsupported method")
}

func (c *cairoSurface) GetRoot() Displayable {
	// Not sure how to throw when error is not part of the interface. :-(
	panic("Unsupported method")
}

func NewCairoSurface(cairo *cairo.Cairo) Surface {
	return &cairoSurface{context: cairo}
}
