package display

type Renderer interface {
	Surface
	Render()
	RenderWithRoot(d Displayable)
}

// Factory that operates over semantic sugar that we use to describe the
// displayable hierarchy.
type renderer struct {
	*SurfaceDelegate
	stack         Stack
	root          Displayable
	renderHandler func(Surface)
}

func (r *renderer) getStack() Stack {
	if r.stack == nil {
		r.stack = NewStack()
	}
	return r.stack
}

func (r *renderer) GetRoot() Displayable {
	return r.root
}

func (r *renderer) Push(d Displayable) error {
	if r.root == nil {
		r.root = d
	}

	s := r.getStack()
	parent := s.Peek()

	if parent != nil {
		parent.AddChild(d)
	}

	err := s.Push(d)
	if err != nil {
		return err
	}

	decl := d.GetDeclaration()
	if decl.Compose != nil {
		decl.Compose()
	} else if decl.ComposeWithUpdate != nil {
		panic("Not yet implemented")
	}

	//if r.root == d {
	// d.Render(r)
	// } else {
	s.Pop()
	// }

	return nil
}

func (r *renderer) Reset() {
	r.stack = NewStack()
	r.root = nil
}

func (r *renderer) RenderWithRoot(root Displayable) {
	r.Reset()
	r.Push(root)
	r.renderHandler(r)
	r.root.Render(r)
}

func (r *renderer) Render() {
	r.Reset()
	r.renderHandler(r)
	r.root.Render(r)
}

func CreateRenderer(s Surface, renderHandler func(s Surface)) Renderer {
	return &renderer{SurfaceDelegate: NewSurfaceDelegate(s), renderHandler: renderHandler}
}
