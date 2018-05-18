package events_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/events"
	"testing"
)

func TestDispatcher(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		instance := events.NewEmitter()
		assert.NotNil(instance)
	})

	t.Run("Adds Handler", func(t *testing.T) {
		var calledWith events.Event
		handler := func(e events.Event) {
			calledWith = e
		}
		instance := events.NewEmitter()
		instance.On("fake-event", handler)
		instance.Emit(events.New("fake-event", instance, "abcd"))
		assert.NotNil(calledWith, "Expected handler to be called")
		assert.Equal(calledWith.Payload(), "abcd", "Received Payload")
		assert.Equal(calledWith.Target(), instance, "received Target")
	})

	t.Run("RemoveHandler", func(t *testing.T) {
		var calledWith events.Event
		handler := func(e events.Event) {
			calledWith = e
		}
		instance := events.NewEmitter()
		remover := instance.On("fake-event", handler)
		remover()
		instance.Emit(events.New("fake-event", nil, nil))
		assert.Nil(calledWith, "Handler was not called")
	})

	t.Run("RemoveAllHandlers", func(t *testing.T) {
		var calledWith events.Event
		handler := func(e events.Event) {
			calledWith = e
		}
		instance := events.NewEmitter()
		instance.On("fake-event", handler)
		found := instance.RemoveAllHandlers()
		assert.True(found, "Expected to find handlers")
		instance.Emit(events.New("fake-event", nil, nil))
		assert.Nil(calledWith, "Handler was not called")

		instance.On("fake-event", handler)
		instance.Emit(events.New("fake-event", nil, nil))
		assert.NotNil(calledWith, "Handler was called")
	})

	t.Run("RemoveAllHandlersFor", func(t *testing.T) {
		var calledWith events.Event
		handler := func(e events.Event) {
			calledWith = e
		}
		instance := events.NewEmitter()
		instance.On("fake-event-1", handler)
		instance.On("fake-event-2", handler)
		found := instance.RemoveAllHandlersFor("fake-event-2")

		assert.True(found, "Expected to find removable handlers")
		instance.Emit(events.New("fake-event-2", nil, nil))
		assert.Nil(calledWith, "Handler was not called")
		instance.Emit(events.New("fake-event-1", nil, nil))
		assert.NotNil(calledWith, "Handler was called")
	})

	t.Run("RemoveAllHandlers returns false if none present", func(t *testing.T) {
		instance := events.NewEmitter()
		assert.False(instance.RemoveAllHandlers(), "Expected no handlers")
	})

	t.Run("RemoveAllHandlersFor returns false if none present", func(t *testing.T) {
		instance := events.NewEmitter()
		assert.False(instance.RemoveAllHandlersFor("no-event"), "Expected no handlers")
	})
}
