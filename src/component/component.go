package component

import (
	"clock"
	"events"
	"ui"
)

// Component is a concrete implementation that is the fundamental building block
// for the display of interactive elements.
//
// The Component definition is spread across all of the files that define
// interfaces that are implemented by it. The master list can be found in
// Displayable.
type Component struct {
	events.EmitterBase

	context      ui.Context
	children     []ui.Displayable
	dirtyNodes   []ui.Displayable
	model        *ui.Model
	parent       ui.Displayable
	traitOptions ui.TraitOptions
	view         ui.RenderHandler
	unsubs       []events.Unsubscriber

	focusedChild          ui.Displayable
	updateableChildrenMap ui.ChildrenTypeMap
	states                map[string][]ui.Option
	currentState          string

	// Typed composition function containers (only one should ever be non-nil)
	composeEmpty                   func()
	composeWithComponent           func(ui.Displayable)
	composeWithContext             func(ui.Context)
	composeWithContextAndComponent func(ui.Context, ui.Displayable)
}

func (c *Component) Clock() clock.Clock {
	return c.Context().Clock()
}

// NewComponent returns a new base component instance as a Displayable.
func New() *Component {
	return &Component{}
}
