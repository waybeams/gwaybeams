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
	t.Run("Call LayoutHandler", func(t *testing.T) {
		root := NewComponent()
		assert.NotNil(t, root)
	})

	t.Run("createStubApp works as expected", func(t *testing.T) {
		root, nodes := createStubApp()
		assert.Equal(t, root.GetId(), "root")
		assert.Equal(t, len(nodes), 6)
		assert.Equal(t, root.GetChildCount(), 3)
	})

	t.Run("DisplayStack LayoutHandler", func(t *testing.T) {
		root, child := createTwoBoxes()

		StackLayout(root)
		assert.Equal(t, child.GetWidth(), 80.0)
		assert.Equal(t, child.GetHeight(), 90.0)
	})

	t.Run("GetLayoutableChildren", func(t *testing.T) {
		t.Run("No children returns empty slice", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := getLayoutableChildren(nodes[3])
			assert.Equal(t, len(children), 0)
		})

		t.Run("Returns layoutable children in general", func(t *testing.T) {
			root, nodes := createDisplayableTree()
			children := getLayoutableChildren(root)
			assert.Equal(t, len(children), 2)
			assert.Equal(t, children[0], nodes[1])
			assert.Equal(t, children[1], nodes[2])
		})

		t.Run("Filters non-layoutable children", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := getLayoutableChildren(nodes[1])
			assert.Equal(t, nodes[1].GetChildCount(), 3)
			assert.Equal(t, len(children), 2)
			assert.Equal(t, children[0], nodes[3])
		})
	})

	t.Run("GetFlexibleChildren", func(t *testing.T) {
		t.Run("Returns non nil slice", func(t *testing.T) {
			root := NewComponent()
			hDelegate := &horizontalDelegate{}
			children := getFlexibleChildren(hDelegate, root)
			if children == nil {
				t.Error("Expected children to not be nil")
			}
		})

		t.Run("No children returns empty slice", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := getFlexibleChildren(hDelegate, nodes[3])
			assert.Equal(t, len(children), 0)
		})

		t.Run("Returns flexible children in general", func(t *testing.T) {
			root, nodes := createDisplayableTree()
			children := getFlexibleChildren(hDelegate, root)
			assert.Equal(t, len(children), 2)
			assert.Equal(t, children[0], nodes[1])
			assert.Equal(t, children[1], nodes[2])
		})

		t.Run("Filters non-flexible children", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			children := getFlexibleChildren(hDelegate, nodes[1])
			assert.Equal(t, nodes[1].GetChildCount(), 3)
			assert.Equal(t, len(children), 1)
			assert.Equal(t, children[0].GetId(), "five")
		})
	})

	t.Run("GetStaticChildren", func(t *testing.T) {
		t.Run("Returns non nil slice", func(t *testing.T) {
			root := NewComponent()
			delegate := &verticalDelegate{}
			children := getStaticChildren(delegate, root)
			if children == nil {
				t.Error("Expected children to be nil")
			}
		})

		t.Run("No children returns empty slice", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			delegate := &verticalDelegate{}
			children := getStaticChildren(delegate, nodes[3])
			assert.Equal(t, len(children), 0)
		})

		t.Run("Returns zero static children if all are flexible", func(t *testing.T) {
			root, _ := createDisplayableTree()
			delegate := &horizontalDelegate{}
			children := getStaticChildren(delegate, root)
			assert.Equal(t, len(children), 0)
		})

		t.Run("Returns only static children", func(t *testing.T) {
			_, nodes := createDisplayableTree()
			delegate := &horizontalDelegate{}
			children := getStaticChildren(delegate, nodes[1])
			assert.Equal(t, len(children), 1)
			assert.Equal(t, children[0].GetId(), "three")
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
			assert.Equal(t, hSize, 20.0)
			vSize := getStaticSize(vDelegate, root)
			assert.Equal(t, vSize, 20.0)
		})
	})

	t.Run("Spread remainder", func(t *testing.T) {
		var root, one, two, three Displayable
		root, _ = HBox(NewBuilder(), Width(152), Children(func(b Builder) {
			one, _ = Box(b, FlexWidth(1))
			two, _ = Box(b, FlexWidth(1))
			three, _ = Box(b, FlexWidth(1))
		}))
		root.Layout()
		assert.Equal(t, one.GetWidth(), 51)
		assert.Equal(t, two.GetWidth(), 51)
		assert.Equal(t, three.GetWidth(), 50)
	})

	t.Run("Basic, nested layout", func(t *testing.T) {
		var root, header, content, footer Displayable
		root, _ = VBox(NewBuilder(), Id("root"), Width(100), Height(300), Children(func(b Builder) {
			header, _ = HBox(b, Id("header"), FlexWidth(1), Height(100), Children(func(b Builder) {
				Box(b, Id("logo"), Width(200), Height(100))
			}))
			content, _ = Box(b, Id("content"), FlexHeight(1), FlexWidth(1))
			footer, _ = Box(b, Id("footer"), Height(80), FlexWidth(1))
		}))
		root.Layout()
		assert.Equal(t, header.GetHeight(), 100)
		assert.Equal(t, footer.GetHeight(), 80)
		assert.Equal(t, content.GetHeight(), 120)
	})
}
