package display

import (
	"assert"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Default state", func(t *testing.T) {
		stack := NewStack()
		assert.NotNil(stack)
		assert.False(stack.HasNext())
	})

	t.Run("HasNext supports Pop", func(t *testing.T) {
		stack := NewStack()
		one := NewSprite()
		stack.Push(one)
		assert.True(stack.HasNext())
	})

	t.Run("Pop returns first element", func(t *testing.T) {
		stack := NewStack()
		one := NewSprite()
		stack.Push(one)
		assert.Equal(stack.Pop(), one)
	})

	t.Run("Pop returns each element", func(t *testing.T) {
		stack := NewStack()
		one := NewSprite()
		two := NewSprite()
		three := NewSprite()
		four := NewSprite()

		stack.Push(one)
		stack.Push(two)
		stack.Push(three)
		stack.Push(four)

		assert.Equal(stack.Peek(), four)
		assert.Equal(stack.Pop(), four)
		assert.True(stack.HasNext())

		assert.Equal(stack.Peek(), three)
		assert.Equal(stack.Pop(), three)
		assert.True(stack.HasNext())

		assert.Equal(stack.Peek(), two)
		assert.Equal(stack.Pop(), two)
		assert.True(stack.HasNext())

		assert.Equal(stack.Peek(), one)
		assert.Equal(stack.Pop(), one)
		assert.False(stack.HasNext())
	})

	t.Run("Peek returns nil if no next element", func(t *testing.T) {
		stack := NewStack()
		assert.Equal(stack.Peek(), nil)
	})
}
