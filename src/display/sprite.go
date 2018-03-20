package display

import (
	"fmt"
	"math"
)

// Concrete Sprite implementation
// Made public for composition, not instantiation.
// Use NewSprite() factory function to create instances.
type Sprite struct {
	children    []Displayable
	parent      Displayable
	declaration *Declaration
}

func (s *Sprite) Declaration(decl *Declaration) {
	s.declaration = decl
}

func (s *Sprite) GetDeclaration() *Declaration {
	if s.declaration == nil {
		fmt.Println("CREATING DECLARATION")
		s.declaration = &Declaration{Options: &Opts{}}
	}
	return s.declaration
}

func (s *Sprite) Width(width float64) {
	s.GetDeclaration().Options.Width = math.Round(width)
}

func (s *Sprite) X(x float64) {
	s.GetDeclaration().Options.X = math.Round(x)
}

func (s *Sprite) GetX() float64 {
	return s.GetDeclaration().Options.X
}

func (s *Sprite) Y(y float64) {
	s.GetDeclaration().Options.Y = math.Round(y)
}

func (s *Sprite) GetY() float64 {
	return s.GetDeclaration().Options.Y
}

func (s *Sprite) GetWidth() float64 {
	return s.GetDeclaration().Options.Width
}

func (s *Sprite) Height(height float64) {
	s.GetDeclaration().Options.Height = math.Round(height)
}

func (s *Sprite) GetFlexWidth() float64 {
	return s.GetDeclaration().Options.FlexWidth
}

func (s *Sprite) GetFlexHeight() float64 {
	return s.GetDeclaration().Options.FlexHeight
}

func (s *Sprite) GetHeight() float64 {
	return s.GetDeclaration().Options.Height
}

func (s *Sprite) setParent(parent Displayable) {
	s.parent = parent
}

func (s *Sprite) GetId() string {
	return s.GetDeclaration().Options.Id
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
}

func (s *Sprite) Render(surface Surface) {
	fmt.Println("Sprite.Render")
	DrawRectangle(surface, s)
}

func NewSprite() Displayable {
	return &Sprite{}
}
