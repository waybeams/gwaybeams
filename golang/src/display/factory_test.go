package display

import (
	"assert"
	"testing"
)

func TestFactory(t *testing.T) {
	instance := &Factory{}

	t.Run("Instantiable", func(t *testing.T) {
		assert.NotNil(instance)
	})

	t.Run("Push nil", func(t *testing.T) {
		instance.Push(nil)
	})
}
