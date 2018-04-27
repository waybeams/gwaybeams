package display

import "events"

// CursorState is how a component responds to cursor movement.
// A cursor may be any pointing device including fingers.
type CursorState int

const (
	CursorActive = iota
	CursorHovered
	CursorPressed
	CursorDisabled
)

type Focusable interface {
	Blur()
	Focus()
	Focused() bool
	IsFocusable() bool
	SetIsFocusable(value bool)
}

func (c *Component) Blur() {
	c.Bubble(NewEvent(events.Blurred, c, nil))
	c.Model().Focused = false
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

func (c *Component) SetIsFocusable(value bool) {
	c.Model().IsFocusable = value
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
