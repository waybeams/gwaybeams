package display

import (
	"assert"
	"fmt"
	"testing"
)

func TestUpdater(t *testing.T) {
	t.Run("Callable", func(t *testing.T) {
		var firstInstance Displayable
		message := "abcd"
		root, _ := VBox(NewUpdater(), Children(func(b Builder) {
			firstInstance, _ = Label(b, Text(message))
		}))
		message = "efgh"
		assert.Equal(t, root.ChildAt(0).Text(), "abcd")
		assert.Equal(t, root.ChildAt(0), firstInstance)

		// In real use-case, the NewUpdater() will be called with
		// only those nodes that have been invalidated.
		root.Builder().Update(root)
		assert.Equal(t, root.ChildCount(), 1)
		assert.Equal(t, root.ChildAt(0).Text(), "efgh")
	})

	t.Run("Removed children", func(t *testing.T) {
		count := 3
		root, _ := VBox(NewUpdater(), Children(func(b Builder) {
			for i := 0; i < count; i++ {
				Button(b, ID(fmt.Sprintf("btn-%v", +i)))
			}
		}))
		assert.Equal(t, root.ChildCount(), 3)
		count = 1
		root.Builder().Update(root)
		assert.Equal(t, root.ChildCount(), 1)
	})

	t.Run("Added children", func(t *testing.T) {
		count := 1
		root, _ := VBox(NewUpdater(), Children(func(b Builder) {
			for i := 0; i < count; i++ {
				Button(b, ID(fmt.Sprintf("btn-%v", +i)))
			}
		}))
		assert.Equal(t, root.ChildCount(), 1)
		count = 3
		root.Builder().Update(root)
		assert.Equal(t, root.ChildCount(), 3)

	})

	t.Run("Reordered same children", func(t *testing.T) {
		ids := []string{"abcd", "efgh", "ijkl"}

		root, _ := VBox(NewUpdater(), Children(func(b Builder) {
			for _, id := range ids {
				Button(b, ID(id))
			}
		}))
		assert.Equal(t, root.ChildAt(0).ID(), "abcd")
		assert.Equal(t, root.ChildAt(1).ID(), "efgh")
		assert.Equal(t, root.ChildAt(2).ID(), "ijkl")

		ids = []string{"efgh", "ijkl", "abcd"}
		root.Builder().Update(root)
		assert.Equal(t, root.ChildCount(), 3)
		assert.Equal(t, root.ChildAt(0).ID(), "efgh")
		assert.Equal(t, root.ChildAt(1).ID(), "ijkl")
		assert.Equal(t, root.ChildAt(2).ID(), "abcd")
	})

	t.Run("Inserted child of different type", func(t *testing.T) {
		constr := Box
		root, _ := Box(NewUpdater(), Children(func(b Builder) {
			constr(b)
		}))
		assert.Equal(t, root.ChildAt(0).TypeName(), "Box")
		constr = Button
		root.Builder().Update(root)
		assert.Equal(t, root.ChildAt(0).TypeName(), "Button")
		assert.Equal(t, root.ChildCount(), 1)
	})
}
