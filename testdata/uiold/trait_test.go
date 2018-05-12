package ui_test

// RED is the global default BgColor for FakeTraitName control.
// var FakeTraitName = control.Define("FakeTraitName", control.New, BgColor(0xff0000ff))

/*
func TestTrait(t *testing.T) {
	t.Skip()
	t.Run("PushTrait", func(t *testing.T) {
		// Outer node does not receive selector value
		root := Box(context.New(), Children(func(c Context) {
			Trait(c, "*", BgColor(0xffcc00ff))
			// Should receive the provided selector value
			Box(c, ID("one"))
			// Override the selector with concrete value
			Box(c, ID("two"), BgColor(0xff00ffff))
		}))

		opts := root.TraitOptions()
		assert.NotNil(t, opts["*"], "Opts collected")

		assert.Equal(t, root.BgColor(), DefaultBgColor, "one bgcolor")
		assert.Equal(t, root.ChildAt(0).BgColor(), 0xffcc00ff, "one bgcolor")
		assert.Equal(t, root.ChildAt(1).BgColor(), 0xff00ffff, "two bgcolor")
	})

	t.Run("Traits applied to control type names", func(t *testing.T) {
		var one, two, three, four Displayable

		red := 0xff0000ff
		blue := 0x00ff00ff
		green := 0x0000ffff

		root := Box(context.New(), Children(func(c Context) {
			one = FakeTraitName(c)
			Box(c, Children(func() {
				// Any FakeTraitName control instances inside of this Box, will have DEFAULT BgColor BLUE.
				Trait(c, "FakeTraitName", BgColor(blue))
				two = FakeTraitName(c)
				// This instance overrides the modified default color
				three = FakeTraitName(c, BgColor(green))
				Box(c, Children(func() {
					// Even nested children pick up the Trait definition
					four = FakeTraitName(c)
				}))
			}))
		}))

		assert.Equal(t, root.ChildCount(), 2)
		assert.Equal(t, one.BgColor(), red, "one")
		assert.Equal(t, two.BgColor(), blue, "two")
		assert.Equal(t, three.BgColor(), green, "three")
		assert.Equal(t, four.BgColor(), blue, "four")
	})

	t.Run("Traits apply by trait names", func(t *testing.T) {
		var child Displayable

		root := Box(context.New(), Children(func(c Context) {
			Trait(c, ".abcd", Width(200))
			Trait(c, ".efgh", Height(100))
			child = Box(c, TraitNames("abcd", "efgh"))
		}))

		assert.NotNil(t, root)
		assert.Equal(t, int(child.Width()), 200)
		assert.Equal(t, int(child.Height()), 100)
	})
}
*/
