package control

import (
	"clock"
	"events"
	"ui"
)

// Control is a concrete implementation that is the fundamental building block
// for the display of interactive elements.
//
// The Control definition is spread across all of the files that define
// interfaces that are implemented by it. The master list can be found in
// Displayable.
type Control struct {
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
	composeEmpty                 func()
	composeWithControl           func(ui.Displayable)
	composeWithContext           func(ui.Context)
	composeWithContextAndControl func(ui.Context, ui.Displayable)
}

func (c *Control) Clock() clock.Clock {
	return c.Context().Clock()
}

// New returns a new base control instance as a Displayable.
func New() *Control {
	return &Control{}
}
