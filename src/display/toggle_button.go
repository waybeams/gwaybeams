package display

type ToggleButtonComponent struct {
	Component
}

func NewToggleButton() Displayable {
	return &ToggleButtonComponent{}
}

var ToggleButton = NewComponentFactoryFrom("ToggleButton", Button)
