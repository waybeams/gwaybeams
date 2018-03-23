package display

import (
	"testing"
)

func TestBox(t *testing.T) {

	t.Run("BoxComponent creation", func(t *testing.T) {
		instance := NewBox()
		if instance == nil {
			t.Error("Expected a new BoxComponent")
		}
	})

	t.Run("AddChild", func(t *testing.T) {
		instance := NewBox()
		one := NewBox()
		two := NewBox()
		instance.AddChild(one)
		instance.AddChild(two)

		if instance.GetChildCount() != 2 {
			t.Error("Expected two children")
		}
	})
}
