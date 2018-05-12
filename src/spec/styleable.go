package spec

const DefaultBgColor = 0xce3262ff
const DefaultFontColor = 0xffffffff
const DefaultFontSize = 24
const DefaultFontFace = "Roboto"
const DefaultStrokeColor = 0xffffffff
const DefaultStrokeSize = 1

// Styleable entities can have their visual styles updated.
type StyleableReader interface {
	BgColor() uint
	FontColor() uint
	FontFace() string
	FontSize() float64
	StrokeColor() uint
	StrokeSize() float64
	Visible() bool
}

type StyleableWriter interface {
	SetBgColor(color uint)
	SetFontColor(color uint)
	SetFontFace(face string)
	SetFontSize(size float64)
	SetStrokeColor(color uint)
	SetStrokeSize(size float64)
	SetVisible(visible bool)
}

type StyleableReadWriter interface {
	StyleableReader
	StyleableWriter
}

func (c *Spec) BgColor() uint {
	return c.bgColor
}

func (c *Spec) FontColor() uint {
	fontColor := c.fontColor
	// Inherit FontColor from nearest parent.
	if fontColor == 0 {
		parent := c.Parent()
		if parent != nil {
			return parent.FontColor()
		}
		return DefaultFontColor
	}
	return fontColor
}

func (c *Spec) FontFace() string {
	fontFace := c.fontFace
	if fontFace == "" {
		parent := c.Parent()
		if parent != nil {
			return parent.FontFace()
		}
		return DefaultFontFace
	}
	return fontFace
}

func (c *Spec) FontSize() float64 {
	fontSize := c.fontSize
	if fontSize == 0 {
		parent := c.Parent()
		if parent != nil {
			return parent.FontSize()
		}
		return DefaultFontSize
	}
	return fontSize
}

func (c *Spec) SetBgColor(color uint) {
	c.bgColor = color
}

func (c *Spec) SetFontFace(face string) {
	c.fontFace = face
}

func (c *Spec) SetFontSize(size float64) {
	c.fontSize = size
}

func (c *Spec) SetFontColor(size uint) {
	c.fontColor = size
}

func (c *Spec) SetStrokeColor(size uint) {
	c.strokeColor = size
}

func (c *Spec) SetStrokeSize(size float64) {
	c.strokeSize = size
}

func (c *Spec) SetVisible(visible bool) {
	// We store the opposite of the boolean because the default value is false.
	c.isInvisible = !visible
}

func (c *Spec) StrokeColor() uint {
	return c.strokeColor
}

func (c *Spec) StrokeSize() float64 {
	return c.strokeSize
}

func (c *Spec) Visible() bool {
	// We return the opposite of the stored value so that the interface reads
	// as inverted. "Visible() == true" by default.
	return !c.isInvisible
}
