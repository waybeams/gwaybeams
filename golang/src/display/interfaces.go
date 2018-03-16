package display

// Composition and structure
type Composable interface {
	// Uniquely identifiable object within a given tree
	Id() int
	Parent() Displayable
	AddChild(child Displayable) int
	setParent(parent Displayable)
}

// Layout and positioning
type Layoutable interface {
	Width(width int)
	GetWidth() int
	Height(height int)
	GetHeight() int
}

// Styling and drawing
type Renderable interface {
	Render(canvas Canvas)
	Styles([]func())
}

// Entities that can be composed, scaled, positioned, and rendered.
type Displayable interface {
	Composable
	Layoutable
	Renderable
}
