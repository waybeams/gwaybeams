package display

import (
	"math"
)

type LayoutDelegate interface {
	ActualSize(d Displayable, size float64)
	GetActualSize(d Displayable) float64
	GetAlign(d Displayable) Alignment
	GetAxis() LayoutAxis
	GetChildrenSize(d Displayable) float64
	GetFixed(d Displayable) float64
	GetFlex(d Displayable) float64 // GetPercent?
	GetIsFlexible(d Displayable) bool
	GetMinSize(d Displayable) float64
	GetPadding(d Displayable) float64
	GetPaddingFirst(d Displayable) float64
	GetPaddingLast(d Displayable) float64
	GetPosition(d Displayable) float64
	GetPreferred(d Displayable) float64
	GetSize(d Displayable) float64
	Position(d Displayable, pos float64)
}

// Delegate for all properties that are used for Horizontal layouts
type horizontalDelegate struct{}

func (h *horizontalDelegate) ActualSize(d Displayable, size float64) {
	d.SetActualWidth(size)
}

func (h *horizontalDelegate) GetActualSize(d Displayable) float64 {
	return d.ActualWidth()
}

func (h *horizontalDelegate) GetAlign(d Displayable) Alignment {
	return d.HAlign()
}

func (h *horizontalDelegate) GetAxis() LayoutAxis {
	return LayoutHorizontal
}

func (h *horizontalDelegate) GetChildrenSize(d Displayable) float64 {
	return 0.0
}

func (h *horizontalDelegate) GetFixed(d Displayable) float64 {
	return d.FixedWidth()
}

func (h *horizontalDelegate) GetFlex(d Displayable) float64 {
	flex := d.FlexWidth()
	if flex == -1 {
		return 0.0
	}
	return flex
}

func (h *horizontalDelegate) GetIsFlexible(d Displayable) bool {
	return d.FlexWidth() > 0.0
}

func (h *horizontalDelegate) GetMinSize(d Displayable) float64 {
	return d.MinWidth()
}

func (h *horizontalDelegate) GetPadding(d Displayable) float64 {
	return math.Max(0, d.HorizontalPadding())
}

func (h *horizontalDelegate) GetPaddingFirst(d Displayable) float64 {
	return math.Max(0, d.PaddingLeft())
}

func (h *horizontalDelegate) GetPaddingLast(d Displayable) float64 {
	return math.Max(0, d.PaddingRight())
}

func (h *horizontalDelegate) GetPosition(d Displayable) float64 {
	return d.X()
}

func (h *horizontalDelegate) GetPreferred(d Displayable) float64 {
	return d.PrefWidth()
}

func (h *horizontalDelegate) GetSize(d Displayable) float64 {
	return d.Width()
}

func (h *horizontalDelegate) GetStaticSize(d Displayable) float64 {
	return h.GetStaticSize(d)
}

func (h *horizontalDelegate) Position(d Displayable, pos float64) {
	d.SetX(pos)
}

// Delegate for all properties that are used for Vertical layouts
type verticalDelegate struct{}

func (v *verticalDelegate) ActualSize(d Displayable, size float64) {
	d.SetActualHeight(size)
}

func (v *verticalDelegate) GetActualSize(d Displayable) float64 {
	return d.ActualHeight()
}

func (v *verticalDelegate) GetAlign(d Displayable) Alignment {
	return d.VAlign()
}

func (v *verticalDelegate) GetAxis() LayoutAxis {
	return LayoutVertical
}

func (v *verticalDelegate) GetChildrenSize(d Displayable) float64 {
	return 0.0
}

func (v *verticalDelegate) GetFixed(d Displayable) float64 {
	return d.FixedHeight()
}

func (v *verticalDelegate) GetFlex(d Displayable) float64 {
	flex := d.FlexHeight()
	if flex == -1 {
		return 0.0
	}
	return flex
}

func (v *verticalDelegate) GetIsFlexible(d Displayable) bool {
	return d.FlexHeight() > 0.0
}

func (v *verticalDelegate) GetMinSize(d Displayable) float64 {
	return d.MinHeight()
}

func (v *verticalDelegate) GetPadding(d Displayable) float64 {
	return math.Max(0, d.VerticalPadding())
}

func (v *verticalDelegate) GetPaddingFirst(d Displayable) float64 {
	return math.Max(0, d.PaddingTop())
}

func (v *verticalDelegate) GetPaddingLast(d Displayable) float64 {
	return math.Max(0, d.PaddingBottom())
}

func (v *verticalDelegate) GetPosition(d Displayable) float64 {
	return d.Y()
}

func (v *verticalDelegate) GetPreferred(d Displayable) float64 {
	return d.PrefHeight()
}

func (v *verticalDelegate) GetSize(d Displayable) float64 {
	return d.Height()
}

func (v *verticalDelegate) GetStaticSize(d Displayable) float64 {
	return v.GetStaticSize(d)
}

func (v *verticalDelegate) Position(d Displayable, pos float64) {
	d.SetY(pos)
}
