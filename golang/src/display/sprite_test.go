package display

import (
	"assert"
	"testing"
)

func TestSprite(t *testing.T) {
	root := NewSprite()
	root.Width(200)
	assert.NotNil(root)
	/*

		t.Run("AddChild", func(t *testing.T) {
			child1 := NewSprite()
			child2 := NewSprite()
			assert.Equal(root.AddChild(child1), 1)
			assert.Equal(root.AddChild(child2), 2)
			assert.Equal(child1.Parent().Id(), root.Id())
			assert.Equal(child2.Parent().Id(), root.Id())
			assert.Nil(root.Parent())
		})
	*/
}
