package control

import (
	"assert"
	"clock"
	"testing"
	. "ui"
	"uiold/context"
	. "ui/controls"
	. "uiold/opts"
)

func TestUpdateable(t *testing.T) {
	t.Run("Render Node", func(t *testing.T) {
		t.Skip()
		textValue := "abcd"

		var one, two, three Displayable
		var rootClosureCallCount, oneClosureCallCount int

		root := Box(context.New(), ID("root"), Children(func(c Context) {
			rootClosureCallCount++
			one = Box(c, ID("one"), Children(func(c Context) {
				oneClosureCallCount++
				two = Box(c, ID("two"), Text(textValue))
				three = Box(c, ID("three"), Text("wxyz"))
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
			t.Error("Expected the inner control to be re-instantiated")
		}

		assert.Equal(t, len(dirtyNodes), 1)
		assert.Equal(t, rootClosureCallCount, 1, "Root closure should NOT have been called again")
		assert.Equal(t, oneClosureCallCount, 2, "inner closure should have run twice")
		assert.Equal(t, one.ChildCount(), 2, "Children are rebuilt")
		assert.Equal(t, two.Text(), "efgh")
		assert.Equal(t, three.Text(), "wxyz")
	})

	t.Run("Does not replace identical control", func(t *testing.T) {
		fakeClock := clock.NewFake()
		root := Box(context.New(context.Clock(fakeClock)), Children(func(c Context) {
			Box(c, Key("abcd"))
		}))

		assert.NotNil(t, root)
	})
}
