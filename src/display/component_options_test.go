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

		t.Run("Accessor fails with second Children application", func(t *testing.T) {
			box, err := Box(NewBuilder(), Children(func() {}), Children(func() {}))

			assert.Nil(t, box)
			assert.NotNil(t, err, "Error should be returned")
			assert.Match(t, "single Compose function", err.Error())
		})
	})
}
