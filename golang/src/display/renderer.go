package display

import "fmt"

type Renderer interface {
	GetRoot() Displayable
	Push(d Displayable) error
}

// Factory that operates over semantic sugar that we use to describe the
// displayable hierarchy.
type renderer struct {
	*SurfaceDelegate
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

	d.Render(f)
	d.RenderChildren(f)
	s.Pop()

	return nil
}

func CreateRenderer(s Surface, renderHandler func(s Surface)) Surface {
	renderContext := &renderer{SurfaceDelegate: NewSurfaceDelegate(s)}
	fmt.Println("RENDER HANDLER PROVIDED!")
	renderHandler(renderContext)

	return renderContext
}
