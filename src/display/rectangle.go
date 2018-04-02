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

	colors := []uint{0xccccccff, 0x333333ff, 0x666666ff, 0x999999ff}

	depthFromRoot := getDepth(0, d)
	index := depthFromRoot % (len(colors) - 1)
	color := colors[index]

	surface.SetFillColor(color)
	surface.SetStrokeWidth(1)
	surface.SetStrokeColor(0x333333ff)
	surface.DrawRectangle(d.GetX(), d.GetY(), d.GetWidth(), d.GetHeight())
	surface.Stroke()
	surface.Fill()
}
