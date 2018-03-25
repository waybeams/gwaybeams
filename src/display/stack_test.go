package display

import (
	"assert"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Default state", func(t *testing.T) {
		stack := NewDisplayStack()
		assert.NotNil(stack)
		assert.False(stack.HasNext())
	})

	t.Run("HasNext supports Pop", func(t *testing.T) {
		stack := NewDisplayStack()
		one := NewComponent()
		stack.Push(one)
		assert.True(stack.HasNext())
	})

	t.Run("Pop returns first element", func(t *testing.T) {
		stack := NewDisplayStack()
		one := NewComponent()
		stack.Push(one)
		assert.TEqual(t, stack.Pop(), one)
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

		assert.TEqual(t, stack.Peek(), four)
		assert.TEqual(t, stack.Pop(), four)
		assert.True(stack.HasNext())

		assert.TEqual(t, stack.Peek(), three)
		assert.TEqual(t, stack.Pop(), three)
		assert.True(stack.HasNext())

		assert.TEqual(t, stack.Peek(), two)
		assert.TEqual(t, stack.Pop(), two)
		assert.True(stack.HasNext())

		assert.TEqual(t, stack.Peek(), one)
		assert.TEqual(t, stack.Pop(), one)
		assert.False(stack.HasNext())
	})

	t.Run("Peek returns nil if no next element", func(t *testing.T) {
		stack := NewDisplayStack()
		assert.TEqual(t, stack.Peek(), nil)
	})

	t.Run("Push does not accept nil value", func(t *testing.T) {
		stack := NewDisplayStack()
		err := stack.Push(nil)
		assert.NotNil(err)
		assert.TEqual(t, err.Error(), "display.DisplayStack does not accept nil entries")
	})
}
