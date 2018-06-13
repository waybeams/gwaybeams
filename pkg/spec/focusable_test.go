package spec_test

import (
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

func TestFocusable(t *testing.T) {
	var createTree = func() (root, one, two spec.ReadWriter) {
		root = ctrl.Box(
			opts.Key("one"),
			opts.Child(ctrl.Box(
				opts.Key("one"),
				opts.Child(ctrl.Box(
					opts.Key("two"),
				)),
			)),
		)
		one = spec.FirstByKey(root, "one")
		two = spec.FirstByKey(root, "two")
		return root, one, two
	}

	t.Run("SetFocusedSpec fails if not root", func(t *testing.T) {
		root, one, two := createTree()

		root.SetFocusedSpec(one)
		assert.Equal(root.FocusedSpec().Key(), "one")
		assert.Equal(one.FocusedSpec().Key(), "one")
		assert.Equal(two.FocusedSpec().Key(), "one")
	})

	t.Run("SetFocusedSpec travels up the tree", func(t *testing.T) {
		root, one, two := createTree()

		two.SetFocusedSpec(one)
		assert.Equal(root.FocusedSpec().Key(), "one")
		assert.Equal(one.FocusedSpec().Key(), "one")
		assert.Equal(two.FocusedSpec().Key(), "one")
	})

	t.Run("Children late bind to FocusedSpec", func(t *testing.T) {
		root, one, two := createTree()

		two.SetFocusedSpec(one)
		assert.Equal(two.FocusedSpec().Key(), "one")
		assert.Equal(root.FocusedSpec().Key(), "one")

		// Ensure that both child and parent do not miss cache clearing
		one.SetFocusedSpec(two)
		assert.Equal(two.FocusedSpec().Key(), "two")
		assert.Equal(root.FocusedSpec().Key(), "two")
	})
}
