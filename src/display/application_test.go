package display

import (
	"assert"
	"testing"
)

func TestApplication(t *testing.T) {
	t.Run("Instantiated", func(t *testing.T) {
		app := NewApplication()
		if app == nil {
			t.Error("Expected application")
		}
	})

	t.Run("Invalidating children kicks off render timer", func(t *testing.T) {
		var one Displayable
		root, _ := Application(NewBuilder(), ID("root"), Children(func(b Builder) {
			one, _ = Box(b, ID("one"))
		}))

		one.Invalidate()
		nodes := root.InvalidNodes()
		assert.Equal(t, len(nodes), 1)
		if !root.ShouldValidate() {
			t.Error("Expected the node to require validation")
		}
	})

}
