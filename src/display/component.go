package display

import (
	"clock"
	"events"
)

// Component is a concrete base component implementation made public for
// composition, not instantiation.
type Component struct {
	EmitterBase

	builder      Builder
	children     []Displayable
	dirtyNodes   []Displayable
	model        *ComponentModel
	parent       Displayable
	traitOptions TraitOptions
	view         RenderHandler
	unsubs       []Unsubscriber

	focusedChild          Focusable
	cursorState           CursorState
	updateableChildrenMap ChildrenTypeMap
	states                map[string][]ComponentOption
	currentState          string

	// Typed composition function containers (only one should ever be non-nil)
	composeEmpty                   func()
	composeWithBuilder             func(Builder)
	composeWithComponent           func(Displayable)
	composeWithBuilderAndComponent func(Builder, Displayable)
}

func (c *Component) Bubble(event Event) {
	c.Emit(event)

	current := c.Parent()
	for current != nil {
		if event.IsCancelled() {
			return
		}
		current.Emit(event)
		current = current.Parent()
	}
}

func (c *Component) Clock() clock.Clock {
	return c.Builder().Clock()
}

// NewComponent returns a new base component instance as a Displayable.
func NewComponent() Displayable {
	c := &Component{}
	c.PushUnsubscriber(c.On(events.Focused, c.focusedHandler))
	c.PushUnsubscriber(c.On(events.Blurred, c.blurredHandler))
	return c
}
