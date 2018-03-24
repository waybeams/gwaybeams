package display

type DisplayableFilter = func(Displayable) bool

type Alignment int

const (
	BottomAlign = iota
	LeftAlign
	RightAlign
	TopAlign
)

// This pattern is probably not the way to go, but I'm having trouble finding a
// reasonable alternative. The problem here is that Layout types will not be
// user-extensible. Box definitions will only be able to refer to the
// Layouts that have been enumerated here. The benefit is that ComponentModel objects
// will remain serializable and simply be a bag of scalars. I'm definitely
// open to suggestions.
type LayoutType int

const (
	// GROSS! I'm sure I've done something wrong here, but the "zero value" for
	// an enum field (check ComponentModel) is 0. This means that not setting the enum will
	// automatically set it to the first value in this list. :barf:
	// DO NOT SORT THESE ALPHABETICALLY!
	StackLayoutType = iota
	// DO NOT SORT
	VFlowLayoutType
	HFlowLayoutType
	RowLayoutType
)

type LayoutDirection int

const (
	Horizontal = iota
	Vertical
)

type Layout func(d Displayable)

type Composable interface {
	Composer(composeFunc interface{}) error
	GetId() string
	GetComposeSimple() func()
	GetComposeWithBuilder() func(Builder)
	GetParent() Displayable
	GetPath() string
	AddChild(child Displayable) int
	GetChildCount() int
	GetChildAt(index int) Displayable
	GetChildren() []Displayable
	GetFilteredChildren(DisplayableFilter) []Displayable
	// TODO(lbayes): This should be capitalized so that external components can implement it.
	setParent(parent Displayable)
}

// Layout and positioning
type Layoutable interface {
	Layout()
	LayoutChildren()

	ActualHeight(height float64)
	ActualWidth(width float64)
	ExcludeFromLayout(bool)
	FlexHeight(int float64)
	FlexWidth(int float64)
	GetActualHeight() float64
	GetActualWidth() float64
	GetExcludeFromLayout() bool
	GetFixedHeight() float64
	GetFixedWidth() float64
	GetFlexHeight() float64
	GetFlexWidth() float64
	GetHAlign() Alignment
	GetHeight() float64
	GetHorizontalPadding() float64
	GetLayoutType() LayoutType
	GetMaxWidth() float64
	GetMinHeight() float64
	GetMinWidth() float64
	GetPadding() float64
	GetPaddingBottom() float64
	GetPaddingLeft() float64
	GetPaddingRight() float64
	GetPaddingTop() float64
	GetPrefHeight() float64
	GetPrefWidth() float64
	GetVAlign() Alignment
	GetVerticalPadding() float64
	GetWidth() float64
	GetX() float64
	GetY() float64
	GetZ() float64
	HAlign(align Alignment)
	Height(height float64)
	LayoutType(layoutType LayoutType)
	MaxHeight(h float64)
	MaxWidth(w float64)
	MinHeight(h float64)
	MinWidth(w float64)
	Padding(value float64)
	PaddingBottom(value float64)
	PaddingLeft(value float64)
	PaddingRight(value float64)
	PaddingTop(value float64)
	PrefHeight(value float64)
	PrefWidth(value float64)
	VAlign(align Alignment)
	Width(width float64)
	X(x float64)
	Y(y float64)
	Z(z float64)
}

// Style and draw methods
type Styleable interface {
	Draw(s Surface)

	GetStyles() StyleDefinition
	GetStylesFor(d Displayable) StyleDefinition
}

// Entities that can be composed, scaled, positioned, and rendered.
type Displayable interface {
	Composable
	Layoutable
	Styleable

	Model(model *ComponentModel)
	GetModel() *ComponentModel

	Title(title string)
	GetTitle() string
}
