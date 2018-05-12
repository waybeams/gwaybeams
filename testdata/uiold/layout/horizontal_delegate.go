package layout

import (
	"math"
	"ui"
)

// Delegate for all properties that are used for Horizontal layouts
type horizontalDelegate struct{}

func (h *horizontalDelegate) ActualSize(d ui.Displayable) float64 {
	return d.ActualWidth()
}

func (h *horizontalDelegate) Align(d ui.Displayable) ui.Alignment {
	return d.HAlign()
}

func (h *horizontalDelegate) Axis() ui.LayoutAxis {
	return ui.LayoutHorizontal
}

func (h *horizontalDelegate) ChildrenSize(d ui.Displayable) float64 {
	return 0.0
}

func (h *horizontalDelegate) Fixed(d ui.Displayable) float64 {
	return d.FixedWidth()
}

func (h *horizontalDelegate) Flex(d ui.Displayable) float64 {
	flex := d.FlexWidth()
	if flex == -1 {
		return 0.0
	}
	return flex
}

func (h *horizontalDelegate) IsFlexible(d ui.Displayable) bool {
	return d.FlexWidth() > 0.0
}

func (h *horizontalDelegate) MinSize(d ui.Displayable) float64 {
	return d.MinWidth()
}

func (h *horizontalDelegate) Padding(d ui.Displayable) float64 {
	return math.Max(0, d.HorizontalPadding())
}

func (h *horizontalDelegate) PaddingFirst(d ui.Displayable) float64 {
	return math.Max(0, d.PaddingLeft())
}

func (h *horizontalDelegate) PaddingLast(d ui.Displayable) float64 {
	return math.Max(0, d.PaddingRight())
}

func (h *horizontalDelegate) Position(d ui.Displayable) float64 {
	return d.X()
}

func (h *horizontalDelegate) Preferred(d ui.Displayable) float64 {
	return d.PrefWidth()
}

func (h *horizontalDelegate) SetActualSize(d ui.Displayable, size float64) {
	d.SetActualWidth(size)
}

func (h *horizontalDelegate) SetPosition(d ui.Displayable, pos float64) {
	d.SetX(pos)
}

func (h *horizontalDelegate) Size(d ui.Displayable) float64 {
	return d.Width()
}

func (h *horizontalDelegate) StaticSize(d ui.Displayable) float64 {
	return h.StaticSize(d)
}
