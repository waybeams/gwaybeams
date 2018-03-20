package display

import (
	"assert"
	"testing"
)

func TestBox(t *testing.T) {
	surface := &FakeSurface{}
	instance := NewBox()
	assert.NotNil(instance)

	t.Run("Box creation", func(t *testing.T) {
		box2 := Box(surface)
		assert.NotNil(box2)
	})

	t.Run("Box caller does not have to hold result", func(t *testing.T) {
		Box(surface)
	})

	t.Run("AddChild", func(t *testing.T) {
		child1 := NewBox()
		child2 := NewBox()
		assert.Equal(instance.AddChild(child1), 1)
		assert.Equal(instance.AddChild(child2), 2)
		assert.Equal(child1.GetParent().GetId(), instance.GetId())
	})
}
