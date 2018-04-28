package display

import (
	"events"
)

type Focusable interface {
	Blur()
	Focus()
	NearestFocusable() Displayable
	Focused() bool
	FocusedChild() Displayable
	IsFocusable() bool
	IsText() bool
	IsTextInput() bool
	SetFocusedChild(child Displayable)
	SetIsFocusable(value bool)
	SetIsText(value bool)
	SetIsTextInput(value bool)
}

func (c *Component) Blur() {
	existingFocused := c.Root().FocusedChild()
	if existingFocused == c {
		c.Root().SetFocusedChild(nil)
	}

	if c.HasState("active") {
		c.SetState("active")
	}
	c.Model().Focused = false
	c.Bubble(NewEvent(events.Blurred, c, nil))
}

func (c *Component) Focus() {
	existingFocused := c.Root().FocusedChild()
	if existingFocused == c {
		return
	} else if existingFocused != nil {
		existingFocused.Blur()
	}
	c.Root().SetFocusedChild(c)

	c.Model().Focused = true
	c.Bubble(NewEvent(events.Focused, c, nil))
	if c.HasState("focused") {
		c.SetState("focused")
	}
}

func (c *Component) Focused() bool {
	return c.Model().Focused
}

func (c *Component) FocusedChild() Displayable {
	parent := c.Parent()
	if parent != nil {
		return parent.FocusedChild()
	}
	return c.focusedChild
}

func (c *Component) IsFocusable() bool {
	return c.Model().IsFocusable
}

func (c *Component) IsText() bool {
	return c.Model().IsText
}

func (c *Component) IsTextInput() bool {
	return c.Model().IsTextInput
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

func (c *Component) SetFocusedChild(child Displayable) {
	if c.Parent() != nil {
		// We're not root, send it up the tree.
		c.Parent().SetFocusedChild(child)
		return
	}
	// Only store the value if we're the Root() node.
	c.focusedChild = child
}

func (c *Component) SetIsFocusable(value bool) {
	c.Model().IsFocusable = value
}

func (c *Component) SetIsText(value bool) {
	c.Model().IsText = value
}

func (c *Component) SetIsTextInput(value bool) {
	c.Model().IsTextInput = value
}
