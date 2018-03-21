package display

import (
	"assert"
	"testing"
)

func TestLayout(t *testing.T) {
	root := NewSprite()

	t.Run("Call Layout", func(t *testing.T) {
		assert.NotNil(root)
	})

	t.Run("directionalDelegate", func(t *testing.T) {
		delegate := DirectionalDelegate(Horizontal)
		if delegate == nil {
			t.Error("Expected DirectionalDelegate to return a function")
		}
	})
}
