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
		instance.Emit(NewEvent("fake-event", instance, "abcd"))
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
		instance.Emit(NewEvent("fake-event", nil, nil))
		assert.Nil(t, calledWith, "Handler was not called")
	})

	t.Run("RemoveAllHandlers", func(t *testing.T) {
		var calledWith Event
		handler := func(e Event) {
			calledWith = e
		}
		instance := NewEmitter()
		instance.AddHandler("fake-event", handler)
		found := instance.RemoveAllHandlers()
		assert.True(t, found, "Expected to find handlers")
		instance.Emit(NewEvent("fake-event", nil, nil))
		assert.Nil(t, calledWith, "Handler was not called")

		instance.AddHandler("fake-event", handler)
		instance.Emit(NewEvent("fake-event", nil, nil))
		assert.NotNil(t, calledWith, "Handler was called")
	})

	t.Run("RemoveAllHandlersFor", func(t *testing.T) {
		var calledWith Event
		handler := func(e Event) {
			calledWith = e
		}
		instance := NewEmitter()
		instance.AddHandler("fake-event-1", handler)
		instance.AddHandler("fake-event-2", handler)
		found := instance.RemoveAllHandlersFor("fake-event-2")

		assert.True(t, found, "Expected to find removable handlers")
		instance.Emit(NewEvent("fake-event-2", nil, nil))
		assert.Nil(t, calledWith, "Handler was not called")
		instance.Emit(NewEvent("fake-event-1", nil, nil))
		assert.NotNil(t, calledWith, "Handler was called")
	})

	t.Run("RemoveAllHandlers returns false if none present", func(t *testing.T) {
		instance := NewEmitter()
		assert.False(t, instance.RemoveAllHandlers(), "Expected no handlers")
	})

	t.Run("RemoveAllHandlersFor returns false if none present", func(t *testing.T) {
		instance := NewEmitter()
		assert.False(t, instance.RemoveAllHandlersFor("no-event"), "Expected no handlers")
	})

	t.Run("Events bubble on Component", func(t *testing.T) {
		var root, one, two, three, four Displayable
		var received []Event
		var receivers []Displayable
		var getHandlerFor = func(d Displayable) EventHandler {
			return func(e Event) {
				receivers = append(receivers, d)
				received = append(received, e)
			}
		}

		root, _ = Box(NewBuilder(), ID("root"), Children(func(b Builder) {
			one, _ = Box(b, ID("one"), Children(func(b Builder) {
				two, _ = Box(b, ID("two"), Children(func(b Builder) {
					three, _ = Box(b, ID("three"), Children(func(b Builder) {
					}))
				}))
			}))
			four, _ = Box(b, ID("four"))
		}))

		root.AddHandler("fake-event", getHandlerFor(root))
		one.AddHandler("fake-event", getHandlerFor(one))
		two.AddHandler("fake-event", getHandlerFor(two))
		three.AddHandler("fake-event", getHandlerFor(three))
		four.AddHandler("fake-event", getHandlerFor(four))

		three.Bubble(NewEvent("fake-event", three, nil))
		four.Emit(NewEvent("fake-event", nil, nil))

		assert.Equal(t, len(received), 5)
		assert.Equal(t, receivers[0].Path(), "/root/one/two/three")
		assert.Equal(t, receivers[1].Path(), "/root/one/two")
		assert.Equal(t, receivers[2].Path(), "/root/one")
		assert.Equal(t, receivers[3].Path(), "/root")
		assert.Equal(t, receivers[4].Path(), "/root/four")
	})

	t.Run("Events can be cancelled", func(t *testing.T) {
		secondCalled := false

		instance := NewEmitter()

		instance.AddHandler("fake-event", func(e Event) {
			e.Cancel()
		})
		instance.AddHandler("fake-event", func(e Event) {
			secondCalled = true
		})
		instance.Emit(NewEvent("fake-event", nil, nil))
		assert.False(t, secondCalled, "Expected Cancel to stop event")
	})
}
