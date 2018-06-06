package layout

import (
	"fmt"
	"github.com/waybeams/waybeams/pkg/spec"
)

// Delegate for all properties that are used for Horizontal layouts
type horizontalDelegate struct{}

func (h *horizontalDelegate) LayoutSpec(c spec.ReadWriter) (updatedSize float64) {
	switch c.LayoutType() {
	case spec.HorizontalFlowLayoutType:
		return FlowOnAxis(h, c)
	case spec.VerticalFlowLayoutType:
		fallthrough
	case spec.StackLayoutType:
		return StackOnAxis(h, c)
	case spec.NoLayoutType:
		return None(h, c)
	default:
		panic(fmt.Sprintf("ERROR: Requested LayoutTypeValue (%v) is not supported:", c.LayoutType()))
	}
}

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

func (h *horizontalDelegate) Flex(d spec.Reader) float64 {
	return d.FlexWidth()
}

func (h *horizontalDelegate) InferredSize(d spec.Reader) float64 {
	return 0
}

func (h *horizontalDelegate) IsFlexible(d spec.Reader) bool {
	return d.FlexWidth() > 0
}

func (h *horizontalDelegate) MaxSize(d spec.Reader) float64 {
	return d.MaxWidth()
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

func (h *horizontalDelegate) PaddingOnAxis(d spec.Reader) float64 {
	return d.HorizontalPadding()
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

func (h *horizontalDelegate) SetChildrenSize(d spec.Writer, size float64) {
	d.SetChildrenWidth(size)
}

func (h *horizontalDelegate) SetContentSize(d spec.Writer, size float64) {
	d.SetContentWidth(size)
}

func (h *horizontalDelegate) SetPosition(d spec.Writer, pos float64) {
	d.SetX(pos)
}

func (h *horizontalDelegate) SetSize(d spec.ReadWriter, size float64) float64 {
	d.SetWidth(size)
	return d.Width()
}

func (h *horizontalDelegate) Size(d spec.Reader) float64 {
	return d.Width()
}

func (h *horizontalDelegate) StaticSize(d spec.Reader) float64 {
	return h.StaticSize(d)
}
