package controls

import (
	"component"
	"events"
	"opts"
	"ui"
	"views"
)

type TextInputComponent struct {
	component.Component

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

func NewTextInput() ui.Displayable {
	return &TextInputComponent{}
}

// Placeholder ComponentOption that only works with TextInputComponent
// instances. This text will appear in the text input whenever the Text field
// is empty.
func Placeholder(text string) ui.Option {
	return func(d ui.Displayable) {
		d.(*TextInputComponent).SetPlaceholder(text)
	}
}

func textInputCharEnteredHandler(e events.Event) {
	instance := e.Target().(ui.Displayable)
	instance.SetText(instance.Text() + string(e.Payload().(rune)))
}

// TextInput is a component that allows the user to input text.
var TextInput = component.Define("TextInput", NewTextInput,
	opts.IsFocusable(true),
	opts.IsTextInput(true),
	opts.BgColor(0xffffffff),
	opts.Padding(5),
	opts.StrokeColor(0x333333ff),
	opts.On(events.CharEntered, textInputCharEnteredHandler),
	opts.OnState("active", opts.StrokeColor(0x333333ff)),
	opts.OnState("focused", opts.StrokeColor(0x0000ffff)),
	opts.OnState("disabled", opts.StrokeColor(0x999999ff)),
	opts.View(views.TextInputView))
