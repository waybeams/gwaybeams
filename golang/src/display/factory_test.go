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

	t.Run("Forwards stack.Push(nil) error", func(t *testing.T) {
		err := instance.Push(nil)
		assert.NotNil(err)
	})

	t.Run("Processes simple args", func(t *testing.T) {
		emptyArgs := []interface{}
		assert.NotNil(emptyArgs)



	})
}
