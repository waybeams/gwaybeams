package control

import (
	"github.com/waybeams/assert"
	"testing"
	. "ui"
	"uiold/context"
	. "ui/controls"
	. "uiold/opts"
)

func TestFocusable(t *testing.T) {
	t.Run("Blurred", func(t *testing.T) {
		instance := Button(context.New(), Blurred())
		assert.False(instance.Focused())
	})

	t.Run("Focused", func(t *testing.T) {
		instance := Button(context.New(), Focused())
		assert.True(instance.Focused())
	})

	t.Run("Unfocuses previously focused elements", func(t *testing.T) {
		instance := VBox(context.New(), Children(func(c Context) {
			Button(c, ID("abcd"))
			Button(c, ID("efgh"))
			Button(c, ID("ijkl"))
			Button(c, ID("mnop"))
		}))

		children := instance.Children()
		abcd := children[0].(Focusable)
		efgh := children[1].(Focusable)
		ijkl := children[2].(Focusable)
		mnop := children[3].(Focusable)

		abcd.Focus()
		assert.True(abcd.Focused())
		assert.False(efgh.Focused())
		assert.False(ijkl.Focused())
		assert.False(mnop.Focused())

		ijkl.Focus()
		assert.False(abcd.Focused())
		assert.False(efgh.Focused())
		assert.True(ijkl.Focused())
		assert.False(mnop.Focused())
	})

	var createTree = func() Displayable {
		root := Box(context.New(), Children(func(c Context) {
			Box(c)
			Box(c)
			Box(c, Children(func() {
				Box(c, ID("uvwx"))
				Button(c, ID("abcd"))
				Box(c, ID("efgh"), IsFocusable(true), Children(func() {
					Box(c, ID("ijkl"), Children(func() {
						Box(c, ID("mnop"))
					}))
				}))
				Button(c, ID("qrst"))
			}))
		}))

		return root
	}

	t.Run("FocusablePath() returns nearest focusable parent", func(t *testing.T) {
		root := createTree()
		child := root.FindControlById("mnop")

		nonFocusable := root.FindControlById("uvwx")
		assert.Equal(nonFocusable.NearestFocusable().Path(), root.Path())

		focusable := root.FindControlById("efgh")
		assert.Equal(focusable.Path(), focusable.NearestFocusable().Path(), "returns self too")

		expected := child.NearestFocusable()
		assert.Equal(focusable.Path(), expected.Path(), "Child returns Focusable grandparent")
	})

	t.Run("Last focusable is blurred", func(t *testing.T) {
		root := createTree()
		abcd := root.FindControlById("abcd")
		qrst := root.FindControlById("qrst")
		abcd.Focus()
		assert.True(abcd.Focused())
		assert.False(qrst.Focused())
		assert.Equal(root.FocusedChild().Path(), abcd.Path())

		qrst.Focus()
		assert.False(abcd.Focused())
		assert.True(qrst.Focused())
		assert.Equal(root.FocusedChild().Path(), qrst.Path())
	})
}
