package display

// Concrete Sprite implementation
// Made public for composition, not instantiation.
// Use NewSprite() factory function to create instances.
type Sprite struct {
	children []Displayable
	id       int
	parent   Displayable

	declaration *Declaration

	height float64
	width  float64
	x      float64
	y      float64
}

func (s *Sprite) Declaration(decl *Declaration) {
	s.declaration = decl
}

func (s *Sprite) GetDeclaration() *Declaration {
	return s.declaration
}

func (s *Sprite) Width(width float64) {
	s.width = width
}

func (s *Sprite) GetX() float64 {
	return s.x
}

func (s *Sprite) GetY() float64 {
	return s.y
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

func (s *Sprite) GetId() int {
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

func (s *Sprite) GetChildCount() int {
	return len(s.children)
}

func (s *Sprite) GetChildAt(index int) Displayable {
	return s.children[index]
}

func (s *Sprite) GetParent() Displayable {
	return s.parent
}

func (s *Sprite) Styles(styles []func()) {
}

func (s *Sprite) GetStyles() []func() {
	return nil
}

// Remove this and just delegate to the opts object
// for state
func (s *Sprite) UpdateState(opts *Opts) {
	s.width = opts.Width
	s.height = opts.Height
	s.x = opts.X
	s.y = opts.Y
}

func NewSprite() Displayable {
	return &Sprite{}
}
