package display

import (
	"assert"
	"testing"
)

func TestTextInput(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		instance, _ := TextInput(NewBuilder(), Placeholder("Hello World"))
		assert.NotNil(t, instance)
	})
}
