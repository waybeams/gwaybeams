package controls

import (
	"assert"
	"clock"
	"github.com/fogleman/ease"
	"testing"
	"time"
	. "ui"
	"ui/context"
	. "ui/opts"
)

func TestTransition(t *testing.T) {

	var createTree = func() (Displayable, clock.FakeClock) {
		fakeClock := clock.NewFake()
		root := Box(context.New(context.Clock(fakeClock)), Children(func(c Context) {
			moveRight := Transition(c,
				X,
				100.0,
				200.0,
				200,
				ease.Linear)
			Box(c, ID("abcd"), moveRight, ExcludeFromLayout(true))
		}))

		return root, fakeClock
	}

	t.Run("Instantiable", func(t *testing.T) {
		root, fakeClock := createTree()
		// Begin listening for enter frame events
		defer root.Context().Destroy()
		go root.Context().Listen()

		child := root.ChildAt(0)

		assert.Equal(t, int(child.X()), 100)
		// I expect enter frames to fire when this happens!
		// But they don't because they're currently implemented by the NanoWindow
		fakeClock.Add(101 * time.Millisecond)
		assert.Equal(t, int(child.X()), 150)
		fakeClock.Add(51 * time.Millisecond)
		assert.Equal(t, int(child.X()), 175)
		fakeClock.Add(51 * time.Millisecond)
		assert.Equal(t, int(child.X()), 200)
		fakeClock.Add(51 * time.Millisecond)
		assert.Equal(t, int(child.X()), 200)
	})

	t.Run("Updateable", func(t *testing.T) {
		t.Skip()
		root, fakeClock := createTree()
		firstChild := root.ChildAt(0)

		// Begin listening for enter frame events
		defer root.Context().Destroy()
		go root.Context().Listen()

		fakeClock.Add(51 * time.Millisecond)
		assert.Equal(t, int(firstChild.X()), 125)
		root.InvalidateChildren()
		fakeClock.Add(51 * time.Millisecond)

		assert.Equal(t, int(firstChild.X()), 125)

		afterRenderChild := root.FindControlById("abcd")
		assert.Equal(t, int(afterRenderChild.X()), 150)
	})
}
