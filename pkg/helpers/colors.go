package helpers

import "strconv"

// UintColorToFloat64 converts a uint color (0-255) to a float representation
// of that same color (0.0 - 1.0)
func UintColorToFloat64(color uint) float64 {
	if color == 0 {
		return 0
	}
	return float64(color) / 255.0
}

// UintColorToFloat32 converts the provided uint color (0-255) to a float
// representation (0.0 - 1.0).
func UintColorToFloat32(color uint) float32 {
	// Could probably call UintColorToFloat64 and coerce result into float32, but
	// not 100% sure at the moment we would not introduce some kind of precision
	// errors. If you know this would be fine, please update and send a PR.
	if color == 0 {
		return 0
	}
	return float32(color) / 255.0
}

// HexIntToRgba separates the red, green, blue and alpha channels from an 8
// character hex value and returns each one as a uint.
func HexIntToRgba(value uint) (r, g, b, a uint) {
	r = (value >> 24) & 0xff
	g = (value >> 16) & 0xff
	b = (value >> 8) & 0xff
	a = (value) & 0xff
	return r, g, b, a
}

// HexIntToRgbaFloat32 separates the red, green, blue and alpha channels from
// an 8 character hex value and returns each one as a uint.
func HexIntToRgbaFloat32(value uint) (r, g, b, a float32) {
	red, green, blue, alpha := HexIntToRgba(value)
	return UintColorToFloat32(red), UintColorToFloat32(green), UintColorToFloat32(blue), UintColorToFloat32(alpha)
}

// HexIntToRgbaFloat64 separates the red, green, blue and alpha channels from a
// 6 character hex value and return each channel as a float64 (0-255 is
// represented as a value from zero to 1).
func HexIntToRgbaFloat64(value uint) (r, g, b, a float64) {
	red, green, blue, alpha := HexIntToRgba(value)
	return UintColorToFloat64(red), UintColorToFloat64(green), UintColorToFloat64(blue), UintColorToFloat64(alpha)
}

// HexIntToRgb separates the red, green and blue channels from a 6 character
// hex value and return each one as a uint.
func HexIntToRgb(value uint) (r, g, b uint) {
	r = (value >> 16) & 0xff
	g = (value >> 8) & 0xff
	b = (value) & 0xff
	return r, g, b
}

// UintToHexString returns the uint value as an 8 character hex color string.
func UintToHexString(value uint) string {
	r, g, b, a := HexIntToRgba(value)
	rs := strconv.FormatInt(int64(r), 16)
	gs := strconv.FormatInt(int64(g), 16)
	bs := strconv.FormatInt(int64(b), 16)
	as := strconv.FormatInt(int64(a), 16)
	return "#" + rs + gs + bs + as
}

// HexIntToRgbFloat64 separates the red, green and blue channels from a 6
// character hex value and returns each channel as a float64 (0-255 is
// represented as a value from zero to 1).
func HexIntToRgbFloat64(value uint) (r, g, b float64) {
	red, green, blue := HexIntToRgb(value)
	return UintColorToFloat64(red), UintColorToFloat64(green), UintColorToFloat64(blue)
}
