package controls_test

import (
	"assert"
	. "controls"
	"ctx"
	"opts"
	"testing"
	"ui"
)

func TestTextInput(t *testing.T) {
	var createTextInput = func(options ...ui.Option) *TextInputComponent {
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
		return TextInput(ctx.NewTestContext(), mergedOptions...).(*TextInputComponent)
	}

	t.Run("Instantiable", func(t *testing.T) {
		instance := createTextInput()
		assert.NotNil(t, instance)
	})

	t.Run("Placeholder text", func(t *testing.T) {
		instance := createTextInput(Placeholder("Hello World"))
		assert.Equal(t, instance.Placeholder(), "Hello World")
	})

	t.Run("Text() uses Placholder() when empty", func(t *testing.T) {
		instance := createTextInput(Placeholder("abcd"))
		assert.Equal(t, instance.Text(), "abcd")
		instance.SetText("efgh")
		assert.Equal(t, instance.Text(), "efgh")
	})
}
