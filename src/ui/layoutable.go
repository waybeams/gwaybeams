package ui

type LayoutAxis int

const (
	LayoutHorizontal = iota
	LayoutVertical
)

// LayoutTypeValue is a serializable enum for selecting a layout scheme.
// This pattern is probably not the way to go, but I'm having trouble finding a
// reasonable alternative. The problem here is that LayoutHandler types will not be
// user-extensible. Box definitions will only be able to refer to the
// Layouts that have been enumerated here. The benefit is that Model objects
// will remain serializable and simply be a bag of scalars. I'm definitely
// open to suggestions.
type LayoutTypeValue int

const (
	StackLayoutType = iota
	VerticalFlowLayoutType
	HorizontalFlowLayoutType
	RowLayoutType
	NoLayoutType
)

// Alignment is used represent alignment of Component children, text or any other
// alignable entities.
type Alignment int

const (
	AlignBottom = iota
	AlignLeft
	AlignRight
	AlignTop
	AlignCenter
)

// LayoutHandler is a concrete implementation of a given layout. These handlers
// are pure functions that accept a Displayable and manage the scale and
// position of the children for that element.
type LayoutHandler func(d Displayable)

// Layoutable is a set of methods for components that can be positions and
// scaled.
type Layoutable interface {
	ActualHeight() float64
	ActualWidth() float64
	ExcludeFromLayout() bool
	FixedHeight() float64
	FixedWidth() float64
	FlexHeight() float64
	FlexWidth() float64
	Gutter() float64
	HAlign() Alignment
	Height() float64
	HorizontalPadding() float64
	Layout()
	LayoutChildren()
	LayoutType() LayoutTypeValue
	MaxHeight() float64
	MaxWidth() float64
	MinHeight() float64
	MinWidth() float64
	Model() *Model
	Padding() float64
	PaddingBottom() float64
	PaddingLeft() float64
	PaddingRight() float64
	PaddingTop() float64
	PrefHeight() float64
	PrefWidth() float64
	SetActualHeight(height float64)
	SetActualWidth(width float64)
	SetExcludeFromLayout(bool)
	SetFlexHeight(int float64)
	SetFlexWidth(int float64)
	SetGutter(value float64)
	SetHAlign(align Alignment)
	SetHeight(height float64)
	SetLayoutType(layoutType LayoutTypeValue)
	SetMaxHeight(h float64)
	SetMaxWidth(w float64)
	SetMinHeight(h float64)
	SetMinWidth(w float64)
	SetModel(model *Model)
	SetPadding(value float64)
	SetPaddingBottom(value float64)
	SetPaddingLeft(value float64)
	SetPaddingRight(value float64)
	SetPaddingTop(value float64)
	SetPrefHeight(value float64)
	SetPrefWidth(value float64)
	SetVAlign(align Alignment)
	SetWidth(width float64)
	SetX(x float64)
	SetY(y float64)
	SetZ(z float64)
	VAlign() Alignment
	VerticalPadding() float64
	Width() float64
	X() float64
	XOffset() float64
	Y() float64
	YOffset() float64
	Z() float64
}
