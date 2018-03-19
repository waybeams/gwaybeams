package display

import "fmt"

type Renderer interface {
	GetRoot() Displayable
	Push(d Displayable) error
}

// Factory that operates over semantic sugar that we use to describe the
// displayable hierarchy.
type renderer struct {
	Surface
	stack Stack
	root  Displayable
}

func (f *renderer) getStack() Stack {
	if f.stack == nil {
		f.stack = NewStack()
	}
	return f.stack
}

func (f *renderer) GetRoot() Displayable {
	return f.root
}

func (f *renderer) Push(d Displayable) error {
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

func CreateRenderer(s Surface) func(func(s Surface)) {

	return func(renderHandler func(s Surface)) {
		fmt.Println("RENDER HANDLER PROVIDED!")
		renderHandler(s)
	}
}
