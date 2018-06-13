package layout_test

import (
	"path/filepath"
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/fakes"
	"github.com/waybeams/waybeams/pkg/layout"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
	surface "github.com/waybeams/waybeams/pkg/surface/fakes"
)

func createStubApp() *spec.Spec {
	root := ctrl.VBox(opts.Key("root"), opts.Width(800), opts.Height(600),
		opts.Child(ctrl.HBox(opts.Key("header"), opts.Padding(5), opts.FlexWidth(1), opts.Height(80),
			opts.Child(ctrl.Box(opts.Key("logo"), opts.Width(50), opts.Height(50))),
			opts.Child(ctrl.Box(opts.Key("content"), opts.FlexWidth(1), opts.FlexHeight(1))),
		)),
		opts.Child(ctrl.Box(opts.Key("body"), opts.Padding(5), opts.FlexWidth(1), opts.FlexHeight(1))),
		opts.Child(ctrl.Box(opts.Key("footer"), opts.FlexWidth(1), opts.Height(60))),
	)

	return root
}

func TestLayout(t *testing.T) {
	var fakeSurface = func() spec.Surface {
		return surface.NewSurfaceFrom(filepath.Join("..", ".."))
	}

	t.Run("createStubApp works as expected", func(t *testing.T) {
		root := createStubApp()

		assert.Equal(root.Key(), "root")
		assert.Equal(root.ChildCount(), 3)
	})

	t.Run("Spread remainder", func(t *testing.T) {
		root := layout.Layout(fakes.Fake(
			opts.Width(152),
			opts.LayoutType(spec.HorizontalFlowLayoutType),
			opts.Child(fakes.Fake(opts.FlexWidth(1))),
			opts.Child(fakes.Fake(opts.FlexWidth(1))),
			opts.Child(fakes.Fake(opts.FlexWidth(1))),
		), fakeSurface())

		assert.Equal(root.ChildAt(0).Width(), 51)
		assert.Equal(root.ChildAt(1).Width(), 51)
		assert.Equal(root.ChildAt(2).Width(), 50)
	})

	t.Run("Stack parent dimensions grow to encapsulate children", func(t *testing.T) {
		root := ctrl.Box(
			opts.Key("root"),
			opts.Width(40),
			opts.Height(45),
			opts.Child(ctrl.Box(
				opts.Key("one"),
				opts.Width(50),
				opts.Height(55),
				opts.Child(ctrl.Box(
					opts.Key("two"),
					opts.Width(60),
					opts.Height(65),
				)),
			)),
		)
		layout.Layout(root, fakeSurface())

		one := spec.FirstByKey(root, "one")
		assert.Equal(one.Width(), 60, "one.W")
		assert.Equal(one.Height(), 65, "one.H")

		assert.Equal(root.Width(), 60, "root.W")
		assert.Equal(root.Height(), 65, "root.H")
	})

	t.Run("Oversized flex values should not break layouts", func(t *testing.T) {
		root := ctrl.VBox(
			opts.Width(100),
			opts.Height(120),
			opts.Child(fakes.Fake(
				opts.Key("one"),
				opts.FlexHeight(3),
				opts.FlexWidth(1),
			)),
			opts.Child(fakes.Fake(
				opts.Key("two"),
				opts.Height(20),
				opts.FlexWidth(1),
			)),
		)
		layout.Layout(root, fakeSurface())

		// Prior to a bug fix where we added math.Floor to flowGetUnitSize, we were getting
		// oversizing containers because of floating point remainders.
		assert.Equal(root.Height(), 120)
		assert.Equal(root.ChildAt(0).Height(), 100)
		assert.Equal(root.ChildAt(1).Height(), 20)
	})

	t.Run("GetFlexibleChildren", func(t *testing.T) {
		t.Run("Scales flex children", func(t *testing.T) {

			root := ctrl.HBox(
				opts.Key("root"),
				opts.Padding(5),
				opts.Width(100),
				opts.Height(110),
				opts.Child(ctrl.Box(
					opts.Key("one"),
					opts.Padding(10),
					opts.FlexWidth(1),
					opts.FlexHeight(1),
				)),
				opts.Child(ctrl.Box(
					opts.Key("two"),
					opts.FlexWidth(1),
					opts.FlexHeight(1),
				)),
			)
			layout.Layout(root, fakeSurface())

			one := spec.FirstByKey(root, "one")
			two := spec.FirstByKey(root, "two")

			assert.Equal(one.Width(), 45, "one width")
			assert.Equal(two.Width(), 45, "two width")
			assert.Equal(one.Height(), 100, "one height")
			assert.Equal(two.Height(), 100, "two height")
		})
	})

	t.Run("Spread remainder", func(t *testing.T) {
		root := ctrl.HBox(
			opts.Width(152),
			opts.Child(ctrl.Box(
				opts.Key("one"),
				opts.FlexWidth(1),
				opts.FlexHeight(1),
			)),
			opts.Child(ctrl.Box(
				opts.Key("two"),
				opts.FlexWidth(1),
				opts.FlexHeight(1),
			)),
			opts.Child(ctrl.Box(
				opts.Key("three"),
				opts.FlexWidth(1),
				opts.FlexHeight(1),
			)),
		)
		layout.Layout(root, fakeSurface())

		one := spec.FirstByKey(root, "one")
		two := spec.FirstByKey(root, "two")
		three := spec.FirstByKey(root, "three")

		assert.Equal(one.Width(), 51)
		assert.Equal(two.Width(), 51)
		assert.Equal(three.Width(), 50)
	})

	t.Run("Basic, nested layout", func(t *testing.T) {
		root := ctrl.VBox(
			opts.Key("root"),
			opts.Width(100),
			opts.Height(300),
			opts.Child(ctrl.HBox(
				opts.Key("header"),
				opts.FlexWidth(1),
				opts.Height(100),
				opts.Child(ctrl.Box(
					opts.Key("logo"),
					opts.Width(200),
					opts.Height(100),
				)),
			)),
			opts.Child(ctrl.Box(
				opts.Key("content"),
				opts.FlexHeight(1),
				opts.FlexWidth(1),
			)),
			opts.Child(ctrl.Box(
				opts.Key("footer"),
				opts.Height(80),
				opts.FlexWidth(1),
			)),
		)

		layout.Layout(root, fakeSurface())
		header := spec.FirstByKey(root, "header")
		footer := spec.FirstByKey(root, "footer")
		content := spec.FirstByKey(root, "content")

		assert.Equal(header.Height(), 100)
		assert.Equal(footer.Height(), 80)
		assert.Equal(content.Height(), 120)
	})

	t.Run("Nested, flexible controls should expand", func(t *testing.T) {
		root := ctrl.Box(
			opts.Key("root"),
			opts.Width(100),
			opts.Child(ctrl.Box(
				opts.Key("one"),
				opts.FlexWidth(1),
				opts.Child(ctrl.Box(
					opts.Key("two"),
					opts.FlexWidth(1),
				)),
			)),
		)
		layout.Layout(root, fakeSurface())

		one := spec.FirstByKey(root, "one")
		two := spec.FirstByKey(root, "two")

		assert.Equal(one.Width(), 100)
		assert.Equal(two.Width(), 100)
	})

	t.Run("Gutter is supported", func(t *testing.T) {
		root := ctrl.VBox(
			opts.Padding(5),
			opts.Gutter(10),
			opts.Child(ctrl.Box(opts.Width(100), opts.Height(20))),
			opts.Child(ctrl.Box(opts.Width(100), opts.Height(20))),
			opts.Child(ctrl.Box(opts.Width(100), opts.Height(20))),
		)

		layout.Layout(root, fakeSurface())

		kids := root.Children()
		one := kids[0]
		two := kids[1]
		three := kids[2]

		assert.Equal(one.Y(), 5)
		assert.Equal(two.Y(), 35)
		assert.Equal(three.Y(), 65)
	})

	t.Run("Layouts with larger children", func(t *testing.T) {
		t.Run("Does not shrink larger parent", func(t *testing.T) {
			root := ctrl.Box(
				opts.Width(50),
				opts.Height(50),
				opts.Child(ctrl.Box(
					opts.Width(10),
					opts.Height(10),
				)),
			)

			assert.Equal(root.Height(), 50)
			assert.Equal(root.Width(), 50)
		})

		t.Run("Vertical", func(t *testing.T) {
			root := ctrl.VBox(
				opts.Gutter(5),
				opts.Padding(5),
				opts.Child(ctrl.Box(opts.Width(20), opts.Height(20))),
				opts.Child(ctrl.Box(opts.Width(20), opts.Height(20))),
			)

			layout.Layout(root, fakeSurface())
			assert.Equal(root.Height(), 55)
			assert.Equal(root.Width(), 30)
		})

		t.Run("Horizontal", func(t *testing.T) {
			root := ctrl.HBox(
				opts.Gutter(5),
				opts.Padding(5),
				opts.Child(ctrl.Box(opts.Width(20), opts.Height(20))),
				opts.Child(ctrl.Box(opts.Width(20), opts.Height(20))),
			)

			layout.Layout(root, fakeSurface())
			assert.Equal(root.Height(), 30)
			assert.Equal(root.Width(), 55)
		})
	})

	t.Run("Align center", func(t *testing.T) {
		root := ctrl.Box(
			opts.HAlign(spec.AlignCenter),
			opts.VAlign(spec.AlignCenter),
			opts.Padding(5),
			opts.Width(60),
			opts.Height(60),
			// This should be positioned in the center even though three blew out the size.
			opts.Child(ctrl.Box(opts.Key("one"), opts.Width(75), opts.Height(75))),
			opts.Child(ctrl.Box(opts.Key("two"), opts.Width(50), opts.Height(50))),
			// Three will blow out the assigned parent dimensions.
			opts.Child(ctrl.Box(opts.Key("three"), opts.Width(25), opts.Height(25))),
		)

		layout.Layout(root, fakeSurface())

		one := spec.FirstByKey(root, "one")
		two := spec.FirstByKey(root, "two")
		three := spec.FirstByKey(root, "three")

		assert.Equal(root.Width(), 85)
		assert.Equal(root.Height(), 85)
		assert.Equal(one.X(), 5)
		assert.Equal(one.Y(), 5)
		assert.Equal(two.X(), 17.5)
		assert.Equal(two.Y(), 17.5)
		assert.Equal(three.X(), 30)
		assert.Equal(three.Y(), 30)
	})

	t.Run("Align last", func(t *testing.T) {
		root := ctrl.Box(
			opts.HAlign(spec.AlignRight),
			opts.VAlign(spec.AlignBottom),
			opts.Padding(5),
			opts.Width(60),
			opts.Height(60),
			// This should be positioned in the center even though three blew out.
			opts.Child(ctrl.Box(
				opts.Key("one"),
				opts.Width(75),
				opts.Height(75),
			)),
			opts.Child(ctrl.Box(
				opts.Key("two"),
				opts.Width(50),
				opts.Height(50),
			)),
			// Three will blow out the assigned parent dimensions.
			opts.Child(ctrl.Box(
				opts.Key("three"),
				opts.Width(25),
				opts.Height(25),
			)),
		)

		layout.Layout(root, fakeSurface())

		one := spec.FirstByKey(root, "one")
		two := spec.FirstByKey(root, "two")
		three := spec.FirstByKey(root, "three")

		assert.Equal(root.Width(), 85)
		assert.Equal(root.Height(), 85)
		assert.Equal(one.X(), 5)
		assert.Equal(one.Y(), 5)
		assert.Equal(two.X(), 30)
		assert.Equal(two.Y(), 30)
		assert.Equal(three.X(), 55)
		assert.Equal(three.Y(), 55)
	})

	t.Run("Distribute space after limit", func(t *testing.T) {
		root := ctrl.VBox(
			opts.Key("root"),
			opts.Width(100),
			opts.Height(100),
			opts.Child(ctrl.Box(
				opts.Key("one"),
				opts.Width(100),
				opts.FlexHeight(1),
				opts.MaxHeight(20),
			)),
			opts.Child(ctrl.Box(
				opts.Key("two"),
				opts.Width(100),
				opts.FlexHeight(1),
				opts.MaxHeight(30),
			)),
			opts.Child(ctrl.Box(
				opts.Key("three"),
				opts.Width(100),
				opts.FlexHeight(1),
			)),
		)

		layout.Layout(root, fakeSurface())

		one := spec.FirstByKey(root, "one")
		two := spec.FirstByKey(root, "two")
		three := spec.FirstByKey(root, "three")

		assert.Equal(one.Height(), 20)
		assert.Equal(two.Height(), 30)
		assert.Equal(three.Height(), 50)
	})

	t.Run("Todo Item Height", func(t *testing.T) {
		root := ctrl.VBox(
			opts.Key("Todo Items"),
			opts.MinHeight(300),
			opts.FlexWidth(1),
			opts.Child(ctrl.HBox(
				opts.Key("item-0"),
				opts.FlexWidth(1),
				opts.Child(ctrl.Button(
					opts.Key("btn"),
					opts.Text("Some Label Value"),
				)),
			)),
			opts.Child(ctrl.HBox(
				opts.Key("item-1"),
				opts.FlexWidth(1),
				opts.Child(ctrl.Button(
					opts.Key("btn"),
					opts.Text("Other Label Value"),
				)),
			)),
		)

		layout.Layout(root, fakeSurface())

		assert.Equal(root.Height(), 300)

		child := spec.FirstByKey(root, "item-0")
		assert.Equal(child.Width(), 173)
		assert.Equal(child.Height(), 34)

		child = spec.FirstByKey(root, "item-1")
		assert.Equal(child.Y(), 34)
		assert.Equal(child.Width(), 170)
		assert.Equal(child.Height(), 34)
	})
}
