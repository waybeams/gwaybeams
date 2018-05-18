package spec

type FocusableReader interface {
	IsFocused() bool
	IsFocusable() bool
	IsText() bool
	IsTextInput() bool
}

type FocusableWriter interface {
	Blur()
	Focus()
	SetIsFocusable(value bool)
	SetIsText(value bool)
	SetIsTextInput(value bool)
}

type FocusableReadWriter interface {
	FocusableReader
	FocusableWriter
}

func (c *Spec) Blur() {
	c.isFocused = false
}

func (c *Spec) Focus() {
	c.isFocused = true
}

func (c *Spec) IsFocused() bool {
	return c.isFocused
}

func (c *Spec) IsFocusable() bool {
	return c.isFocusable
}

func (c *Spec) IsText() bool {
	return c.isText
}

func (c *Spec) IsTextInput() bool {
	return c.isTextInput
}

func (c *Spec) SetIsFocusable(value bool) {
	c.isFocusable = value
}

func (c *Spec) SetIsText(value bool) {
	c.isText = value
}

func (c *Spec) SetIsTextInput(value bool) {
	c.isTextInput = value
}
