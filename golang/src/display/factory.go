package display

// Display declaration is a normalized bag of values built from the
// semantic sugar that describes the hierarchy.
type Declaration struct {
	options     Opts
	displayable Displayable
	compose     func()
}

// Receive the slice of arbitrary, untyped arguments from a factory function
// and convert them into a Declaration or an error.
func ProcessArgs(args []interface{}) (*Declaration, error) {
	return nil, nil
}

// Factory that operates over semantic sugar that we use to describe the
// displayable hierarchy.
type Factory struct {
	stack Stack
}

func (f *Factory) getStack() Stack {
	if f.stack == nil {
		f.stack = NewStack()
	}
	return f.stack
}

func (f *Factory) Push(d Displayable) error {
	s := f.getStack()

	if !s.HasNext() {
		err := s.Push(d)
		if err != nil {
			return err
		}
	}

	return nil
}
