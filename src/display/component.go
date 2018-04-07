package display

import (
	"errors"
	"fmt"
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
	builder                        Builder
	children                       []Displayable
	parent                         Displayable
	model                          *ComponentModel
	composeEmpty                   func()
	composeWithBuilder             func(Builder)
	composeWithComponent           func(Displayable)
	composeWithBuilderAndComponent func(Builder, Displayable)
	traitOptions                   TraitOptions
	view                           RenderHandler
}

func (c *Component) GetID() string {
	model := c.GetModel()
	if model.ID == "" {
		model.ID = xid.New().String()
	}

	return model.ID
}

func (c *Component) GetTypeName() string {
	return c.GetModel().TypeName
}

func (c *Component) TypeName(name string) {
	c.GetModel().TypeName = name
}

func (c *Component) Invalidate() {
}

func (c *Component) PushTrait(selector string, opts ...ComponentOption) error {
	traitOptions := c.GetTraitOptions()
	if traitOptions[selector] != nil {
		return errors.New("duplicate trait selector found with:" + selector)
	}
	traitOptions[selector] = opts
	return nil
}

func (c *Component) GetTraitOptions() TraitOptions {
	if c.traitOptions == nil {
		c.traitOptions = make(map[string][]ComponentOption)
	}
	return c.traitOptions
}

func (c *Component) Composer(composer interface{}) error {
	switch composer.(type) {
	case func():
		c.composeEmpty = composer.(func())
	case func(Builder):
		c.composeWithBuilder = composer.(func(Builder))
	case func(Displayable):
		c.composeWithComponent = composer.(func(Displayable))
	case func(Builder, Displayable):
		c.composeWithBuilderAndComponent = composer.(func(Builder, Displayable))
	default:
		return errors.New("Component.Composer() called with unexpected signature")
	}
	return nil
}

func (c *Component) GetIsContainedBy(node Displayable) bool {
	current := c.GetParent()
	for current != nil {
		if current == node {
			return true
		}
		current = current.GetParent()
	}

	return false
}

func (c *Component) GetComposeEmpty() func() {
	return c.composeEmpty
}

func (c *Component) GetComposeWithBuilder() func(Builder) {
	return c.composeWithBuilder
}

func (c *Component) GetComposeWithComponent() func(Displayable) {
	return c.composeWithComponent
}

func (c *Component) GetComposeWithBuilderAndComponent() func(Builder, Displayable) {
	return c.composeWithBuilderAndComponent
}

func (c *Component) LayoutType(layoutType LayoutTypeValue) {
	c.GetModel().LayoutType = layoutType
}

func (c *Component) GetLayoutType() LayoutTypeValue {
	return c.GetModel().LayoutType
}

func (c *Component) GetLayout() LayoutHandler {
	switch c.GetLayoutType() {
	case StackLayoutType:
		return StackLayout
	case HorizontalFlowLayoutType:
		return HorizontalFlowLayout
	case VerticalFlowLayoutType:
		return VerticalFlowLayout
	default:
		log.Printf("ERROR: Requested LayoutTypeValue (%v) is not supported", c.GetLayoutType())
		return nil
	}
}

func (c *Component) Model(model *ComponentModel) {
	c.model = model
}

func (c *Component) GetModel() *ComponentModel {
	if c.model == nil {
		c.model = &ComponentModel{}
	}
	return c.model
}

func (c *Component) X(x float64) {
	c.GetModel().X = math.Round(x)
}

func (c *Component) GetX() float64 {
	return c.GetModel().X
}

func (c *Component) Y(y float64) {
	c.GetModel().Y = math.Round(y)
}

func (c *Component) Z(z float64) {
	c.GetModel().Z = math.Round(z)
}

func (c *Component) GetY() float64 {
	return c.GetModel().Y
}

func (c *Component) GetZ() float64 {
	return c.GetModel().Z
}

func (c *Component) HAlign(value Alignment) {
	c.GetModel().HAlign = value
}

func (c *Component) Width(w float64) {
	model := c.GetModel()
	if model.Width != w {
		model.Width = -1
		c.ActualWidth(w)
	}
}

func (c *Component) Height(h float64) {
	model := c.GetModel()
	if model.Height != h {
		model.Height = -1
		c.ActualHeight(h)
	}
}

func (c *Component) WidthInBounds(w float64) float64 {
	min := c.GetMinWidth()
	max := c.GetMaxWidth()
	width := w

	if min > -1 {
		width = math.Max(min, width)
	}

	if max > -1 {
		width = math.Min(max, width)
	}
	return width
}

func (c *Component) HeightInBounds(h float64) float64 {
	min := c.GetMinHeight()
	max := c.GetMaxHeight()

	height := math.Round(h)

	if min > -1 {
		height = math.Max(min, height)
	}

	if max > -1 {
		height = math.Min(max, height)
	}
	return height
}

func (c *Component) GetWidth() float64 {
	model := c.GetModel()
	if model.ActualWidth == -1 {
		prefWidth := c.GetPrefWidth()
		if prefWidth > -1 {
			return prefWidth
		}
		inBounds := c.WidthInBounds(model.Width)
		if inBounds > -1.0 {
			return inBounds
		}
		return 0
	}
	return model.ActualWidth
}

func (c *Component) GetHeight() float64 {
	model := c.GetModel()
	if model.ActualHeight == -1 {
		prefHeight := c.GetPrefHeight()
		if prefHeight > -1 {
			return prefHeight
		}
		inBounds := c.HeightInBounds(model.Height)
		if inBounds > -1 {
			return inBounds
		}
		return 0
	}
	return model.ActualHeight
}

func (c *Component) GetFixedWidth() float64 {
	return c.GetModel().Width
}

func (c *Component) GetFixedHeight() float64 {
	return c.GetModel().Height
}

func (c *Component) PrefWidth(value float64) {
	c.GetModel().PrefWidth = value
}

func (c *Component) PrefHeight(value float64) {
	c.GetModel().PrefHeight = value
}

func (c *Component) GetPrefWidth() float64 {
	return c.GetModel().PrefWidth
}

func (c *Component) GetPrefHeight() float64 {
	return c.GetModel().PrefHeight
}

func (c *Component) ActualWidth(width float64) {
	inBounds := c.WidthInBounds(width)
	model := c.GetModel()
	model.ActualWidth = inBounds
	if model.Width != -1 && model.Width != width {
		model.Width = width
	}
}

func (c *Component) ActualHeight(height float64) {
	inBounds := c.HeightInBounds(height)
	model := c.GetModel()
	model.ActualHeight = inBounds
	if model.Height != -1 && model.Height != height {
		model.Height = height
	}
}

func (c *Component) GetInferredMinWidth() float64 {
	result := 0.0
	for _, child := range c.children {
		if !child.GetExcludeFromLayout() {
			result = math.Max(result, child.GetMinWidth())
		}
	}
	return result + c.GetHorizontalPadding()
}

func (c *Component) GetInferredMinHeight() float64 {
	result := 0.0
	for _, child := range c.children {
		if !child.GetExcludeFromLayout() {
			result = math.Max(result, child.GetMinHeight())
		}
	}
	return result + c.GetHorizontalPadding()
}

func (c *Component) ExcludeFromLayout(value bool) {
	c.GetModel().ExcludeFromLayout = value
}

func (c *Component) GetActualWidth() float64 {
	model := c.GetModel()

	if model.Width > -1 {
		return model.Width
	} else if model.ActualWidth > -1 {
		return model.ActualWidth
	}
	prefWidth := c.GetPrefWidth()
	if prefWidth > -1 {
		return prefWidth
	}

	return c.GetMinWidth()
}

func (c *Component) GetActualHeight() float64 {
	model := c.GetModel()

	if model.Height > -1 {
		return model.Height
	} else if model.ActualHeight > -1 {
		return model.ActualHeight
	}
	prefHeight := c.GetPrefHeight()
	if prefHeight > -1 {
		return prefHeight
	}

	return c.GetMinHeight()
}

func (c *Component) GetHAlign() Alignment {
	return c.GetModel().HAlign
}

func (c *Component) GetVAlign() Alignment {
	return c.GetModel().VAlign
}

func (c *Component) MinWidth(min float64) {
	c.GetModel().MinWidth = min
	// Ensure we're not already too small for the new min
	if c.GetActualWidth() < min {
		c.ActualWidth(min)
	}
}

func (c *Component) MinHeight(min float64) {
	c.GetModel().MinHeight = min
	// Ensure we're not already too small for the new min
	if c.GetActualHeight() < min {
		c.ActualHeight(min)
	}
}

func (c *Component) GetMinWidth() float64 {
	model := c.GetModel()
	width := model.Width
	minWidth := model.MinWidth
	result := -1.0

	if width > -1.0 {
		result = width
	}
	if minWidth > -1.0 {
		result = minWidth
	}

	inferredMinWidth := c.GetInferredMinWidth()
	if inferredMinWidth > 0 {
		return math.Max(result, inferredMinWidth)
	}
	return result
}

func (c *Component) GetMinHeight() float64 {
	model := c.GetModel()
	height := model.Height
	minHeight := model.MinHeight
	result := -1.0

	if height > -1.0 {
		result = height
	}
	if minHeight > -1.0 {
		result = minHeight
	}

	inferredMinHeight := c.GetInferredMinHeight()
	if inferredMinHeight > 0.0 {
		return math.Max(result, inferredMinHeight)
	}
	return result
}

func (c *Component) MaxWidth(max float64) {
	if c.GetWidth() > max {
		c.Width(max)
	}
	c.GetModel().MaxWidth = max
}

func (c *Component) MaxHeight(max float64) {
	if c.GetHeight() > max {
		c.Height(max)
	}
	c.GetModel().MaxHeight = max
}

func (c *Component) GetMaxWidth() float64 {
	return c.GetModel().MaxWidth
}

func (c *Component) GetMaxHeight() float64 {
	return c.GetModel().MaxHeight
}

func (c *Component) GetExcludeFromLayout() bool {
	return c.GetModel().ExcludeFromLayout
}

func (c *Component) FlexWidth(value float64) {
	c.GetModel().FlexWidth = value
}

func (c *Component) FlexHeight(value float64) {
	c.GetModel().FlexHeight = value
}

func (c *Component) GetFlexWidth() float64 {
	return c.GetModel().FlexWidth
}

func (c *Component) GetFlexHeight() float64 {
	return c.GetModel().FlexHeight
}

func (c *Component) Padding(value float64) {
	c.GetModel().Padding = value
}

func (c *Component) PaddingBottom(value float64) {
	c.GetModel().PaddingBottom = value
}

func (c *Component) PaddingLeft(value float64) {
	c.GetModel().PaddingLeft = value
}

func (c *Component) PaddingRight(value float64) {
	c.GetModel().PaddingRight = value
}

func (c *Component) PaddingTop(value float64) {
	c.GetModel().PaddingTop = value
}

func (c *Component) GetPadding() float64 {
	return c.GetModel().Padding
}

func (c *Component) VAlign(value Alignment) {
	c.GetModel().VAlign = value
}

func (c *Component) GetHorizontalPadding() float64 {
	return c.GetPaddingLeft() + c.GetPaddingRight()
}

func (c *Component) GetVerticalPadding() float64 {
	return c.GetPaddingTop() + c.GetPaddingBottom()
}

func (c *Component) getPaddingForSide(getter func() float64) float64 {
	model := c.GetModel()
	if getter() == -1.0 {
		if model.Padding > -1.0 {
			return model.Padding
		}
		return -1.0
	}
	return getter()
}

func (c *Component) GetPaddingLeft() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.GetModel().PaddingLeft
	})
}

func (c *Component) GetPaddingRight() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.GetModel().PaddingRight
	})
}

func (c *Component) GetPaddingBottom() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.GetModel().PaddingBottom
	})
}

func (c *Component) GetPaddingTop() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.GetModel().PaddingTop
	})
}

func (c *Component) setParent(parent Displayable) {
	c.parent = parent
}

func (c *Component) AddChild(child Displayable) int {
	if c.children == nil {
		c.children = make([]Displayable, 0)
	}

	c.children = append(c.children, child)
	child.setParent(c)
	return len(c.children)
}

func (c *Component) GetBuilder() Builder {
	if c.parent != nil {
		return c.parent.GetBuilder()
	}
	return c.builder
}

func (c *Component) GetChildCount() int {
	return len(c.children)
}

func (c *Component) GetChildAt(index int) Displayable {
	return c.children[index]
}

func (c *Component) GetChildren() []Displayable {
	return append([]Displayable{}, c.children...)
}

func (c *Component) GetFilteredChildren(filter DisplayableFilter) []Displayable {
	result := make([]Displayable, 0)
	for _, child := range c.children {
		if filter(child) {
			result = append(result, child)
		}
	}
	return result
}

func (c *Component) GetYOffset() float64 {
	offset := c.GetY()
	parent := c.GetParent()
	if parent != nil {
		offset = offset + parent.GetYOffset()
	}
	return math.Max(0.0, offset)
}

func (c *Component) GetXOffset() float64 {
	offset := c.GetX()
	parent := c.GetParent()
	if parent != nil {
		offset = offset + parent.GetXOffset()
	}
	return math.Max(0.0, offset)
}

func (c *Component) GetPath() string {
	parent := c.GetParent()
	localPath := "/" + c.GetID()

	if parent != nil {
		return parent.GetPath() + localPath
	}
	return localPath

}

func (c *Component) GetParent() Displayable {
	return c.parent
}

func (c *Component) LayoutChildren() {
	for _, child := range c.children {
		child.Layout()
	}
}

func (c *Component) Layout() {
	c.LayoutChildren()
	c.GetLayout()(c)
}

func (c *Component) View(view RenderHandler) {
	c.view = view
}

func (c *Component) GetView() RenderHandler {
	if c.view == nil {
		return c.GetDefaultView()
	}
	return c.view
}

func (c *Component) GetDefaultView() RenderHandler {
	return RectangleView

}

func (c *Component) DrawChildren(surface Surface) {

	childSurface := surface.GetOffsetSurfaceFor(c)
	for _, child := range c.children {
		// Create an surface delegate that includes an appropriate offset
		// for each child and send that to the Child's Draw() method.
		child.Draw(childSurface)
	}
}

func (c *Component) Draw(surface Surface) {
	fmt.Println("DRAW NOW")
	c.GetView()(surface, c)
	c.DrawChildren(surface)
}

func (c *Component) Text(text string) {
	c.GetModel().Text = text
}

func (c *Component) GetText() string {
	return c.GetModel().Text
}

func (c *Component) Title(title string) {
	c.GetModel().Title = title
}

func (c *Component) GetTitle() string {
	return c.GetModel().Title
}

/* STYLE ATTRIBUTES */

func (c *Component) BgColor(color int) {
	c.GetModel().BgColor = color
}

func (c *Component) FontFace(face string) {
	c.GetModel().FontFace = face
}

func (c *Component) FontSize(size int) {
	c.GetModel().FontSize = size
}

func (c *Component) GetBgColor() int {
	bgColor := c.GetModel().BgColor
	if bgColor == -1 {
		if c.parent != nil {
			return c.parent.GetBgColor()
		}
		return DefaultBgColor
	}
	return bgColor
}

func (c *Component) GetFontColor() int {
	fontColor := c.GetModel().FontColor
	if fontColor == -1 {
		if c.parent != nil {
			return c.parent.GetFontColor()
		}
		return DefaultFontColor
	}
	return fontColor
}

func (c *Component) FontColor(size int) {
	c.GetModel().FontColor = size
}

func (c *Component) GetFontFace() string {
	fontFace := c.GetModel().FontFace
	if fontFace == "" {
		if c.parent != nil {
			return c.parent.GetFontFace()
		}
		return DefaultFontFace
	}
	return fontFace
}

func (c *Component) GetFontSize() int {
	fontSize := c.GetModel().FontSize
	if fontSize == -1 {
		if c.parent != nil {
			return c.parent.GetFontSize()
		}
		return DefaultFontSize
	}
	return fontSize
}

func (c *Component) GetStrokeColor() int {
	strokeColor := c.GetModel().StrokeColor
	if strokeColor == -1 {
		if c.parent != nil {
			return c.parent.GetStrokeColor()
		}
		return DefaultStrokeColor
	}
	return strokeColor
}

func (c *Component) StrokeColor(size int) {
	c.GetModel().StrokeColor = size
}

func (c *Component) GetStrokeSize() int {
	strokeSize := c.GetModel().StrokeSize
	if strokeSize == -1 {
		if c.parent != nil {
			return c.parent.GetStrokeSize()
		}
		return DefaultStrokeSize
	}
	return strokeSize
}

func (c *Component) StrokeSize(size int) {
	c.GetModel().StrokeSize = size
}

func (c *Component) OnClick(handler EventHandler) {
	// TODO(lbayes): Design event system, rather than just callbacks
}

func (c *Component) Click() {
	// Trigger OnClick handler(s)
}

// NewComponent returns a new base component instance as a Displayable.
func NewComponent() Displayable {
	return &Component{}
}
