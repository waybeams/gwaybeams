package display

// Concrete Sprite implementation
// Made public mainly for composition, not instantiation.
// Use NewSprite() factory function to create instances.
type Sprite struct {
	children []Displayable
	id       int
	parent   Displayable

	height float64
	width  float64
	x      float64
	y      float64
}

func (s *Sprite) Width(width float64) {
	s.width = width
}

func (s *Sprite) GetWidth() float64 {
	return s.width
}

func (s *Sprite) Height(height float64) {
	s.height = height
}

func (s *Sprite) GetHeight() float64 {
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

func (s *Sprite) Render(surface Surface) {
	s.RenderChildren(surface)
}

func (s *Sprite) RenderChildren(surface Surface) {
	if s.children != nil {
		for _, child := range s.children {
			child.Render(surface)
		}
	}
}

func (s *Sprite) Styles(styles []func()) {
}

func (s *Sprite) GetStyles() []func() {
	return nil
}

func (s *Sprite) UpdateState(opts *Opts) {
	s.width = opts.Width
	s.height = opts.Height
	s.x = opts.X
	s.y = opts.Y
}

func NewSprite() Displayable {
	return &Sprite{}
}
