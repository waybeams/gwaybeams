package display

import (
	"assert"
	"testing"
)

func TestSprite(t *testing.T) {

	t.Run("AddChild", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		root.Width(200)
		assert.Equal(root.AddChild(one), 1)
		assert.Equal(root.AddChild(two), 2)
		assert.Equal(one.GetParent().GetId(), root.GetId())
		assert.Equal(two.GetParent().GetId(), root.GetId())
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

	t.Run("GetChildren returns empty list", func(t *testing.T) {
		root := NewSprite()
		children := root.GetChildren()

		if children == nil {
			t.Error("GetChildren should not return nil")
		}

		assert.Equal(len(children), 0)
	})

	t.Run("GetChildren returns new list", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		three := NewSprite()

		root.AddChild(one)
		root.AddChild(two)
		root.AddChild(three)

		children := root.GetChildren()
		assert.Equal(len(children), 3)
	})
}
