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
				t.Error("Did not expect renderScheduler")
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
				t.Error("Did not expect renderScheduler")
			}
		})

		t.Run("RenderScheduler", func(t *testing.T) {
			var calledWith Displayable
			composer := func(d Displayable) {
				calledWith = d
			}
			root, err := Box(NewBuilder(), Children(composer))
			if err != nil {
				t.Error(err)
			}
			if calledWith == nil {
				t.Error("Expected call with scheduler")
			}
			if root.GetComposeWithComponent() == nil {
				t.Error("Expected ComposeWithScheduler to be configured")
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
				// nodes := b.GetInvalidNodes()
				assert.Equal(t, returned.GetID(), box.GetID())
			})

			t.Run("Builder prunes nested dirty nodes", func(t *testing.T) {
				var root, one, two, three Displayable
				var b Builder

				var setUp = func() {
					b = NewBuilder()
					Box(b, ID("root"), Children(func(d Displayable) {
						root = d
						Box(b, ID("one"), Children(func(d Displayable) {
							one = d
							Box(b, ID("two"), Children(func(d Displayable) {
								two = d
							}))
							Box(b, ID("three"), Children(func(d Displayable) {
								three = d
							}))
						}))
					}))

					nodes := b.GetInvalidNodes()
					assert.Equal(t, len(nodes), 0)
				}

				// Ensure we do not leave a timer running that will leak memory
				// and attempt to render our invalidated nodes.
				var tearDown = func() {
					b.Destroy()
				}

				t.Run("parent hides children", func(t *testing.T) {
					t.Skip()
					defer tearDown()
					setUp()

					three.Invalidate()
					two.Invalidate()
					one.Invalidate()
					root.Invalidate()

					nodes := b.GetInvalidNodes()
					assert.Equal(t, len(nodes), 1, "hide children")
					assert.Equal(t, nodes[0].GetID(), "root")
				})

				t.Run("Siblings are sorted fifo", func(t *testing.T) {
					t.Skip()
					defer tearDown()
					setUp()

					three.Invalidate()
					two.Invalidate()
					nodes := b.GetInvalidNodes()
					assert.Equal(t, len(nodes), 2, "Expected two")
					assert.Equal(t, nodes[0].GetID(), "three")
					assert.Equal(t, nodes[1].GetID(), "two")
				})
			})
		})
	})
}
