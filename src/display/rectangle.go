package display

func getDepth(sum int, d Displayable) int {
	parent := d.GetParent()
	if parent == nil {
		return sum
	}
	sum++
	return getDepth(sum, parent)
}

func DrawRectangle(surface Surface, d Displayable) {
	surface.SetFillColor(uint(d.GetBgColor()))
	surface.SetStrokeWidth(float64(d.GetStrokeSize()))
	surface.SetStrokeColor(uint(d.GetStrokeColor()))
	surface.DrawRectangle(d.GetX(), d.GetY(), d.GetWidth(), d.GetHeight())
	surface.Stroke()
	surface.Fill()
}
