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

	t.Run("Traits applied to component type names", func(t *testing.T) {
		var one, two, three, four Displayable

		red := 0xff0000ff
		blue := 0x00ff00ff
		green := 0x0000ffff

		// RED is the global default BgColor for FakeTraitName components.
		FakeTraitName := NewComponentFactory("FakeTraitName", NewComponent, BgColor(red))

		root, _ := Box(NewBuilder(), Children(func(b Builder) {
			one, _ = FakeTraitName(b)
			Box(b, Children(func() {
				// Any FakeTraitName component instances inside of this Box, will have DEFAULT BgColor BLUE.
				Trait(b, "FakeTraitName", BgColor(blue))
				two, _ = FakeTraitName(b)
				// This instance overrides the modified default color
				three, _ = FakeTraitName(b, BgColor(green))
				Box(b, Children(func() {
					// Even nested children pick up the Trait definition
					four, _ = FakeTraitName(b)
				}))
			}))
		}))

		assert.Equal(t, root.GetChildCount(), 2)
		assert.Equal(t, one.GetBgColor(), red, "one")
		assert.Equal(t, two.GetBgColor(), blue, "two")
		assert.Equal(t, three.GetBgColor(), green, "three")
		assert.Equal(t, four.GetBgColor(), blue, "four")
	})
}
