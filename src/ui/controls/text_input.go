package controls

import (
	"events"
	"ui"
	"ui/control"
	"ui/opts"
	"views"
)

const PlaceholderKey = "TextInput.Placeholder"
const TextKey = "TextInput.Text"

// Placeholder Option that only works with TextInputControl
// instances. This text will appear in the text input whenever the Text field
// is empty.
func Placeholder(text string) ui.Option {
	return func(d ui.Displayable) {
		d.SetData(PlaceholderKey, text)
	}
}

func createConfiguredHandler() events.EventHandler {
	return func(e events.Event) {
		control := e.Target().(ui.Displayable)
		userText := control.DataAsString(TextKey)

		if userText == "" && control.State() != "focused" {
			placeholder := control.DataAsString(PlaceholderKey)
			if placeholder != "" {
				userText = placeholder
			}
		}
		control.SetText(userText)
	}
}

func textInputCharEnteredHandler(e events.Event) {
	control := e.Target().(ui.Displayable)
	textValue := control.DataAsString(TextKey) + string(e.Payload().(rune))

	control.SetData(TextKey, textValue)
	control.SetText(textValue)
	control.Invalidate()
}

// TextInput is a control that allows the user to input text.
var TextInput = control.Define("TextInput",
	control.New,
	opts.OnConfigured(CreateLabelMeasureHandler(func(d ui.Displayable) string {
		text := d.Text()
		if text == "" {
			return d.DataAsString(TextKey)
		}
		return text
	})),
	opts.OnConfigured(createConfiguredHandler()),
	opts.LayoutType(ui.NoLayoutType),
	opts.IsFocusable(true),
	opts.IsTextInput(true),
	opts.BgColor(0xffffffff),
	opts.Padding(5),
	opts.StrokeColor(0x333333ff),
	opts.On(events.CharEntered, textInputCharEnteredHandler),
	opts.OnState("active", opts.StrokeColor(0x333333ff)),
	opts.OnState("focused", opts.StrokeColor(0x0000ffff)),
	opts.OnState("disabled", opts.StrokeColor(0x999999ff)),
	opts.View(views.LabelView))
