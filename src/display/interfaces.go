package display

type Composable interface {
	GetId() string
	GetParent() Displayable
	AddChild(child Displayable) int
	GetChildCount() int
	GetChildAt(index int) Displayable
	setParent(parent Displayable)
}

// Layout and positioning
type Layoutable interface {
	GetHeight() float64
	GetWidth() float64
	GetX() float64
	X(x float64)
	GetY() float64
	Y(y float64)
	Height(height float64)
	Width(width float64)

	GetFlexWidth() int
	GetFlexHeight() int
}

// Styling and drawing
type Renderable interface {
	Declaration(decl *Declaration)
	GetDeclaration() *Declaration

	RenderChildren(s Surface)
	Render(s Surface)
	// GetLayout() func()
	// GetStyles() []func()
}

// Entities that can be composed, scaled, positioned, and rendered.
type Displayable interface {
	Composable
	Layoutable
	Renderable
}
