package opts

import (
	"github.com/waybeams/assert"
	"events"
	"testing"
	. "ui"
	"uiold/context"
	"ui/controls"
	. "uiold/optspts"
)

func TestControlOptions(t *testing.T) {
	t.Run("Children", func(t *testing.T) {

		t.Run("Simple composer", func(t *testing.T) {
			box := controls.Box(context.New(), Children(func(c Context) {
				controls.Box(c)
			}))

			assert.Equal(box.ChildCount(), 1)
		})

		t.Run("Last received compose function is used", func(t *testing.T) {
			var first, second bool
			controls.Box(context.New(), Children(func() { first = true }), Children(func() { second = true }))

			assert.False(first, "Did not expect first composer to get called")
			assert.True(second, "Expected second Children handler")
		})
	})

	t.Run("OnConfigured", func(t *testing.T) {
		t.Run("Single sub", func(t *testing.T) {
			var calledWith events.Event
			var configuredHandler = func(e events.Event) {
				calledWith = e
			}
			box := controls.Box(context.New(), OnConfigured(configuredHandler))
			assert.NotNil(calledWith, "Expected event")
			if calledWith != nil {
				assert.Equal(box, calledWith.Target())
			}
		})
	})
}
