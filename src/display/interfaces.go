package display

type DisplayableFilter = func(Displayable) bool

type Composable interface {
	GetId() string
	GetParent() Displayable
	AddChild(child Displayable) int
	GetChildCount() int
	GetChildAt(index int) Displayable
	GetChildren() []Displayable
	GetFilteredChildren(DisplayableFilter) []Displayable
	setParent(parent Displayable)
}

// Layout and positioning
type Layoutable interface {
	ActualHeight(h float64)
	ActualWidth(w float64)
	GetActualHeight() float64
	GetActualWidth() float64
	GetExcludeFromLayout() bool
	GetFlexHeight() float64
	GetFlexWidth() float64
	GetHeight() float64
	GetLayout() Layout
	GetMaxWidth() float64
	GetMinHeight() float64
	GetMinWidth() float64
	GetPrefHeight() float64
	GetPrefWidth() float64
	GetWidth() float64
	GetX() float64
	GetY() float64
	GetZ() float64
	Height(height float64)
	Layout(layout Layout)
	MaxHeight(h float64)
	MaxWidth(w float64)
	MinHeight(h float64)
	MinWidth(w float64)
	PrefHeight(h float64)
	PrefWidth(w float64)
	Width(width float64)
	X(x float64)
	Y(y float64)
	Z(z float64)
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

	Title(title string)
	GetTitle() string
}
