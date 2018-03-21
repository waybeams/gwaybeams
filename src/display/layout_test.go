package display

import (
	"assert"
	"testing"
)

func createDisplayableTree() (Displayable, []Displayable) {
	root := NewSprite()
	one := NewSpriteWithOpts(&Opts{FlexWidth: 1})
	two := NewSpriteWithOpts(&Opts{FlexWidth: 2})
	three := NewSpriteWithOpts(&Opts{Id: "three"})
	four := NewSpriteWithOpts(&Opts{Id: "four", ExcludeFromLayout: true})
	five := NewSpriteWithOpts(&Opts{Id: "five", FlexWidth: 1})

	root.AddChild(one)
	root.AddChild(two)

	one.AddChild(three)
	one.AddChild(four)
	one.AddChild(five)

	return root, []Displayable{root, one, two, three, four, five}
}

func TestLayout(t *testing.T) {
	root := NewSprite()

	t.Run("Call Layout", func(t *testing.T) {
		assert.NotNil(root)
	})

	t.Run("GetLayoutableChildren", func(t *testing.T) {
		t.Run("No children returns empty slice", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := GetLayoutableChildren(nodes[3])
			assert.Equal(len(children), 0)
		})

		t.Run("Returns layoutable children in general", func(t *testing.T) {
			root, nodes := createDisplayableTree()
			children := GetLayoutableChildren(root)
			assert.Equal(len(children), 2)
			assert.Equal(children[0], nodes[1])
			assert.Equal(children[1], nodes[2])
		})

		t.Run("Filters non-layoutable children", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := GetLayoutableChildren(nodes[1])
			assert.Equal(nodes[1].GetChildCount(), 3)
			assert.Equal(len(children), 2)
			assert.Equal(children[0], nodes[3])
		})
	})

	t.Run("GetFlexibleChildren", func(t *testing.T) {
		t.Run("Returns non nil slice", func(t *testing.T) {
			root = NewSprite()
			children := GetFlexibleChildren(root)
			if children == nil {
				t.Error("Expected children to not be nil")
			}
		})

		t.Run("No children returns empty slice", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := GetFlexibleChildren(nodes[3])
			assert.Equal(len(children), 0)
		})

		t.Run("Returns flexible children in general", func(t *testing.T) {
			root, nodes := createDisplayableTree()
			children := GetFlexibleChildren(root)
			assert.Equal(len(children), 2)
			assert.Equal(children[0], nodes[1])
			assert.Equal(children[1], nodes[2])
		})

		t.Run("Filters non-flexible children", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := GetFlexibleChildren(nodes[1])
			assert.Equal(nodes[1].GetChildCount(), 3)
			assert.Equal(len(children), 1)
			assert.Equal(children[0].GetId(), "five")
		})
	})

	t.Run("GetStaticChildren", func(t *testing.T) {
		t.Run("Returns non nil slice", func(t *testing.T) {
			root = NewSprite()
			children := GetStaticChildren(root)
			if children == nil {
				t.Error("Expected children to not be nil")
			}
		})

		t.Run("No children returns empty slice", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := GetStaticChildren(nodes[3])
			assert.Equal(len(children), 0)
		})

		t.Run("Returns zero static children if all are flexible", func(t *testing.T) {
			root, _ := createDisplayableTree()
			children := GetStaticChildren(root)
			assert.Equal(len(children), 0)
		})

		t.Run("Returns only static children", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := GetStaticChildren(nodes[1])
			assert.Equal(len(children), 1)
			assert.Equal(children[0].GetId(), "three")
		})
	})

	t.Run("directionalDelegate", func(t *testing.T) {
		delegate := DirectionalDelegate(Horizontal)
		if delegate == nil {
			t.Error("Expected DirectionalDelegate to return a function")
		}
	})

	t.Run("horizontalDelegate", func(t *testing.T) {
		t.Run("StaticSize kids", func(t *testing.T) {
			root := NewSprite()
			one := NewSpriteWithOpts(&Opts{Width: 10, Height: 10})
			two := NewSpriteWithOpts(&Opts{FlexWidth: 1, FlexHeight: 1})
			three := NewSpriteWithOpts(&Opts{Width: 10, Height: 10})
			root.AddChild(one)
			root.AddChild(two)
			root.AddChild(three)

			hDelegate := &horizontalDelegate{}
			vDelegate := &horizontalDelegate{}

			hSize := GetStaticSize(hDelegate, root)
			assert.Equal(hSize, 20.0)
			vSize := GetStaticSize(vDelegate, root)
			assert.Equal(vSize, 20.0)
		})
	})
}
