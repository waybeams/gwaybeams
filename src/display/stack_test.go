package display

import (
	"assert"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Default state", func(t *testing.T) {
		stack := NewDisplayStack()
		assert.NotNil(t, stack)
		assert.False(t, stack.HasNext())
	})

	t.Run("HasNext supports Pop", func(t *testing.T) {
		stack := NewDisplayStack()
		one := NewComponent()
		stack.Push(one)
		assert.True(t, stack.HasNext())
	})

	t.Run("Pop returns first element", func(t *testing.T) {
		stack := NewDisplayStack()
		one := NewComponent()
		stack.Push(one)
		assert.Equal(t, stack.Pop(), one)
	})

	t.Run("Pop returns each element", func(t *testing.T) {
		stack := NewDisplayStack()
		one := NewComponent()
		two := NewComponent()
		three := NewComponent()
		four := NewComponent()

		stack.Push(one)
		stack.Push(two)
		stack.Push(three)
		stack.Push(four)

		assert.Equal(t, stack.Peek(), four)
		assert.Equal(t, stack.Pop(), four)
		assert.True(t, stack.HasNext())

		assert.Equal(t, stack.Peek(), three)
		assert.Equal(t, stack.Pop(), three)
		assert.True(t, stack.HasNext())

		assert.Equal(t, stack.Peek(), two)
		assert.Equal(t, stack.Pop(), two)
		assert.True(t, stack.HasNext())

		assert.Equal(t, stack.Peek(), one)
		assert.Equal(t, stack.Pop(), one)
		assert.False(t, stack.HasNext())
	})

	t.Run("Peek returns nil if no next element", func(t *testing.T) {
		stack := NewDisplayStack()
		assert.Equal(t, stack.Peek(), nil)
	})

	t.Run("Push does not accept nil value", func(t *testing.T) {
		stack := NewDisplayStack()
		err := stack.Push(nil)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "display.DisplayStack does not accept nil entries")
	})
}
