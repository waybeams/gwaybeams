package display

const DefaultBgColor = 0x999999ff
const DefaultFontColor = 0x111111ff
const DefaultFontSize = 12
const DefaultFontFace = "Roboto"
const DefaultStrokeColor = 0x333333ff
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

func (c *Component) BgColor() int {
	bgColor := c.Model().BgColor
	if bgColor == -1 {
		if c.parent != nil {
			return c.parent.BgColor()
		}
		return DefaultBgColor
	}

	return bgColor
}

func (c *Component) FontColor() int {
	fontColor := c.Model().FontColor
	if fontColor == -1 {
		if c.parent != nil {
			return c.parent.FontColor()
		}
		return DefaultFontColor
	}
	return fontColor
}

func (c *Component) FontFace() string {
	fontFace := c.Model().FontFace
	if fontFace == "" {
		if c.parent != nil {
			return c.parent.FontFace()
		}
		return DefaultFontFace
	}
	return fontFace
}

func (c *Component) FontSize() int {
	fontSize := c.Model().FontSize
	if fontSize == -1 {
		if c.parent != nil {
			return c.parent.FontSize()
		}
		return DefaultFontSize
	}
	return fontSize
}

func (c *Component) Gutter() float64 {
	return c.Model().Gutter
}

func (c *Component) SetBgColor(color int) {
	c.Model().BgColor = color
}

func (c *Component) SetFontFace(face string) {
	c.Model().FontFace = face
}

func (c *Component) SetFontSize(size int) {
	c.Model().FontSize = size
}

func (c *Component) SetGutter(gutter float64) {
	c.Model().Gutter = gutter
}

func (c *Component) SetFontColor(size int) {
	c.Model().FontColor = size
}

func (c *Component) SetStrokeColor(size int) {
	c.Model().StrokeColor = size
}

func (c *Component) SetStrokeSize(size int) {
	c.Model().StrokeSize = size
}

func (c *Component) SetVisible(visible bool) {
	c.Model().Visible = visible
}

func (c *Component) StrokeColor() int {
	strokeColor := c.Model().StrokeColor
	if strokeColor == -1 {
		if c.parent != nil {
			return c.parent.StrokeColor()
		}
		return DefaultStrokeColor
	}
	return strokeColor
}

func (c *Component) StrokeSize() int {
	strokeSize := c.Model().StrokeSize
	if strokeSize == -1 {
		if c.parent != nil {
			return c.parent.StrokeSize()
		}
		return DefaultStrokeSize
	}
	return strokeSize
}

func (c *Component) Visible() bool {
	return c.Model().Visible
}
