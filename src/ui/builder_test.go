package ui_test

import (
	"assert"
	"fmt"
	"testing"
	. "ui"
	"ui/context"
	. "ui/controls"
	. "ui/opts"
)

func TestBuilder(t *testing.T) {
	t.Run("Instantiated", func(t *testing.T) {
		b := NewBuilder()
		if b == nil {
			t.Error("Expected builder instance")
		}
	})

	t.Run("Compose function can request an instance of the Builder", func(t *testing.T) {
		var child Displayable
		var wasCalled = false
		Box(context.New(), Children(func(c Context) {
			wasCalled = true
			if c == nil {
				t.Error("Expected builder to be returned to first child")
			}
			child = Box(c, ID("one"))
		}))
		if !wasCalled {
			t.Error("Inner composition function was not called")
		}
	})

	t.Run("Builds provided elements", func(t *testing.T) {
		sprite := Box(context.New(), Width(200), Height(100))
		if sprite == nil {
			t.Error("Expected root displayable to be returned")
		}
		if sprite.Width() != 200.0 {
			t.Errorf("Expected sprite width to be set but was %f", sprite.Width())
		}
		if sprite.Height() != 100 {
			t.Errorf("Expected sprite height to be set but was %f", sprite.Height())
		}
	})

	t.Run("Composer", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			var wasCalled = false
			composer := func() {
				wasCalled = true
			}
			root := Box(context.New(), Children(composer))
			if !wasCalled {
				t.Error("Expected composer to be called")
			}
			if root.GetComposeEmpty() == nil {
				t.Error("Expected compose empty")
			}
			if root.GetComposeWithContext() != nil {
				t.Error("Did not expect builder")
			}
			if root.GetComposeWithControl() != nil {
				t.Error("Did not expect renderControl")
			}
		})

		t.Run("Builder", func(t *testing.T) {
			var calledWith Context
			composer := func(c Context) {
				calledWith = c
			}
			root := Box(context.New(), Children(composer))
			if calledWith == nil {
				t.Error("Expected builder in call")
			}
			if root.GetComposeWithContext() == nil {
				t.Error("Expected node configured")
			}
			if root.GetComposeEmpty() != nil {
				t.Error("Did not expect default value")
			}
			if root.GetComposeWithControl() != nil {
				t.Error("Did not expect renderControl")
			}
		})

		t.Run("Composition", func(t *testing.T) {
			var calledWith Displayable
			composer := func(d Displayable) {
				calledWith = d
			}
			root := Box(context.New(), Children(composer))
			if calledWith == nil {
				t.Error("Expected call with control")
			}
			if root.GetComposeWithControl() == nil {
				t.Error("Expected ComposeWithControl to be configured")
			}
			if root.GetComposeEmpty() != nil {
				t.Error("Did not expect default value")
			}
			if root.GetComposeWithContext() != nil {
				t.Error("Did not expect builder")
			}
		})

		t.Run("Displayable", func(t *testing.T) {
			t.Run("returned when requested", func(t *testing.T) {
				var returned Displayable
				b := context.New()
				box := Box(b, ID("abcd"), Children(func(d Displayable) {
					returned = d
				}))
				assert.Equal(t, returned.ID(), box.ID())
			})

		})

		t.Run("Listen", func(t *testing.T) {
			root := Box(context.New())
			defer root.Context().Destroy()
			go root.Context().Listen()
		})
	})

	t.Run("Updates", func(t *testing.T) {

		t.Run("Callable", func(t *testing.T) {
			var firstInstance Displayable
			message := "abcd"
			root := VBox(context.New(), Children(func(c Context) {
				firstInstance = Label(c, Text(message))
			}))
			message = "efgh"
			assert.Equal(t, root.ChildAt(0).Text(), "abcd")
			assert.Equal(t, root.ChildAt(0), firstInstance)

			// In real use-case, the context.New() will be called with
			// only those nodes that have been invalidated.
			root.Context().Builder().Update(root)

			assert.Equal(t, root.ChildCount(), 1)
			assert.Equal(t, root.ChildAt(0).Text(), "efgh")
		})

		t.Run("Removed children", func(t *testing.T) {
			count := 3
			root := VBox(context.New(), Children(func(c Context) {
				for i := 0; i < count; i++ {
					Button(c, ID(fmt.Sprintf("btn-%v", +i)))
				}
			}))
			assert.Equal(t, root.ChildCount(), 3)
			count = 1
			root.Context().Builder().Update(root)
			assert.Equal(t, root.ChildCount(), 1)
		})

		t.Run("Added children", func(t *testing.T) {
			count := 1
			root := VBox(context.New(), Children(func(c Context) {
				for i := 0; i < count; i++ {
					Button(c, ID(fmt.Sprintf("btn-%v", +i)))
				}
			}))
			assert.Equal(t, root.ChildCount(), 1)
			count = 3
			root.Context().Builder().Update(root)
			assert.Equal(t, root.ChildCount(), 3)

		})

		t.Run("Reordered same children", func(t *testing.T) {
			ids := []string{"abcd", "efgh", "ijkl"}

			root := VBox(context.New(), Children(func(c Context) {
				for _, id := range ids {
					Button(c, ID(id))
				}
			}))
			assert.Equal(t, root.ChildAt(0).ID(), "abcd")
			assert.Equal(t, root.ChildAt(1).ID(), "efgh")
			assert.Equal(t, root.ChildAt(2).ID(), "ijkl")

			ids = []string{"efgh", "ijkl", "abcd"}
			root.Context().Builder().Update(root)
			assert.Equal(t, root.ChildCount(), 3)
			assert.Equal(t, root.ChildAt(0).ID(), "efgh")
			assert.Equal(t, root.ChildAt(1).ID(), "ijkl")
			assert.Equal(t, root.ChildAt(2).ID(), "abcd")
		})

		t.Run("Inserted child of different type", func(t *testing.T) {
			constr := Box
			root := Box(context.New(), Children(func(c Context) {
				constr(c)
			}))
			assert.Equal(t, root.ChildAt(0).TypeName(), "Box")
			constr = Button
			root.Context().Builder().Update(root)
			assert.Equal(t, root.ChildAt(0).TypeName(), "Button")
			assert.Equal(t, root.ChildCount(), 1)
		})
	})
}
