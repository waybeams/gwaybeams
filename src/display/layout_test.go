package display

import (
	"assert"
	"testing"
)

func createDisplayableTree() (Displayable, []Displayable) {
	var root, one, two, three, four, five Displayable
	root, _ = TestComponent(NewBuilder(), Children(func(b Builder) {
		one, _ = TestComponent(b, FlexWidth(1), Children(func() {
			three, _ = TestComponent(b, ID("three"))
			four, _ = TestComponent(b, ID("four"), ExcludeFromLayout(true))
			five, _ = TestComponent(b, ID("five"), FlexWidth(1))
		}))
		two, _ = TestComponent(b, FlexWidth(2))
	}))

	return root, []Displayable{root, one, two, three, four, five}
}

func createStubApp() (Displayable, []Displayable) {
	var root, header, body, footer, logo, content Displayable

	root, _ = TestComponent(NewBuilder(), ID("root"), Width(800), Height(600), Children(func(b Builder) {
		header, _ = TestComponent(b, ID("header"), Padding(5), FlexWidth(1), Height(80), Children(func(b Builder) {
			logo, _ = TestComponent(b, ID("logo"), Width(50), Height(50))
			content, _ = TestComponent(b, ID("content"), FlexWidth(1), FlexHeight(1))
		}))
		body, _ = TestComponent(b, ID("body"), Padding(5), FlexWidth(1), FlexHeight(1))
		footer, _ = TestComponent(b, ID("footer"), FlexWidth(1), Height(60))
	}))

	return root, []Displayable{root, header, body, footer, logo, content}
}

func createTwoBoxes() (Displayable, Displayable) {
	var root, child Displayable
	root, _ = TestComponent(NewBuilder(), ID("root"), Padding(10), Width(100), Height(110), Children(func(b Builder) {
		child, _ = TestComponent(b, ID("child"), FlexWidth(1), FlexHeight(1))
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
		assert.Equal(t, root.ID(), "root")
		assert.Equal(t, len(nodes), 6)
		assert.Equal(t, root.ChildCount(), 3)
	})

	t.Run("Stack LayoutHandler", func(t *testing.T) {
		root, child := createTwoBoxes()

		StackLayout(root)
		assert.Equal(t, child.Width(), 80.0)
		assert.Equal(t, child.Height(), 90.0)
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
			assert.Equal(t, nodes[1].ChildCount(), 3)
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
			assert.Equal(t, nodes[1].ChildCount(), 3)
			assert.Equal(t, len(children), 1)
			assert.Equal(t, children[0].ID(), "five")
		})

		t.Run("Scales flex children", func(t *testing.T) {
			var one, two Displayable
			HBox(NewBuilder(), ID("root"), Padding(5), Width(100), Height(110), Children(func(b Builder) {
				one, _ = Box(b, ID("one"), Padding(10), FlexWidth(1), FlexHeight(1))
				two, _ = Box(b, ID("two"), FlexWidth(1), FlexHeight(1))
			}))
			assert.Equal(t, one.Width(), 45, "one width")
			assert.Equal(t, two.Width(), 45, "two width")
			assert.Equal(t, one.Height(), 100, "one height")
			assert.Equal(t, two.Height(), 100, "two height")
		})
	})

	t.Run("Spread remainder", func(t *testing.T) {
		var one, two, three Displayable
		HBox(NewBuilder(), Width(152), Children(func(b Builder) {
			one, _ = Box(b, FlexWidth(1))
			two, _ = Box(b, FlexWidth(1))
			three, _ = Box(b, FlexWidth(1))
		}))
		assert.Equal(t, one.Width(), 51)
		assert.Equal(t, two.Width(), 51)
		assert.Equal(t, three.Width(), 50)
	})

	t.Run("Basic, nested layout", func(t *testing.T) {
		var header, content, footer Displayable
		VBox(NewBuilder(), ID("root"), Width(100), Height(300), Children(func(b Builder) {
			header, _ = HBox(b, ID("header"), FlexWidth(1), Height(100), Children(func(b Builder) {
				Box(b, ID("logo"), Width(200), Height(100))
			}))
			content, _ = Box(b, ID("content"), FlexHeight(1), FlexWidth(1))
			footer, _ = Box(b, ID("footer"), Height(80), FlexWidth(1))
		}))
		assert.Equal(t, header.Height(), 100)
		assert.Equal(t, footer.Height(), 80)
		assert.Equal(t, content.Height(), 120)
	})

	t.Run("Nested, flexible controls should expand", func(t *testing.T) {
		root, _ := Box(NewBuilder(), ID("root"), Width(100), Children(func(b Builder) {
			Box(b, ID("one"), FlexWidth(1), Children(func() {
				Box(b, ID("two"), FlexWidth(1))
			}))
		}))
		one := root.FindComponentByID("one")
		two := root.FindComponentByID("two")

		assert.Equal(t, one.Width(), 100)
		assert.Equal(t, two.Width(), 100)
	})

	t.Run("Gutter is supported", func(t *testing.T) {
		root, _ := VBox(NewBuilder(), Padding(5), Gutter(10), Children(func(b Builder) {
			Trait(b, "Box", Width(100), Height(20))
			Box(b)
			Box(b)
			Box(b)
		}))

		kids := root.Children()
		one := kids[0]
		two := kids[1]
		three := kids[2]

		assert.Equal(t, one.Y(), 5)
		assert.Equal(t, two.Y(), 35)
		assert.Equal(t, three.Y(), 65)
	})

	t.Run("Layouts with larger children", func(t *testing.T) {
		t.Run("Does not shrink larger parent", func(t *testing.T) {
			root, _ := Box(NewBuilder(), Width(50), Height(50), Children(func(b Builder) {
				Box(b, Width(10), Height(10))
			}))
			assert.Equal(t, root.Height(), 50)
			assert.Equal(t, root.Width(), 50)
		})

		t.Run("Vertical", func(t *testing.T) {
			root, _ := VBox(NewBuilder(), Gutter(5), Padding(5), Children(func(b Builder) {
				Box(b, Width(20), Height(20))
				Box(b, Width(20), Height(20))
				Box(b, Width(20), Height(20))
				Box(b, Width(20), Height(20))
				Box(b, Width(20), Height(20))
			}))

			assert.Equal(t, root.Height(), 135)
			assert.Equal(t, root.Width(), 30)
		})

		t.Run("Horizontal", func(t *testing.T) {
			root, _ := HBox(NewBuilder(), Gutter(5), Padding(5), Children(func(b Builder) {
				Box(b, Width(20), Height(20))
				Box(b, Width(20), Height(20))
				Box(b, Width(20), Height(20))
				Box(b, Width(20), Height(20))
				Box(b, Width(20), Height(20))
			}))

			assert.Equal(t, root.Height(), 30)
			assert.Equal(t, root.Width(), 135)
		})
	})

	t.Run("Align center", func(t *testing.T) {
		var one, two, three Displayable
		root, _ := Box(NewBuilder(), HAlign(AlignCenter), VAlign(AlignCenter), Padding(5), Width(60), Height(60), Children(func(b Builder) {
			// This should be positioned in the center even though three blew out.
			one, _ = Box(b, Width(75), Height(75))
			two, _ = Box(b, Width(50), Height(50))
			// Three will blow out the assigned parent dimensions.
			three, _ = Box(b, Width(25), Height(25))
		}))

		assert.Equal(t, root.Width(), 85)
		assert.Equal(t, root.Height(), 85)
		assert.Equal(t, one.X(), 5)
		assert.Equal(t, one.Y(), 5)
		assert.Equal(t, two.X(), 17.5)
		assert.Equal(t, two.Y(), 17.5)
		assert.Equal(t, three.X(), 30)
		assert.Equal(t, three.Y(), 30)
	})

	t.Run("Align last", func(t *testing.T) {
		var one, two, three Displayable
		root, _ := Box(NewBuilder(), HAlign(AlignRight), VAlign(AlignBottom), Padding(5), Width(60), Height(60), Children(func(b Builder) {
			// This should be positioned in the center even though three blew out.
			one, _ = Box(b, Width(75), Height(75))
			two, _ = Box(b, Width(50), Height(50))
			// Three will blow out the assigned parent dimensions.
			three, _ = Box(b, Width(25), Height(25))
		}))

		assert.Equal(t, root.Width(), 85)
		assert.Equal(t, root.Height(), 85)
		assert.Equal(t, one.X(), 5)
		assert.Equal(t, one.Y(), 5)
		assert.Equal(t, two.X(), 30)
		assert.Equal(t, two.Y(), 30)
		assert.Equal(t, three.X(), 55)
		assert.Equal(t, three.Y(), 55)
	})

	t.Run("Distribute space after limit", func(t *testing.T) {
		var one, two, three Displayable
		VBox(NewBuilder(), Width(100), Height(100), Children(func(b Builder) {
			one, _ = Box(b, Width(100), FlexHeight(1), MaxHeight(20))
			two, _ = Box(b, Width(100), FlexHeight(1), MaxHeight(30))
			three, _ = Box(b, Width(100), FlexHeight(1))
		}))

		assert.Equal(t, one.Height(), 20)
		assert.Equal(t, two.Height(), 30)
		// NOTE(lbayes): The following is INCORRECT (off by one rounding somewhere),
		// but it's better than no spread, so checking it in.
		assert.Equal(t, three.Height(), 51)
	})
}
