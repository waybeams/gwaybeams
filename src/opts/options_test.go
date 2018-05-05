package opts_test

import (
	"assert"
	"controls"
	"ctx"
	. "opts"
	"testing"
	. "ui"
)

func TestComponentOptions(t *testing.T) {
	t.Run("Children", func(t *testing.T) {

		t.Run("Simple composer", func(t *testing.T) {
			box := controls.Box(ctx.New(), Children(func(c Context) {
				controls.Box(c)
			}))

			assert.Equal(t, box.ChildCount(), 1)
		})

		t.Run("Last received compose function is used", func(t *testing.T) {
			var first, second bool
			controls.Box(ctx.New(), Children(func() { first = true }), Children(func() { second = true }))

			assert.False(t, first, "Did not expect first composer to get called")
			assert.True(t, second, "Expected second Children handler")
		})
	})
}
