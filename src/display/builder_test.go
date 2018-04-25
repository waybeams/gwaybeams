package display

import (
	"assert"
	"fmt"
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
			root, _ := Box(NewBuilder(), Children(composer))
			if !wasCalled {
				t.Error("Expected composer to be called")
			}
			if root.GetComposeEmpty() == nil {
				t.Error("Expected compose empty")
			}
			if root.GetComposeWithBuilder() != nil {
				t.Error("Did not expect builder")
			}
			if root.GetComposeWithComponent() != nil {
				t.Error("Did not expect renderComponent")
			}
		})

		t.Run("Builder", func(t *testing.T) {
			var calledWith Builder
			composer := func(b Builder) {
				calledWith = b
			}
			root, _ := Box(NewBuilder(), Children(composer))
			if calledWith == nil {
				t.Error("Expected builder in call")
			}
			if root.GetComposeWithBuilder() == nil {
				t.Error("Expected node configured")
			}
			if root.GetComposeEmpty() != nil {
				t.Error("Did not expect default value")
			}
			if root.GetComposeWithComponent() != nil {
				t.Error("Did not expect renderComponent")
			}
		})

		t.Run("Composition", func(t *testing.T) {
			var calledWith Displayable
			composer := func(d Displayable) {
				calledWith = d
			}
			root, err := Box(NewBuilder(), Children(composer))
			if err != nil {
				t.Error(err)
			}
			if calledWith == nil {
				t.Error("Expected call with component")
			}
			if root.GetComposeWithComponent() == nil {
				t.Error("Expected ComposeWithComponent to be configured")
			}
			if root.GetComposeEmpty() != nil {
				t.Error("Did not expect default value")
			}
			if root.GetComposeWithBuilder() != nil {
				t.Error("Did not expect builder")
			}
		})

		t.Run("Displayable", func(t *testing.T) {
			t.Run("returned when requested", func(t *testing.T) {
				var returned Displayable
				b := NewBuilder()
				box, _ := Box(b, ID("abcd"), Children(func(d Displayable) {
					returned = d
				}))
				assert.Equal(t, returned.ID(), box.ID())
			})

		})

		t.Run("Listen", func(t *testing.T) {
			root, _ := Box(NewBuilder())
			defer root.Builder().Destroy()
			go root.Builder().Listen()
		})
	})

	t.Run("Updates", func(t *testing.T) {

		t.Run("Callable", func(t *testing.T) {
			var firstInstance Displayable
			message := "abcd"
			root, _ := VBox(NewBuilder(), Children(func(b Builder) {
				firstInstance, _ = Label(b, Text(message))
			}))
			message = "efgh"
			assert.Equal(t, root.ChildAt(0).Text(), "abcd")
			assert.Equal(t, root.ChildAt(0), firstInstance)

			// In real use-case, the NewBuilder() will be called with
			// only those nodes that have been invalidated.
			root.Builder().Update(root)

			assert.Equal(t, root.ChildCount(), 1)
			assert.Equal(t, root.ChildAt(0).Text(), "efgh")
		})

		t.Run("Removed children", func(t *testing.T) {
			count := 3
			root, _ := VBox(NewBuilder(), Children(func(b Builder) {
				for i := 0; i < count; i++ {
					Button(b, ID(fmt.Sprintf("btn-%v", +i)))
				}
			}))
			assert.Equal(t, root.ChildCount(), 3)
			count = 1
			root.Builder().Update(root)
			assert.Equal(t, root.ChildCount(), 1)
		})

		t.Run("Added children", func(t *testing.T) {
			count := 1
			root, _ := VBox(NewBuilder(), Children(func(b Builder) {
				for i := 0; i < count; i++ {
					Button(b, ID(fmt.Sprintf("btn-%v", +i)))
				}
			}))
			assert.Equal(t, root.ChildCount(), 1)
			count = 3
			root.Builder().Update(root)
			assert.Equal(t, root.ChildCount(), 3)

		})

		t.Run("Reordered same children", func(t *testing.T) {
			ids := []string{"abcd", "efgh", "ijkl"}

			root, _ := VBox(NewBuilder(), Children(func(b Builder) {
				for _, id := range ids {
					Button(b, ID(id))
				}
			}))
			assert.Equal(t, root.ChildAt(0).ID(), "abcd")
			assert.Equal(t, root.ChildAt(1).ID(), "efgh")
			assert.Equal(t, root.ChildAt(2).ID(), "ijkl")

			ids = []string{"efgh", "ijkl", "abcd"}
			root.Builder().Update(root)
			assert.Equal(t, root.ChildCount(), 3)
			assert.Equal(t, root.ChildAt(0).ID(), "efgh")
			assert.Equal(t, root.ChildAt(1).ID(), "ijkl")
			assert.Equal(t, root.ChildAt(2).ID(), "abcd")
		})

		t.Run("Inserted child of different type", func(t *testing.T) {
			constr := Box
			root, _ := Box(NewBuilder(), Children(func(b Builder) {
				constr(b)
			}))
			assert.Equal(t, root.ChildAt(0).TypeName(), "Box")
			constr = Button
			root.Builder().Update(root)
			assert.Equal(t, root.ChildAt(0).TypeName(), "Button")
			assert.Equal(t, root.ChildCount(), 1)
		})
	})
}
