package display

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Run("Instantiated", func(t *testing.T) {
		builder := NewBuilder()
		if builder == nil {
			t.Error("Expected builder instance")
		}
	})

	t.Run("Compose function can request an instance of the Builder", func(t *testing.T) {
		var child Displayable
		var wasCalled = false
		var childError error
		Box(NewBuilder(), Children(func(b Builder) {
			wasCalled = true
			if b == nil {
				t.Error("Expected builder to be returned to first child")
			}
			child, childError = Box(b, ID("one"))
		}))
		if !wasCalled {
			t.Error("Inner composition function was not called")
		}
		if childError != nil {
			t.Error(childError)
		} else if child == nil {
			t.Error("Child was not created and no error was thrown")
		}
	})

	t.Run("Builds provided elements", func(t *testing.T) {
		sprite, err := Box(NewBuilder(), Width(200), Height(100))
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
