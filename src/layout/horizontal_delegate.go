package layout

import (
	"spec"
)

// Delegate for all properties that are used for Horizontal layouts
type horizontalDelegate struct{}

func (h *horizontalDelegate) ActualSize(d spec.Reader) float64 {
	return d.ActualWidth()
}

func (h *horizontalDelegate) Align(d spec.Reader) spec.Alignment {
	return d.HAlign()
}

func (h *horizontalDelegate) Axis() spec.LayoutAxis {
	return spec.LayoutHorizontal
}

func (h *horizontalDelegate) ChildrenSize(d spec.Reader) float64 {
	return 0
}

func (h *horizontalDelegate) Fixed(d spec.Reader) float64 {
	return d.FixedWidth()
}

func (h *horizontalDelegate) Flex(d spec.Reader) float64 {
	return d.FlexWidth()
}

func (h *horizontalDelegate) InferredSize(d spec.Reader) float64 {
	return 0
}

func (h *horizontalDelegate) IsFlexible(d spec.Reader) bool {
	return d.FlexWidth() > 0
}

func (h *horizontalDelegate) MinSize(d spec.Reader) float64 {
	return d.MinWidth()
}

func (h *horizontalDelegate) Padding(d spec.Reader) float64 {
	return d.HorizontalPadding()
}

func (h *horizontalDelegate) PaddingFirst(d spec.Reader) float64 {
	return d.PaddingLeft()
}

func (h *horizontalDelegate) PaddingLast(d spec.Reader) float64 {
	return d.PaddingRight()
}

func (h *horizontalDelegate) Position(d spec.Reader) float64 {
	return d.X()
}

func (h *horizontalDelegate) Preferred(d spec.Reader) float64 {
	return d.PrefWidth()
}

func (h *horizontalDelegate) SetActualSize(d spec.Writer, size float64) {
	d.SetActualWidth(size)
}

func (h *horizontalDelegate) SetPosition(d spec.Writer, pos float64) {
	d.SetX(pos)
}

func (h *horizontalDelegate) Size(d spec.Reader) float64 {
	return d.Width()
}

func (h *horizontalDelegate) StaticSize(d spec.Reader) float64 {
	return h.StaticSize(d)
}
