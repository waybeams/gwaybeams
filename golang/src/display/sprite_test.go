package display

import (
	"assert"
	"testing"
)

func TestSprite(t *testing.T) {

	t.Run("AddChild", func(t *testing.T) {
		root := NewSprite()
		root.Width(200)
		child1 := NewSprite()
		child2 := NewSprite()
		assert.Equal(root.AddChild(child1), 1)
		assert.Equal(root.AddChild(child2), 2)
		assert.Equal(child1.Parent().Id(), root.Id())
		assert.Equal(child2.Parent().Id(), root.Id())
		assert.Nil(root.Parent())
	})

	t.Run("ChildCount", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		three := NewSprite()
		root.AddChild(one)
		one.AddChild(two)
		one.AddChild(three)
		assert.Equal(root.ChildCount(), 1)
		assert.Equal(root.ChildAt(0), one)

		assert.Equal(one.ChildCount(), 2)
		assert.Equal(one.ChildAt(0), two)
		assert.Equal(one.ChildAt(1), three)
	})
}
