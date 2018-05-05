package component

import (
	. "ui"
)

// DisplayableDefinition is a function that, when called will instantiate,
// configure and attach a new Displayable instance to a new or existing tree.
type DisplayableDefinition func(c Context, options ...Option) Displayable

// Define creates a new Displayable Definition that can be used later to
// instantiate and attach Displayables.
func Define(typeName string, constr interface{}, specOpts ...Option) DisplayableDefinition {
	return func(c Context, instanceOpts ...Option) Displayable {
		// Create a builder if we weren't provided with one. This makes tests much, much
		// more readable, but it not be expected
		if c == nil {
			panic("component.Define() requires a Context as first argument, try Component(ctx.New()) or in the parent closure, add a (b Builder) argument and forward it to the child nodes")
		}

		// Instantiate the component from the provided factory function and coerce
		// into Displayable. This should panic for non-Displayables.
		var instance Displayable
		switch constr.(type) {
		case func() Displayable:
			instance = constr.(func() Displayable)()
		case func() *Component:
			instance = constr.(func() *Component)()
		default:
			panic("Custom component constructors must explicitly return ui.Displayable")
		}

		// TODO(lbayes): Maybe trait name or bag name or some other thing?
		instance.SetTypeName(typeName)
		// TODO(lbayes): SetContext() instead.
		instance.SetContext(c)

		// Apply all default, selected and provided options to the component instance.
		options := append([]Option{}, specOpts...)
		options = append(options, instanceOpts...)
		for _, opt := range options {
			opt(instance)
		}

		// TODO(lbayes): Turn Traits back on when done with rework!
		// traitOpts := TraitOptionsFor(fake, b.Peek())

		// Send the instance to the provided builder for tree placement, and
		// full option application.
		c.Builder().Push(instance, options...)

		// Everything worked great, return the instance.
		return instance
	}
}
