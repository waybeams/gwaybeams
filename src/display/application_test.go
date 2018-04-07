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
		nodes := root.GetInvalidNodes()
		assert.Equal(t, len(nodes), 1)
		if !root.ShouldValidate() {
			t.Error("Expected the node to require validation")
		}
	})

	t.Run("Render Node", func(t *testing.T) {
		textValue := "abcd"

		var one, two, three Displayable
		var rootClosureCallCount, oneClosureCallCount int

		root, _ := Application(NewBuilder(), ID("root"), Children(func(b Builder) {
			rootClosureCallCount++
			one, _ = Box(b, ID("one"), Children(func(b Builder) {
				oneClosureCallCount++
				two, _ = Box(b, ID("two"), Text(textValue))
				three, _ = Box(b, ID("three"), Text("wxyz"))
			}))
		}))
		assert.Equal(t, rootClosureCallCount, 1)
		assert.Equal(t, oneClosureCallCount, 1)
		assert.NotNil(t, root)
		assert.Equal(t, two.GetText(), "abcd")
		assert.Equal(t, three.GetText(), "wxyz")

		firstInstanceOfTwo := two
		// Update a derived value
		textValue = "efgh"
		// Invalidate a nested child
		two.Invalidate()
		// Run validation from Root
		root.Validate()

		if firstInstanceOfTwo == two {
			t.Error("Expected the inner component to be re-instantiated")
		}

		assert.Equal(t, rootClosureCallCount, 1, "Root closure should NOT have been called again")
		assert.Equal(t, oneClosureCallCount, 2, "inner closure should have run twice")
		assert.Equal(t, one.GetChildCount(), 2, "Children are rebuilt")
		assert.Equal(t, two.GetText(), "efgh")
		assert.Equal(t, three.GetText(), "wxyz")
	})
}
