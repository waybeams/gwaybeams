package layout

import (
	"fmt"
	"github.com/waybeams/waybeams/pkg/spec"
)

// Delegate for all properties that are used for Vertical layouts
type verticalDelegate struct{}

func (v *verticalDelegate) LayoutSpec(c spec.ReadWriter) (updatedSize float64) {
	switch c.LayoutType() {
	case spec.VerticalFlowLayoutType:
		return FlowOnAxis(v, c)
	case spec.HorizontalFlowLayoutType:
		fallthrough
	case spec.StackLayoutType:
		return StackOnAxis(v, c)
	case spec.NoLayoutType:
		return None(v, c)
	default:
		panic(fmt.Sprintf("ERROR: Requested LayoutTypeValue (%v) is not supported:", c.LayoutType()))
	}
}

func (v *verticalDelegate) ActualSize(d spec.Reader) float64 {
	return d.ActualHeight()
}

func (v *verticalDelegate) Align(d spec.Reader) spec.Alignment {
	return d.VAlign()
}

func (v *verticalDelegate) Axis() spec.LayoutAxis {
	return spec.LayoutVertical
}

func (v *verticalDelegate) ChildrenSize(d spec.Reader) float64 {
	return 0
}

func (v *verticalDelegate) Flex(d spec.Reader) float64 {
	return d.FlexHeight()
}

func (v *verticalDelegate) InferredSize(d spec.Reader) float64 {
	return 0
}

func (v *verticalDelegate) IsFlexible(d spec.Reader) bool {
	return d.FlexHeight() > 0.0
}

func (v *verticalDelegate) MaxSize(d spec.Reader) float64 {
	return d.MaxHeight()
}

func (v *verticalDelegate) MinSize(d spec.Reader) float64 {
	return d.MinHeight()
}

func (v *verticalDelegate) Padding(d spec.Reader) float64 {
	return d.VerticalPadding()
}

func (v *verticalDelegate) PaddingFirst(d spec.Reader) float64 {
	return d.PaddingTop()
}

func (v *verticalDelegate) PaddingLast(d spec.Reader) float64 {
	return d.PaddingBottom()
}

func (v *verticalDelegate) PaddingOnAxis(d spec.Reader) float64 {
	return d.VerticalPadding()
}

func (v *verticalDelegate) Position(d spec.Reader) float64 {
	return d.Y()
}

func (v *verticalDelegate) Preferred(d spec.Reader) float64 {
	return d.PrefHeight()
}

func (v *verticalDelegate) SetActualSize(d spec.Writer, size float64) {
	d.SetActualHeight(size)
}

func (v *verticalDelegate) SetChildrenSize(d spec.Writer, size float64) {
	d.SetChildrenHeight(size)
}

func (v *verticalDelegate) SetContentSize(d spec.Writer, size float64) {
	d.SetContentHeight(size)
}

func (v *verticalDelegate) SetPosition(d spec.Writer, pos float64) {
	d.SetY(pos)
}

func (v *verticalDelegate) SetSize(d spec.ReadWriter, size float64) float64 {
	d.SetHeight(size)
	return d.Height()
}

func (v *verticalDelegate) Size(d spec.Reader) float64 {
	return d.Height()
}

func (v *verticalDelegate) StaticSize(d spec.Reader) float64 {
	return v.StaticSize(d)
}
