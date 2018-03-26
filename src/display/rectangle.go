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
	colors := []int{0xffcc00, 0xccff00, 0xcc00ff, 0x00ccff}

	depthFromRoot := getDepth(0, d)
	index := depthFromRoot % (len(colors) - 1)
	r, g, b := HexIntToRgb(colors[index])

	surface.MoveTo(d.GetX(), d.GetY())
	surface.SetRgba(float64(r), float64(g), float64(b), 1)

	surface.DrawRectangle(d.GetX(), d.GetY(), d.GetWidth(), d.GetHeight())
	surface.FillPreserve()

	surface.SetLineWidth(1)
	surface.SetRgba(0, 0, 0, 1)
	surface.Stroke()
}
