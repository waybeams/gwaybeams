package display

import (
	"errors"
	"github.com/rs/xid"
	"log"
	"math"
)

const DefaultBgColor = 0x999999ff
const DefaultFontColor = 0x111111ff
const DefaultFontSize = 12
const DefaultFontFace = "sans"
const DefaultStrokeColor = 0x333333ff
const DefaultStrokeSize = 2

type TraitOptions map[string][]ComponentOption

// Component is a concrete base component implementation made public for
// composition, not instantiation.
type Component struct {
	children           []Displayable
	parent             Displayable
	model              *ComponentModel
	composeSimple      func()
	composeWithBuilder func(Builder)
	traitOptions       TraitOptions
}

func (s *Component) GetID() string {
	model := s.GetModel()
	if model.ID == "" {
		model.ID = xid.New().String()
	}

	return model.ID
}

func (s *Component) PushTrait(selector string, opts ...ComponentOption) error {
	traitOptions := s.GetTraitOptions()
	if traitOptions[selector] != nil {
		return errors.New("duplicate trait selector found with:" + selector)
	}
	traitOptions[selector] = opts
	return nil
}

func (s *Component) GetTraitOptions() TraitOptions {
	if s.traitOptions == nil {
		s.traitOptions = make(map[string][]ComponentOption)
	}
	return s.traitOptions
}

func (s *Component) Composer(composer interface{}) error {
	switch composer.(type) {
	case func():
		s.composeSimple = composer.(func())
	case func(Builder):
		s.composeWithBuilder = composer.(func(Builder))
	default:
		return errors.New("Component.Composer() called with unexpected signature")
	}
	return nil
}

func (s *Component) GetComposeSimple() func() {
	return s.composeSimple
}

func (s *Component) GetComposeWithBuilder() func(Builder) {
	return s.composeWithBuilder
}

func (s *Component) LayoutType(layoutType LayoutTypeValue) {
	s.GetModel().LayoutType = layoutType
}

func (s *Component) GetLayoutType() LayoutTypeValue {
	return s.GetModel().LayoutType
}

func (s *Component) GetLayout() LayoutHandler {
	switch s.GetLayoutType() {
	case StackLayoutType:
		return StackLayout
	case HorizontalFlowLayoutType:
		return HorizontalFlowLayout
	case VerticalFlowLayoutType:
		return VerticalFlowLayout
	default:
		log.Printf("ERROR: Requested LayoutTypeValue (%v) is not supported", s.GetLayoutType())
		return nil
	}
}

func (s *Component) Model(model *ComponentModel) {
	s.model = model
}

func (s *Component) GetModel() *ComponentModel {
	if s.model == nil {
		s.model = &ComponentModel{}
	}
	return s.model
}

func (s *Component) X(x float64) {
	s.GetModel().X = math.Round(x)
}

func (s *Component) GetX() float64 {
	return s.GetModel().X
}

func (s *Component) Y(y float64) {
	s.GetModel().Y = math.Round(y)
}

func (s *Component) Z(z float64) {
	s.GetModel().Z = math.Round(z)
}

func (s *Component) GetY() float64 {
	return s.GetModel().Y
}

func (s *Component) GetZ() float64 {
	return s.GetModel().Z
}

func (s *Component) HAlign(value Alignment) {
	s.GetModel().HAlign = value
}

func (s *Component) Width(w float64) {
	model := s.GetModel()
	if model.Width != w {
		model.Width = -1
		s.ActualWidth(w)
	}
}

func (s *Component) Height(h float64) {
	model := s.GetModel()
	if model.Height != h {
		model.Height = -1
		s.ActualHeight(h)
	}
}

func (s *Component) WidthInBounds(w float64) float64 {
	min := s.GetMinWidth()
	max := s.GetMaxWidth()
	width := w

	if min > -1 {
		width = math.Max(min, width)
	}

	if max > -1 {
		width = math.Min(max, width)
	}
	return width
}

func (s *Component) HeightInBounds(h float64) float64 {
	min := s.GetMinHeight()
	max := s.GetMaxHeight()

	height := math.Round(h)

	if min > -1 {
		height = math.Max(min, height)
	}

	if max > -1 {
		height = math.Min(max, height)
	}
	return height
}

func (s *Component) GetWidth() float64 {
	model := s.GetModel()
	if model.ActualWidth == -1 {
		prefWidth := s.GetPrefWidth()
		if prefWidth > -1 {
			return prefWidth
		}
		inBounds := s.WidthInBounds(model.Width)
		if inBounds > -1.0 {
			return inBounds
		}
		return 0
	}
	return model.ActualWidth
}

func (s *Component) GetHeight() float64 {
	model := s.GetModel()
	if model.ActualHeight == -1 {
		prefHeight := s.GetPrefHeight()
		if prefHeight > -1 {
			return prefHeight
		}
		inBounds := s.HeightInBounds(model.Height)
		if inBounds > -1 {
			return inBounds
		}
		return 0
	}
	return model.ActualHeight
}

func (s *Component) GetFixedWidth() float64 {
	return s.GetModel().Width
}

func (s *Component) GetFixedHeight() float64 {
	return s.GetModel().Height
}

func (s *Component) PrefWidth(value float64) {
	s.GetModel().PrefWidth = value
}

func (s *Component) PrefHeight(value float64) {
	s.GetModel().PrefHeight = value
}

func (s *Component) GetPrefWidth() float64 {
	return s.GetModel().PrefWidth
}

func (s *Component) GetPrefHeight() float64 {
	return s.GetModel().PrefHeight
}

func (s *Component) ActualWidth(width float64) {
	inBounds := s.WidthInBounds(width)
	model := s.GetModel()
	model.ActualWidth = inBounds
	if model.Width != -1 && model.Width != width {
		model.Width = width
	}
}

func (s *Component) ActualHeight(height float64) {
	inBounds := s.HeightInBounds(height)
	model := s.GetModel()
	model.ActualHeight = inBounds
	if model.Height != -1 && model.Height != height {
		model.Height = height
	}
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

func (s *Component) ExcludeFromLayout(value bool) {
	s.GetModel().ExcludeFromLayout = value
}

func (s *Component) GetActualWidth() float64 {
	model := s.GetModel()

	if model.Width > -1 {
		return model.Width
	} else if model.ActualWidth > -1 {
		return model.ActualWidth
	}
	prefWidth := s.GetPrefWidth()
	if prefWidth > -1 {
		return prefWidth
	}

	return s.GetMinWidth()
}

func (s *Component) GetActualHeight() float64 {
	model := s.GetModel()

	if model.Height > -1 {
		return model.Height
	} else if model.ActualHeight > -1 {
		return model.ActualHeight
	}
	prefHeight := s.GetPrefHeight()
	if prefHeight > -1 {
		return prefHeight
	}

	return s.GetMinHeight()
}

func (s *Component) GetHAlign() Alignment {
	return s.GetModel().HAlign
}

func (s *Component) GetVAlign() Alignment {
	return s.GetModel().VAlign
}

func (s *Component) MinWidth(min float64) {
	s.GetModel().MinWidth = min
	// Ensure we're not already too small for the new min
	if s.GetActualWidth() < min {
		s.ActualWidth(min)
	}
}

func (s *Component) MinHeight(min float64) {
	s.GetModel().MinHeight = min
	// Ensure we're not already too small for the new min
	if s.GetActualHeight() < min {
		s.ActualHeight(min)
	}
}

func (s *Component) GetMinWidth() float64 {
	model := s.GetModel()
	width := model.Width
	minWidth := model.MinWidth
	result := -1.0

	if width > -1.0 {
		result = width
	}
	if minWidth > -1.0 {
		result = minWidth
	}

	inferredMinWidth := s.GetInferredMinWidth()
	if inferredMinWidth > 0 {
		return math.Max(result, inferredMinWidth)
	}
	return result
}

func (s *Component) GetMinHeight() float64 {
	model := s.GetModel()
	height := model.Height
	minHeight := model.MinHeight
	result := -1.0

	if height > -1.0 {
		result = height
	}
	if minHeight > -1.0 {
		result = minHeight
	}

	inferredMinHeight := s.GetInferredMinHeight()
	if inferredMinHeight > 0.0 {
		return math.Max(result, inferredMinHeight)
	}
	return result
}

func (s *Component) MaxWidth(max float64) {
	if s.GetWidth() > max {
		s.Width(max)
	}
	s.GetModel().MaxWidth = max
}

func (s *Component) MaxHeight(max float64) {
	if s.GetHeight() > max {
		s.Height(max)
	}
	s.GetModel().MaxHeight = max
}

func (s *Component) GetMaxWidth() float64 {
	return s.GetModel().MaxWidth
}

func (s *Component) GetMaxHeight() float64 {
	return s.GetModel().MaxHeight
}

func (s *Component) GetExcludeFromLayout() bool {
	return s.GetModel().ExcludeFromLayout
}

func (s *Component) FlexWidth(value float64) {
	s.GetModel().FlexWidth = value
}

func (s *Component) FlexHeight(value float64) {
	s.GetModel().FlexHeight = value
}

func (s *Component) GetFlexWidth() float64 {
	return s.GetModel().FlexWidth
}

func (s *Component) GetFlexHeight() float64 {
	return s.GetModel().FlexHeight
}

func (s *Component) Padding(value float64) {
	s.GetModel().Padding = value
}

func (s *Component) PaddingBottom(value float64) {
	s.GetModel().PaddingBottom = value
}

func (s *Component) PaddingLeft(value float64) {
	s.GetModel().PaddingLeft = value
}

func (s *Component) PaddingRight(value float64) {
	s.GetModel().PaddingRight = value
}

func (s *Component) PaddingTop(value float64) {
	s.GetModel().PaddingTop = value
}

func (s *Component) GetPadding() float64 {
	return s.GetModel().Padding
}

func (s *Component) VAlign(value Alignment) {
	s.GetModel().VAlign = value
}

func (s *Component) GetHorizontalPadding() float64 {
	return s.GetPaddingLeft() + s.GetPaddingRight()
}

func (s *Component) GetVerticalPadding() float64 {
	return s.GetPaddingTop() + s.GetPaddingBottom()
}

func (s *Component) getPaddingForSide(getter func() float64) float64 {
	model := s.GetModel()
	if getter() == -1.0 {
		if model.Padding > -1.0 {
			return model.Padding
		}
		return -1.0
	}
	return getter()
}

func (s *Component) GetPaddingLeft() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetModel().PaddingLeft
	})
}

func (s *Component) GetPaddingRight() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetModel().PaddingRight
	})
}

func (s *Component) GetPaddingBottom() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetModel().PaddingBottom
	})
}

func (s *Component) GetPaddingTop() float64 {
	return s.getPaddingForSide(func() float64 {
		return s.GetModel().PaddingTop
	})
}

func (s *Component) setParent(parent Displayable) {
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

func (s *Component) GetYOffset() float64 {
	offset := s.GetY()
	parent := s.GetParent()
	if parent != nil {
		offset = offset + parent.GetYOffset()
	}
	return math.Max(0.0, offset)
}

func (s *Component) GetXOffset() float64 {
	offset := s.GetX()
	parent := s.GetParent()
	if parent != nil {
		offset = offset + parent.GetXOffset()
	}
	return math.Max(0.0, offset)
}

func (s *Component) GetPath() string {
	parent := s.GetParent()
	localPath := "/" + s.GetID()

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
	s.LayoutChildren()
	s.GetLayout()(s)
}

func (s *Component) DrawChildren(surface Surface) {

	childSurface := surface.GetOffsetSurfaceFor(s)
	for _, child := range s.children {
		// Create an surface delegate that includes an appropriate offset
		// for each child and send that to the Child's Draw() method.
		child.Draw(childSurface)
	}
}

func (s *Component) Draw(surface Surface) {
	// Note: Only doing this here while implementing layouts and styles.
	// Will eventually compose read-only views that pull values from the
	// Displayable and draw them onto a surface.
	//
	// Ordering here is important though, as children need to be drawn
	// over the parents (for now anyway). Eventually, we can probably get
	// smarter with not drawing fully occluded entities.
	DrawRectangle(surface, s)
	s.DrawChildren(surface)
}

func (s *Component) Text(text string) {
	s.GetModel().Text = text
}

func (s *Component) GetText() string {
	return s.GetModel().Text
}

func (s *Component) Title(title string) {
	s.GetModel().Title = title
}

func (s *Component) GetTitle() string {
	return s.GetModel().Title
}

/* STYLE ATTRIBUTES */

func (s *Component) BgColor(color int) {
	s.GetModel().BgColor = color
}

func (s *Component) FontFace(face string) {
	s.GetModel().FontFace = face
}

func (s *Component) FontSize(size int) {
	s.GetModel().FontSize = size
}

func (s *Component) GetBgColor() int {
	bgColor := s.GetModel().BgColor
	if bgColor == -1 {
		if s.parent != nil {
			return s.parent.GetBgColor()
		}
		return DefaultBgColor
	}
	return bgColor
}

func (s *Component) GetFontColor() int {
	fontColor := s.GetModel().FontColor
	if fontColor == -1 {
		if s.parent != nil {
			return s.parent.GetFontColor()
		}
		return DefaultFontColor
	}
	return fontColor
}

func (s *Component) FontColor(size int) {
	s.GetModel().FontColor = size
}

func (s *Component) GetFontFace() string {
	fontFace := s.GetModel().FontFace
	if fontFace == "" {
		if s.parent != nil {
			return s.parent.GetFontFace()
		}
		return DefaultFontFace
	}
	return fontFace
}

func (s *Component) GetFontSize() int {
	fontSize := s.GetModel().FontSize
	if fontSize == -1 {
		if s.parent != nil {
			return s.parent.GetFontSize()
		}
		return DefaultFontSize
	}
	return fontSize
}

func (s *Component) GetStrokeColor() int {
	strokeColor := s.GetModel().StrokeColor
	if strokeColor == -1 {
		if s.parent != nil {
			return s.parent.GetStrokeColor()
		}
		return DefaultStrokeColor
	}
	return strokeColor
}

func (s *Component) StrokeColor(size int) {
	s.GetModel().StrokeColor = size
}

func (s *Component) GetStrokeSize() int {
	strokeSize := s.GetModel().StrokeSize
	if strokeSize == -1 {
		if s.parent != nil {
			return s.parent.GetStrokeSize()
		}
		return DefaultStrokeSize
	}
	return strokeSize
}

func (s *Component) StrokeSize(size int) {
	s.GetModel().StrokeSize = size
}

// NewComponent returns a new base component instance as a Displayable.
func NewComponent() Displayable {
	return &Component{}
}
