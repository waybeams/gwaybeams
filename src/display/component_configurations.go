package display

// Box is a base component with a Stack layout.
var Box = NewComponentFactory("Box", NewComponent, LayoutType(StackLayoutType))

// HBox is a base component with a horizontal flow layout.
var HBox = NewComponentFactory("HBox", NewComponent, LayoutType(HorizontalFlowLayoutType))

// VBox is a base component with a vertical flow layout.
var VBox = NewComponentFactory("VBox", NewComponent, LayoutType(VerticalFlowLayoutType))

// Label is a component with a text title that is rendered over the background.
var Label = NewComponentFactory("Label", NewComponent, View(LabelView))

// Spacer provides an invisible component that can take up arbitrary dimensions
var Spacer = NewComponentFactory("Spacer", NewComponent)

// Button is a stub component pending implementation.
var Button = NewComponentFactory("Button", NewComponent, IsFocusable(true))

// RadioGroup is a stub component pending implementation.
var RadioGroup = NewComponentFactory("RadioGroup", NewComponent, IsFocusable(true))

// Checkbox is a stub component pending implementation.
var Checkbox = NewComponentFactory("Checkbox", NewComponent, IsFocusable(true), LayoutType(HorizontalFlowLayoutType))

