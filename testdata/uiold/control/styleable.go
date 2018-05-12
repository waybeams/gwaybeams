package control

import "ui"

func (c *Control) BgColor() int {
	bgColor := c.Model().BgColor
	if bgColor == -1 {
		if c.parent != nil {
			return c.parent.BgColor()
		}
		return ui.DefaultBgColor
	}

	return bgColor
}

func (c *Control) FontColor() int {
	fontColor := c.Model().FontColor
	if fontColor == -1 {
		if c.parent != nil {
			return c.parent.FontColor()
		}
		return ui.DefaultFontColor
	}
	return fontColor
}

func (c *Control) FontFace() string {
	fontFace := c.Model().FontFace
	if fontFace == "" {
		if c.parent != nil {
			return c.parent.FontFace()
		}
		return ui.DefaultFontFace
	}
	return fontFace
}

func (c *Control) FontSize() int {
	fontSize := c.Model().FontSize
	if fontSize == -1 {
		if c.parent != nil {
			return c.parent.FontSize()
		}
		return ui.DefaultFontSize
	}
	return fontSize
}

func (c *Control) Gutter() float64 {
	return c.Model().Gutter
}

func (c *Control) SetBgColor(color int) {
	c.Model().BgColor = color
}

func (c *Control) SetFontFace(face string) {
	c.Model().FontFace = face
}

func (c *Control) SetFontSize(size int) {
	c.Model().FontSize = size
}

func (c *Control) SetGutter(gutter float64) {
	c.Model().Gutter = gutter
}

func (c *Control) SetFontColor(size int) {
	c.Model().FontColor = size
}

func (c *Control) SetStrokeColor(size int) {
	c.Model().StrokeColor = size
}

func (c *Control) SetStrokeSize(size int) {
	c.Model().StrokeSize = size
}

func (c *Control) SetVisible(visible bool) {
	c.Model().Visible = visible
}

func (c *Control) StrokeColor() int {
	strokeColor := c.Model().StrokeColor
	if strokeColor == -1 {
		if c.parent != nil {
			return c.parent.StrokeColor()
		}
		return ui.DefaultStrokeColor
	}
	return strokeColor
}

func (c *Control) StrokeSize() int {
	strokeSize := c.Model().StrokeSize
	if strokeSize == -1 {
		if c.parent != nil {
			return c.parent.StrokeSize()
		}
		return ui.DefaultStrokeSize
	}
	return strokeSize
}

func (c *Control) Visible() bool {
	return c.Model().Visible
}
