package spec

type FocusableReader interface {
	FocusedSpec() ReadWriter
	IsFocusable() bool
	IsText() bool
	IsTextInput() bool
}

type FocusableWriter interface {
	SetFocusedSpec(spec ReadWriter)
	SetIsFocusable(value bool)
	SetIsText(value bool)
	SetIsTextInput(value bool)
}

type FocusableReadWriter interface {
	FocusableReader
	FocusableWriter
}

func (c *Spec) FocusedSpec() ReadWriter {
	if c.Parent() == nil {
		return c.focusedSpec
	}
	return Root(c).FocusedSpec()
}

func (c *Spec) SetFocusedSpec(spec ReadWriter) {
	if c.Parent() == nil {
		c.focusedSpec = spec
		return
	}
	Root(c).SetFocusedSpec(spec)
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
