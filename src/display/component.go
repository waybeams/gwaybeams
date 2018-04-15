package display

import (
	"clock"
	"errors"
	"fmt"
	"log"
	"math"
)

const DefaultBgColor = 0x999999ff
const DefaultFontColor = 0x111111ff
const DefaultFontSize = 12
const DefaultFontFace = "sans"
const DefaultStrokeColor = 0x333333ff
const DefaultStrokeSize = 1

type TraitOptions map[string][]ComponentOption

// Component is a concrete base component implementation made public for
// composition, not instantiation.
type Component struct {
	EmitterBase

	builder                        Builder
	children                       []Displayable
	parent                         Displayable
	model                          *ComponentModel
	composeEmpty                   func()
	composeWithBuilder             func(Builder)
	composeWithComponent           func(Displayable)
	composeWithBuilderAndComponent func(Builder, Displayable)
	dirtyNodes                     []Displayable
	traitOptions                   TraitOptions
	view                           RenderHandler
}

func (c *Component) ID() string {
	return c.Model().ID
}

func (c *Component) Key() string {
	key := c.Model().Key

	if key != "" {
		return key
	}

	return c.ID()
}

func (c *Component) Bubble(event Event) {
	c.Emit(event)

	current := c.Parent()
	for current != nil {
		if event.IsCancelled() {
			return
		}
		current.Emit(event)
		current = current.Parent()
	}
}

func (c *Component) SetTypeName(name string) {
	c.Model().TypeName = name
}

func (c *Component) TypeName() string {
	return c.Model().TypeName
}

func (c *Component) SetTraitNames(names ...string) {
	c.Model().TraitNames = names
}

func (c *Component) TraitNames() []string {
	return c.Model().TraitNames
}

// Root returns a outermost Displayable in the current tree.
func (c *Component) Root() Displayable {
	parent := c.Parent()
	if parent != nil {
		return parent.Root()
	}
	return c
}

func (c *Component) Invalidate() {
	// NOTE(lbayes): This is not desired behavior, but it's what we've got right now.
	if c.Parent() != nil {
		c.Parent().InvalidateChildren()
	}
}

func (c *Component) InvalidateChildren() {
	c.InvalidateChildrenFor(c)
}

// TODO(lbayes): Rename to something less confusing
func (c *Component) RecomposeChildren() []Displayable {
	nodes := c.InvalidNodes()
	b := c.Builder()
	for _, node := range nodes {
		err := b.UpdateChildren(node)
		if err != nil {
			panic(err)
		}
		// node.Layout()
	}
	c.dirtyNodes = []Displayable{}
	return nodes
}

func (c *Component) InvalidateChildrenFor(d Displayable) {
	// Late binding to find root at the time of invalidation.
	if c.Parent() != nil {
		c.Root().InvalidateChildrenFor(d)
		return
	}
	c.dirtyNodes = append(c.dirtyNodes, d)
}

func (c *Component) ShouldRecompose() bool {
	return len(c.dirtyNodes) > 0
}

func (c *Component) InvalidNodes() []Displayable {
	nodes := c.dirtyNodes
	results := []Displayable{}
	for nIndex, node := range nodes {
		ancestorFound := false
		for aIndex, possibleAncestor := range nodes {
			if aIndex != nIndex && node.IsContainedBy(possibleAncestor) {
				ancestorFound = true
				break
			}
		}
		if !ancestorFound {
			results = append(results, node)
		}
	}

	return results
}

func (c *Component) PushTrait(selector string, opts ...ComponentOption) error {
	traitOptions := c.TraitOptions()
	if traitOptions[selector] != nil {
		return errors.New("duplicate trait selector found with:" + selector)
	}
	traitOptions[selector] = opts
	return nil
}

func (c *Component) TraitOptions() TraitOptions {
	if c.traitOptions == nil {
		c.traitOptions = make(map[string][]ComponentOption)
	}
	return c.traitOptions
}

func (c *Component) Composer(composer interface{}) error {
	// Ensure we do not already have a compose function assigned
	if composer != nil &&
		(c.composeEmpty != nil ||
			c.composeWithBuilder != nil ||
			c.composeWithBuilderAndComponent != nil ||
			c.composeWithComponent != nil) {
		return errors.New("Components can only accept a single Compose function")
	}

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

func (c *Component) IsContainedBy(node Displayable) bool {
	current := c.Parent()
	for current != nil {
		if current == node {
			return true
		}
		current = current.Parent()
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

func (c *Component) SetLayoutType(layoutType LayoutTypeValue) {
	c.Model().LayoutType = layoutType
}

func (c *Component) LayoutType() LayoutTypeValue {
	return c.Model().LayoutType
}

func (c *Component) Layout() {
	c.GetLayout()(c)
	c.LayoutChildren()
}

func (c *Component) LayoutChildren() {
	for _, child := range c.Children() {
		child.Layout()
	}
}

// NOTE(lbayes): There's a naming conflict. Layout() is used above as a verb
// and here as a noun.
func (c *Component) GetLayout() LayoutHandler {
	switch c.LayoutType() {
	case StackLayoutType:
		return StackLayout
	case HorizontalFlowLayoutType:
		return HorizontalFlowLayout
	case VerticalFlowLayoutType:
		return VerticalFlowLayout
	default:
		log.Fatal("ERROR: Requested LayoutTypeValue (%v) is not supported:", c.LayoutType())
		return nil
	}
}

func (c *Component) SetModel(model *ComponentModel) {
	c.model = model
}

func (c *Component) Model() *ComponentModel {
	if c.model == nil {
		c.model = &ComponentModel{}
	}
	return c.model
}

func (c *Component) SetX(x float64) {
	c.Model().X = x
}

func (c *Component) SetY(y float64) {
	c.Model().Y = y
}

func (c *Component) SetZ(z float64) {
	c.Model().Z = z
}

func (c *Component) X() float64 {
	return c.Model().X
}

func (c *Component) Y() float64 {
	return c.Model().Y
}

func (c *Component) Z() float64 {
	return c.Model().Z
}

func (c *Component) SetHAlign(value Alignment) {
	c.Model().HAlign = value
}

func (c *Component) HAlign() Alignment {
	return c.Model().HAlign
}

func (c *Component) VAlign() Alignment {
	return c.Model().VAlign
}

func (c *Component) SetVAlign(value Alignment) {
	c.Model().VAlign = value
}

func (c *Component) SetWidth(w float64) {
	model := c.Model()
	if model.Width != w {
		model.Width = -1
		c.SetActualWidth(w)
	}
}

func (c *Component) SetHeight(h float64) {
	model := c.Model()
	if model.Height != h {
		model.Height = -1
		c.SetActualHeight(h)
	}
}

func (c *Component) WidthInBounds(w float64) float64 {
	min := c.MinWidth()
	max := c.MaxWidth()
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
	min := c.MinHeight()
	max := c.MaxHeight()

	height := math.Round(h)

	if min > -1 {
		height = math.Max(min, height)
	}

	if max > -1 {
		height = math.Min(max, height)
	}
	return height
}

func (c *Component) Width() float64 {
	model := c.Model()
	if model.ActualWidth == -1 {
		prefWidth := c.PrefWidth()
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

func (c *Component) Height() float64 {
	model := c.Model()
	if model.ActualHeight == -1 {
		prefHeight := c.PrefHeight()
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

func (c *Component) FixedWidth() float64 {
	return c.Model().Width
}

func (c *Component) FixedHeight() float64 {
	return c.Model().Height
}

func (c *Component) SetPrefWidth(value float64) {
	c.Model().PrefWidth = value
}

func (c *Component) SetPrefHeight(value float64) {
	c.Model().PrefHeight = value
}

func (c *Component) PrefWidth() float64 {
	return c.Model().PrefWidth
}

func (c *Component) PrefHeight() float64 {
	return c.Model().PrefHeight
}

func (c *Component) SetActualWidth(width float64) {
	inBounds := c.WidthInBounds(width)
	model := c.Model()
	model.ActualWidth = inBounds
	if model.Width != -1 && model.Width != width {
		model.Width = width
	}
}

func (c *Component) SetActualHeight(height float64) {
	inBounds := c.HeightInBounds(height)
	model := c.Model()
	model.ActualHeight = inBounds
	if model.Height != -1 && model.Height != height {
		model.Height = height
	}
}

func (c *Component) InferredMinWidth() float64 {
	result := 0.0
	for _, child := range c.Children() {
		if !child.ExcludeFromLayout() {
			result = math.Max(result, child.MinWidth())
		}
	}
	return result + c.HorizontalPadding()
}

func (c *Component) InferredMinHeight() float64 {
	result := 0.0
	for _, child := range c.Children() {
		if !child.ExcludeFromLayout() {
			result = math.Max(result, child.MinHeight())
		}
	}
	return result + c.HorizontalPadding()
}

func (c *Component) SetExcludeFromLayout(value bool) {
	c.Model().ExcludeFromLayout = value
}

func (c *Component) ActualWidth() float64 {
	model := c.Model()

	if model.Width > -1 {
		return model.Width
	} else if model.ActualWidth > -1 {
		return model.ActualWidth
	}
	prefWidth := c.PrefWidth()
	if prefWidth > -1 {
		return prefWidth
	}

	return c.MinWidth()
}

func (c *Component) ActualHeight() float64 {
	model := c.Model()

	if model.Height > -1 {
		return model.Height
	} else if model.ActualHeight > -1 {
		return model.ActualHeight
	}
	prefHeight := c.PrefHeight()
	if prefHeight > -1 {
		return prefHeight
	}

	return c.MinHeight()
}

func (c *Component) SetMinWidth(min float64) {
	c.Model().MinWidth = min
	// Ensure we're not already too small for the new min
	if c.ActualWidth() < min {
		c.SetActualWidth(min)
	}
}

func (c *Component) SetMinHeight(min float64) {
	c.Model().MinHeight = min
	// Ensure we're not already too small for the new min
	if c.ActualHeight() < min {
		c.SetActualHeight(min)
	}
}

func (c *Component) MinWidth() float64 {
	model := c.Model()
	width := model.Width
	minWidth := model.MinWidth
	result := -1.0

	if width > -1.0 {
		result = width
	}
	if minWidth > -1.0 {
		result = minWidth
	}

	inferredMinWidth := c.InferredMinWidth()
	if inferredMinWidth > 0 {
		return math.Max(result, inferredMinWidth)
	}
	return result
}

func (c *Component) MinHeight() float64 {
	model := c.Model()
	height := model.Height
	minHeight := model.MinHeight
	result := -1.0

	if height > -1.0 {
		result = height
	}
	if minHeight > -1.0 {
		result = minHeight
	}

	inferredMinHeight := c.InferredMinHeight()
	if inferredMinHeight > 0.0 {
		return math.Max(result, inferredMinHeight)
	}
	return result
}

func (c *Component) SetMaxWidth(max float64) {
	if c.Width() > max {
		c.SetWidth(max)
	}
	c.Model().MaxWidth = max
}

func (c *Component) SetMaxHeight(max float64) {
	if c.Height() > max {
		c.SetHeight(max)
	}
	c.Model().MaxHeight = max
}

func (c *Component) MaxWidth() float64 {
	return c.Model().MaxWidth
}

func (c *Component) MaxHeight() float64 {
	return c.Model().MaxHeight
}

func (c *Component) ExcludeFromLayout() bool {
	return c.Model().ExcludeFromLayout
}

func (c *Component) SetFlexWidth(value float64) {
	c.Model().FlexWidth = value
}

func (c *Component) SetFlexHeight(value float64) {
	c.Model().FlexHeight = value
}

func (c *Component) FlexWidth() float64 {
	return c.Model().FlexWidth
}

func (c *Component) FlexHeight() float64 {
	return c.Model().FlexHeight
}

func (c *Component) SetPadding(value float64) {
	c.Model().Padding = value
}

func (c *Component) SetPaddingBottom(value float64) {
	c.Model().PaddingBottom = value
}

func (c *Component) SetPaddingLeft(value float64) {
	c.Model().PaddingLeft = value
}

func (c *Component) SetPaddingRight(value float64) {
	c.Model().PaddingRight = value
}

func (c *Component) SetPaddingTop(value float64) {
	c.Model().PaddingTop = value
}

func (c *Component) Padding() float64 {
	return c.Model().Padding
}

func (c *Component) HorizontalPadding() float64 {
	return c.PaddingLeft() + c.PaddingRight()
}

func (c *Component) VerticalPadding() float64 {
	return c.PaddingTop() + c.PaddingBottom()
}

func (c *Component) getPaddingForSide(getter func() float64) float64 {
	model := c.Model()
	if getter() == -1.0 {
		if model.Padding > -1.0 {
			return model.Padding
		}
		return -1.0
	}
	return getter()
}

func (c *Component) PaddingLeft() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingLeft
	})
}

func (c *Component) PaddingRight() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingRight
	})
}

func (c *Component) PaddingBottom() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingBottom
	})
}

func (c *Component) PaddingTop() float64 {
	return c.getPaddingForSide(func() float64 {
		return c.Model().PaddingTop
	})
}

func (c *Component) SetParent(parent Displayable) {
	c.parent = parent
}

func (c *Component) AddChild(child Displayable) int {
	c.children = append(c.Children(), child)
	child.SetParent(c)
	return len(c.children)
}

func (c *Component) SetBuilder(b Builder) {
	// NOTE(lbayes): This method is called on temporary components
	// that are never going to be added to a valid tree. Be sure
	// we do not mutate the state of the Builder for any reason
	// from this call.
	c.builder = b
}

func (c *Component) Builder() Builder {
	if c.parent != nil {
		return c.parent.Builder()
	}
	return c.builder
}

func (c *Component) Clock() clock.Clock {
	return c.Builder().Clock()
}

func (c *Component) ChildCount() int {
	return len(c.Children())
}

func (c *Component) FindComponentByID(id string) Displayable {
	if id == c.ID() {
		return c
	}
	for _, child := range c.Children() {
		result := child.FindComponentByID(id)
		if result != nil {
			return result
		}
	}
	return nil
}

func (c *Component) RemoveChild(toRemove Displayable) int {
	children := c.Children()
	for index, child := range children {
		if child == toRemove {
			c.children = append(children[:index], children[index+1:]...)
			c.onChildRemoved(child)
			return index
		}
	}

	return -1
}

func (c *Component) RemoveAllChildren() {
	children := c.Children()
	c.children = make([]Displayable, 0)
	for _, child := range children {
		c.onChildRemoved(child)
	}
}

func (c *Component) onChildRemoved(child Displayable) {
	child.SetParent(nil)
}

func (c *Component) ChildAt(index int) Displayable {
	return c.Children()[index]
}

func (c *Component) Children() []Displayable {
	if c.children == nil {
		c.children = make([]Displayable, 0)
	}
	return c.children
}

func (c *Component) GetFilteredChildren(filter DisplayableFilter) []Displayable {
	result := []Displayable{}
	kids := c.Children()
	for _, child := range kids {
		if filter(child) {
			result = append(result, child)
		}
	}
	return result
}

func (c *Component) YOffset() float64 {
	offset := c.Y()
	parent := c.Parent()
	if parent != nil {
		offset = offset + parent.YOffset()
	}
	return math.Max(0.0, offset)
}

func (c *Component) XOffset() float64 {
	offset := c.X()
	parent := c.Parent()
	if parent != nil {
		offset = offset + parent.XOffset()
	}
	return math.Max(0.0, offset)
}

func (c *Component) pathPart() string {
	// Try ID first
	id := c.ID()
	if id != "" {
		return c.ID()
	}

	// Empty ID, try Key
	key := c.Key()
	if key != "" {
		return c.Key()
	}

	parent := c.Parent()
	if parent != nil {
		siblings := parent.Children()
		for index, child := range siblings {
			if child == c {
				return fmt.Sprintf("%v%v", c.TypeName(), index)
			}
		}
	}

	// Empty ID and Key, and Parent just use TypeName
	return c.TypeName()
}

func (c *Component) Path() string {
	parent := c.Parent()
	localPath := "/" + c.pathPart()

	if parent != nil {
		return parent.Path() + localPath
	}
	return localPath

}

func (c *Component) Parent() Displayable {
	return c.parent
}

func (c *Component) SetView(view RenderHandler) {
	c.view = view
}

func (c *Component) View() RenderHandler {
	if c.view == nil {
		return c.GetDefaultView()
	}
	return c.view
}

func (c *Component) GetDefaultView() RenderHandler {
	return RectangleView
}

func (c *Component) DrawChildren(surface Surface) {
	for _, child := range c.Children() {
		// Create an surface delegate that includes an appropriate offset
		// for each child and send that to the Child's Draw() method.
		child.Draw(surface)
	}
}

func (c *Component) Draw(surface Surface) {
	local := surface.GetOffsetSurfaceFor(c)
	c.View()(local, c)
	c.DrawChildren(surface)
}

func (c *Component) SetText(text string) {
	c.Model().Text = text
}

func (c *Component) Text() string {
	return c.Model().Text
}

func (c *Component) SetTitle(title string) {
	c.Model().Title = title
}

func (c *Component) Title() string {
	return c.Model().Title
}

/* STYLE ATTRIBUTES */

func (c *Component) SetBgColor(color int) {
	c.Model().BgColor = color
}

func (c *Component) SetFontFace(face string) {
	c.Model().FontFace = face
}

func (c *Component) SetFontSize(size int) {
	c.Model().FontSize = size
}

func (c *Component) BgColor() int {
	bgColor := c.Model().BgColor
	if bgColor == -1 {
		if c.parent != nil {
			return c.parent.BgColor()
		}
		return DefaultBgColor
	}
	return bgColor
}

func (c *Component) FontColor() int {
	fontColor := c.Model().FontColor
	if fontColor == -1 {
		if c.parent != nil {
			return c.parent.FontColor()
		}
		return DefaultFontColor
	}
	return fontColor
}

func (c *Component) SetFontColor(size int) {
	c.Model().FontColor = size
}

func (c *Component) FontFace() string {
	fontFace := c.Model().FontFace
	if fontFace == "" {
		if c.parent != nil {
			return c.parent.FontFace()
		}
		return DefaultFontFace
	}
	return fontFace
}

func (c *Component) FontSize() int {
	fontSize := c.Model().FontSize
	if fontSize == -1 {
		if c.parent != nil {
			return c.parent.FontSize()
		}
		return DefaultFontSize
	}
	return fontSize
}

func (c *Component) StrokeColor() int {
	strokeColor := c.Model().StrokeColor
	if strokeColor == -1 {
		if c.parent != nil {
			return c.parent.StrokeColor()
		}
		return DefaultStrokeColor
	}
	return strokeColor
}

func (c *Component) SetStrokeColor(size int) {
	c.Model().StrokeColor = size
}

func (c *Component) StrokeSize() int {
	strokeSize := c.Model().StrokeSize
	if strokeSize == -1 {
		if c.parent != nil {
			return c.parent.StrokeSize()
		}
		return DefaultStrokeSize
	}
	return strokeSize
}

func (c *Component) SetStrokeSize(size int) {
	c.Model().StrokeSize = size
}

// NewComponent returns a new base component instance as a Displayable.
func NewComponent() Displayable {
	return &Component{}
}
