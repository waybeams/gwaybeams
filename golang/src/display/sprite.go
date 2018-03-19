package display

import "fmt"

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
	return s.declaration.Options.X
}

func (s *Sprite) GetY() float64 {
	return s.declaration.Options.Y
}

func (s *Sprite) GetWidth() float64 {
	return s.declaration.Options.Width
}

func (s *Sprite) Height(height float64) {
	s.height = height
}

func (s *Sprite) GetHeight() float64 {
	return s.declaration.Options.Height
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

func (s *Sprite) RenderChildren(surface Surface) {
	fmt.Println("Sprite.RenderChildren")
}

func (s *Sprite) Render(surface Surface) {
	fmt.Println("Sprite.Render")
	DrawRectangle(surface, s)
}

func NewSprite() Displayable {
	return &Sprite{}
}
