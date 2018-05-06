package ui

const DefaultBgColor = 0xce3262ff
const DefaultFontColor = 0xffffffff
const DefaultFontSize = 24
const DefaultFontFace = "Roboto"
const DefaultStrokeColor = 0xffffffff
const DefaultStrokeSize = 1

// Styleable entities can have their visual styles updated.
type Styleable interface {
	BgColor() int
	FontColor() int
	FontFace() string
	FontSize() int
	SetBgColor(color int)
	SetFontColor(color int)
	SetFontFace(face string)
	SetFontSize(size int)
	SetStrokeColor(color int)
	SetStrokeSize(size int)
	SetVisible(visible bool)
	StrokeColor() int
	StrokeSize() int
	Visible() bool
}
