package control

import (
	"assert"
	"events"
	"testing"
	. "ui"
	"uiold/context"
	. "ui/controls"
	. "uiold/opts"
)

func TestDisplayable(t *testing.T) {
	t.Run("Traits", func(t *testing.T) {
		box := Box(context.New(), TraitNames("abcd", "efgh"))

		names := box.TraitNames()
		assert.Equal(t, len(names), 2)
		assert.Equal(t, names[0], "abcd")
		assert.Equal(t, names[1], "efgh")
	})

	t.Run("PushUnsub", func(t *testing.T) {
		var callCount int
		var handler = func(e events.Event) {
			callCount++
		}
		root := Box(context.New(), On("foo", handler), On("foo", handler))

		root.Emit(events.New("foo", nil, nil))
		assert.Equal(t, callCount, 2)

		root.UnsubAll()
		callCount = 0

		root.Emit(events.New("foo", nil, nil))
		assert.Equal(t, callCount, 0)
	})

	t.Run("Data", func(t *testing.T) {
		t.Run("scalar", func(t *testing.T) {
			root := Box(context.New(), Data("abcd", 1234))
			assert.Equal(t, root.Data("abcd"), 1234)
		})

		t.Run("Coerces empty data string", func(t *testing.T) {
			root := Box(context.New())
			value := root.DataAsString("unused-key")
			assert.Equal(t, value, "")
		})
	})

	t.Run("InvalidNodes", func(t *testing.T) {
		t.Run("root invalidates", func(t *testing.T) {
			root := Box(context.New(), ID("abcd"))
			root.Invalidate()
			nodes := root.InvalidNodes()
			assert.Equal(t, len(nodes), 1)
		})

		t.Run("root can update", func(t *testing.T) {
			root := Box(context.New(), ID("abcd"))
			root.Invalidate()
			root.Context().Builder().Update(root)
			assert.Equal(t, root.ID(), "abcd")
		})

		t.Run("surprisingly invalidates the PARENT", func(t *testing.T) {
			root := Box(context.New(), ID("root"), Children(func(c Context) {
				Box(c, ID("abcd"))
			}))

			root.FirstChild().Invalidate()
			assert.Equal(t, root.InvalidNodes()[0].ID(), "root")
		})
	})

	t.Run("Text", func(t *testing.T) {
		root := Box(context.New(), Text("abcd"))
		assert.Equal(t, root.Text(), "abcd")
	})

	t.Run("Title", func(t *testing.T) {
		root := Box(context.New(), Title("abcd"))
		assert.Equal(t, root.Title(), "abcd")
	})
}
