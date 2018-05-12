package controls

import (
	"opts"
	"spec"
)

var VBoxOptions = []spec.Option{
	opts.LayoutType(spec.VerticalFlowLayoutType),
	// spec.FlexHeight(1),
	// spec.FlexWidth(1),
}

var HBoxOptions = []spec.Option{
	opts.LayoutType(spec.VerticalFlowLayoutType),
	// spec.FlexHeight(1),
	// spec.FlexWidth(1),
}

var BoxOptions = []spec.Option{
	opts.LayoutType(spec.StackLayoutType),
	// spec.FlexHeight(1),
	// spec.FlexWidth(1),
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
