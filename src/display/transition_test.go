package display

import (
	"assert"
	"github.com/fogleman/ease"
	"testing"
)

func TestTransition(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		var tweenX ComponentOption
		root, _ := Box(NewBuilder(), Children(func(b Builder) {
			tweenX = Transition(X,
				110.0,
				200.0,
				200,
				ease.OutCubic)
			Box(b, tweenX)

		}))

		child := root.ChildAt(0)
		assert.NotNil(t, tweenX, "Expected tween creation")
		assert.Equal(t, int(child.X()), 110)
	})
}
