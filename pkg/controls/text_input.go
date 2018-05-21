package controls

import (
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/views"
)

const PlaceholderKey = "TextInput.Placeholder"
const TextKey = "TextInput.Text"

type TextInputSpec struct {
	LabelSpec

	placeholder string
}

func (t *TextInputSpec) Placeholder() string {
	return t.placeholder
}

// TextInput is a control that allows the user to input text.
var TextInput = func(options ...spec.Option) spec.ReadWriter {
	input := &TextInputSpec{}
	var charEnteredHandler = func(e events.Event) {
		ctrl := e.Target().(spec.ReadWriter)
		updatedText := ctrl.Text() + e.Payload().(string)
		ctrl.SetText(updatedText)
		ctrl.Emit(events.New(events.TextChanged, ctrl, updatedText))
	}

	defaults := []spec.Option{
		opts.BgColor(0xfefefeff),
		opts.HAlign(spec.AlignLeft),
		opts.IsFocusable(true),
		opts.IsMeasured(true),
		opts.IsTextInput(true),
		opts.LayoutType(spec.StackLayoutType),
		opts.On(events.CharEntered, charEnteredHandler),
		opts.SpecName("Label"),
		opts.StrokeSize(1),
		opts.VAlign(spec.AlignTop),
		opts.View(views.LabelView),
		opts.OnState("active", opts.StrokeColor(0x666666ff)),
		opts.OnState("focused", opts.StrokeColor(0x44d9e6ff)),
		opts.On(events.Blurred, opts.OptionsHandler(opts.SetState("active"))),
		opts.On(events.Focused, opts.OptionsHandler(opts.SetState("focused"))),
	}

	spec.ApplyAll(input, defaults, options)

	// We should add a placeholder Label as a child.
	if input.Text() == "" && input.Placeholder() != "" {
		// Create a bag of options and then apply them to the input instance.
		opts.Bag(
			opts.Child(Label(
				opts.FontColor(0x666666ff),
				opts.Key("TextInput.Placeholder"),
				opts.Padding(10),
				opts.Text(input.Placeholder()),
			)),
			opts.IsMeasured(false),
		)(input)
	}

	return input
}

// Placeholder Option that only works with TextInputSpec instances. This text
// will appear in the text input whenever the Text property is empty.
func Placeholder(text string) spec.Option {
	return func(d spec.ReadWriter) {
		d.(*TextInputSpec).placeholder = text
	}
}
