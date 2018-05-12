package controls

import (
	"opts"
	"spec"
)

var VBoxOptions = []spec.Option{
	opts.LayoutType(spec.VerticalFlowLayoutType),
}

var HBoxOptions = []spec.Option{
	opts.LayoutType(spec.HorizontalFlowLayoutType),
	opts.VAlign(spec.AlignBottom),
}

var BoxOptions = []spec.Option{
	opts.LayoutType(spec.StackLayoutType),
}

var SpacerOptions = []spec.Option{
	opts.FlexHeight(1),
	opts.FlexWidth(1),
}

func VBox(options ...spec.Option) *spec.Spec {
	box := spec.New()
	spec.ApplyAll(box, VBoxOptions, options)
	return box
}

func HBox(options ...spec.Option) *spec.Spec {
	box := spec.New()
	spec.ApplyAll(box, HBoxOptions, options)
	return box
}

func Box(options ...spec.Option) *spec.Spec {
	box := spec.New()
	spec.ApplyAll(box, BoxOptions, options)
	return box
}

func Spacer(options ...spec.Option) *spec.Spec {
	spacer := spec.New()
	spec.ApplyAll(spacer, SpacerOptions, options)
	return spacer
}
