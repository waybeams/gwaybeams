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
	return s.GetOptions().Id
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

func (s *Sprite) GetOptions() *Opts {
	return s.GetDeclaration().Options
}

func (s *Sprite) GetDeclaration() *Declaration {
	if s.declaration == nil {
		s.declaration = &Declaration{Options: &Opts{}}
	}
	return s.declaration
}

func (s *Sprite) X(x float64) {
	s.GetOptions().X = math.Round(x)
}

func (s *Sprite) GetX() float64 {
	return s.GetOptions().X
}

func (s *Sprite) Y(y float64) {
	s.GetOptions().Y = math.Round(y)
}

func (s *Sprite) Z(z float64) {
	s.GetOptions().Z = math.Round(z)
}

func (s *Sprite) GetY() float64 {
	return s.GetOptions().Y
}

func (s *Sprite) GetZ() float64 {
	return s.GetOptions().Z
}

func (s *Sprite) Width(w float64) {
	s.GetOptions().Width = w
}

func (s *Sprite) WidthInBounds(w float64) float64 {
	return 0.0
}

func (s *Sprite) HeightInBounds(h float64) float64 {
	return 0.0
}

func (s *Sprite) GetWidth() float64 {
	return s.GetOptions().Width
}

func (s *Sprite) Height(h float64) {
	s.GetOptions().Height = h
}

func (s *Sprite) GetHeight() float64 {
	return s.GetOptions().Height
}

func (s *Sprite) GetPrefWidth() float64 {
	return s.GetOptions().PrefWidth
}

func (s *Sprite) GetPrefHeight() float64 {
	return s.GetOptions().PrefHeight
}

func (s *Sprite) GetActualWidth() float64 {
	return s.GetOptions().ActualWidth
}

func (s *Sprite) GetActualHeight() float64 {
	return s.GetOptions().ActualHeight
}

func (s *Sprite) MinWidth(w float64) {
	s.GetOptions().MinWidth = w
}

func (s *Sprite) GetMinWidth() float64 {
	return s.GetOptions().MinWidth
}

func (s *Sprite) MinHeight(h float64) {
	s.GetOptions().MinHeight = h
}

func (s *Sprite) GetMinHeight() float64 {
	return s.GetOptions().MinHeight
}

func (s *Sprite) MaxWidth(w float64) {
	s.GetOptions().MaxWidth = w
}

func (s *Sprite) GetMaxWidth() float64 {
	return s.GetOptions().MaxWidth
}

func (s *Sprite) MaxHeight(h float64) {
	s.GetOptions().MaxHeight = h
}

func (s *Sprite) GetMaxHeight() float64 {
	return s.GetOptions().MaxHeight
}

func (s *Sprite) GetExcludeFromLayout() bool {
	return s.GetOptions().ExcludeFromLayout
}

func (s *Sprite) GetFlexWidth() float64 {
	return s.GetOptions().FlexWidth
}

func (s *Sprite) GetFlexHeight() float64 {
	return s.GetOptions().FlexHeight
}

func (s *Sprite) GetHorizontalPadding() float64 {
	return s.GetOptions().Padding
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
	s.GetOptions().Title = title
}

func (s *Sprite) GetTitle() string {
	return s.GetOptions().Title
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
