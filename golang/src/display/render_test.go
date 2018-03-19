package display

import (
	"assert"
	"testing"
)

func TestRender(t *testing.T) {
	t.Run("Render single node", func(t *testing.T) {
		root := NewSprite()
		assert.NotNil(root)
	})
}
