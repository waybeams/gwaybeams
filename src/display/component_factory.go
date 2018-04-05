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
		X(0),
		Y(0),
		Z(0),

		/* styles */
		FontColor(-1),
		FontFace(""),
		FontSize(-1),
		BgColor(-1),
		StrokeSize(-1),
		StrokeColor(-1),
	}
}

// NewComponentFactory returns a component factory for the provided component.
// This factory will configure the instantiated component instance with the
// provided default values.
func NewComponentFactory(c componentConstructor, factoryOpts ...ComponentOption) ComponentFactory {
	return func(b Builder, instanceOpts ...ComponentOption) (Displayable, error) {
		// Create a builder if we weren't provided with one. This makes tests much, much
		// more readable, but it not be expected
		if b == nil {
			return nil, errors.New("component factory requires a Builder instance, try Component(NewBuilder()) or in the parent closure, add a (b Builder) argument and forward it to the child nodes")
		}

		// Instantiate the component from the provided factory function.
		instance := c()

		traitOpts := OptionsFor(instance, b.Peek())
		// Apply all default, selected and provided options to the component instance.
		options := append([]ComponentOption{}, DefaultComponentOpts...)
		options = append(options, factoryOpts...)
		options = append(options, traitOpts...)
		options = append(options, instanceOpts...)

		// Send the instance to the provided builder for tree placement.
		b.Push(instance, options...)

		// Everything worked great, return the instance.
		return instance, nil
	}
}
