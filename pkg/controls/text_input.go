package controls

import (
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

const PlaceholderKey = "TextInput.Placeholder"
const TextKey = "TextInput.Text"

// Placeholder Option that only works with TextInputControl
// instances. This text will appear in the text input whenever the Text field
// is empty.
func Placeholder(text string) spec.Option {
	return func(d spec.ReadWriter) {
		// d.SetData(PlaceholderKey, text)
	}
}

/*
func createConfiguredHandler() events.EventHandler {
	return func(e events.Event) {
		control := e.Target().(spec.ReadWriter)
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
	control := e.Target().(spec.ReadWriter)
	textValue := control.Text() + string(e.Payload().(rune))

	control.SetText(textValue)
	// control.Invalidate()
}
*/

// TextInput is a control that allows the user to input text.
var TextInput = func(options ...spec.Option) spec.ReadWriter {
	input := Label()
	var charEnteredHandler = func(e events.Event) {
		ctrl := e.Target().(spec.ReadWriter)
		updatedText := ctrl.Text() + e.Payload().(string)
		ctrl.SetText(updatedText)
		ctrl.Emit(events.New(events.TextChanged, ctrl, updatedText))
	}

	defaults := []spec.Option{
		opts.IsFocusable(true),
		opts.IsTextInput(true),
		opts.BgColor(0xfefefeff),
		opts.StrokeColor(0x666666ff),
		opts.StrokeSize(1),
		opts.On(events.CharEntered, charEnteredHandler),
	}

	spec.ApplyAll(input, defaults, options)
	return input
}

/*
	control.New,
	opts.OnConfigured(CreateLabelMeasureHandler(func(d spec.ReadWriter) string {
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
*/
