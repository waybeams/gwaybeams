package display

type FocusableComponent struct {
	Component

	focused  bool
	selected bool
}

func (t *FocusableComponent) SetSelected(value bool) {
	t.selected = value
}

func (t *FocusableComponent) Selected() bool {
	return t.selected
}

func (t *FocusableComponent) Focus() {
	t.focused = true
}

func (t *FocusableComponent) Blur() {
	t.focused = false
}

func (t *FocusableComponent) IsFocused() bool {
	return t.focused
}

type Focusable interface {
	Selected() bool
	SetSelected(value bool)
	Focus()
	Blur()
	IsFocused() bool
}

// Focusable component options
func Selected(value bool) ComponentOption {
	return func(d Displayable) error {
		d.(Focusable).SetSelected(value)
		return nil
	}
}

func Focused() ComponentOption {
	return func(d Displayable) error {
		d.(Focusable).Focus()
		return nil
	}
}

func Blurred() ComponentOption {
	return func(d Displayable) error {
		d.(Focusable).Blur()
		return nil
	}
}
