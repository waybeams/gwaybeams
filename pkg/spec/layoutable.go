package spec

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
	NoLayoutType = iota
	StackLayoutType
	VerticalFlowLayoutType
	HorizontalFlowLayoutType
	RowLayoutType
)

// Alignment is used represent alignment of Spec children, text or any other
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
type LayoutHandler func(node ReadWriter) (minWidth, minHeight float64)

type ResizableWriter interface {
	SetHeight(height float64)
	SetWidth(width float64)
}

type ResizableReader interface {
	Width() float64
	Height() float64
}

type LayoutableWriter interface {
	ResizableWriter

	SetActualHeight(value float64)
	SetActualWidth(value float64)
	SetChildrenHeight(height float64)
	SetChildrenWidth(width float64)
	SetContentHeight(height float64)
	SetContentWidth(width float64)
	SetExcludeFromLayout(bool)
	SetFlexHeight(int float64)
	SetFlexWidth(int float64)
	SetGutter(value float64)
	SetHAlign(align Alignment)
	SetIsMeasured(measured bool)
	SetLayoutType(layoutType LayoutTypeValue)
	SetMaxHeight(h float64)
	SetMaxWidth(w float64)
	SetMinHeight(h float64)
	SetMinWidth(w float64)
	SetPadding(value float64)
	SetPaddingBottom(value float64)
	SetPaddingLeft(value float64)
	SetPaddingRight(value float64)
	SetPaddingTop(value float64)
	SetPrefHeight(value float64)
	SetPrefWidth(value float64)
	SetTextX(value float64)
	SetTextY(value float64)
	SetVAlign(align Alignment)
	SetX(x float64)
	SetY(y float64)
}

type LayoutableReader interface {
	ResizableReader

	ActualHeight() float64
	ActualWidth() float64
	ChildrenHeight() float64
	ChildrenWidth() float64
	ContentHeight() float64
	ContentWidth() float64
	ExcludeFromLayout() bool
	FlexHeight() float64
	FlexWidth() float64
	Gutter() float64
	HAlign() Alignment
	IsMeasured() bool
	HorizontalPadding() float64
	LayoutType() LayoutTypeValue
	MaxHeight() float64
	MaxWidth() float64
	Measure(s Surface)
	MinHeight() float64
	MinWidth() float64
	PaddingBottom() float64
	PaddingLeft() float64
	PaddingRight() float64
	PaddingTop() float64
	PrefHeight() float64
	PrefWidth() float64
	TextX() float64
	TextY() float64
	VAlign() Alignment
	VerticalPadding() float64
	X() float64
	XOffset() float64
	Y() float64
	YOffset() float64
}

type LayoutableReadWriter interface {
	LayoutableReader
	LayoutableWriter
}

func (c *Spec) ActualHeight() float64 {
	return c.actualHeight
}

func (c *Spec) ActualWidth() float64 {
	return c.actualWidth
}

func (c *Spec) SetLayoutType(layoutType LayoutTypeValue) {
	c.layoutType = layoutType
}

func (c *Spec) LayoutType() LayoutTypeValue {
	return c.layoutType
}

func (c *Spec) IsMeasured() bool {
	return c.isMeasured
}

func (c *Spec) Measure(s Surface) {
	// noop
}

func (c *Spec) Gutter() float64 {
	return c.gutter
}

func (c *Spec) SetIsMeasured(measured bool) {
	c.isMeasured = measured
}

func (c *Spec) SetX(x float64) {
	c.x = x
}

func (c *Spec) SetY(y float64) {
	c.y = y
}

func (c *Spec) SetTextX(x float64) {
	c.textX = x
}

func (c *Spec) SetTextY(y float64) {
	c.textY = y
}

func (c *Spec) TextX() float64 {
	return (c.X() + c.PaddingLeft()) - c.textX
}

func (c *Spec) TextY() float64 {
	return (c.Y() + c.PaddingTop()) - c.textY
}

func (c *Spec) X() float64 {
	return c.x
}

func (c *Spec) Y() float64 {
	return c.y
}

func (c *Spec) SetHAlign(value Alignment) {
	c.hAlign = value
}

func (c *Spec) HAlign() Alignment {
	return c.hAlign
}

func (c *Spec) VAlign() Alignment {
	return c.vAlign
}

func (c *Spec) SetVAlign(value Alignment) {
	c.vAlign = value
}

func (c *Spec) SetGutter(gutter float64) {
	c.gutter = gutter
}

func (c *Spec) SetWidth(w float64) {
	c.width = w
}

func (c *Spec) SetHeight(h float64) {
	c.height = h
}

func (c *Spec) ChildrenWidth() float64 {
	return c.childrenWidth
}

func (c *Spec) ChildrenHeight() float64 {
	return c.childrenHeight
}

func (c *Spec) SetChildrenWidth(size float64) {
	c.childrenWidth = size
}

func (c *Spec) SetChildrenHeight(size float64) {
	c.childrenHeight = size
}

func (c *Spec) ContentWidth() float64 {
	return c.contentWidth
}

func (c *Spec) ContentHeight() float64 {
	return c.contentHeight
}

func (c *Spec) SetContentWidth(size float64) {
	c.contentWidth = size
}

func (c *Spec) SetContentHeight(size float64) {
	c.contentHeight = size
}

func (c *Spec) Width() float64 {
	actual := c.width
	padding := c.HorizontalPadding()

	if c.maxWidth > 0 && actual > c.maxWidth {
		actual = c.maxWidth
	}
	if c.minWidth > 0 && actual < c.minWidth {
		actual = c.minWidth
	}
	if c.childrenWidth > 0 && actual < c.childrenWidth+padding {
		actual = c.childrenWidth + padding
	}
	if c.contentWidth > 0 && actual < c.contentWidth+padding {
		actual = c.contentWidth + padding
	}
	return actual
}

func (c *Spec) Height() float64 {
	actual := c.height
	padding := c.VerticalPadding()

	if c.maxHeight > 0 && actual > c.maxHeight {
		actual = c.maxHeight
	}
	if c.minHeight > 0 && actual < c.minHeight {
		actual = c.minHeight
	}
	if c.childrenHeight > 0 && actual < c.childrenHeight+padding {
		actual = c.childrenHeight + padding
	}
	if c.contentHeight > 0 && actual < c.contentHeight+padding {
		actual = c.contentHeight + padding
	}
	return actual
}

func (c *Spec) SetPrefWidth(value float64) {
	c.prefWidth = value
}

func (c *Spec) SetPrefHeight(value float64) {
	c.prefHeight = value
}

func (c *Spec) PrefWidth() float64 {
	return c.prefWidth
}

func (c *Spec) PrefHeight() float64 {
	return c.prefHeight
}

func (c *Spec) SetActualWidth(width float64) {
	c.actualWidth = width
}

func (c *Spec) SetActualHeight(height float64) {
	c.actualHeight = height
}

func (c *Spec) SetExcludeFromLayout(value bool) {
	c.excludeFromLayout = value
}

func (c *Spec) SetMinWidth(min float64) {
	c.minWidth = min
}

func (c *Spec) SetMinHeight(min float64) {
	c.minHeight = min
}

func (c *Spec) MinWidth() float64 {
	return c.minWidth
}

func (c *Spec) MinHeight() float64 {
	return c.minHeight
}

func (c *Spec) SetMaxWidth(max float64) {
	c.maxWidth = max
}

func (c *Spec) SetMaxHeight(max float64) {
	c.maxHeight = max
}

func (c *Spec) MaxWidth() float64 {
	return c.maxWidth
}

func (c *Spec) MaxHeight() float64 {
	return c.maxHeight
}

func (c *Spec) ExcludeFromLayout() bool {
	return c.excludeFromLayout
}

func (c *Spec) SetFlexWidth(value float64) {
	c.flexWidth = value
}

func (c *Spec) SetFlexHeight(value float64) {
	c.flexHeight = value
}

func (c *Spec) FlexWidth() float64 {
	return c.flexWidth
}

func (c *Spec) FlexHeight() float64 {
	return c.flexHeight
}

func (c *Spec) SetPadding(value float64) {
	c.paddingBottom = value
	c.paddingLeft = value
	c.paddingRight = value
	c.paddingTop = value
}

func (c *Spec) SetPaddingBottom(value float64) {
	c.paddingBottom = value
}

func (c *Spec) SetPaddingLeft(value float64) {
	c.paddingLeft = value
}

func (c *Spec) SetPaddingRight(value float64) {
	c.paddingRight = value
}

func (c *Spec) SetPaddingTop(value float64) {
	c.paddingTop = value
}

func (c *Spec) HorizontalPadding() float64 {
	return c.PaddingLeft() + c.PaddingRight()
}

func (c *Spec) VerticalPadding() float64 {
	return c.PaddingTop() + c.PaddingBottom()
}

func (c *Spec) PaddingLeft() float64 {
	return c.paddingLeft
}

func (c *Spec) PaddingRight() float64 {
	return c.paddingRight
}

func (c *Spec) PaddingBottom() float64 {
	return c.paddingBottom
}

func (c *Spec) PaddingTop() float64 {
	return c.paddingTop
}

func (c *Spec) YOffset() float64 {
	offset := c.Y()
	parent := c.Parent()
	if parent != nil {
		offset = offset + parent.YOffset()
	}
	if offset > 0 {
		return offset
	}
	return 0
}

func (c *Spec) XOffset() float64 {
	offset := c.X()
	parent := c.Parent()
	if parent != nil {
		offset = offset + parent.XOffset()
	}
	if offset > 0 {
		return offset
	}
	return 0
}
