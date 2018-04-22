package display

import (
	"assert"
	"testing"
)

func TestInteractiveComponent(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		instance, _ := Interactive(NewBuilder())
		assert.NotNil(t, instance)
		assert.False(t, instance.Selected())
	})

	t.Run("Selected", func(t *testing.T) {
		instance, _ := Interactive(NewBuilder(), Selected(true))
		assert.True(t, instance.Selected())
	})

	t.Run("Blurred", func(t *testing.T) {
		instance, _ := Interactive(NewBuilder(), Blurred())
		assert.False(t, instance.Focused())
	})

	t.Run("Focused", func(t *testing.T) {
		instance, _ := Interactive(NewBuilder(), Focused())
		assert.True(t, instance.Focused())
	})

	t.Run("Unfocuses previously focused elements", func(t *testing.T) {
		instance, _ := VBox(NewBuilder(), Children(func(b Builder) {
			Interactive(b, ID("abcd"))
			Interactive(b, ID("efgh"))
			Interactive(b, ID("ijkl"))
			Interactive(b, ID("mnop"))
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
