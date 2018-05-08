package controls

import (
	"events"
	"ui"
	"ui/control"
	"ui/opts"
	"views"
)

type TextInputControl struct {
	LabelControl

	placeholder string
}

func (t *TextInputControl) SetPlaceholder(text string) {
	t.placeholder = text
}

func (t *TextInputControl) Placeholder() string {
	return t.placeholder
}

func (t *TextInputControl) Text() string {
	text := t.Model().Text
	if text == "" {
		return t.Placeholder()
	}
	return text
}

func NewTextInput() ui.Displayable {
	return &TextInputControl{}
}

// Placeholder Option that only works with TextInputControl
// instances. This text will appear in the text input whenever the Text field
// is empty.
func Placeholder(text string) ui.Option {
	return func(d ui.Displayable) {
		d.(*TextInputControl).SetPlaceholder(text)
	}
}

func textInputCharEnteredHandler(e events.Event) {
	instance := e.Target().(ui.Displayable)
	instance.SetText(instance.Text() + string(e.Payload().(rune)))
	instance.Invalidate()
}

// TextInput is a control that allows the user to input text.
var TextInput = control.Define("TextInput", NewTextInput,
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
