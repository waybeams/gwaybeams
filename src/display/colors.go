package display

func HexIntToRgb(value uint) (r, g, b uint) {
	r = (value >> 16) & 0xff
	g = (value >> 8) & 0xff
	b = (value) & 0xff
	return r, g, b
}
