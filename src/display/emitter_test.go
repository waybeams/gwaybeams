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

	t.Run("RemoveAllHandlers", func(t *testing.T) {
		var calledWith Event
		handler := func(e Event) {
			calledWith = e
		}
		instance := NewEmitter()
		instance.AddHandler("fake-event", handler)
		found := instance.RemoveAllHandlers()
		assert.True(t, found, "Expected to find handlers")
		instance.Emit("fake-event", nil)
		assert.Nil(t, calledWith, "Handler was not called")

		instance.AddHandler("fake-event", handler)
		instance.Emit("fake-event", nil)
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
		instance.Emit("fake-event-2", nil)
		assert.Nil(t, calledWith, "Handler was not called")
		instance.Emit("fake-event-1", nil)
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

	/*
		t.Run("AddCaptureHandler", func(t *testing.T) {
			var calledWith Event
			handler := func(e Event) {
				calledWith = e
			}
			instance := NewEmitter()
			instance.AddCaptureHandler("fake-event", handler)
			instance.Emit("fake-event", "abcd")
			assert.Nil(t, calledWith, "Emit should not bubble!")

			instance.Bubble("fake-event", "abcd")
		})

		/*
			t.Run("Events bubble", func(t *testing.T) {
				var root, one, two, three Displayable
				root, _ = Box(NewBuilder(), Children(func(b Builder) {
					one, _ = Box(b, Children(func(b Builder) {
						two, _ = Box(b, Children(func(b Builder) {
							three, _ = Box(b, Children(func(b Builder) {
							}))
						}))
					}))
				}))

				var rootHandler = func(e Event) {
				}
				var oneHandler = func(e Event) {
				}
				var twoHandler = func(e Event) {
				}
				var threeHandler = func(e Event) {
				}
				root.AddHandler("fake-event", rootHandler)
				one.AddHandler("fake-event", oneHandler)
				two.AddHandler("fake-event", twoHandler)
				three.AddHandler("fake-event", threeHandler)
			})
	*/
}
