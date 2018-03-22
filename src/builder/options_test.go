package builder

import (
	"display"
	"testing"
)

type FakeComponent struct {
	display.SpriteComponent
}

func NewFake() display.Displayable {
	return &FakeComponent{}
}

// Create a new factory using our component creation function reference.
var Fake = NewComponentFactory(NewFake)

func TestComponentFactory(t *testing.T) {

	t.Run("Custom type", func(t *testing.T) {
		b, _ := New()
		Fake(b)
	})
}
