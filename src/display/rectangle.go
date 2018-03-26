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
	colors := []uint{0x333333, 0x666666, 0x999999, 0xcccccc}

	depthFromRoot := getDepth(0, d)
	index := depthFromRoot % (len(colors) - 1)
	r, g, b := HexIntToRgb(colors[index])

	surface.MoveTo(d.GetX(), d.GetY())
	surface.SetRgba(r, g, b, 255)

	surface.DrawRectangle(d.GetX(), d.GetY(), d.GetWidth(), d.GetHeight())
	surface.FillPreserve()

	surface.SetLineWidth(1)
	surface.SetRgba(0x33, 0x33, 0x33, 0xff)
	surface.Stroke()
}
