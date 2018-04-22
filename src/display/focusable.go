package display

import (
	"events"
)

type Focusable interface {
	Selected() bool
	SetSelected(value bool)
	Focus()
	Blur()
	Focused() bool
}

type InteractiveComponent struct {
	Component

	focused  bool
	selected bool
}

func (t *InteractiveComponent) SetSelected(value bool) {
	t.selected = value
}

func (t *InteractiveComponent) Selected() bool {
	return t.selected
}

func (t *InteractiveComponent) Focus() {
	t.Bubble(NewEvent(events.Focused, t, nil))
	t.focused = true
}

func (t *InteractiveComponent) Blur() {
	t.Bubble(NewEvent(events.Blurred, t, nil))
	t.focused = false
}

func (t *InteractiveComponent) Focused() bool {
	return t.focused
}

func NewInteractiveComponent() Displayable {
	return &InteractiveComponent{}
}

var factory = NewComponentFactory("Interactive", NewInteractiveComponent)

func Interactive(b Builder, options ...ComponentOption) (*InteractiveComponent, error) {
	instance, err := factory(b, options...)

	return instance.(*InteractiveComponent), err
}

// Options for Focusable components

func Blurred() ComponentOption {
	return func(d Displayable) error {
		d.(Focusable).Blur()
		return nil
	}
}

func Focused() ComponentOption {
	return func(d Displayable) error {
		d.(Focusable).Focus()
		return nil
	}
}

func OnClick(handler EventHandler) ComponentOption {
	return func(d Displayable) error {
		d.PushUnsubscriber(d.On(events.Clicked, handler))
		return nil
	}
}

// Focusable component options
func Selected(value bool) ComponentOption {
	return func(d Displayable) error {
		d.(Focusable).SetSelected(value)
		return nil
	}
}
