package display

import (
	"events"
)

type TextInputComponent struct {
	Component

	placeholder string
}

func (t *TextInputComponent) SetPlaceholder(text string) {
	t.placeholder = text
}

func (t *TextInputComponent) Placeholder() string {
	return t.placeholder
}

func (t *TextInputComponent) Text() string {
	text := t.Model().Text
	if text == "" {
		return t.Placeholder()
	}
	return text
}

func NewTextInput() Displayable {
	return &TextInputComponent{}
}

// Placeholder ComponentOption that only works with TextInputComponent
// instances. This text will appear in the text input whenever the Text field
// is empty.
func Placeholder(text string) ComponentOption {
	return func(d Displayable) error {
		d.(*TextInputComponent).SetPlaceholder(text)
		return nil
	}
}

func textInputCharEnteredHandler(e Event) {
	instance := e.Target().(Displayable)
	instance.SetText(instance.Text() + string(e.Payload().(rune)))
}

// TextInput is a component that allows the user to input text.
var TextInput = NewComponentFactory("TextInput", NewTextInput,
	IsFocusable(true),
	IsTextInput(true),
	BgColor(0xffffffff),
	Padding(5),
	StrokeColor(0x333333ff),
	On(events.CharEntered, textInputCharEnteredHandler),
	OnState("active", StrokeColor(0x333333ff)),
	OnState("focused", StrokeColor(0x0000ffff)),
	OnState("disabled", StrokeColor(0x999999ff)),
	View(TextInputView))
