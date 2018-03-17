package display

type Rectangle struct {
	Sprite
}

func (c *Rectangle) Render(surface Surface) {
	surface.MakeRectangle(float64(c.x), float64(c.y), float64(c.width), float64(c.height))
	// surface.MakeRectangle(10, 10, 200, 300)

	surface.SetRgba(1, 1, 0, 1)
	surface.FillPreserve()

	surface.SetLineWidth(1)
	surface.SetRgba(0, 0, 0, 1)
	surface.Stroke()
}

func NewRectangle() Displayable {
	return &Rectangle{}
}
