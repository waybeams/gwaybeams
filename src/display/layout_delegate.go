package display

import (
	"math"
)

type LayoutDelegate interface {
	ActualSize(d Displayable) float64
	Align(d Displayable) Alignment
	Axis() LayoutAxis
	ChildrenSize(d Displayable) float64
	Fixed(d Displayable) float64
	Flex(d Displayable) float64 // GetPercent?
	IsFlexible(d Displayable) bool
	MinSize(d Displayable) float64
	Padding(d Displayable) float64
	PaddingFirst(d Displayable) float64
	PaddingLast(d Displayable) float64
	Position(d Displayable) float64
	Preferred(d Displayable) float64
	SetActualSize(d Displayable, size float64)
	SetPosition(d Displayable, pos float64)
	Size(d Displayable) float64
}

// Delegate for all properties that are used for Horizontal layouts
type horizontalDelegate struct{}

func (h *horizontalDelegate) ActualSize(d Displayable) float64 {
	return d.ActualWidth()
}

func (h *horizontalDelegate) Align(d Displayable) Alignment {
	return d.HAlign()
}

func (h *horizontalDelegate) Axis() LayoutAxis {
	return LayoutHorizontal
}

func (h *horizontalDelegate) ChildrenSize(d Displayable) float64 {
	return 0.0
}

func (h *horizontalDelegate) Fixed(d Displayable) float64 {
	return d.FixedWidth()
}

func (h *horizontalDelegate) Flex(d Displayable) float64 {
	flex := d.FlexWidth()
	if flex == -1 {
		return 0.0
	}
	return flex
}

func (h *horizontalDelegate) IsFlexible(d Displayable) bool {
	return d.FlexWidth() > 0.0
}

func (h *horizontalDelegate) MinSize(d Displayable) float64 {
	return d.MinWidth()
}

func (h *horizontalDelegate) Padding(d Displayable) float64 {
	return math.Max(0, d.HorizontalPadding())
}

func (h *horizontalDelegate) PaddingFirst(d Displayable) float64 {
	return math.Max(0, d.PaddingLeft())
}

func (h *horizontalDelegate) PaddingLast(d Displayable) float64 {
	return math.Max(0, d.PaddingRight())
}

func (h *horizontalDelegate) Position(d Displayable) float64 {
	return d.X()
}

func (h *horizontalDelegate) Preferred(d Displayable) float64 {
	return d.PrefWidth()
}

func (h *horizontalDelegate) SetActualSize(d Displayable, size float64) {
	d.SetActualWidth(size)
}

func (h *horizontalDelegate) SetPosition(d Displayable, pos float64) {
	d.SetX(pos)
}

func (h *horizontalDelegate) Size(d Displayable) float64 {
	return d.Width()
}

func (h *horizontalDelegate) StaticSize(d Displayable) float64 {
	return h.StaticSize(d)
}

// Delegate for all properties that are used for Vertical layouts
type verticalDelegate struct{}

func (v *verticalDelegate) ActualSize(d Displayable) float64 {
	return d.ActualHeight()
}

func (v *verticalDelegate) Align(d Displayable) Alignment {
	return d.VAlign()
}

func (v *verticalDelegate) Axis() LayoutAxis {
	return LayoutVertical
}

func (v *verticalDelegate) ChildrenSize(d Displayable) float64 {
	return 0.0
}

func (v *verticalDelegate) Fixed(d Displayable) float64 {
	return d.FixedHeight()
}

func (v *verticalDelegate) Flex(d Displayable) float64 {
	flex := d.FlexHeight()
	if flex == -1 {
		return 0.0
	}
	return flex
}

func (v *verticalDelegate) IsFlexible(d Displayable) bool {
	return d.FlexHeight() > 0.0
}

func (v *verticalDelegate) MinSize(d Displayable) float64 {
	return d.MinHeight()
}

func (v *verticalDelegate) Padding(d Displayable) float64 {
	return math.Max(0, d.VerticalPadding())
}

func (v *verticalDelegate) PaddingFirst(d Displayable) float64 {
	return math.Max(0, d.PaddingTop())
}

func (v *verticalDelegate) PaddingLast(d Displayable) float64 {
	return math.Max(0, d.PaddingBottom())
}

func (v *verticalDelegate) Position(d Displayable) float64 {
	return d.Y()
}

func (v *verticalDelegate) Preferred(d Displayable) float64 {
	return d.PrefHeight()
}

func (v *verticalDelegate) SetActualSize(d Displayable, size float64) {
	d.SetActualHeight(size)
}

func (v *verticalDelegate) SetPosition(d Displayable, pos float64) {
	d.SetY(pos)
}

func (v *verticalDelegate) Size(d Displayable) float64 {
	return d.Height()
}

func (v *verticalDelegate) StaticSize(d Displayable) float64 {
	return v.StaticSize(d)
}
