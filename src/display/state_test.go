package display

import (
	"assert"
	"testing"
)

func TestState(t *testing.T) {
	t.Run("Default state is automatically applied by builder", func(t *testing.T) {
		box, _ := Box(NewBuilder(),
			AddState("default", Width(100)))

		assert.Equal(t, box.Width(), 100)
	})

	t.Run("Component applies options", func(t *testing.T) {
		box, _ := Box(NewBuilder(),
			AddState("default", Width(100)),
			AddState("foo", Width(200)),
			AddState("bar", Width(300)))

		assert.Equal(t, box.Width(), 100)
		box.SetState("foo", nil)
		assert.Equal(t, box.Width(), 200)
		box.SetState("bar")
		assert.Equal(t, box.Width(), 300)
		box.SetState("default")
		assert.Equal(t, box.Width(), 100)
	})

	t.Run("Clobbers existing states", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Width(100),
			AddState("foo", Width(200)),
			AddState("bar", Width(300)))

		assert.Equal(t, box.Width(), 100)
		box.SetState("foo", nil)
		assert.Equal(t, box.Width(), 200)
		box.SetState("bar")
		assert.Equal(t, box.Width(), 300)
	})

	t.Run("Works with children too", func(t *testing.T) {
		root, _ := Box(NewBuilder(),
			FlexWidth(1), FlexHeight(1), BgColor(0x00ff00ff),
			AddState("default", Children(func(b Builder) {
				Label(b, Text("Hello World"))
			})),
			AddState("foo", Children(func(b Builder) {
				Label(b, Text("Goodbye World"))
			})),
			AddState("bar", Children(func() {})))

		// Verify the default state
		assert.Equal(t, root.ChildAt(0).Text(), "Hello World")

		// Update to different state that mutates Children
		root.SetState("foo")
		root.Builder().Update(root)
		assert.Equal(t, root.ChildAt(0).Text(), "Goodbye World")

		// Clear the children
		root.SetState("bar")
		root.Builder().Update(root)
		assert.Equal(t, root.ChildCount(), 0)
	})
}
