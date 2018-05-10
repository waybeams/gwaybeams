package controls

import (
	"assert"
	"events"
	"testing"
	"ui"
	"uiold/context"
	"uiold/opts"
)

func TestTextInput(t *testing.T) {
	var createTextInput = func(options ...ui.Option) ui.Displayable {
		defaultOptions := []ui.Option{
			opts.BgColor(0xffffffff),
			opts.StrokeSize(2),
			opts.StrokeColor(0x222222ff),
			opts.FontFace("Roboto"),
			opts.FontSize(24),
			opts.FontColor(0x333333ff),
			opts.Width(400),
			opts.Height(80),
		}
		mergedOptions := append(defaultOptions, options...)
		return TextInput(context.NewTestContext(), mergedOptions...)
	}

	t.Run("No placeholder or Text", func(t *testing.T) {
		instance := createTextInput()
		assert.Equal(t, instance.Text(), "")
	})

	t.Run("Placeholder removed on focus", func(t *testing.T) {
		instance := createTextInput(Placeholder("Hello"))
		instance.Focus()
		instance.Emit(events.New(events.Configured, instance, nil))
		assert.Equal(t, instance.Text(), "")
	})

	t.Run("Placeholder text", func(t *testing.T) {
		instance := createTextInput(Placeholder("Hello World"))
		assert.Equal(t, instance.Text(), "Hello World")
	})

	t.Run("Text() uses Placholder() when empty", func(t *testing.T) {
		instance := createTextInput(Placeholder("abcd"))
		assert.Equal(t, instance.Text(), "abcd")
		instance.SetText("efgh")
		assert.Equal(t, instance.Text(), "efgh")
	})

	t.Run("Key inputs increment text", func(t *testing.T) {
		instance := createTextInput(Placeholder("default"))
		instance.Emit(events.New(events.CharEntered, instance, rune('B')))

		assert.Equal(t, instance.Text(), "B")
		instance.Emit(events.New(events.CharEntered, instance, rune('Y')))
		assert.Equal(t, instance.Text(), "BY")
		instance.Emit(events.New(events.CharEntered, instance, rune('E')))
		assert.Equal(t, instance.Text(), "BYE")

		// Clear the user-entered text:
		instance.SetData("TextInput.Text", "")
		instance.Emit(events.New(events.Configured, instance, nil)) // :barf:
		assert.Equal(t, instance.Text(), "default")
	})
}
