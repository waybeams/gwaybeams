package display

type Circle struct {
	BaseComponent
}

func (c *Circle) Draw(surface Surface) {
	surface.SetLineWidth(3.0)
	surface.DrawRectangle(10, 10, 200, 300)
	surface.Stroke()

	// surface.SetLineWidth(3.0)
	// surface.SetRgba(1, 1, 1, 1)
	// surface.Arc(xc, yc, radius, angle1, angle2)
	// surface.Stroke()
}

func NewCircle() Displayable {
	return &Circle{}
}
