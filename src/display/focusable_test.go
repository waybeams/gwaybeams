package display

import (
	"assert"
	"testing"
)

func TestFocusable(t *testing.T) {
	t.Run("Blurred", func(t *testing.T) {
		instance, _ := Button(NewBuilder(), Blurred())
		assert.False(t, instance.Focused())
	})

	t.Run("Focused", func(t *testing.T) {
		instance, _ := Button(NewBuilder(), Focused())
		assert.True(t, instance.Focused())
	})

	t.Run("Unfocuses previously focused elements", func(t *testing.T) {
		instance, _ := VBox(NewBuilder(), Children(func(b Builder) {
			Button(b, ID("abcd"))
			Button(b, ID("efgh"))
			Button(b, ID("ijkl"))
			Button(b, ID("mnop"))
		}))

		children := instance.Children()
		abcd := children[0].(Focusable)
		efgh := children[1].(Focusable)
		ijkl := children[2].(Focusable)
		mnop := children[3].(Focusable)

		abcd.Focus()
		assert.True(t, abcd.Focused())
		assert.False(t, efgh.Focused())
		assert.False(t, ijkl.Focused())
		assert.False(t, mnop.Focused())

		ijkl.Focus()
		assert.False(t, abcd.Focused())
		assert.False(t, efgh.Focused())
		assert.True(t, ijkl.Focused())
		assert.False(t, mnop.Focused())
	})

	t.Run("FocusablePath() returns nearest focusable parent", func(t *testing.T) {
		var child Displayable
		root, _ := Box(NewBuilder(), Children(func(b Builder) {
			Box(b)
			Box(b)
			Box(b, Children(func() {
				Box(b, ID("abcd"))
				Box(b, ID("efgh"), IsFocusable(true), Children(func() {
					Box(b, Children(func() {
						child, _ = Box(b)
					}))
				}))
			}))
		}))

		nonFocusable := root.FindComponentByID("abcd")
		assert.Equal(t, nonFocusable.NearestFocusable(), root)

		focusable := root.FindComponentByID("efgh")
		assert.Equal(t, focusable, focusable.NearestFocusable(), "returns self too")

		expected := child.NearestFocusable()
		assert.Equal(t, focusable, expected, "Child returns Focusable grandparent")
	})
}
