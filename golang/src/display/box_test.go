package display

import (
	"assert"
	"testing"
)

func TestBox(t *testing.T) {
	instance := NewBox()
	assert.NotNil(instance)

	t.Run("AddChild", func(t *testing.T) {
		child1 := NewBox()
		child2 := NewBox()
		assert.Equal(instance.AddChild(child1), 1)
		assert.Equal(instance.AddChild(child2), 2)
		assert.Equal(child1.Parent().Id(), instance.Id())
	})
}
