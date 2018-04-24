package display

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

// TextInput is a component that allows the user to input text.
var TextInput = NewComponentFactory("TextInput", NewTextInput,
	IsFocusable(true),
	BgColor(0xeeeeeeff),
	StrokeColor(0xddddddff),
	View(TextInputView))
