package display

import (
	"assert"
	"clock"
	"github.com/fogleman/ease"
	"testing"
	"time"
)

func TestTransition(t *testing.T) {

	var createTree = func() (Displayable, clock.FakeClock) {
		fakeClock := clock.NewFake()
		root, _ := Box(NewBuilderUsing(fakeClock), Children(func(b Builder) {
			moveRight := Transition(b,
				X,
				100.0,
				200.0,
				200,
				ease.Linear)
			Box(b, moveRight, ExcludeFromLayout(true))
		}))

		// Begin listening for enter frame events
		defer root.Builder().Destroy()
		go root.Builder().Listen()

		return root, fakeClock
	}

	t.Run("Instantiable", func(t *testing.T) {
		root, fakeClock := createTree()
		child := root.ChildAt(0)
		root.Layout()

		assert.Equal(t, int(child.X()), 100)
		// I expect enter frames to fire when this happens!
		// But they don't because they're currently implemented by the NanoWindow
		fakeClock.Add(101 * time.Millisecond)
		assert.Equal(t, int(child.X()), 150)

		// fakeClock.Add(51 * time.Millisecond)
		// assert.Equal(t, int(child.X()), 175)
		// fakeClock.Add(51 * time.Millisecond)
		// assert.Equal(t, int(child.X()), 200)
		// fakeClock.Add(51 * time.Millisecond)
		// assert.Equal(t, int(child.X()), 200)
	})

	t.Run("Updateable", func(t *testing.T) {
		t.Skip()
		fakeClock := clock.NewFake()
		var tweenX ComponentOption
		var composerCallCount int
		root, _ := Box(NewBuilderUsing(fakeClock), Children(func(b Builder) {
			composerCallCount++
			tweenX = Transition(b,
				X,
				100.0,
				200.0,
				200,
				ease.Linear)
			Box(b, ID("abcd"), tweenX, ExcludeFromLayout(true))
		}))
		child := root.ChildAt(0)
		assert.NotNil(t, child, "Expected child")

		defer root.Builder().Destroy()
		go root.Builder().Listen()

		assert.Equal(t, composerCallCount, 1)
		root.Layout()
		fakeClock.Add(51 * time.Millisecond)
		assert.Equal(t, int(child.X()), 125)
		root.InvalidateChildren()
		fakeClock.Add(51 * time.Millisecond)
		assert.Equal(t, composerCallCount, 2)

		child = root.FindComponentByID("abcd")
		// This assertion is a very subtle test (probably too subtle).
		// We're asserting that the pointer that we received before the Invalidate()
		// is still pointing to the actively attached component that we think it is.
		// We're asserting that the transition has been retained and advanced in time.
		assert.Equal(t, int(child.X()), 150)
		assert.True(t, false)
	})
}
