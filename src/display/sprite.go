package display

import (
	"math"
)

// Concrete Sprite implementation
// Made public for composition, not instantiation.
// Use NewSprite() factory function to create instances.
type Sprite struct {
	children    []Displayable
	parent      Displayable
	declaration *Declaration
	layout      Layout
}

func (s *Sprite) GetId() string {
	return s.GetDeclaration().Options.Id
}

func (s *Sprite) Layout(layout Layout) {
	s.layout = layout
}

func (s *Sprite) GetLayout() Layout {
	return s.layout
}

func (s *Sprite) Declaration(decl *Declaration) {
	s.declaration = decl
}

func (s *Sprite) GetDeclaration() *Declaration {
	if s.declaration == nil {
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

func (s *Sprite) Z(z float64) {
	s.GetDeclaration().Options.Z = math.Round(z)
}

func (s *Sprite) GetY() float64 {
	return s.GetDeclaration().Options.Y
}

func (s *Sprite) GetZ() float64 {
	return s.GetDeclaration().Options.Z
}

func (s *Sprite) GetWidth() float64 {
	return s.GetDeclaration().Options.Width
}

func (s *Sprite) Height(height float64) {
	s.GetDeclaration().Options.Height = math.Round(height)
}

func (s *Sprite) GetExcludeFromLayout() bool {
	return s.GetDeclaration().Options.ExcludeFromLayout
}

func (s *Sprite) GetFlexWidth() int {
	return s.GetDeclaration().Options.FlexWidth
}

func (s *Sprite) GetFlexHeight() int {
	return s.GetDeclaration().Options.FlexHeight
}

func (s *Sprite) GetHeight() float64 {
	return s.GetDeclaration().Options.Height
}

func (s *Sprite) setParent(parent Displayable) {
	s.parent = parent
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

func (s *Sprite) GetChildren() []Displayable {
	return append([]Displayable{}, s.children...)
}

func (s *Sprite) GetFilteredChildren(filter DisplayableFilter) []Displayable {
	result := make([]Displayable, 0)
	for _, child := range s.children {
		if filter(child) {
			result = append(result, child)
		}
	}
	return result
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
	DrawRectangle(surface, s)
}

func (s *Sprite) Title(title string) {
	s.GetDeclaration().Options.Title = title
}

func (s *Sprite) GetTitle() string {
	return s.GetDeclaration().Options.Title
}

func NewSpriteWithOpts(opts *Opts) *Sprite {
	instance := NewSprite()
	args := []interface{}{opts}
	decl, _ := NewDeclaration(args)
	instance.Declaration(decl)
	return instance
}

func NewSprite() *Sprite {
	return &Sprite{}
}
