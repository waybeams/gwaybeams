package comp

import "ui"

func (c *Component) BgColor() int {
	bgColor := c.Model().BgColor
	if bgColor == -1 {
		if c.parent != nil {
			return c.parent.BgColor()
		}
		return ui.DefaultBgColor
	}

	return bgColor
}

func (c *Component) FontColor() int {
	fontColor := c.Model().FontColor
	if fontColor == -1 {
		if c.parent != nil {
			return c.parent.FontColor()
		}
		return ui.DefaultFontColor
	}
	return fontColor
}

func (c *Component) FontFace() string {
	fontFace := c.Model().FontFace
	if fontFace == "" {
		if c.parent != nil {
			return c.parent.FontFace()
		}
		return ui.DefaultFontFace
	}
	return fontFace
}

func (c *Component) FontSize() int {
	fontSize := c.Model().FontSize
	if fontSize == -1 {
		if c.parent != nil {
			return c.parent.FontSize()
		}
		return ui.DefaultFontSize
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
		return ui.DefaultStrokeColor
	}
	return strokeColor
}

func (c *Component) StrokeSize() int {
	strokeSize := c.Model().StrokeSize
	if strokeSize == -1 {
		if c.parent != nil {
			return c.parent.StrokeSize()
		}
		return ui.DefaultStrokeSize
	}
	return strokeSize
}

func (c *Component) Visible() bool {
	return c.Model().Visible
}
