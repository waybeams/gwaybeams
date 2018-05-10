package layout

import "spec"

type Delegate interface {
	ActualSize(d spec.Reader) float64
	Align(d spec.Reader) spec.Alignment
	Axis() spec.LayoutAxis
	ChildrenSize(d spec.Reader) float64
	Fixed(d spec.Reader) float64
	Flex(d spec.Reader) float64 // GetPercent?
	InferredSize(d spec.Reader) float64
	IsFlexible(d spec.Reader) bool
	MinSize(d spec.Reader) float64
	Padding(d spec.Reader) float64
	PaddingFirst(d spec.Reader) float64
	PaddingLast(d spec.Reader) float64
	Position(d spec.Reader) float64
	Preferred(d spec.Reader) float64
	SetActualSize(d spec.Writer, size float64)
	SetPosition(d spec.Writer, pos float64)
	Size(d spec.Reader) float64
}
