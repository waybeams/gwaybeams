package display

import (
	"assert"
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Run("Instantiated", func(t *testing.T) {
		builder := NewBuilder()
		if builder == nil {
			t.Error("Expected builder instance")
		}
	})

	t.Run("Returns error when more than one root node is provided", func(t *testing.T) {
		builder := NewGlfwBuilder()
		box, err := builder.Build(func(b Builder) {
			Sprite(b)
			Sprite(b)
		})
		if err == nil {
			t.Error("Expected an error from builder")
		}
		assert.ErrorMatch("single root node", err)

		if box != nil {
			t.Errorf("Expected nil result with error state")
		}
	})

	t.Run("Builds provided elements", func(t *testing.T) {
		builder := NewBuilder()
		sprite, err := builder.Build(func(b Builder) {
			Sprite(b, Width(200), Height(100))
		})
		if err != nil {
			t.Error(err)
		}
		if sprite == nil {
			t.Error("Expected root displayable to be returned")
		}
		if sprite.GetWidth() != 200.0 {
			t.Errorf("Expected sprite width to be set but was %f", sprite.GetWidth())
		}
		if sprite.GetHeight() != 100 {
			t.Errorf("Expected sprite height to be set but was %f", sprite.GetHeight())
		}
	})
}
