package display

import (
	"assert"
	"testing"
)

func TestSprite(t *testing.T) {
	instance := NewSprite()
	assert.NotNil(instance)

	t.Run("AddChild", func(t *testing.T) {
		child1 := NewSprite()
		child2 := NewSprite()
		assert.Equal(instance.AddChild(child1), 1)
		assert.Equal(instance.AddChild(child2), 2)
		assert.Equal(child1.Parent().Id(), instance.Id())
		assert.Equal(child2.Parent().Id(), instance.Id())
		assert.Nil(instance.Parent())
	})
}
