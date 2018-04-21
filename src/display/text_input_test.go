package display

import (
	"assert"
	"testing"
)

func TestTextInput(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		instance, _ := TextInput(NewBuilder())
		assert.NotNil(t, instance)
	})

	t.Run("Placeholder text", func(t *testing.T) {
		instance, _ := TextInput(NewBuilder(), Placeholder("Hello World"))

		// Coerce to the concrete type.
		textInput := instance.(*TextInputComponent)
		assert.Equal(t, textInput.Placeholder(), "Hello World")
	})
}
