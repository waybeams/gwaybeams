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
		t.Run("single", func(t *testing.T) {
			t.Skip()
			root, _ := Box(NewBuilder(), ID("abcd"))
			root.Invalidate()
			nodes := root.InvalidNodes()
			assert.Equal(t, len(nodes), 1)
		})

		t.Run("one child and not another", func(t *testing.T) {
			t.Skip()
			root, _ := Box(NewBuilder(), ID("root"), Children(func(b Builder) {
				Box(b, ID("abcd"))
				Box(b, ID("efgh"))
			}))

			root.FirstChild().Invalidate()
			assert.Equal(t, root.InvalidNodes()[0].ID(), "abcd")
		})

	})
}
