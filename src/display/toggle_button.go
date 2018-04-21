package display

type ToggleButtonComponent struct {
	FocusableComponent
}

func NewToggleButton() Displayable {
	return &ToggleButtonComponent{}
}

var ToggleButton = NewComponentFactory("ToggleButton", NewToggleButton)
