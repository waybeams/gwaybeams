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
		assert.Equal(child1.GetParent().GetId(), root.GetId())
		assert.Equal(child2.GetParent().GetId(), root.GetId())
		assert.Nil(root.GetParent())
	})

	t.Run("GetChildCount", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		three := NewSprite()
		root.AddChild(one)
		one.AddChild(two)
		one.AddChild(three)
		assert.Equal(root.GetChildCount(), 1)
		assert.Equal(root.GetChildAt(0), one)

		assert.Equal(one.GetChildCount(), 2)
		assert.Equal(one.GetChildAt(0), two)
		assert.Equal(one.GetChildAt(1), three)
	})
}
