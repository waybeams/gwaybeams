package controls_test

import (
	"assert"
	"controls"
	"events"
	"opts"
	"spec"
	"testing"
)

type inputModel struct {
	Text string
}

func (i *inputModel) SetText(text string) {
	i.Text = text
}

func TestTextInput(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		instance := controls.TextInput(opts.Text("Hello World"))
		assert.Equal(t, instance.Text(), "Hello World")
	})

	t.Run("Renders from model through re-renders", func(t *testing.T) {
		// Define an external model.
		model := &inputModel{Text: "abcd"}

		var create = func(model *inputModel) spec.ReadWriter {
			return controls.TextInput(
				opts.Text(model.Text),
				opts.BindStringPayloadTo(events.TextChanged, model.SetText),
			)
		}

		instance := create(model)
		assert.Equal(t, instance.Text(), "abcd")

		instance.Emit(events.New(events.CharEntered, instance, "Q"))
		assert.Equal(t, model.Text, "abcdQ")

		instance = create(model)
		instance.Emit(events.New(events.CharEntered, instance, "R"))
		assert.Equal(t, model.Text, "abcdQR")

		instance = create(model)
		instance.Emit(events.New(events.CharEntered, instance, "S"))
		assert.Equal(t, model.Text, "abcdQRS")

		instance = create(model)
		instance.Emit(events.New(events.CharEntered, instance, "T"))
		assert.Equal(t, model.Text, "abcdQRST")
	})
}
