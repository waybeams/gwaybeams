package display

import (
	"clock"
	"events"
)

// Component is a concrete implementation that is the fundamental building block
// for the display of interactive elements.
//
// The Component definition is spread across all of the files that define
// interfaces that are implemented by it. The master list can be found in
// Displayable.
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

func (c *Component) Clock() clock.Clock {
	return c.Builder().Clock()
}

// NewComponent returns a new base component instance as a Displayable.
func NewComponent() Displayable {
	c := &Component{}
	c.PushUnsub(c.On(events.Focused, c.focusedHandler))
	c.PushUnsub(c.On(events.Blurred, c.blurredHandler))
	return c
}
