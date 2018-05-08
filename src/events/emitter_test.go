package events_test

import (
	"assert"
	"events"
	"testing"
	. "ui"
	"ui/context"
	. "ui/controls"
	. "ui/opts"
)

func TestDispatcher(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		instance := events.NewEmitter()
		assert.NotNil(t, instance)
	})

	t.Run("Adds Handler", func(t *testing.T) {
		var calledWith events.Event
		handler := func(e events.Event) {
			calledWith = e
		}
		instance := events.NewEmitter()
		instance.On("fake-event", handler)
		instance.Emit(events.New("fake-event", instance, "abcd"))
		assert.NotNil(t, calledWith, "Expected handler to be called")
		assert.Equal(t, calledWith.Payload(), "abcd", "Received Payload")
		assert.Equal(t, calledWith.Target(), instance, "received Target")
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
		assert.Nil(t, calledWith, "Handler was not called")
	})

	t.Run("RemoveAllHandlers", func(t *testing.T) {
		var calledWith events.Event
		handler := func(e events.Event) {
			calledWith = e
		}
		instance := events.NewEmitter()
		instance.On("fake-event", handler)
		found := instance.RemoveAllHandlers()
		assert.True(t, found, "Expected to find handlers")
		instance.Emit(events.New("fake-event", nil, nil))
		assert.Nil(t, calledWith, "Handler was not called")

		instance.On("fake-event", handler)
		instance.Emit(events.New("fake-event", nil, nil))
		assert.NotNil(t, calledWith, "Handler was called")
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

		assert.True(t, found, "Expected to find removable handlers")
		instance.Emit(events.New("fake-event-2", nil, nil))
		assert.Nil(t, calledWith, "Handler was not called")
		instance.Emit(events.New("fake-event-1", nil, nil))
		assert.NotNil(t, calledWith, "Handler was called")
	})

	t.Run("RemoveAllHandlers returns false if none present", func(t *testing.T) {
		instance := events.NewEmitter()
		assert.False(t, instance.RemoveAllHandlers(), "Expected no handlers")
	})

	t.Run("RemoveAllHandlersFor returns false if none present", func(t *testing.T) {
		instance := events.NewEmitter()
		assert.False(t, instance.RemoveAllHandlersFor("no-event"), "Expected no handlers")
	})

	t.Run("Events bubble on Component", func(t *testing.T) {
		var root, one, two, three, four Displayable
		var received []events.Event
		var receivers []Displayable
		var getHandlerFor = func(d Displayable) events.EventHandler {
			return func(e events.Event) {
				receivers = append(receivers, d)
				received = append(received, e)
			}
		}

		root = Box(context.New(), ID("root"), Children(func(c Context) {
			one = Box(c, ID("one"), Children(func() {
				two = Box(c, ID("two"), Children(func() {
					three = Box(c, ID("three"))
				}))
			}))
			four = Box(c, ID("four"))
		}))

		root.On("fake-event", getHandlerFor(root))
		one.On("fake-event", getHandlerFor(one))
		two.On("fake-event", getHandlerFor(two))
		three.On("fake-event", getHandlerFor(three))
		four.On("fake-event", getHandlerFor(four))

		three.Bubble(events.New("fake-event", three, nil))
		four.Emit(events.New("fake-event", nil, nil))

		assert.Equal(t, len(received), 5)
		assert.Equal(t, receivers[0].Path(), "/root/one/two/three")
		assert.Equal(t, receivers[1].Path(), "/root/one/two")
		assert.Equal(t, receivers[2].Path(), "/root/one")
		assert.Equal(t, receivers[3].Path(), "/root")
		assert.Equal(t, receivers[4].Path(), "/root/four")
	})

	t.Run("Events can be cancelled", func(t *testing.T) {
		secondCalled := false

		instance := events.NewEmitter()

		instance.On("fake-event", func(e events.Event) {
			e.Cancel()
		})
		instance.On("fake-event", func(e events.Event) {
			secondCalled = true
		})
		instance.Emit(events.New("fake-event", nil, nil))
		assert.False(t, secondCalled, "Expected Cancel to stop event")
	})
}
