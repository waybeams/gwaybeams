package display

import (
	"assert"
	"testing"
)

func createDisplayableTree() (Displayable, []Displayable) {
	root := NewComponent()
	one := NewComponentWithOpts(&ComponentModel{FlexWidth: 1})
	two := NewComponentWithOpts(&ComponentModel{FlexWidth: 2})
	three := NewComponentWithOpts(&ComponentModel{Id: "three"})
	four := NewComponentWithOpts(&ComponentModel{Id: "four", ExcludeFromLayout: true})
	five := NewComponentWithOpts(&ComponentModel{Id: "five", FlexWidth: 1})

	root.AddChild(one)
	root.AddChild(two)

	one.AddChild(three)
	one.AddChild(four)
	one.AddChild(five)

	return root, []Displayable{root, one, two, three, four, five}
}

func createStubApp() (Displayable, []Displayable) {
	root := NewComponentWithOpts(&ComponentModel{Id: "root", Width: 800, Height: 600})
	header := NewComponentWithOpts(&ComponentModel{Id: "header", Padding: 5, FlexWidth: 1, Height: 80})
	body := NewComponentWithOpts(&ComponentModel{Id: "body", Padding: 5, FlexWidth: 1, FlexHeight: 1})
	footer := NewComponentWithOpts(&ComponentModel{Id: "footer", FlexWidth: 1, Height: 60})

	logo := NewComponentWithOpts(&ComponentModel{Id: "logo", Width: 50, Height: 50})
	content := NewComponentWithOpts(&ComponentModel{Id: "content", FlexWidth: 1, FlexHeight: 1})

	header.AddChild(logo)
	body.AddChild(content)

	root.AddChild(header)
	root.AddChild(body)
	root.AddChild(footer)

	return root, []Displayable{root, header, body, footer, logo, content}
}

func createTwoBoxes() (Displayable, Displayable) {
	root := NewComponentWithOpts(&ComponentModel{Id: "root", Padding: 10, Width: 100, Height: 110})
	child := NewComponentWithOpts(&ComponentModel{Id: "child", FlexWidth: 1, FlexHeight: 1})
	root.AddChild(child)
	return root, child
}

func TestLayout(t *testing.T) {
	root := NewComponent()

	t.Run("Call Layout", func(t *testing.T) {
		assert.NotNil(root)
	})

	t.Run("createStubApp works as expected", func(t *testing.T) {
		root, nodes := createStubApp()
		assert.Equal(root.GetId(), "root")
		assert.Equal(len(nodes), 6)
		assert.Equal(root.GetChildCount(), 3)
	})

	t.Run("DisplayStack Layout", func(t *testing.T) {
		root, child := createTwoBoxes()

		StackLayout(root)
		assert.Equal(child.GetWidth(), 80.0)
		assert.Equal(child.GetHeight(), 90.0)
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
			root = NewComponent()
			hDelegate := &horizontalDelegate{}
			children := GetFlexibleChildren(hDelegate, root)
			if children == nil {
				t.Error("Expected children to not be nil")
			}
		})

		t.Run("No children returns empty slice", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := GetFlexibleChildren(hDelegate, nodes[3])
			assert.Equal(len(children), 0)
		})

		t.Run("Returns flexible children in general", func(t *testing.T) {
			root, nodes := createDisplayableTree()
			children := GetFlexibleChildren(hDelegate, root)
			assert.Equal(len(children), 2)
			assert.Equal(children[0], nodes[1])
			assert.Equal(children[1], nodes[2])
		})

		t.Run("Filters non-flexible children", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := GetFlexibleChildren(hDelegate, nodes[1])
			assert.Equal(nodes[1].GetChildCount(), 3)
			assert.Equal(len(children), 1)
			assert.Equal(children[0].GetId(), "five")
		})
	})

	t.Run("GetStaticChildren", func(t *testing.T) {
		t.Run("Returns non nil slice", func(t *testing.T) {
			root = NewComponent()
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

	t.Run("horizontalDelegate", func(t *testing.T) {
		t.Run("StaticSize kids", func(t *testing.T) {
			root := NewComponent()
			one := NewComponentWithOpts(&ComponentModel{Width: 10, Height: 10})
			two := NewComponentWithOpts(&ComponentModel{FlexWidth: 1, FlexHeight: 1})
			three := NewComponentWithOpts(&ComponentModel{Width: 10, Height: 10})
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
