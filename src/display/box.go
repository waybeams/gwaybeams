package display

// Box is a base component with a Stack layout
var Box = NewComponentFactory(NewComponent, LayoutType(StackLayoutType))

// HBox is a base component with a horizontal flow layout
var HBox = NewComponentFactory(NewComponent, LayoutType(HorizontalFlowLayoutType))

// VBox is a base component with a vertical flow layout
var VBox = NewComponentFactory(NewComponent, LayoutType(VerticalFlowLayoutType))
