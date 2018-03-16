package display

// Entities that can participate in the renderable surface and layout scheme
// must implement this interface.
type Displayable interface {
	// Identity
	Id() int

	// Composition
	Parent() Displayable
	AddChild(child Displayable) int
	setParent(parent Displayable)

	// Layout
	SetWidth(width int)
	GetWidth() int
	SetHeight(height int)
	GetHeight() int
}
