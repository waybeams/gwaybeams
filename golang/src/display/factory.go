package display

import "errors"

// Display declaration is a normalized bag of values built from the
// semantic sugar that describes the hierarchy.
type Declaration struct {
	Options           *Opts
	Data              interface{}
	Display           Displayable
	Compose           func()
	ComposeWithUpdate func(func())
}

// Receive the slice of arbitrary, untyped arguments from a factory function
// and convert them into a Declaration or an error.
func ProcessArgs(args []interface{}) (decl *Declaration, err error) {
	decl = &Declaration{}

	if len(args) > 3 {
		return nil, errors.New("Too many arguments sent to ProcessArgs for component factory")
	}

	for _, entry := range args {
		switch entry.(type) {
		case *Opts:
			decl.Options = entry.(*Opts)
		case func():
			decl.Compose = entry.(func())
		case func(func()):
			decl.ComposeWithUpdate = entry.(func(func()))
		default:
			decl.Data = entry
		}
	}

	if decl.Compose != nil && decl.ComposeWithUpdate != nil {
		return nil, errors.New("Only one composition function allowed")
	}

	return decl, nil
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
