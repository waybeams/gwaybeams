package display

type Circle struct {
	Sprite
}

func (c *Circle) Render(surface Surface) {
	surface.SetLineWidth(3.0)
	surface.MakeRectangle(10, 10, 200, 300)
	surface.Stroke()

	// surface.SetLineWidth(3.0)
	// surface.SetRgba(1, 1, 1, 1)
	// surface.Arc(xc, yc, radius, angle1, angle2)
	// surface.Stroke()
}

func NewCircle() Displayable {
	return &Circle{}
}
