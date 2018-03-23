package display

import (
	"github.com/rs/xid"
	"log"
	"math"
)

// Concrete Base component implementation
// Made public for composition, not instantiation.
// Use NewComponent() factory function to create instances.
type Component struct {
	children        []Displayable
	parent          Displayable
	declaration     *Declaration
	styles          StyleDefinition
	stylesAreDefalt bool
}

func (s *Component) GetId() string {
	opts := s.GetOptions()
	if opts.Id == "" {
		opts.Id = xid.New().String()
	}

	return opts.Id
}

func (s *Component) LayoutType(layoutType LayoutType) {
	s.GetOptions().LayoutType = layoutType
}

func (s *Component) GetLayoutType() LayoutType {
	return s.GetOptions().LayoutType
}

func (s *Component) GetLayout() Layout {
	switch s.GetLayoutType() {
	case StackLayoutType:
		return StackLayout
	default:
		log.Printf("ERROR: Requested LayoutType (%v) is not supported", s.GetLayoutType())
		return nil
	}
}

func (s *Component) Styles(styles StyleDefinition) {
	s.styles = styles
}

func (s *Component) GetStylesFor(d Displayable) StyleDefinition {
	log.Println("STYLES?:", s.styles, s.parent)
	if s.styles == nil {
		if s.parent == nil {
			s.styles = NewDefaultStyleDefinition()
			s.stylesAreDefalt = true
		} else {
			return s.parent.GetStylesFor(d)
		}
	}
	return s.styles
}

func (s *Component) GetStyles() StyleDefinition {
	return s.GetStylesFor(s)
}

func (s *Component) Declaration(decl *Declaration) {
	s.declaration = decl
}

func (s *Component) GetOptions() *Opts {
	return s.GetDeclaration().Options
}

func (s *Component) GetDeclaration() *Declaration {
	if s.declaration == nil {
		s.declaration = &Declaration{Options: &Opts{}}
	}
	return s.declaration
}

func (s *Component) X(x float64) {
	s.GetOptions().X = math.Round(x)
}

func (s *Component) GetX() float64 {
	return s.GetOptions().X
}

func (s *Component) Y(y float64) {
	s.GetOptions().Y = math.Round(y)
}

func (s *Component) Z(z float64) {
	s.GetOptions().Z = math.Round(z)
}

func (s *Component) GetY() float64 {
	return s.GetOptions().Y
}

func (s *Component) GetZ() float64 {
	return s.GetOptions().Z
}

func (s *Component) HAlign(value Alignment) {
	s.GetOptions().HAlign = value
}

func (s *Component) Width(w float64) {
	opts := s.GetOptions()
	if opts.Width != w {
		opts.Width = -1
		s.ActualWidth(w)
		opts.Width = opts.ActualWidth
	}
}

func (s *Component) WidthInBounds(w float64) float64 {
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

func (s *Component) HeightInBounds(h float64) float64 {
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

func (s *Component) GetWidth() float64 {
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

func (s *Component) Height(h float64) {
	opts := s.GetOptions()
	if opts.Height != h {
		opts.Height = -1
		s.ActualHeight(h)
		opts.Height = opts.ActualHeight
	}
}

func (s *Component) GetHeight() float64 {
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

func (s *Component) GetFixedWidth() float64 {
	return s.GetWidth()
}

func (s *Component) GetFixedHeight() float64 {
	return s.GetHeight()
}

func (s *Component) GetPrefWidth() float64 {
	return s.GetOptions().PrefWidth
}

func (s *Component) GetPrefHeight() float64 {
	return s.GetOptions().PrefHeight
}

func (s *Component) ActualWidth(width float64) {
	s.GetOptions().ActualWidth = s.WidthInBounds(width)
}

func (s *Component) GetInferredMinWidth() float64 {
	result := 0.0
	for _, child := range s.children {
		if !child.GetExcludeFromLayout() {
			result = math.Max(result, child.GetMinWidth())
		}
	}
	return result + s.GetHorizontalPadding()
}

func (s *Component) GetInferredMinHeight() float64 {
	result := 0.0
	for _, child := range s.children {
		if !child.GetExcludeFromLayout() {
			result = math.Max(result, child.GetMinHeight())
		}
	}
	return result + s.GetHorizontalPadding()
}

func (s *Component) ActualHeight(height float64) {
	s.GetOptions().ActualHeight = s.HeightInBounds(height)
}

func (s *Component) ExcludeFromLayout(value bool) {
	s.GetOptions().ExcludeFromLayout = value
}

func (s *Component) GetActualWidth() float64 {
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

func (s *Component) GetActualHeight() float64 {
	return s.GetOptions().ActualHeight
}

func (s *Component) GetHAlign() Alignment {
	return s.GetOptions().HAlign
}

func (s *Component) GetVAlign() Alignment {
	return s.GetOptions().VAlign
}

func (s *Component) MinWidth(min float64) {
	s.GetOptions().MinWidth = min
	// Ensure we're not already too small for the new min
	if s.GetActualWidth() < min {
		s.ActualWidth(min)
	}
}

func (s *Component) GetMinWidth() float64 {
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

func (s *Component) MinHeight(h float64) {
	s.GetOptions().MinHeight = h
}

func (s *Component) GetMinHeight() float64 {
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

func (s *Component) MaxWidth(w float64) {
	s.GetOptions().MaxWidth = w
}

func (s *Component) GetMaxWidth() float64 {
	return s.GetOptions().MaxWidth
}

func (s *Component) MaxHeight(h float64) {
	s.GetOptions().MaxHeight = h
}

func (s *Component) GetMaxHeight() float64 {
	return s.GetOptions().MaxHeight
}

func (s *Component) GetExcludeFromLayout() bool {
	return s.GetOptions().ExcludeFromLayout
}

func (s *Component) FlexWidth(value float64) {
	s.GetOptions().FlexWidth = value
}

func (s *Component) FlexHeight(value float64) {
	s.GetOptions().FlexHeight = value
}

func (s *Component) GetFlexWidth() float64 {
	return s.GetOptions().FlexWidth
}

func (s *Component) GetFlexHeight() float64 {
	return s.GetOptions().FlexHeight
}

func (s *Component) Padding(value float64) {
	s.GetOptions().Padding = value
}

func (s *Component) PaddingBottom(value float64) {
	s.GetOptions().PaddingBottom = value
}

func (s *Component) PaddingLeft(value float64) {
	s.GetOptions().PaddingLeft = value
}

func (s *Component) PaddingRight(value float64) {
	s.GetOptions().PaddingRight = value
}

func (s *Component) PaddingTop(value float64) {
	s.GetOptions().PaddingTop = value
}

func (s *Component) GetPadding() float64 {
	return s.GetOptions().Padding
}

func (s *Component) VAlign(value Alignment) {
	s.GetOptions().VAlign = value
}

func (s *Component) GetHorizontalPadding() float64 {
	return s.GetPaddingLeft() + s.GetPaddingRight()
}

func (s *Component) GetVerticalPadding() float64 {
	return s.GetPaddingTop() + s.GetPaddingBottom()
}

func (s *Component) getPaddingForSide(getter func() float64) float64 {
	opts := s.GetOptions()
	if getter() == -1 {
		if opts.Padding > 0 {
			return opts.Padding
		}
		return 0
	}
	return getter()
}

func (s *Component) GetPaddingLeft() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingLeft
	})
}

func (s *Component) GetPaddingRight() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingRight
	})
}

func (s *Component) GetPaddingBottom() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingBottom
	})
}

func (s *Component) GetPaddingTop() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingTop
	})
}

func (s *Component) setParent(parent Displayable) {
	if s.stylesAreDefalt && s.parent == nil {
		s.stylesAreDefalt = false
		s.styles = nil
	}

	s.parent = parent
}

func (s *Component) AddChild(child Displayable) int {
	if s.children == nil {
		s.children = make([]Displayable, 0)
	}

	s.children = append(s.children, child)
	child.setParent(s)
	return len(s.children)
}

func (s *Component) GetChildCount() int {
	return len(s.children)
}

func (s *Component) GetChildAt(index int) Displayable {
	return s.children[index]
}

func (s *Component) GetChildren() []Displayable {
	return append([]Displayable{}, s.children...)
}

func (s *Component) GetFilteredChildren(filter DisplayableFilter) []Displayable {
	result := make([]Displayable, 0)
	for _, child := range s.children {
		if filter(child) {
			result = append(result, child)
		}
	}
	return result
}

func (s *Component) GetPath() string {
	parent := s.GetParent()
	localPath := "/" + s.GetId()

	if parent != nil {
		return parent.GetPath() + localPath
	}
	return localPath

}

func (s *Component) GetParent() Displayable {
	return s.parent
}

func (s *Component) LayoutChildren() {
	for _, child := range s.children {
		child.Layout()
	}
}

func (s *Component) Layout() {
	s.GetLayout()(s)
	s.LayoutChildren()
}

func (s *Component) Draw(surface Surface) {
	DrawRectangle(surface, s)
	for _, child := range s.children {
		child.Draw(surface)
	}
}

func (s *Component) Title(title string) {
	s.GetOptions().Title = title
}

func (s *Component) GetTitle() string {
	return s.GetOptions().Title
}

func NewComponentWithOpts(opts *Opts) Displayable {
	instance := NewComponent()
	args := []interface{}{opts}
	decl, _ := NewDeclaration(args)
	instance.Declaration(decl)
	return instance
}

func NewComponent() Displayable {
	return &Component{}
}
