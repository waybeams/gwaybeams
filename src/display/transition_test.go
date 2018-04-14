package display

import (
	"assert"
	"clock"
	"github.com/fogleman/ease"
	"testing"
	"time"
)

func TestTransition(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		fakeClock := clock.NewFake()
		var tweenX ComponentOption
		root, _ := Box(NewBuilderUsing(fakeClock), Children(func(b Builder) {
			tweenX = Transition(b,
				X,
				100.0,
				200.0,
				200,
				ease.Linear)
			Box(b, tweenX, ExcludeFromLayout(true))
		}))

		// Begin listening for enter frame events
		defer root.Builder().Destroy()
		go root.Builder().Listen()

		child := root.ChildAt(0)
		assert.NotNil(t, tweenX, "Expected tween creation")

		root.Layout()
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
}
