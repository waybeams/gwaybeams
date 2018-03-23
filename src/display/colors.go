package display

func HexIntToRgb(value int) (r, g, b int) {
	r = (value >> 16) & 0xff
	g = (value >> 8) & 0xff
	b = (value) & 0xff
	return r, g, b
}
