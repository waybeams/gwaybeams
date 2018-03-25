package display

import (
	"assert"
	"testing"
)

func createDisplayableTree() (Displayable, []Displayable) {
	var root, one, two, three, four, five Displayable
	root, _ = TestComponent(NewBuilder(), Children(func(b Builder) {
		one, _ = TestComponent(b, FlexWidth(1), Children(func() {
			three, _ = TestComponent(b, Id("three"))
			four, _ = TestComponent(b, Id("four"), ExcludeFromLayout(true))
			five, _ = TestComponent(b, Id("five"), FlexWidth(1))
		}))
		two, _ = TestComponent(b, FlexWidth(2))
	}))

	return root, []Displayable{root, one, two, three, four, five}
}

func createStubApp() (Displayable, []Displayable) {
	var root, header, body, footer, logo, content Displayable

	root, _ = TestComponent(NewBuilder(), Id("root"), Width(800), Height(600), Children(func(b Builder) {
		header, _ = TestComponent(b, Id("header"), Padding(5), FlexWidth(1), Height(80), Children(func(b Builder) {
			logo, _ = TestComponent(b, Id("logo"), Width(50), Height(50))
			content, _ = TestComponent(b, Id("content"), FlexWidth(1), FlexHeight(1))
		}))
		body, _ = TestComponent(b, Id("body"), Padding(5), FlexWidth(1), FlexHeight(1))
		footer, _ = TestComponent(b, Id("footer"), FlexWidth(1), Height(60))
	}))

	return root, []Displayable{root, header, body, footer, logo, content}
}

func createTwoBoxes() (Displayable, Displayable) {
	var root, child Displayable
	root, _ = TestComponent(NewBuilder(), Id("root"), Padding(10), Width(100), Height(110), Children(func(b Builder) {
		child, _ = TestComponent(b, Id("child"), FlexWidth(1), FlexHeight(1))
	}))
	return root, child
}

func TestLayout(t *testing.T) {
	root := NewComponent()

	t.Run("Call LayoutHandler", func(t *testing.T) {
		assert.NotNil(root)
	})

	t.Run("createStubApp works as expected", func(t *testing.T) {
		root, nodes := createStubApp()
		assert.TEqual(t, root.GetId(), "root")
		assert.TEqual(t, len(nodes), 6)
		assert.TEqual(t, root.GetChildCount(), 3)
	})

	t.Run("DisplayStack LayoutHandler", func(t *testing.T) {
		root, child := createTwoBoxes()

		StackLayout(root)
		assert.TEqual(t, child.GetWidth(), 80.0)
		assert.TEqual(t, child.GetHeight(), 90.0)
	})

	t.Run("GetLayoutableChildren", func(t *testing.T) {
		t.Run("No children returns empty slice", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := getLayoutableChildren(nodes[3])
			assert.TEqual(t, len(children), 0)
		})

		t.Run("Returns layoutable children in general", func(t *testing.T) {
			root, nodes := createDisplayableTree()
			children := getLayoutableChildren(root)
			assert.TEqual(t, len(children), 2)
			assert.TEqual(t, children[0], nodes[1])
			assert.TEqual(t, children[1], nodes[2])
		})

		t.Run("Filters non-layoutable children", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := getLayoutableChildren(nodes[1])
			assert.TEqual(t, nodes[1].GetChildCount(), 3)
			assert.TEqual(t, len(children), 2)
			assert.TEqual(t, children[0], nodes[3])
		})
	})

	t.Run("GetFlexibleChildren", func(t *testing.T) {
		t.Run("Returns non nil slice", func(t *testing.T) {
			root = NewComponent()
			hDelegate := &horizontalDelegate{}
			children := getFlexibleChildren(hDelegate, root)
			if children == nil {
				t.Error("Expected children to not be nil")
			}
		})

		t.Run("No children returns empty slice", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := getFlexibleChildren(hDelegate, nodes[3])
			assert.TEqual(t, len(children), 0)
		})

		t.Run("Returns flexible children in general", func(t *testing.T) {
			root, nodes := createDisplayableTree()
			children := getFlexibleChildren(hDelegate, root)
			assert.TEqual(t, len(children), 2)
			assert.TEqual(t, children[0], nodes[1])
			assert.TEqual(t, children[1], nodes[2])
		})

		t.Run("Filters non-flexible children", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := getFlexibleChildren(hDelegate, nodes[1])
			assert.TEqual(t, nodes[1].GetChildCount(), 3)
			assert.TEqual(t, len(children), 1)
			assert.TEqual(t, children[0].GetId(), "five")
		})
	})

	t.Run("GetStaticChildren", func(t *testing.T) {
		t.Run("Returns non nil slice", func(t *testing.T) {
			root = NewComponent()
			children := getStaticChildren(root)
			if children == nil {
				t.Error("Expected children to not be nil")
			}
		})

		t.Run("No children returns empty slice", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := getStaticChildren(nodes[3])
			assert.TEqual(t, len(children), 0)
		})

		t.Run("Returns zero static children if all are flexible", func(t *testing.T) {
			root, _ := createDisplayableTree()
			children := getStaticChildren(root)
			assert.TEqual(t, len(children), 0)
		})

		t.Run("Returns only static children", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := getStaticChildren(nodes[1])
			assert.TEqual(t, len(children), 1)
			assert.TEqual(t, children[0].GetId(), "three")
		})
	})

	t.Run("horizontalDelegate", func(t *testing.T) {
		t.Run("StaticSize kids", func(t *testing.T) {
			var root, one, two, three Displayable
			root, _ = TestComponent(NewBuilder(), Children(func(b Builder) {
				one, _ = TestComponent(b, Width(10), Height(10))
				two, _ = TestComponent(b, FlexWidth(1), FlexHeight(1))
				three, _ = TestComponent(b, Width(10), Height(10))
			}))

			hDelegate := &horizontalDelegate{}
			vDelegate := &horizontalDelegate{}

			hSize := getStaticSize(hDelegate, root)
			assert.TEqual(t, hSize, 20.0)
			vSize := getStaticSize(vDelegate, root)
			assert.TEqual(t, vSize, 20.0)
		})
	})

	t.Run("Basic, nested layout", func(t *testing.T) {
		t.Skip()
		var root, header, two, three Displayable
		root, _ = VBox(NewBuilder(), Id("root"), Width(100), Height(100), Children(func(b Builder) {
			header, _ = HBox(b, Id("header"), Children(func(b Builder) {
				Box(b, Id("logo"), Width(200), Height(100))
			}))
			two, _ = Box(b, Id("two"), FlexHeight(1), FlexWidth(1))
			three, _ = Box(b, Id("two"), FlexHeight(1), FlexWidth(1))
		}))
		root.Layout()
		if header.GetHeight() != 33 {
			t.Error("Box should share the space vertically")
		}
	})
}
