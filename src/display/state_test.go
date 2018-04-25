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
		box.SetState("foo")
		box.Builder().Update(box)
		assert.Equal(t, box.Width(), 200)
		box.SetState("bar")
		box.Builder().Update(box)
		assert.Equal(t, box.Width(), 300)
		box.SetState("default")
		box.Builder().Update(box)
		assert.Equal(t, box.Width(), 100)
	})

	t.Run("Clobbers existing states", func(t *testing.T) {
		box, _ := Box(NewBuilder(),
			AddState("default", Width(100)),
			AddState("foo", Width(200)),
			AddState("bar", Width(300)))

		assert.Equal(t, box.Width(), 100)
		box.SetState("foo")
		box.Builder().Update(box)
		assert.Equal(t, box.Width(), 200)
		box.SetState("bar")
		box.Builder().Update(box)
		assert.Equal(t, box.Width(), 300)
	})

	t.Run("Works with children too", func(t *testing.T) {
		root, _ := Box(NewBuilder(),
			FlexWidth(1), FlexHeight(1), BgColor(0x00ff00ff),
			AddState("default", Children(func(b Builder) {
				Label(b, Text("Hello World"))
			})),
			AddState("goodbye", Children(func(b Builder) {
				Label(b, Text("Goodbye World"))
			})),
			AddState("empty", Children(func() {})))

		// Verify the default state
		assert.Equal(t, root.ChildAt(0).Text(), "Hello World")

		// Update to different state that mutates Children
		root.SetState("goodbye")
		root.Builder().Update(root)
		assert.Equal(t, root.ChildAt(0).Text(), "Goodbye World")

		// Clear the children
		root.SetState("empty")
		root.Builder().Update(root)
		assert.Equal(t, root.ChildCount(), 0)
	})

	t.Run("undefined default state reruns all initial options", func(t *testing.T) {
		root, _ := Box(NewBuilder(),
			AddState("default", Width(10)),
			AddState("wider", Width(20)))

		assert.Equal(t, root.Width(), 10, "Default state")
		root.SetState("wider")
		root.Builder().Update(root)
		assert.Equal(t, root.Width(), 20, "wide state")
		root.SetState("default")
		root.Builder().Update(root)
		assert.Equal(t, root.Width(), 10)
	})
}
