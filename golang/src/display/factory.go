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
// and convert them into a well-formed Declaration or return an error.
// Callers can provide an array of objects that include at most 3 entries.
// These entries can include zero or one Opts object, user-typed data struct,
// and zero or one of either a func() or func(func()) callback that will
// compose children on the declared Displayable.
func ProcessArgs(args []interface{}) (decl *Declaration, err error) {
	decl = &Declaration{}

	if len(args) > 3 {
		return nil, errors.New("Too many arguments sent to ProcessArgs for component factory")
	}

	for _, entry := range args {
		switch entry.(type) {
		case *Opts:
			if decl.Options != nil {
				return nil, errors.New("Only one Opts object expected")
			}
			decl.Options = entry.(*Opts)
		case func():
			if decl.Compose != nil {
				return nil, errors.New("Only one Compose function expected")
			}
			decl.Compose = entry.(func())
		case func(func()):
			if decl.ComposeWithUpdate != nil {
				return nil, errors.New("Only one ComposeWithUpdate function expected")
			}
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
