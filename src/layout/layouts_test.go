package layout_test

import (
	"assert"
	"controls"
	"fakes"
	"layout"
	. "opts"
	"spec"
	"surface"
	"testing"
)

func createStubApp() *spec.Spec {
	root := controls.VBox(Key("root"), Width(800), Height(600),
		Child(controls.HBox(Key("header"), Padding(5), FlexWidth(1), Height(80),
			Child(controls.Box(Key("logo"), Width(50), Height(50))),
			Child(controls.Box(Key("content"), FlexWidth(1), FlexHeight(1))),
		)),
		Child(controls.Box(Key("body"), Padding(5), FlexWidth(1), FlexHeight(1))),
		Child(controls.Box(Key("footer"), FlexWidth(1), Height(60))),
	)

	return root
}

/*
func createTwoBoxes() *fakes.FakeSpec {
	return fakes.Fake(Key("root"), Padding(10), Width(100), Height(110),
		Child(fakes.Fake(Key("child"), FlexWidth(1), FlexHeight(1))),
	)
}
*/

func TestLayout(t *testing.T) {
	t.Run("createStubApp works as expected", func(t *testing.T) {
		root := createStubApp()

		assert.Equal(t, root.Key(), "root")
		assert.Equal(t, root.ChildCount(), 3)
	})

	t.Run("Spread remainder", func(t *testing.T) {
		root := layout.Layout(fakes.Fake(
			Width(152),
			LayoutType(spec.HorizontalFlowLayoutType),
			Child(fakes.Fake(FlexWidth(1))),
			Child(fakes.Fake(FlexWidth(1))),
			Child(fakes.Fake(FlexWidth(1))),
		), surface.NewFake())

		assert.Equal(t, root.ChildAt(0).Width(), 51)
		assert.Equal(t, root.ChildAt(1).Width(), 51)
		assert.Equal(t, root.ChildAt(2).Width(), 50)
	})

	t.Run("Parent dimensions grow to encapsulate children", func(t *testing.T) {
		root := controls.VBox(
			Key("root"),
			Width(40),
			Height(45),
			Child(controls.VBox(
				Key("one"),
				Width(50),
				Height(55),
				Child(controls.Box(
					Key("two"),
					Width(60),
					Height(65),
				)),
			)),
		)
		layout.Layout(root, surface.NewFake())

		one := spec.FirstByKey(root, "one")
		assert.Equal(t, root.Width(), 60, "root.W")
		assert.Equal(t, root.Height(), 65, "root.H")

		assert.Equal(t, one.Width(), 60, "one.W")
		assert.Equal(t, one.Height(), 65, "one.H")
	})

	t.Run("Oversized flex values should not break layouts", func(t *testing.T) {
		t.Skip()
		root := controls.VBox(
			Width(100),
			Height(100),
			Child(fakes.Fake(
				Key("one"),
				FlexHeight(3),
				FlexWidth(1),
			)),
			Child(fakes.Fake(
				Key("two"),
				Height(20),
				FlexWidth(1),
			)),
		)

		root.SetHeight(120)

		// Prior to a bug fix where we added math.Floor to flowGetUnitSize, we were getting
		// oversizing containers because of floating point remainders.
		assert.Equal(t, root.Height(), 120)
		assert.Equal(t, root.ChildAt(0).Height(), 100)
		assert.Equal(t, root.ChildAt(1).Height(), 20)
	})
}

/*

func createStubApp() (Displayable, []Displayable) {
	var root, header, body, footer, logo, content Displayable

	root = TestControl(context.New(), ID("root"), Width(800), Height(600), Children(func(c Context) {
		header = TestControl(c, ID("header"), Padding(5), FlexWidth(1), Height(80), Children(func(c Context) {
			logo = TestControl(c, ID("logo"), Width(50), Height(50))
			content = TestControl(c, ID("content"), FlexWidth(1), FlexHeight(1))
		}))
		body = TestControl(c, ID("body"), Padding(5), FlexWidth(1), FlexHeight(1))
		footer = TestControl(c, ID("footer"), FlexWidth(1), Height(60))
	}))

	return root, []Displayable{root, header, body, footer, logo, content}
}

func createTwoBoxes() (Displayable, Displayable) {
	var root, child Displayable
	root = TestControl(context.New(), ID("root"), Padding(10), Width(100), Height(110), Children(func(c Context) {
		child = TestControl(c, ID("child"), FlexWidth(1), FlexHeight(1))
	}))
	return root, child
}

func TestLayout(t *testing.T) {
	t.Run("Call LayoutHandler", func(t *testing.T) {
		root := control.New()
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

	t.Run("GetFlexibleChildren", func(t *testing.T) {
		t.Run("Scales flex children", func(t *testing.T) {
			var one, two Displayable
			HBox(context.New(), ID("root"), Padding(5), Width(100), Height(110), Children(func(c Context) {
				one = Box(c, ID("one"), Padding(10), FlexWidth(1), FlexHeight(1))
				two = Box(c, ID("two"), FlexWidth(1), FlexHeight(1))
			}))
			assert.Equal(t, one.Width(), 45, "one width")
			assert.Equal(t, two.Width(), 45, "two width")
			assert.Equal(t, one.Height(), 100, "one height")
			assert.Equal(t, two.Height(), 100, "two height")
		})
	})

	t.Run("Spread remainder", func(t *testing.T) {
		var one, two, three Displayable
		HBox(context.New(), Width(152), Children(func(c Context) {
			one = Box(c, FlexWidth(1))
			two = Box(c, FlexWidth(1))
			three = Box(c, FlexWidth(1))
		}))
		assert.Equal(t, one.Width(), 51)
		assert.Equal(t, two.Width(), 51)
		assert.Equal(t, three.Width(), 50)
	})

	t.Run("Basic, nested layout", func(t *testing.T) {
		var header, content, footer Displayable
		VBox(context.New(), ID("root"), Width(100), Height(300), Children(func(c Context) {
			header = HBox(c, ID("header"), FlexWidth(1), Height(100), Children(func(c Context) {
				Box(c, ID("logo"), Width(200), Height(100))
			}))
			content = Box(c, ID("content"), FlexHeight(1), FlexWidth(1))
			footer = Box(c, ID("footer"), Height(80), FlexWidth(1))
		}))
		assert.Equal(t, header.Height(), 100)
		assert.Equal(t, footer.Height(), 80)
		assert.Equal(t, content.Height(), 120)
	})

	t.Run("Nested, flexible controls should expand", func(t *testing.T) {
		root := Box(context.New(), ID("root"), Width(100), Children(func(c Context) {
			Box(c, ID("one"), FlexWidth(1), Children(func() {
				Box(c, ID("two"), FlexWidth(1))
			}))
		}))
		one := root.FindControlById("one")
		two := root.FindControlById("two")

		assert.Equal(t, one.Width(), 100)
		assert.Equal(t, two.Width(), 100)
	})

	t.Run("Gutter is supported", func(t *testing.T) {
		root := VBox(context.New(), Padding(5), Gutter(10), Children(func(c Context) {
			Box(c, Width(100), Height(20))
			Box(c, Width(100), Height(20))
			Box(c, Width(100), Height(20))
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
			root := Box(context.New(), Width(50), Height(50), Children(func(c Context) {
				Box(c, Width(10), Height(10))
			}))
			assert.Equal(t, root.Height(), 50)
			assert.Equal(t, root.Width(), 50)
		})

		t.Run("Vertical", func(t *testing.T) {
			root := VBox(context.New(), Gutter(5), Padding(5), Children(func(c Context) {
				Box(c, Width(20), Height(20))
				Box(c, Width(20), Height(20))
				Box(c, Width(20), Height(20))
				Box(c, Width(20), Height(20))
				Box(c, Width(20), Height(20))
			}))

			assert.Equal(t, root.Height(), 135)
			assert.Equal(t, root.Width(), 30)
		})

		t.Run("Horizontal", func(t *testing.T) {
			root := HBox(context.New(), Gutter(5), Padding(5), Children(func(c Context) {
				Box(c, Width(20), Height(20))
				Box(c, Width(20), Height(20))
				Box(c, Width(20), Height(20))
				Box(c, Width(20), Height(20))
				Box(c, Width(20), Height(20))
			}))

			assert.Equal(t, root.Height(), 30)
			assert.Equal(t, root.Width(), 135)
		})
	})

	t.Run("Align center", func(t *testing.T) {
		var one, two, three Displayable
		root := Box(context.New(), HAlign(AlignCenter), VAlign(AlignCenter), Padding(5), Width(60), Height(60), Children(func(c Context) {
			// This should be positioned in the center even though three blew out.
			one = Box(c, Width(75), Height(75))
			two = Box(c, Width(50), Height(50))
			// Three will blow out the assigned parent dimensions.
			three = Box(c, Width(25), Height(25))
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
		root := Box(context.New(), HAlign(AlignRight), VAlign(AlignBottom), Padding(5), Width(60), Height(60), Children(func(c Context) {
			// This should be positioned in the center even though three blew out.
			one = Box(c, Width(75), Height(75))
			two = Box(c, Width(50), Height(50))
			// Three will blow out the assigned parent dimensions.
			three = Box(c, Width(25), Height(25))
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
		VBox(context.New(), Width(100), Height(100), Children(func(c Context) {
			one = Box(c, Width(100), FlexHeight(1), MaxHeight(20))
			two = Box(c, Width(100), FlexHeight(1), MaxHeight(30))
			three = Box(c, Width(100), FlexHeight(1))
		}))

		assert.Equal(t, one.Height(), 20)
		assert.Equal(t, two.Height(), 30)
		// NOTE(lbayes): The following is INCORRECT (off by one rounding somewhere),
		// but it's better than no spread, so checking it in.
		assert.Equal(t, three.Height(), 51)
	})

}

*/
