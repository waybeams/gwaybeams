package display

import (
	"github.com/rs/xid"
	"log"
	"math"
)

// Concrete Base component implementation
// Made public for composition, not instantiation.
// Use NewComponent() factory function to create instances.
type BaseComponent struct {
	children        []Displayable
	parent          Displayable
	declaration     *Declaration
	styles          StyleDefinition
	stylesAreDefalt bool
}

func (s *BaseComponent) GetId() string {
	opts := s.GetOptions()
	if opts.Id == "" {
		opts.Id = xid.New().String()
	}

	return opts.Id
}

func (s *BaseComponent) LayoutType(layoutType LayoutType) {
	s.GetOptions().LayoutType = layoutType
}

func (s *BaseComponent) GetLayoutType() LayoutType {
	return s.GetOptions().LayoutType
}

func (s *BaseComponent) GetLayout() Layout {
	switch s.GetLayoutType() {
	case StackLayoutType:
		return StackLayout
	default:
		log.Printf("ERROR: Requested LayoutType (%v) is not supported", s.GetLayoutType())
		return nil
	}
}

func (s *BaseComponent) Styles(styles StyleDefinition) {
	s.styles = styles
}

func (s *BaseComponent) GetStylesFor(d Displayable) StyleDefinition {
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

func (s *BaseComponent) GetStyles() StyleDefinition {
	return s.GetStylesFor(s)
}

func (s *BaseComponent) Declaration(decl *Declaration) {
	s.declaration = decl
}

func (s *BaseComponent) GetOptions() *Opts {
	return s.GetDeclaration().Options
}

func (s *BaseComponent) GetDeclaration() *Declaration {
	if s.declaration == nil {
		s.declaration = &Declaration{Options: &Opts{}}
	}
	return s.declaration
}

func (s *BaseComponent) X(x float64) {
	s.GetOptions().X = math.Round(x)
}

func (s *BaseComponent) GetX() float64 {
	return s.GetOptions().X
}

func (s *BaseComponent) Y(y float64) {
	s.GetOptions().Y = math.Round(y)
}

func (s *BaseComponent) Z(z float64) {
	s.GetOptions().Z = math.Round(z)
}

func (s *BaseComponent) GetY() float64 {
	return s.GetOptions().Y
}

func (s *BaseComponent) GetZ() float64 {
	return s.GetOptions().Z
}

func (s *BaseComponent) HAlign(value Alignment) {
	s.GetOptions().HAlign = value
}

func (s *BaseComponent) Width(w float64) {
	opts := s.GetOptions()
	if opts.Width != w {
		opts.Width = -1
		s.ActualWidth(w)
		opts.Width = opts.ActualWidth
	}
}

func (s *BaseComponent) WidthInBounds(w float64) float64 {
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

func (s *BaseComponent) HeightInBounds(h float64) float64 {
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

func (s *BaseComponent) GetWidth() float64 {
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

func (s *BaseComponent) Height(h float64) {
	opts := s.GetOptions()
	if opts.Height != h {
		opts.Height = -1
		s.ActualHeight(h)
		opts.Height = opts.ActualHeight
	}
}

func (s *BaseComponent) GetHeight() float64 {
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

func (s *BaseComponent) GetFixedWidth() float64 {
	return s.GetWidth()
}

func (s *BaseComponent) GetFixedHeight() float64 {
	return s.GetHeight()
}

func (s *BaseComponent) GetPrefWidth() float64 {
	return s.GetOptions().PrefWidth
}

func (s *BaseComponent) GetPrefHeight() float64 {
	return s.GetOptions().PrefHeight
}

func (s *BaseComponent) ActualWidth(width float64) {
	s.GetOptions().ActualWidth = s.WidthInBounds(width)
}

func (s *BaseComponent) GetInferredMinWidth() float64 {
	result := 0.0
	for _, child := range s.children {
		if !child.GetExcludeFromLayout() {
			result = math.Max(result, child.GetMinWidth())
		}
	}
	return result + s.GetHorizontalPadding()
}

func (s *BaseComponent) GetInferredMinHeight() float64 {
	result := 0.0
	for _, child := range s.children {
		if !child.GetExcludeFromLayout() {
			result = math.Max(result, child.GetMinHeight())
		}
	}
	return result + s.GetHorizontalPadding()
}

func (s *BaseComponent) ActualHeight(height float64) {
	s.GetOptions().ActualHeight = s.HeightInBounds(height)
}

func (s *BaseComponent) ExcludeFromLayout(value bool) {
	s.GetOptions().ExcludeFromLayout = value
}

func (s *BaseComponent) GetActualWidth() float64 {
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

func (s *BaseComponent) GetActualHeight() float64 {
	return s.GetOptions().ActualHeight
}

func (s *BaseComponent) GetHAlign() Alignment {
	return s.GetOptions().HAlign
}

func (s *BaseComponent) GetVAlign() Alignment {
	return s.GetOptions().VAlign
}

func (s *BaseComponent) MinWidth(min float64) {
	s.GetOptions().MinWidth = min
	// Ensure we're not already too small for the new min
	if s.GetActualWidth() < min {
		s.ActualWidth(min)
	}
}

func (s *BaseComponent) GetMinWidth() float64 {
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

func (s *BaseComponent) MinHeight(h float64) {
	s.GetOptions().MinHeight = h
}

func (s *BaseComponent) GetMinHeight() float64 {
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

func (s *BaseComponent) MaxWidth(w float64) {
	s.GetOptions().MaxWidth = w
}

func (s *BaseComponent) GetMaxWidth() float64 {
	return s.GetOptions().MaxWidth
}

func (s *BaseComponent) MaxHeight(h float64) {
	s.GetOptions().MaxHeight = h
}

func (s *BaseComponent) GetMaxHeight() float64 {
	return s.GetOptions().MaxHeight
}

func (s *BaseComponent) GetExcludeFromLayout() bool {
	return s.GetOptions().ExcludeFromLayout
}

func (s *BaseComponent) FlexWidth(value float64) {
	s.GetOptions().FlexWidth = value
}

func (s *BaseComponent) FlexHeight(value float64) {
	s.GetOptions().FlexHeight = value
}

func (s *BaseComponent) GetFlexWidth() float64 {
	return s.GetOptions().FlexWidth
}

func (s *BaseComponent) GetFlexHeight() float64 {
	return s.GetOptions().FlexHeight
}

func (s *BaseComponent) Padding(value float64) {
	s.GetOptions().Padding = value
}

func (s *BaseComponent) PaddingBottom(value float64) {
	s.GetOptions().PaddingBottom = value
}

func (s *BaseComponent) PaddingLeft(value float64) {
	s.GetOptions().PaddingLeft = value
}

func (s *BaseComponent) PaddingRight(value float64) {
	s.GetOptions().PaddingRight = value
}

func (s *BaseComponent) PaddingTop(value float64) {
	s.GetOptions().PaddingTop = value
}

func (s *BaseComponent) GetPadding() float64 {
	return s.GetOptions().Padding
}

func (s *BaseComponent) VAlign(value Alignment) {
	s.GetOptions().VAlign = value
}

func (s *BaseComponent) GetHorizontalPadding() float64 {
	return s.GetPaddingLeft() + s.GetPaddingRight()
}

func (s *BaseComponent) GetVerticalPadding() float64 {
	return s.GetPaddingTop() + s.GetPaddingBottom()
}

func (s *BaseComponent) getPaddingForSide(getter func() float64) float64 {
	opts := s.GetOptions()
	if getter() == -1 {
		if opts.Padding > 0 {
			return opts.Padding
		}
		return 0
	}
	return getter()
}

func (s *BaseComponent) GetPaddingLeft() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingLeft
	})
}

func (s *BaseComponent) GetPaddingRight() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingRight
	})
}

func (s *BaseComponent) GetPaddingBottom() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingBottom
	})
}

func (s *BaseComponent) GetPaddingTop() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetOptions().PaddingTop
	})
}

func (s *BaseComponent) setParent(parent Displayable) {
	if s.stylesAreDefalt && s.parent == nil {
		s.stylesAreDefalt = false
		s.styles = nil
	}

	s.parent = parent
}

func (s *BaseComponent) AddChild(child Displayable) int {
	if s.children == nil {
		s.children = make([]Displayable, 0)
	}

	s.children = append(s.children, child)
	child.setParent(s)
	return len(s.children)
}

func (s *BaseComponent) GetChildCount() int {
	return len(s.children)
}

func (s *BaseComponent) GetChildAt(index int) Displayable {
	return s.children[index]
}

func (s *BaseComponent) GetChildren() []Displayable {
	return append([]Displayable{}, s.children...)
}

func (s *BaseComponent) GetFilteredChildren(filter DisplayableFilter) []Displayable {
	result := make([]Displayable, 0)
	for _, child := range s.children {
		if filter(child) {
			result = append(result, child)
		}
	}
	return result
}

func (s *BaseComponent) GetPath() string {
	parent := s.GetParent()
	localPath := "/" + s.GetId()

	if parent != nil {
		return parent.GetPath() + localPath
	}
	return localPath

}

func (s *BaseComponent) GetParent() Displayable {
	return s.parent
}

func (s *BaseComponent) LayoutChildren() {
	for _, child := range s.children {
		child.Layout()
	}
}

func (s *BaseComponent) Layout() {
	s.GetLayout()(s)
	s.LayoutChildren()
}

func (s *BaseComponent) Draw(surface Surface) {
	DrawRectangle(surface, s)
	for _, child := range s.children {
		child.Draw(surface)
	}
}

func (s *BaseComponent) Title(title string) {
	s.GetOptions().Title = title
}

func (s *BaseComponent) GetTitle() string {
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
	return &BaseComponent{}
}

// Named access for builder integration
var Component = NewComponentFactory(NewComponent)
