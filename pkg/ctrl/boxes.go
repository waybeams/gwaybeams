package ctrl

import (
	"github.com/waybeams/waybeams/pkg/spec"
)

func VBox(options ...spec.Option) *spec.Spec {
	box := spec.New()
	box.SetSpecName("VBox")
	box.SetLayoutType(spec.VerticalFlowLayoutType)
	spec.Apply(box, options...)
	return box
}

func HBox(options ...spec.Option) *spec.Spec {
	box := spec.New()
	box.SetSpecName("HBox")
	box.SetLayoutType(spec.HorizontalFlowLayoutType)
	spec.Apply(box, options...)
	return box
}

func Box(options ...spec.Option) *spec.Spec {
	box := spec.New()
	box.SetSpecName("Box")
	box.SetLayoutType(spec.StackLayoutType)
	spec.Apply(box, options...)
	return box
}

func Spacer(options ...spec.Option) *spec.Spec {
	spacer := spec.New()
	spacer.SetSpecName("Spacer")
	spacer.SetFlexHeight(1)
	spacer.SetFlexWidth(1)
	spec.Apply(spacer, options...)
	return spacer
}
