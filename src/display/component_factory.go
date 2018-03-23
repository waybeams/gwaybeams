package display

type newComponent (func() Displayable)
type innerComponentFactory (func(b Builder, opts ...ComponentOption) (Displayable, error))

type ComponentFactory (func(c newComponent) innerComponentFactory)

// Returns a component factory that will properly accept options and register a
// component with the Builder.
//
// Usage:
//   var Box = NewComponentFactory(NewComponent)
//
// Callers can then:
//   sprite, err := Box(FlexWidth(1), MaxWidth(100), MinWidth(10))
//
func NewComponentFactory(c newComponent) innerComponentFactory {
	return func(b Builder, opts ...ComponentOption) (Displayable, error) {
		// Instantiate the component from the provided factory function.
		instance := c()
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
