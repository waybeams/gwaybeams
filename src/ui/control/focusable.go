package control

import (
	"events"
	"ui"
)

func (c *Control) Blur() {
	existingFocused := c.Root().FocusedChild()
	if existingFocused == c {
		c.Root().SetFocusedChild(nil)
	}

	if c.HasState("active") {
		c.SetState("active")
	}
	c.Model().Focused = false
	c.Bubble(events.New(events.Blurred, c, nil))
}

func (c *Control) Focus() {
	existingFocused := c.Root().FocusedChild()
	if existingFocused == c {
		return
	} else if existingFocused != nil {
		existingFocused.Blur()
	}
	c.Root().SetFocusedChild(c)

	c.Model().Focused = true
	c.Bubble(events.New(events.Focused, c, nil))
	if c.HasState("focused") {
		c.SetState("focused")
	}
	c.Invalidate()
}

func (c *Control) Focused() bool {
	return c.Model().Focused
}

func (c *Control) FocusedChild() ui.Displayable {
	parent := c.Parent()
	if parent != nil {
		return parent.FocusedChild()
	}
	return c.focusedChild
}

func (c *Control) IsFocusable() bool {
	return c.Model().IsFocusable
}

func (c *Control) IsText() bool {
	return c.Model().IsText
}

func (c *Control) IsTextInput() bool {
	return c.Model().IsTextInput
}

func (c *Control) NearestFocusable() ui.Displayable {
	var candidate ui.Displayable = c
	for candidate != nil {
		parent := candidate.Parent()
		if parent == nil || candidate.IsFocusable() {
			return candidate
		}
		candidate = candidate.Parent()
	}
	return nil
}

func (c *Control) SetFocusedChild(child ui.Displayable) {
	if c.Parent() != nil {
		// We're not root, send it up the tree.
		c.Parent().SetFocusedChild(child)
		return
	}
	// Only store the value if we're the Root() node.
	c.focusedChild = child
}

func (c *Control) SetIsFocusable(value bool) {
	c.Model().IsFocusable = value
}

func (c *Control) SetIsText(value bool) {
	c.Model().IsText = value
}

func (c *Control) SetIsTextInput(value bool) {
	c.Model().IsTextInput = value
}
