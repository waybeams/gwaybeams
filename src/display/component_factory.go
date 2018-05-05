package display

import (
	"errors"
)

type componentConstructor (func() Displayable)
type ComponentFactory (func(b Builder, opts ...ComponentOption) (Displayable, error))

var knownTypes map[string]bool

// applyDefaults will apply default component options values. The numeric defaults being set
// to -1 rather than 0 allows the layout engine to more readily determine
// developer intent by answering the question, "Has this value been explicitly
// set?"
func applyDefaults(d Displayable) error {
	model := d.Model()
	model.ActualHeight = -1
	model.ActualWidth = -1
	model.FlexHeight = -1
	model.FlexWidth = -1
	model.HAlign = AlignLeft
	model.VAlign = AlignTop
	model.Height = -1
	model.Width = -1
	model.MaxHeight = -1
	model.MaxWidth = -1
	model.MinHeight = -1
	model.MinWidth = -1
	model.Padding = -1
	model.PaddingBottom = -1
	model.PaddingLeft = -1
	model.PaddingRight = -1
	model.PaddingTop = -1
	model.PrefHeight = -1
	model.PrefWidth = -1
	model.X = 0
	model.Y = 0
	model.Z = 0
	model.LayoutType = NoLayoutType
	model.FontColor = -1
	model.FontSize = -1
	model.BgColor = -1
	model.StrokeSize = -1
	model.StrokeColor = -1
	return nil
}

// NewComponentFactory returns a component factory for the provided component.
// This factory will configure the instantiated component instance with the
// provided values.
func NewComponentFactory(typeName string, c componentConstructor, factoryOpts ...ComponentOption) ComponentFactory {
	return func(b Builder, instanceOpts ...ComponentOption) (Displayable, error) {
		// Create a builder if we weren't provided with one. This makes tests much, much
		// more readable, but it not be expected
		if b == nil {
			return nil, errors.New("component factory requires a Builder instance, try Component(NewBuilder()) or in the parent closure, add a (b Builder) argument and forward it to the child nodes")
		}

		// Instantiate the component from the provided factory function.
		instance := c()
		instance.SetTypeName(typeName)

		// We cannot figure out which traits should be applied to the component until
		// we have applied all other known options to the component instance.
		// Additionally, we expect some options (instanceOpts here) to be applied
		// AFTER traits are applied.
		// For these reasons, we're going ahead with instantiating a temporary instance
		// that we will apply all these options to simply to figure out which traits
		// should be included. The real component will be configured by the builder,
		// and any option errors will be propagated from there.
		// This looks like it's very inefficient and will potentially double the cost of
		// every component node, keep looking here for bottlenecks when profiling.
		// For now, I'm working on correctness, we'll need to work on fast(er) in a
		// separate step.
		fake := c()
		fake.SetTypeName(typeName)
		fake.SetBuilder(b)
		// Apply all default, selected and provided options to the component instance.
		earlyOpts := append([]ComponentOption{}, applyDefaults)
		earlyOpts = append(earlyOpts, factoryOpts...)
		tempOpts := append([]ComponentOption{}, earlyOpts...)
		tempOpts = append(tempOpts, instanceOpts...)
		for _, opt := range tempOpts {
			opt(fake)
		}
		traitOpts := TraitOptionsFor(fake, b.Peek())

		// Apply all default, selected and provided options to the component instance.
		options := append(earlyOpts, traitOpts...)
		options = append(options, instanceOpts...)

		// Send the instance to the provided builder for tree placement, and
		// full option application.
		b.Push(instance, options...)

		err := b.LastError()
		if err != nil {
			return nil, err
		}

		// Everything worked great, return the instance.
		return instance, nil
	}
}

// NewComponentFactoryFrom returns a new factory from an existing factory, with
// provided attribute modifications.
func NewComponentFactoryFrom(typeName string, f ComponentFactory, factoryOpts ...ComponentOption) ComponentFactory {
	return func(b Builder, instanceOpts ...ComponentOption) (Displayable, error) {
		// traitOpts := OptionsFor(instance, b.Peek())
		options := append([]ComponentOption{}, applyDefaults)
		options = append(options, factoryOpts...)
		options = append(options, instanceOpts...)
		instance, err := f(b, options...)

		return instance, err
	}
}
