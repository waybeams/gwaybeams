package display

import (
	"errors"
)

type componentConstructor (func() Displayable)
type ComponentFactory (func(b Builder, opts ...ComponentOption) (Displayable, error))

var DefaultComponentOpts []ComponentOption

// Initialize default component options values. The numeric defaults being set
// to -1 rather than 0 allows the layout engine to more readily determine
// developer intent by answering the question, "Has this value been explicitly
// set?"
func init() {
	DefaultComponentOpts = []ComponentOption{
		ActualHeight(-1),
		ActualWidth(-1),
		FlexHeight(-1),
		FlexWidth(-1),
		HAlign(LeftAlign),
		Height(-1),
		LayoutType(StackLayoutType),
		MaxHeight(-1),
		MaxWidth(-1),
		MinHeight(-1),
		MinWidth(-1),
		Padding(-1),
		PaddingBottom(-1),
		PaddingLeft(-1),
		PaddingRight(-1),
		PaddingTop(-1),
		PrefHeight(-1),
		PrefWidth(-1),
		VAlign(TopAlign),
		Width(-1),
		X(-1),
		Y(-1),
		Z(-1),
	}
}

// Returns a component factory for the base component. This factory will
// configure the instantiated component instance with default values that are
// not the same as Golang defaults (e.g., generally, numerics are -1 rather
// than 0. This allows the layout engine (and others) to more easily determin
// user intent when scaling and moving components.
func NewComponentFactory(c componentConstructor, defaultOpts ...ComponentOption) ComponentFactory {
	return func(b Builder, opts ...ComponentOption) (Displayable, error) {
		// Create a builder if we weren't provided with one. This makes tests much, much
		// more readable, but it not be expected
		if b == nil {
			return nil, errors.New("Compnent factory requires a Builder instance, try Component(NewBuilder()) or in the parent closure, add a (b Builder) argument and forward it to the child nodes.")
		}
		instance := c()
		defaults := append(DefaultComponentOpts, defaultOpts...)
		opts = append(defaults, opts...)
		// Instantiate the component from the provided factory function.
		// Apply all provided options to the component instance.
		for _, opt := range opts {
			err := opt(instance)
			if err != nil {
				// If an option error is found, bail with it.
				return nil, err
			}
		}

		// Send the instance to the provided builder for tree placement.
		b.Push(instance)
		// Everything worked great, return the instance.
		return instance, nil
	}
}

// Admittedly odd, poorly readable scheme for inheriting a set of default
// options from some other, known component factory.
// The idea here is that one factory may set up a set of defaults and another
// factory and add to that set. Here is a contrived example:
//
// Box := NewComponentFactory
// LimitedBox := NewComponentFactoryFrom(Box, MaxWidth(200), MaxHeight(400))
//
// limited := LimitedBox(NewBuilder(), Width(300), Height(500))
//
// if limited.GetHeight() == 400 {
//     log.Printf("It worked!, the max height default was respected")
// }
func NewComponentFactoryFrom(baseFactory ComponentFactory, defaultOpts ...ComponentOption) ComponentFactory {
	return func(b Builder, opts ...ComponentOption) (Displayable, error) {
		opts = append(defaultOpts, opts...)
		instance, err := baseFactory(b, opts...)
		if err != nil {
			return nil, err
		}
		return instance, nil
	}
}
