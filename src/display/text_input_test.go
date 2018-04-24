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
		assert.Equal(t, instance.(*TextInputComponent).Placeholder(), "Hello World")
	})

	t.Run("Text() uses Placholder() when empty", func(t *testing.T) {
		instance, _ := TextInput(NewBuilder(), Placeholder("abcd"))
		assert.Equal(t, instance.Text(), "abcd")
		instance.SetText("efgh")
		assert.Equal(t, instance.Text(), "efgh")
	})
}
