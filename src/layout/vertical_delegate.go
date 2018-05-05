package layout

import (
	"math"
	"ui"
)

// Delegate for all properties that are used for Vertical layouts
type verticalDelegate struct{}

func (v *verticalDelegate) ActualSize(d ui.Displayable) float64 {
	return d.ActualHeight()
}

func (v *verticalDelegate) Align(d ui.Displayable) ui.Alignment {
	return d.VAlign()
}

func (v *verticalDelegate) Axis() ui.LayoutAxis {
	return ui.LayoutVertical
}

func (v *verticalDelegate) ChildrenSize(d ui.Displayable) float64 {
	return 0.0
}

func (v *verticalDelegate) Fixed(d ui.Displayable) float64 {
	return d.FixedHeight()
}

func (v *verticalDelegate) Flex(d ui.Displayable) float64 {
	flex := d.FlexHeight()
	if flex == -1 {
		return 0.0
	}
	return flex
}

func (v *verticalDelegate) IsFlexible(d ui.Displayable) bool {
	return d.FlexHeight() > 0.0
}

func (v *verticalDelegate) MinSize(d ui.Displayable) float64 {
	return d.MinHeight()
}

func (v *verticalDelegate) Padding(d ui.Displayable) float64 {
	return math.Max(0, d.VerticalPadding())
}

func (v *verticalDelegate) PaddingFirst(d ui.Displayable) float64 {
	return math.Max(0, d.PaddingTop())
}

func (v *verticalDelegate) PaddingLast(d ui.Displayable) float64 {
	return math.Max(0, d.PaddingBottom())
}

func (v *verticalDelegate) Position(d ui.Displayable) float64 {
	return d.Y()
}

func (v *verticalDelegate) Preferred(d ui.Displayable) float64 {
	return d.PrefHeight()
}

func (v *verticalDelegate) SetActualSize(d ui.Displayable, size float64) {
	d.SetActualHeight(size)
}

func (v *verticalDelegate) SetPosition(d ui.Displayable, pos float64) {
	d.SetY(pos)
}

func (v *verticalDelegate) Size(d ui.Displayable) float64 {
	return d.Height()
}

func (v *verticalDelegate) StaticSize(d ui.Displayable) float64 {
	return v.StaticSize(d)
}
