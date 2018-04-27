package display

import (
	"assert"
	"testing"
)

func TestDisplayable(t *testing.T) {
	t.Run("Traits", func(t *testing.T) {
		box, _ := Box(NewBuilder(), TraitNames("abcd", "efgh"))

		names := box.TraitNames()
		assert.Equal(t, len(names), 2)
		assert.Equal(t, names[0], "abcd")
		assert.Equal(t, names[1], "efgh")
	})

	t.Run("PushUnsub", func(t *testing.T) {
		var callCount int
		var handler = func(e Event) {
			callCount++
		}
		root, _ := Box(NewBuilder(), On("foo", handler), On("foo", handler))

		root.Emit(NewEvent("foo", nil, nil))
		assert.Equal(t, callCount, 2)

		root.UnsubAll()
		callCount = 0

		root.Emit(NewEvent("foo", nil, nil))
		assert.Equal(t, callCount, 0)
	})

	t.Run("Data", func(t *testing.T) {
		t.Run("scalar", func(t *testing.T) {
			root, _ := Box(NewBuilder(), Data(1234))
			assert.Equal(t, root.Data(), 1234)
		})
	})

	t.Run("InvalidNodes", func(t *testing.T) {
		t.Run("root invalidates", func(t *testing.T) {
			root, _ := Box(NewBuilder(), ID("abcd"))
			root.Invalidate()
			nodes := root.InvalidNodes()
			assert.Equal(t, len(nodes), 1)
		})

		t.Run("root can update", func(t *testing.T) {
			root, _ := Box(NewBuilder(), ID("abcd"))
			root.Invalidate()
			root.Builder().Update(root)
			assert.Equal(t, root.ID(), "abcd")
		})

		t.Run("surprisingly invalidates the PARENT", func(t *testing.T) {
			root, _ := Box(NewBuilder(), ID("root"), Children(func(b Builder) {
				Box(b, ID("abcd"))
			}))

			root.FirstChild().Invalidate()
			assert.Equal(t, root.InvalidNodes()[0].ID(), "root")
		})
	})

	t.Run("Text", func(t *testing.T) {
		root, _ := Box(NewBuilder(), Text("abcd"))
		assert.Equal(t, root.Text(), "abcd")
	})

	t.Run("Title", func(t *testing.T) {
		root, _ := Box(NewBuilder(), Title("abcd"))
		assert.Equal(t, root.Title(), "abcd")
	})
}
