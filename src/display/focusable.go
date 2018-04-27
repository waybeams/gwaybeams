package display

import "events"

type Focusable interface {
	Blur()
	Focus()
	NearestFocusable() Displayable
	Focused() bool
	IsFocusable() bool
	IsText() bool
	SetIsFocusable(value bool)
	SetIsText(value bool)
}

func (c *Component) Blur() {
	c.Bubble(NewEvent(events.Blurred, c, nil))
	c.Model().Focused = false
}

func (c *Component) NearestFocusable() Displayable {
	var candidate Displayable = c
	for candidate != nil {
		parent := candidate.Parent()
		if parent == nil || candidate.IsFocusable() {
			return candidate
		}
		candidate = candidate.Parent()
	}
	return nil
}

func (c *Component) Focus() {
	c.Bubble(NewEvent(events.Focused, c, nil))
	c.Model().Focused = true
}

func (c *Component) Focused() bool {
	return c.Model().Focused
}

func (c *Component) IsFocusable() bool {
	return c.Model().IsFocusable
}

func (c *Component) IsText() bool {
	return c.Model().IsText
}

func (c *Component) SetIsFocusable(value bool) {
	c.Model().IsFocusable = value
}

func (c *Component) SetIsText(value bool) {
	c.Model().IsText = value
}

func (c *Component) focusedHandler(e Event) {
	if c.Parent() == nil {
		if c.focusedChild != nil {
			c.focusedChild.Blur()
		}
		c.focusedChild = e.Target().(Displayable)
	}
}

func (c *Component) blurredHandler(e Event) {
	if c.Parent() == nil && c.focusedChild != nil {
		c.focusedChild = nil
	}
}
