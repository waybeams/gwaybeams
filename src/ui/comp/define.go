package comp

import (
	. "ui"
)

// DisplayableDefinition is a function that, when called will instantiate,
// configure and attach a new Displayable instance to a new or existing tree.
type DisplayableDefinition func(c Context, options ...Option) Displayable

// getInstance returns a new instace of the component that is created by
// the provided creation function. If you have an idiomatic creation function
// that returns a concrete type, it must be wrapped with a function that
// returns a Displayable when calling Define.
func getInstance(constr interface{}) Displayable {
	// Instantiate the component from the provided factory function and coerce
	// into Displayable. This should panic for non-Displayables.
	var instance Displayable
	switch constr.(type) {
	case func() *Component:
		instance = constr.(func() *Component)()
	default:
		instance = constr.(func() Displayable)()
	}

	return instance
}

// Define creates a new Displayable Definition that can be used later to
// instantiate and attach Displayables.
func Define(typeName string, constr interface{}, specOpts ...Option) DisplayableDefinition {
	return func(c Context, instanceOpts ...Option) Displayable {
		// Create a builder if we weren't provided with one. This makes tests much, much
		// more readable, but it not be expected
		if c == nil {
			panic("comp.Define() requires a Context as first argument, try Component(context.New()) or in the parent closure, add a (b Builder) argument and forward it to the child nodes")
		}

		instance := getInstance(constr)
		// TODO(lbayes): Maybe trait name or bag name or some other thing?
		instance.SetTypeName(typeName)
		// TODO(lbayes): SetContext() instead.
		instance.SetContext(c)

		// Apply all default, selected and provided options to the component instance.
		options := append([]Option{}, specOpts...)
		options = append(options, instanceOpts...)

		// Send the instance to the provided builder for tree placement, and
		// full option application.
		c.Builder().Push(instance, options...)

		// Everything worked great, return the instance.
		return instance
	}
}
