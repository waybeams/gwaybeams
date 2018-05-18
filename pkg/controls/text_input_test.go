package controls_test

import (
	"github.com/waybeams/waybeams/pkg/controls"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
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
		assert.Equal(instance.Text(), "Hello World")
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
		assert.Equal(instance.Text(), "abcd")

		instance.Emit(events.New(events.CharEntered, instance, "Q"))
		assert.Equal(model.Text, "abcdQ")

		instance = create(model)
		instance.Emit(events.New(events.CharEntered, instance, "R"))
		assert.Equal(model.Text, "abcdQR")

		instance = create(model)
		instance.Emit(events.New(events.CharEntered, instance, "S"))
		assert.Equal(model.Text, "abcdQRS")

		instance = create(model)
		instance.Emit(events.New(events.CharEntered, instance, "T"))
		assert.Equal(model.Text, "abcdQRST")
	})
}
