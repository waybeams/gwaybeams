package display

import (
	"assert"
	"clock"
	"testing"
)

func TestUpdateable(t *testing.T) {
	t.Run("Render Node", func(t *testing.T) {
		t.Skip()
		textValue := "abcd"

		var one, two, three Displayable
		var rootClosureCallCount, oneClosureCallCount int

		root, _ := Box(NewBuilder(), ID("root"), Children(func(b Builder) {
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
		assert.Equal(t, two.Text(), "abcd")
		assert.Equal(t, three.Text(), "wxyz")

		firstInstanceOfTwo := two
		// Update a derived value
		textValue = "efgh"
		// Invalidate a nested child
		one.InvalidateChildren()
		// Run validation from Root
		dirtyNodes := root.RecomposeChildren()

		if firstInstanceOfTwo == two {
			t.Error("Expected the inner component to be re-instantiated")
		}

		assert.Equal(t, len(dirtyNodes), 1)
		assert.Equal(t, rootClosureCallCount, 1, "Root closure should NOT have been called again")
		assert.Equal(t, oneClosureCallCount, 2, "inner closure should have run twice")
		assert.Equal(t, one.ChildCount(), 2, "Children are rebuilt")
		assert.Equal(t, two.Text(), "efgh")
		assert.Equal(t, three.Text(), "wxyz")
	})

	t.Run("Does not replace identical components", func(t *testing.T) {
		fakeClock := clock.NewFake()
		root, _ := Box(NewBuilderUsing(fakeClock), Children(func(b Builder) {
			Box(b, Key("abcd"))
		}))

		assert.NotNil(t, root)
	})
}
