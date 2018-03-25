package display

var Box = NewComponentFactory(NewComponent, LayoutType(StackLayoutType))
var HBox = NewComponentFactory(NewComponent, LayoutType(HorizontalFlowLayoutType))
var VBox = NewComponentFactory(NewComponent, LayoutType(VerticalFlowLayoutType))
