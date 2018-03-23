package display

import (
	"github.com/rs/xid"
	"log"
	"math"
)

// Concrete Sprite implementation
// Made public for composition, not instantiation.
// Use NewSprite() factory function to create instances.
type SpriteComponent struct {
	children    []Displayable
	parent      Displayable
	declaration *Declaration
}

func (s *SpriteComponent) GetId() string {
	opts := s.GetOptions()
	if opts.Id == "" {
		opts.Id = xid.New().String()
	}

	return opts.Id
}

func (s *SpriteComponent) LayoutType(layoutType LayoutType) {
	s.GetOptions().LayoutType = layoutType
}

func (s *SpriteComponent) GetLayoutType() LayoutType {
	return s.GetOptions().LayoutType
}

func (s *SpriteComponent) GetLayout() Layout {
	switch s.GetLayoutType() {
	case StackLayoutType:
		return StackLayout
	default:
		log.Printf("ERROR: Requested LayoutType (%v) is not supported", s.GetLayoutType())
		return nil
	}
}

func (s *SpriteComponent) Declaration(decl *Declaration) {
	s.declaration = decl
}

func (s *SpriteComponent) GetOptions() *Opts {
	return s.GetDeclaration().Options
}

func (s *SpriteComponent) GetDeclaration() *Declaration {
	if s.declaration == nil {
		s.declaration = &Declaration{Options: &Opts{}}
	}
	return s.declaration
}

func (s *SpriteComponent) X(x float64) {
	s.GetOptions().X = math.Round(x)
}

func (s *SpriteComponent) GetX() float64 {
	return s.GetOptions().X
}

func (s *SpriteComponent) Y(y float64) {
	s.GetOptions().Y = math.Round(y)
}

func (s *SpriteComponent) Z(z float64) {
	s.GetOptions().Z = math.Round(z)
}

func (s *SpriteComponent) GetY() float64 {
	return s.GetOptions().Y
}

func (s *SpriteComponent) GetZ() float64 {
	return s.GetOptions().Z
}

func (s *SpriteComponent) HAlign(value Alignment) {
	s.GetOptions().HAlign = value
}

func (s *SpriteComponent) Width(w float64) {
	opts := s.GetOptions()
	if opts.Width != w {
		opts.Width = -1
		s.ActualWidth(w)
		opts.Width = opts.ActualWidth
	}
}

func (s *SpriteComponent) WidthInBounds(w float64) float64 {
	min := s.GetMinWidth()
	max := s.GetMaxWidth()

	width := math.Round(w)
	if min > 0 {
		width = math.Max(min, width)
	}

	if max > 0 {
		width = math.Min(max, width)
	}
	return width
}

func (s *SpriteComponent) HeightInBounds(h float64) float64 {
	min := s.GetMinHeight()
	max := s.GetMaxHeight()

	height := math.Round(h)
	if min > 0 {
		height = math.Max(min, height)
	}

	if max > 0 {
		height = math.Min(max, height)
	}
	return height
}

func (s *SpriteComponent) GetWidth() float64 {
	opts := s.GetOptions()
	if opts.ActualWidth == 0 {
		prefWidth := s.GetPrefWidth()
		if prefWidth > 0 {
			return prefWidth
		}
		return s.GetMinWidth()
	}
	return opts.ActualWidth
}

func (s *SpriteComponent) Height(h float64) {
	opts := s.GetOptions()
	if opts.Height != h {
		opts.Height = -1
		s.ActualHeight(h)
		opts.Height = opts.ActualHeight
	}
}

func (s *SpriteComponent) GetHeight() float64 {
	opts := s.GetOptions()
	if opts.ActualHeight == 0 {
		prefHeight := s.GetPrefHeight()
		if prefHeight > 0 {
			return prefHeight
		}
		return s.GetMinHeight()
	}
	return opts.ActualHeight
}

func (s *SpriteComponent) GetFixedWidth() float64 {
	return s.GetWidth()
}

func (s *SpriteComponent) GetFixedHeight() float64 {
	return s.GetHeight()
}

func (s *SpriteComponent) GetPrefWidth() float64 {
	return s.GetOptions().PrefWidth
}

func (s *SpriteComponent) GetPrefHeight() float64 {
	return s.GetOptions().PrefHeight
}

func (s *SpriteComponent) ActualWidth(width float64) {
	s.GetOptions().ActualWidth = s.WidthInBounds(width)
}

func (s *SpriteComponent) GetInferredMinWidth() float64 {
	result := 0.0
	for _, child := range s.children {
		if !child.GetExcludeFromLayout() {
			result = math.Max(result, child.GetMinWidth())
		}
	}
	return result + s.GetHorizontalPadding()
}

func (s *SpriteComponent) GetInferredMinHeight() float64 {
	result := 0.0
	for _, child := range s.children {
		if !child.GetExcludeFromLayout() {
			result = math.Max(result, child.GetMinHeight())
		}
	}
	return result + s.GetHorizontalPadding()
}

func (s *SpriteComponent) ActualHeight(height float64) {
	s.GetOptions().ActualHeight = s.HeightInBounds(height)
}

func (s *SpriteComponent) ExcludeFromLayout(value bool) {
	s.GetOptions().ExcludeFromLayout = value
}

func (s *SpriteComponent) GetActualWidth() float64 {
	opts := s.GetOptions()

	if opts.Width > 0 {
		return opts.Width
	} else if opts.ActualWidth > 0 {
		return opts.ActualWidth
	}
	prefWidth := s.GetPrefWidth()
	if prefWidth > 0 {
		return prefWidth
	}

	return s.GetMinWidth()
}

func (s *SpriteComponent) GetActualHeight() float64 {
	return s.GetOptions().ActualHeight
}

func (s *SpriteComponent) GetHAlign() Alignment {
	return s.GetOptions().HAlign
}

func (s *SpriteComponent) GetVAlign() Alignment {
	return s.GetOptions().VAlign
}

func (s *SpriteComponent) MinWidth(min float64) {
	s.GetOptions().MinWidth = min
	// Ensure we're not already too small for the new min
	if s.GetActualWidth() < min {
		s.ActualWidth(min)
	}
}

func (s *SpriteComponent) GetMinWidth() float64 {
	opts := s.GetOptions()
	width := opts.Width
	minWidth := opts.MinWidth
	result := 0.0

	if width > 0 {
		result = width
	}
	if minWidth > 0 {
		result = minWidth
	}
	return math.Max(result, s.GetInferredMinWidth())
}

func (s *SpriteComponent) MinHeight(h float64) {
	s.GetOptions().MinHeight = h
}

func (s *SpriteComponent) GetMinHeight() float64 {
	opts := s.GetOptions()
	height := opts.Height
	minHeight := opts.MinHeight
	result := 0.0

	if height > 0 {
		result = height
	}
	if minHeight > 0 {
		result = minHeight
	}
	return math.Max(result, s.GetInferredMinHeight())
}

func (s *SpriteComponent) MaxWidth(w float64) {
	s.GetOptions().MaxWidth = w
}

func (s *SpriteComponent) GetMaxWidth() float64 {
	return s.GetOptions().MaxWidth
}

func (s *SpriteComponent) MaxHeight(h float64) {
	s.GetOptions().MaxHeight = h
}

func (s *SpriteComponent) GetMaxHeight() float64 {
	return s.GetOptions().MaxHeight
}

func (s *SpriteComponent) GetExcludeFromLayout() bool {
	return s.GetOptions().ExcludeFromLayout
}

func (s *SpriteComponent) FlexWidth(value float64) {
	s.GetOptions().FlexWidth = value
}

func (s *SpriteComponent) FlexHeight(value float64) {
	s.GetOptions().FlexHeight = value
}

func (s *SpriteComponent) GetFlexWidth() float64 {
	return s.GetOptions().FlexWidth
}

func (s *SpriteComponent) GetFlexHeight() float64 {
	return s.GetOptions().FlexHeight
}

func (s *SpriteComponent) Padding(value float64) {
	s.GetOptions().Padding = value
}

func (s *SpriteComponent) PaddingBottom(value float64) {
	s.GetOptions().PaddingBottom = value
}

func (s *SpriteComponent) PaddingLeft(value float64) {
	s.GetOptions().PaddingLeft = value
}

func (s *SpriteComponent) PaddingRight(value float64) {
	s.GetOptions().PaddingRight = value
}

func (s *SpriteComponent) PaddingTop(value float64) {
	s.GetOptions().PaddingTop = value
}

func (s *SpriteComponent) GetPadding() float64 {
	return s.GetOptions().Padding
}

func (s *SpriteComponent) VAlign(value Alignment) {
	s.GetOptions().VAlign = value
}

func (s *SpriteComponent) GetHorizontalPadding() float64 {
	return s.GetPaddingLeft() + s.GetPaddingRight()
}

func (s *SpriteComponent) GetVerticalPadding() float64 {
	return s.GetPaddingTop() + s.GetPaddingBottom()
}

func (s *SpriteComponent) getPaddingForSide(getter func() float64) float64 {
	opts := s.GetOptions()
	if getter() == -1 {
		if opts.Padding > 0 {
			return opts.Padding
		}
		return 0
	}
	return getter()
}

func (s *SpriteComponent) GetPaddingLeft() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingLeft
	})
}

func (s *SpriteComponent) GetPaddingRight() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingRight
	})
}

func (s *SpriteComponent) GetPaddingBottom() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingBottom
	})
}

func (s *SpriteComponent) GetPaddingTop() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingTop
	})
}

func (s *SpriteComponent) setParent(parent Displayable) {
	s.parent = parent
}

func (s *SpriteComponent) AddChild(child Displayable) int {
	if s.children == nil {
		s.children = make([]Displayable, 0)
	}

	s.children = append(s.children, child)
	child.setParent(s)
	return len(s.children)
}

func (s *SpriteComponent) GetChildCount() int {
	return len(s.children)
}

func (s *SpriteComponent) GetChildAt(index int) Displayable {
	return s.children[index]
}

func (s *SpriteComponent) GetChildren() []Displayable {
	return append([]Displayable{}, s.children...)
}

func (s *SpriteComponent) GetFilteredChildren(filter DisplayableFilter) []Displayable {
	result := make([]Displayable, 0)
	for _, child := range s.children {
		if filter(child) {
			result = append(result, child)
		}
	}
	return result
}

func (s *SpriteComponent) GetPath() string {
	parent := s.GetParent()
	localPath := "/" + s.GetId()

	if parent != nil {
		return parent.GetPath() + localPath
	}
	return localPath

}

func (s *SpriteComponent) GetParent() Displayable {
	return s.parent
}

func (s *SpriteComponent) Styles(styles []func()) {
}

func (s *SpriteComponent) GetStyles() []func() {
	return nil
}

func (s *SpriteComponent) RenderChildren() {
	for _, child := range s.children {
		child.Render()
	}
}

func (s *SpriteComponent) Render() {
	s.GetLayout()(s)
	s.RenderChildren()
}

func (s *SpriteComponent) Draw(surface Surface) {
	DrawRectangle(surface, s)
	for _, child := range s.children {
		child.Draw(surface)
	}
}

func (s *SpriteComponent) Title(title string) {
	s.GetOptions().Title = title
}

func (s *SpriteComponent) GetTitle() string {
	return s.GetOptions().Title
}

func NewSpriteWithOpts(opts *Opts) Displayable {
	instance := NewSprite()
	args := []interface{}{opts}
	decl, _ := NewDeclaration(args)
	instance.Declaration(decl)
	return instance
}

func NewSprite() Displayable {
	return &SpriteComponent{}
}

// Named access for builder integration
var Sprite = NewComponentFactory(NewSprite)
