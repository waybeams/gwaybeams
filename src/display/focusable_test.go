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
}
