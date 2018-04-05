package display

// Trait is a concrete factory function that builds a bag of ComponentOptions
// and applies them to all Selected Components before applying
// instance-specified options.
func Trait(b Builder, selector string, opts ...ComponentOption) error {
	component := b.Peek()
	if component == nil {
		panic("Trait definition must be nested inside of a component")
	}
	component.PushTrait(selector, opts...)
	return nil
}
