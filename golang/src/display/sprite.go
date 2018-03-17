package display

// Concrete Sprite implementation
// Made public mainly for composition, not instantiation.
// Use NewSprite() factory function to create instances.
type Sprite struct {
	children []Displayable
	id       int
	parent   Displayable
	height   int
	width    int
}

func (s *Sprite) Width(width int) {
	s.width = width
}

func (s *Sprite) GetWidth() int {
	return s.width
}

func (s *Sprite) Height(height int) {
	s.height = height
}

func (s *Sprite) GetHeight() int {
	return s.height
}

func (s *Sprite) setParent(parent Displayable) {
	s.parent = parent
}

func (s *Sprite) Id() int {
	return s.id
}

func (s *Sprite) AddChild(child Displayable) int {
	if s.children == nil {
		s.children = make([]Displayable, 0)
	}

	s.children = append(s.children, child)
	child.setParent(s)
	return len(s.children)
}

func (s *Sprite) Parent() Displayable {
	return s.parent
}

func (s *Sprite) Render(canvas Canvas) {
}

func (s *Sprite) Styles(styles []func()) {
}

func (s *Sprite) GetStyles() []func() {
	return nil
}

func NewSprite() Displayable {
	return &Sprite{}
}
