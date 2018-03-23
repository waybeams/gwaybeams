package display

type VBoxComponent struct {
	BoxComponent
}

func NewVBox() Displayable {
	return &VBoxComponent{}
}

// Debating whether this belongs in this file, or if they should all be
// defined in component_factory.go, or maybe someplace else?
// This is the hook that is used within the Builder context.
var VBox = NewComponentFactory(NewVBox)
