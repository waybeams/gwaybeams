package display

type Factory interface {
	GetRoot() Displayable
	Push(d Displayable) error
}

// Factory that operates over semantic sugar that we use to describe the
// displayable hierarchy.
type factory struct {
	stack Stack
	root  Displayable
}

func (f *factory) getStack() Stack {
	if f.stack == nil {
		f.stack = NewStack()
	}
	return f.stack
}

func (f *factory) GetRoot() Displayable {
	return f.root
}

func (f *factory) Push(d Displayable) error {
	if f.stack == nil {
		f.root = d
	}

	s := f.getStack()

	if !s.HasNext() {
		err := s.Push(d)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewFactory() Factory {
	return &factory{}
}
