package layout

import "ui"

type LayoutDelegate interface {
	ActualSize(d ui.Displayable) float64
	Align(d ui.Displayable) ui.Alignment
	Axis() ui.LayoutAxis
	ChildrenSize(d ui.Displayable) float64
	Fixed(d ui.Displayable) float64
	Flex(d ui.Displayable) float64 // GetPercent?
	IsFlexible(d ui.Displayable) bool
	MinSize(d ui.Displayable) float64
	Padding(d ui.Displayable) float64
	PaddingFirst(d ui.Displayable) float64
	PaddingLast(d ui.Displayable) float64
	Position(d ui.Displayable) float64
	Preferred(d ui.Displayable) float64
	SetActualSize(d ui.Displayable, size float64)
	SetPosition(d ui.Displayable, pos float64)
	Size(d ui.Displayable) float64
}
