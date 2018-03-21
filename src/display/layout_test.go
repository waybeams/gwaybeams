package display

import (
	"assert"
	"testing"
)

func createDisplayableTree() (Displayable, Displayable, Displayable, Displayable) {
	root := NewSprite()
	one := NewSpriteWithOpts(&Opts{FlexWidth: 1})
	two := NewSpriteWithOpts(&Opts{FlexWidth: 2})
	three := NewSprite()
	four := NewSpriteWithOpts(&Opts{ExcludeFromLayout: true})
	five := NewSpriteWithOpts(&Opts{FlexWidth: 1, Id: "five"})

	root.AddChild(one)
	root.AddChild(two)

	one.AddChild(three)
	one.AddChild(four)
	one.AddChild(five)

	return root, one, two, three
}

func TestLayout(t *testing.T) {
	root := NewSprite()

	t.Run("Call Layout", func(t *testing.T) {
		assert.NotNil(root)
	})

	t.Run("GetLayoutableChildren", func(t *testing.T) {
		t.Run("No children returns empty slice", func(t *testing.T) {
			_, _, _, three := createDisplayableTree()
			children := GetLayoutableChildren(three)
			assert.Equal(len(children), 0)
		})

		t.Run("Returns layoutable children in general", func(t *testing.T) {
			root, one, two, _ := createDisplayableTree()
			children := GetLayoutableChildren(root)
			assert.Equal(len(children), 2)
			assert.Equal(children[0], one)
			assert.Equal(children[1], two)
		})

		t.Run("Filters non-layoutable children", func(t *testing.T) {
			_, one, _, three := createDisplayableTree()
			children := GetLayoutableChildren(one)
			assert.Equal(one.GetChildCount(), 3)
			assert.Equal(len(children), 2)
			assert.Equal(children[0], three)
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
			_, _, _, three := createDisplayableTree()
			children := GetFlexibleChildren(three)
			assert.Equal(len(children), 0)
		})

		t.Run("Returns flexible children in general", func(t *testing.T) {
			root, one, two, _ := createDisplayableTree()
			children := GetFlexibleChildren(root)
			assert.Equal(len(children), 2)
			assert.Equal(children[0], one)
			assert.Equal(children[1], two)
		})

		t.Run("Filters non-flexible children", func(t *testing.T) {
			_, one, _, _ := createDisplayableTree()
			children := GetFlexibleChildren(one)
			assert.Equal(one.GetChildCount(), 3)
			assert.Equal(len(children), 1)
			assert.Equal(children[0].GetId(), "five")
		})
	})

	t.Run("directionalDelegate", func(t *testing.T) {
		delegate := DirectionalDelegate(Horizontal)
		if delegate == nil {
			t.Error("Expected DirectionalDelegate to return a function")
		}
	})
}
