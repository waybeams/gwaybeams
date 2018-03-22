package display

type Option (func(d Displayable) error)

type NewComponent (func() Displayable)
type ComponentFactory (func(b Builder, opts ...Option) (Displayable, error))

func NewComponentFactory(c NewComponent) ComponentFactory {
	return func(b Builder, opts ...Option) (Displayable, error) {
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

var Sprite = NewComponentFactory(NewSprite)

func FlexWidth(value float64) Option {
	return func(d Displayable) error {
		d.FlexWidth(value)
		return nil
	}
}

func FlexHeight(value float64) Option {
	return func(d Displayable) error {
		d.FlexHeight(value)
		return nil
	}
}
