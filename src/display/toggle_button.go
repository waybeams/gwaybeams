package display

type ToggleButtonComponent struct {
	InteractiveComponent
}

func NewToggleButton() Displayable {
	return &ToggleButtonComponent{}
}

var ToggleButton = NewComponentFactory("ToggleButton", NewToggleButton)
