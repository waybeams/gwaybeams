package display

import (
	"assert"
	"testing"
)

func TestTrait(t *testing.T) {
	t.Run("PushTrait", func(t *testing.T) {
		// Outer node does not receive selector value
		root, _ := Box(NewBuilder(), Children(func(b Builder) {
			Trait(b, "*", BgColor(0xffcc00ff))
			// Should receive the provided selector value
			Box(b, ID("one"))
			// Override the selector with concrete value
			Box(b, ID("two"), BgColor(0xff00ffff))
		}))

		opts := root.GetTraitOptions()
		assert.NotNil(t, opts["*"], "Opts collected")

		assert.Equal(t, root.GetBgColor(), DefaultBgColor, "one bgcolor")
		assert.Equal(t, root.GetChildAt(0).GetBgColor(), 0xffcc00ff, "one bgcolor")
		assert.Equal(t, root.GetChildAt(1).GetBgColor(), 0xff00ffff, "two bgcolor")
	})
}
