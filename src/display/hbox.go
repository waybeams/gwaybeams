package display

type HBoxComponent struct {
	BoxComponent
}

func NewHBox() Displayable {
	return &HBoxComponent{}
}

// Debating whether this belongs in this file, or if they should all be
// defined in component_factory.go, or maybe someplace else?
// This is the hook that is used within the Builder context.
var HBox = NewComponentFactory(NewHBox)
