package display

import (
	"assert"
	"testing"
)

func TestRectangle(t *testing.T) {
	t.Run("Rectangle draws", func(t *testing.T) {
		root := NewSprite()
		assert.NotNil(root)
	})
}
