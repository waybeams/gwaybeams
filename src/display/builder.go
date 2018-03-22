package display

type Builder interface {
	Build()
	BuildWithRoot(d Displayable)
	GetRoot() Displayable
}

// Factory that operates over semantic sugar that we use to describe the
// displayable hierarchy.
type builder struct {
	*SurfaceDelegate
	stack        Stack
	root         Displayable
	buildHandler func(Surface)
}

func (b *builder) getStack() Stack {
	if b.stack == nil {
		b.stack = NewStack()
	}
	return b.stack
}

func (b *builder) GetRoot() Displayable {
	return b.root
}

func (b *builder) Push(d Displayable) error {
	if b.root == nil {
		b.root = d
	}

	s := b.getStack()
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

	//if b.root == d {
	// d.builder(r)
	// } else {
	s.Pop()
	// }

	return nil
}

func (b *builder) Reset() {
	b.stack = NewStack()
	b.root = nil
}

func (b *builder) BuildWithRoot(root Displayable) {
	b.Reset()
	b.Push(root)
	b.buildHandler(b)
	b.root.Render(b)
}

func (b *builder) Build() {
	b.Reset()
	b.buildHandler(b)
	b.root.Render(b)
}

func CreateBuilder(s Surface, buildHandler func(s Surface)) Builder {
	return &builder{SurfaceDelegate: NewSurfaceDelegate(s), buildHandler: buildHandler}
}
