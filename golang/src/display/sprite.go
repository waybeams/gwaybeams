package display

var lastId = 0

// Concrete Sprite implementation
type sprite struct {
	children []Displayable
	id       int
	parent   Displayable
	height   int
	width    int
}

func (s *sprite) SetWidth(width int) {
	s.width = width
}

func (s *sprite) GetWidth() int {
	return s.width
}

func (s *sprite) SetHeight(height int) {
	s.height = height
}

func (s *sprite) GetHeight() int {
	return s.height
}

func (s *sprite) setParent(parent Displayable) {
	s.parent = parent
}

func (s *sprite) Id() int {
	return s.id
}

func (s *sprite) AddChild(child Displayable) int {
	if s.children == nil {
		s.children = make([]Displayable, 0)
	}

	s.children = append(s.children, child)
	child.setParent(s)
	return len(s.children)
}

func (s *sprite) Parent() Displayable {
	return s.parent
}

func NewSprite() Displayable {
	lastId++
	return &sprite{
		id: lastId,
	}
}
