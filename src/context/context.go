package context

import (
	"clock"
	"display"
)

//-----------------------------------------------------------------------------
// Context Interfaces

type Context interface {
	Builder() display.Builder
	Clock() clock.Clock
}

type ContextOption func(c *baseContext)

//-----------------------------------------------------------------------------
// Context Implementation

// baseContext is a concrete implementation of the Context interface that
// is required in order to instantiate any component.
type baseContext struct {
	builder display.Builder
	clock   clock.Clock
}

// Builder returns the configured (or default) Builder instance.
func (b *baseContext) Builder() display.Builder {
	return b.builder
}

// Clock returns the configured (or default) Clock instance.
func (b *baseContext) Clock() clock.Clock {
	return b.clock
}

//-----------------------------------------------------------------------------
// Context Constructor

// NewContext creates and returns a configured context
func New(options ...ContextOption) *baseContext {
	instance := &baseContext{}
	for _, option := range options {
		option(instance)
	}
	if instance.builder == nil {
		instance.builder = display.NewBuilder()
	}
	if instance.clock == nil {
		instance.clock = clock.New()
	}
	return instance
}

//-----------------------------------------------------------------------------
// Context Options Below

// Builder ContextOption
func Builder(b display.Builder) ContextOption {
	return func(c *baseContext) {
		c.builder = b
	}
}

// Clock ContextOption
func Clock(c clock.Clock) ContextOption {
	return func(context *baseContext) {
		context.clock = c
	}
}
