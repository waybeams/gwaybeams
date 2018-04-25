package display

import (
	"assert"
	"testing"
)

func TestComponentOptions(t *testing.T) {
	t.Run("Children", func(t *testing.T) {

		t.Run("Simple composer", func(t *testing.T) {
			box, _ := Box(NewBuilder(), Children(func(b Builder) {
				Box(b)
			}))

			assert.Equal(t, box.ChildCount(), 1)
		})

		t.Run("Last received compose function is used", func(t *testing.T) {
			var first, second bool
			Box(NewBuilder(), Children(func() { first = true }), Children(func() { second = true }))

			assert.False(t, first, "Did not expect first composer to get called")
			assert.True(t, second, "Expected second Children handler")
		})
	})
}
