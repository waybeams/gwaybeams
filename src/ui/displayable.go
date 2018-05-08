package ui

import "events"

const DefaultFrameRate = 60
const DefaultWindowWidth = 800
const DefaultWindowHeight = 600
const DefaultWindowTitle = "Default Title"

// RenderHandler is a function type that will draw control state onto the provided
// Surface.
type RenderHandler func(s Surface, d Displayable) error

// Displayable entities can be composed, scaled, positioned, and drawn.
// This is the uber-interface for all Visual Elements in Waybeams.
type Displayable interface {
	events.Emitter
	Composable
	Layoutable
	Styleable
	Focusable
	Stateful
	Updateable

	// Text and Title are both kind of weird for the general
	// control case... Need to think more about this.
	Data(key string) interface{}
	DataAsString(key string) string
	Draw(s Surface)
	InvalidNodes() []Displayable
	Invalidate()
	InvalidateChildren()
	InvalidateChildrenFor(d Displayable)
	PushTrait(sel string, opts ...Option) error
	PushUnsub(events.Unsubscriber)
	SetData(key string, data interface{})
	SetText(text string)
	SetTitle(title string)
	SetView(view RenderHandler)
	Text() string
	Title() string
	TraitOptions() TraitOptions
	UnsubAll()
	View() RenderHandler
}
