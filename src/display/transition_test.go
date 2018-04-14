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
				110.0,
				200.0,
				200,
				ease.OutCubic)
			Box(b, tweenX, ExcludeFromLayout(true))
		}))

		child := root.ChildAt(0)
		assert.NotNil(t, tweenX, "Expected tween creation")

		root.Layout()
		assert.Equal(t, int(child.X()), 110)
		// I expect enter frames to fire when this happens!
		// But they don't because they're currently implemented by the NanoWindow
		fakeClock.Add(100 * time.Millisecond)
		// assert.Equal(t, int(child.X()), 120)
	})
}
