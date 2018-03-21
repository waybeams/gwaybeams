package display

import (
	"assert"
	"strings"
	"testing"
)

func TestSprite(t *testing.T) {

	t.Run("Generated Id", func(t *testing.T) {
		root := NewSprite()
		assert.Equal(len(root.GetId()), 20)
	})

	t.Run("Provided Id", func(t *testing.T) {
		root := NewSpriteWithOpts(&Opts{Id: "root"})
		assert.Equal(root.GetId(), "root")
	})

	t.Run("GetPath for root", func(t *testing.T) {
		root := NewSpriteWithOpts(&Opts{Id: "root"})
		assert.Equal(root.GetPath(), "/root")
	})

	t.Run("GetPath with depth", func(t *testing.T) {
		root := NewSpriteWithOpts(&Opts{Id: "root"})
		one := NewSpriteWithOpts(&Opts{Id: "one"})
		two := NewSpriteWithOpts(&Opts{Id: "two"})
		three := NewSpriteWithOpts(&Opts{Id: "three"})
		root.AddChild(one)
		one.AddChild(two)
		two.AddChild(three)
		assert.Equal(one.GetPath(), "/root/one")
		assert.Equal(two.GetPath(), "/root/one/two")
		assert.Equal(three.GetPath(), "/root/one/two/three")
	})

	t.Run("Padding", func(t *testing.T) {
		t.Run("Applying Padding spreads to all four sides", func(t *testing.T) {
			root := NewSpriteWithOpts(&Opts{Padding: 10})

			assert.Equal(root.GetHorizontalPadding(), 20.0)
			assert.Equal(root.GetVerticalPadding(), 20.0)

			assert.Equal(root.GetPaddingBottom(), 10.0)
			assert.Equal(root.GetPaddingLeft(), 10.0)
			assert.Equal(root.GetPaddingRight(), 10.0)
			assert.Equal(root.GetPaddingTop(), 10.0)
		})

		t.Run("PaddingTop overrides Padding", func(t *testing.T) {
			root := NewSpriteWithOpts(&Opts{Padding: 10, PaddingTop: 5})
			assert.Equal(root.GetPaddingTop(), 5.0)
			assert.Equal(root.GetPaddingBottom(), 10.0)
			assert.Equal(root.GetPadding(), 10.0)
		})
	})

	t.Run("PrefWidth default value", func(t *testing.T) {
		one := NewSprite()
		assert.Equal(0.0, one.GetPrefWidth())
	})

	t.Run("PrefWidth Opts value", func(t *testing.T) {
		one := NewSpriteWithOpts(&Opts{PrefWidth: 200})
		assert.Equal(200.0, one.GetPrefWidth())
	})

	t.Run("AddChild", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		root.Width(200)
		assert.Equal(root.AddChild(one), 1)
		assert.Equal(root.AddChild(two), 2)
		assert.Equal(one.GetParent().GetId(), root.GetId())
		assert.Equal(two.GetParent().GetId(), root.GetId())
		assert.Nil(root.GetParent())
	})

	t.Run("GetChildCount", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		three := NewSprite()
		root.AddChild(one)
		one.AddChild(two)
		one.AddChild(three)

		assert.Equal(root.GetChildCount(), 1)
		assert.Equal(root.GetChildAt(0), one)

		assert.Equal(one.GetChildCount(), 2)
		assert.Equal(one.GetChildAt(0), two)
		assert.Equal(one.GetChildAt(1), three)
	})

	t.Run("GetFilteredChildren", func(t *testing.T) {
		createTree := func() (Displayable, []Displayable) {

			root := NewSprite()
			one := NewSpriteWithOpts(&Opts{Id: "a-t-one"})
			two := NewSpriteWithOpts(&Opts{Id: "a-t-two"})
			three := NewSpriteWithOpts(&Opts{Id: "b-t-three"})
			four := NewSpriteWithOpts(&Opts{Id: "b-t-four"})

			root.AddChild(one)
			root.AddChild(two)
			root.AddChild(three)
			root.AddChild(four)

			return root, []Displayable{one, two, three, four}
		}

		allKids := func(d Displayable) bool {
			return strings.Index(d.GetId(), "-t-") > -1
		}

		bKids := func(d Displayable) bool {
			return strings.Index(d.GetId(), "b-") > -1
		}

		t.Run("returns Empty slice", func(t *testing.T) {
			root := NewSprite()
			filtered := root.GetFilteredChildren(allKids)
			assert.Equal(len(filtered), 0)
		})

		t.Run("returns all matched children in simple match", func(t *testing.T) {
			root, _ := createTree()
			filtered := root.GetFilteredChildren(allKids)
			assert.Equal(len(filtered), 4)
		})

		t.Run("returns all matched children in harder match", func(t *testing.T) {
			root, _ := createTree()
			filtered := root.GetFilteredChildren(bKids)
			assert.Equal(len(filtered), 2)
			assert.Equal(filtered[0].GetId(), "b-t-three")
			assert.Equal(filtered[1].GetId(), "b-t-four")
		})
	})

	t.Run("GetChildren returns empty list", func(t *testing.T) {
		root := NewSprite()
		children := root.GetChildren()

		if children == nil {
			t.Error("GetChildren should not return nil")
		}

		assert.Equal(len(children), 0)
	})

	t.Run("GetChildren returns new list", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		three := NewSprite()

		root.AddChild(one)
		root.AddChild(two)
		root.AddChild(three)

		children := root.GetChildren()
		assert.Equal(len(children), 3)
	})
}
