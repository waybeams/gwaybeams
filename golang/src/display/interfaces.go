package display

type Composable interface {
	Id() int
	Parent() Displayable
	AddChild(child Displayable) int
	ChildCount() int
	ChildAt(index int) Displayable
	setParent(parent Displayable)
}

// Layout and positioning
type Layoutable interface {
	GetHeight() float64
	GetWidth() float64
	GetX() float64
	GetY() float64
	Height(height float64)
	UpdateState(opts *Opts)
	Width(width float64)
}

// Styling and drawing
type Renderable interface {
	Declaration(decl *Declaration)
	GetDeclaration() *Declaration

	// GetLayout() func()
	// GetStyles() []func()
}

// Entities that can be composed, scaled, positioned, and rendered.
type Displayable interface {
	Composable
	Layoutable
	Renderable
}
