package display

import (
	"testing"
)

type FakeComponent struct {
	SpriteComponent
}

func NewFake() Displayable {
	return &FakeComponent{}
}

// Create a new factory using our component creation function reference.
var Fake = NewComponentFactory(NewFake)

func TestComponentFactory(t *testing.T) {

	t.Run("Custom type", func(t *testing.T) {
		b := NewBuilder()
		Fake(b)
	})
}
