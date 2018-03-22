package display

import (
	"github.com/rs/xid"
	"log"
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

func (s *Sprite) GetId() string {
	opts := s.GetOptions()
	if opts.Id == "" {
		opts.Id = xid.New().String()
	}

	return opts.Id
}

func (s *Sprite) LayoutType(layoutType LayoutType) {
	s.GetOptions().LayoutType = layoutType
}

func (s *Sprite) GetLayoutType() LayoutType {
	return s.GetOptions().LayoutType
}

func (s *Sprite) GetLayout() Layout {
	switch s.GetLayoutType() {
	case StackLayoutType:
		return StackLayout
	default:
		log.Printf("ERROR: Requested LayoutType (%v) is not supported", s.GetLayoutType())
		return nil
	}
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
	opts := s.GetOptions()
	if opts.Width != w {
		opts.Width = -1
		s.ActualWidth(w)
		opts.Width = opts.ActualWidth
	}
}

func (s *Sprite) WidthInBounds(w float64) float64 {
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

func (s *Sprite) HeightInBounds(h float64) float64 {
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

func (s *Sprite) GetWidth() float64 {
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

func (s *Sprite) Height(h float64) {
	opts := s.GetOptions()
	if opts.Height != h {
		opts.Height = -1
		s.ActualHeight(h)
		opts.Height = opts.ActualHeight
	}
}

func (s *Sprite) GetHeight() float64 {
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

func (s *Sprite) GetFixedWidth() float64 {
	return s.GetWidth()
}

func (s *Sprite) GetFixedHeight() float64 {
	return s.GetHeight()
}

func (s *Sprite) GetPrefWidth() float64 {
	return s.GetOptions().PrefWidth
}

func (s *Sprite) GetPrefHeight() float64 {
	return s.GetOptions().PrefHeight
}

func (s *Sprite) ActualWidth(width float64) {
	s.GetOptions().ActualWidth = s.WidthInBounds(width)
}

func (s *Sprite) GetInferredMinWidth() float64 {
	result := 0.0
	for _, child := range s.children {
		if !child.GetExcludeFromLayout() {
			result = math.Max(result, child.GetMinWidth())
		}
	}
	return result + s.GetHorizontalPadding()
}

func (s *Sprite) GetInferredMinHeight() float64 {
	result := 0.0
	for _, child := range s.children {
		if !child.GetExcludeFromLayout() {
			result = math.Max(result, child.GetMinHeight())
		}
	}
	return result + s.GetHorizontalPadding()
}

func (s *Sprite) ActualHeight(height float64) {
	s.GetOptions().ActualHeight = s.HeightInBounds(height)
}

func (s *Sprite) GetActualWidth() float64 {
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

func (s *Sprite) GetActualHeight() float64 {
	return s.GetOptions().ActualHeight
}

func (s *Sprite) GetHAlign() Alignment {
	return s.GetOptions().HAlign
}

func (s *Sprite) GetVAlign() Alignment {
	return s.GetOptions().VAlign
}

func (s *Sprite) MinWidth(min float64) {
	s.GetOptions().MinWidth = min
	// Ensure we're not already too small for the new min
	if s.GetActualWidth() < min {
		s.ActualWidth(min)
	}
}

func (s *Sprite) GetMinWidth() float64 {
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

func (s *Sprite) MinHeight(h float64) {
	s.GetOptions().MinHeight = h
}

func (s *Sprite) GetMinHeight() float64 {
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

func (s *Sprite) GetPadding() float64 {
	return s.GetOptions().Padding
}

func (s *Sprite) GetHorizontalPadding() float64 {
	return s.GetPaddingLeft() + s.GetPaddingRight()
}

func (s *Sprite) GetVerticalPadding() float64 {
	return s.GetPaddingTop() + s.GetPaddingBottom()
}

func (s *Sprite) getPaddingForSide(getter func() float64) float64 {
	opts := s.GetOptions()
	if getter() == -1 {
		if opts.Padding > 0 {
			return opts.Padding
		}
		return 0
	}
	return getter()
}

func (s *Sprite) GetPaddingLeft() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingLeft
	})
}

func (s *Sprite) GetPaddingRight() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingRight
	})
}

func (s *Sprite) GetPaddingBottom() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingBottom
	})
}

func (s *Sprite) GetPaddingTop() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingTop
	})
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

func (s *Sprite) GetPath() string {
	parent := s.GetParent()
	localPath := "/" + s.GetId()

	if parent != nil {
		return parent.GetPath() + localPath
	}
	return localPath

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
	for _, child := range s.children {
		log.Printf("RENDERING CHILD:", child)
		child.Render(surface)
	}
}

func (s *Sprite) Render(surface Surface) {
	s.GetLayout()(s)
	DrawRectangle(surface, s)
	s.RenderChildren(surface)
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
