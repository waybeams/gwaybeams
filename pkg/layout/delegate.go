package layout

import "github.com/waybeams/waybeams/pkg/spec"

// Delegate Spec fields with axis free semantics.
type Delegate interface {
	ActualSize(d spec.Reader) float64
	Align(d spec.Reader) spec.Alignment
	Flex(d spec.Reader) float64 // GetPercent?
	IsFlexible(d spec.Reader) bool
	LayoutSpec(c spec.ReadWriter) (updatedSize float64)
	MaxSize(d spec.Reader) float64
	MinSize(d spec.Reader) float64
	Padding(d spec.Reader) float64
	PaddingFirst(d spec.Reader) float64
	PaddingLast(d spec.Reader) float64
	PaddingOnAxis(d spec.Reader) float64
	Position(d spec.Reader) float64
	SetActualSize(d spec.Writer, size float64)
	SetChildrenSize(d spec.Writer, size float64)
	SetContentSize(d spec.Writer, size float64)
	SetPosition(d spec.Writer, pos float64)
	SetSize(d spec.ReadWriter, size float64) float64
	Size(d spec.Reader) float64

	/*
		Axis() spec.LayoutAxis
		ChildrenSize(d spec.Reader) float64
		InferredSize(d spec.Reader) float64
		Position(d spec.Reader) float64
		Preferred(d spec.Reader) float64
	*/
}
