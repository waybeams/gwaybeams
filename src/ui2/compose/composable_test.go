package compose_test

import (
	"assert"
	"testing"
	"ui2/compose"
)

func TestComposable(t *testing.T) {

	t.Run("Children", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			root := compose.New()
			assert.Equal(t, len(compose.Children(root)), 0)
		})

		t.Run("single", func(t *testing.T) {
			root := compose.New()
			one := compose.New()
			compose.AddChild(root, one)
			assert.Equal(t, len(compose.Children(root)), 1)
		})
	})

	t.Run("AddChild", func(t *testing.T) {
		t.Run("single", func(t *testing.T) {
			root := compose.New()
			one := compose.New()
			compose.AddChild(root, one)
			assert.Equal(t, compose.ChildCount(root), 1)
			assert.Equal(t, compose.Parent(one), root)
		})

		t.Run("Cannot add to self", func(t *testing.T) {
			root := compose.New()

			assert.Panic(t, "Cannot add Composite to self", func() {
				compose.AddChild(root, root)
			})
		})
	})

	t.Run("ChildAt", func(t *testing.T) {
		root := compose.New()
		one := compose.New()
		two := compose.New()
		three := compose.New()
		compose.AddChild(root, one)
		compose.AddChild(root, two)
		compose.AddChild(root, three)

		assert.Equal(t, compose.ChildCount(root), 3)
		assert.Equal(t, compose.ChildAt(root, 0), one)
		assert.Equal(t, compose.ChildAt(root, 1), two)
		assert.Equal(t, compose.ChildAt(root, 2), three)
	})

	t.Run("Root", func(t *testing.T) {
		root := compose.New()
		one := compose.New()
		two := compose.New()
		three := compose.New()
		compose.AddChild(root, one)
		compose.AddChild(one, two)
		compose.AddChild(two, three)

		assert.Equal(t, root, compose.Root(root))
		assert.Equal(t, root, compose.Root(one))
		assert.Equal(t, root, compose.Root(two))
		assert.Equal(t, root, compose.Root(three))
	})
}
