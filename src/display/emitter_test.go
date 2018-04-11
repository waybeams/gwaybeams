package display

import (
	"assert"
	"testing"
)

func TestDispatcher(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		instance := NewEmitter()
		assert.NotNil(t, instance)
	})

	t.Run("Adds Handler", func(t *testing.T) {
		var calledWith Event
		handler := func(e Event) {
			calledWith = e
		}
		instance := NewEmitter()
		instance.AddHandler("fake-event", handler)
		instance.Emit("fake-event", "abcd")
		assert.NotNil(t, calledWith, "Expected handler to be called")
		assert.Equal(t, calledWith.Payload(), "abcd", "Received Payload")
	})

	t.Run("RemoveHandler", func(t *testing.T) {
		var calledWith Event
		handler := func(e Event) {
			calledWith = e

		}
		instance := NewEmitter()
		remover := instance.AddHandler("fake-event", handler)
		remover()
		instance.Emit("fake-event", nil)
		assert.Nil(t, calledWith, "Handler was not called")
	})
}
