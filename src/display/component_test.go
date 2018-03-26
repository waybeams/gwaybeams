package display

import (
	"assert"
	"strings"
	"testing"
)

func TestBaseComponent(t *testing.T) {
	t.Run("Generated Id", func(t *testing.T) {
		root := NewComponent()
		assert.Equal(t, len(root.GetId()), 20)
	})

	t.Run("Default Size", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		assert.Equal(t, box.GetFixedWidth(), -1)
		assert.Equal(t, box.GetFixedHeight(), -1)
	})

	t.Run("Default Size after Layout", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		box.Layout()
		if box.GetWidth() != 0 {
			t.Errorf("Expected width to be 0 but was %v", box.GetWidth())
		}
		if box.GetHeight() != 0 {
			t.Errorf("Expected height to be 0 but was %v", box.GetHAlign())

		}
	})

	t.Run("Default styles", func(t *testing.T) {
		// Create a new Box that will eventually become a child of another
		one := NewComponent()
		// Retrive styles (creating a default styles object)
		styles := one.GetStyles()

		if styles.GetFontSize() != DefaultStyleFontSize {
			t.Error("Expected to create styles from first request on root node")
		}

		t.Run("are removed when component is added to a parent", func(t *testing.T) {
			parent := NewComponent()
			parentStyles := parent.GetStyles()
			parentStyles.FontSize(11)

			two := NewComponent()
			parent.AddChild(one)
			parent.AddChild(two)

			oneStyles := one.GetStyles()
			if oneStyles.GetFontSize() != 11 {
				t.Error("Expected component to discard default font style, and defer to parent configuration")
			}

			twoStyles := one.GetStyles()
			if twoStyles.GetFontSize() != 11 {
				t.Error("Expected new component to pull styles from parent")
			}
		})
	})

	t.Run("Provided Id", func(t *testing.T) {
		root, _ := Box(NewBuilder(), Id("root"))
		assert.Equal(t, root.GetId(), "root")
	})

	t.Run("GetPath for root", func(t *testing.T) {
		root, _ := Box(NewBuilder(), Id("root"))
		assert.Equal(t, root.GetPath(), "/root")
	})

	t.Run("GetLayoutType default value", func(t *testing.T) {
		root, _ := Box(NewBuilder())
		if root.GetLayoutType() != StackLayoutType {
			t.Errorf("Expected %v but got %v", StackLayoutType, root.GetLayoutType())
		}
	})

	t.Run("MinHeight becomes unset Height", func(t *testing.T) {
		box, _ := Box(NewBuilder(), MinHeight(20))
		assert.Equal(t, box.GetHeight(), 20.0)
	})

	t.Run("MinWidth becomes unset Width", func(t *testing.T) {
		box, _ := Box(NewBuilder(), MinWidth(20))
		assert.Equal(t, box.GetWidth(), 20.0)
	})

	t.Run("MinHeight replaces existing Height", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		box.Height(10)
		box.MinHeight(20)
		assert.Equal(t, box.GetHeight(), 20.0)
	})

	t.Run("MinWidth replaces existing Width", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		box.Width(10)
		box.MinWidth(20)
		assert.Equal(t, box.GetWidth(), 20.0)
	})

	t.Run("MaxWidth constrained Width", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Width(50), MaxWidth(40))
		assert.Equal(t, box.GetWidth(), 40.0)
	})

	t.Run("MaxHeight constrained Height", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Height(51), MaxHeight(41))
		assert.Equal(t, box.GetHeight(), 41.0)
	})

	t.Run("MinWidth might expand actual", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Width(10), Height(11), MinWidth(20), MinHeight(21))

		assert.Equal(t, box.GetWidth(), 20.0)
		assert.Equal(t, box.GetHeight(), 21.0)
	})

	t.Run("WidthInBounds", func(t *testing.T) {
		box, _ := Box(NewBuilder(), MinWidth(10), MaxWidth(20), Width(15))
		box.Width(21)
		assert.Equal(t, box.GetWidth(), 20.0)
		box.Width(9)
		assert.Equal(t, box.GetWidth(), 10.0)
		box.Width(16)
		assert.Equal(t, box.GetWidth(), 16.0)
	})

	t.Run("Padding", func(t *testing.T) {

		t.Run("DefaultPadding", func(t *testing.T) {
			box, err := Box(NewBuilder())
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, box.GetPadding(), -1, "Default Padding")
			assert.Equal(t, box.GetPaddingBottom(), -1, "Default PaddingBottom")
			assert.Equal(t, box.GetPaddingTop(), -1, "Default PaddingTop")
			assert.Equal(t, box.GetPaddingLeft(), -1, "Default PaddingLeft")
			assert.Equal(t, box.GetPaddingRight(), -1, "Default PaddingRight")

			assert.Equal(t, box.GetMinWidth(), -1, "GetMinWidth")
			assert.Equal(t, box.GetMinHeight(), -1, "GetMinWidth")

			assert.Equal(t, box.GetWidth(), 0, "Width")
		})

		t.Run("Override side padding", func(t *testing.T) {
			box, err := Box(NewBuilder(), Padding(10))
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, box.GetPadding(), 10, "Default Padding")
			assert.Equal(t, box.GetPaddingBottom(), 10, "Default PaddingBottom")
			assert.Equal(t, box.GetPaddingTop(), 10, "Default PaddingTop")
			assert.Equal(t, box.GetPaddingLeft(), 10, "Default PaddingLeft")
			assert.Equal(t, box.GetPaddingRight(), 10, "Default PaddingRight")
		})

		t.Run("Interacts with GetMinWidth()", func(t *testing.T) {
			box, err := Box(NewBuilder(), Padding(10))
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, box.GetMinWidth(), 20, "GetMinWidth")
			assert.Equal(t, box.GetMinHeight(), 20, "GetMinWidth")
		})
	})

	t.Run("WidthInBounds from Child expansion plus Padding", func(t *testing.T) {
		box, err := Box(NewBuilder(), Padding(10), Width(30), Height(20), Children(func(b Builder) {
			Box(b, MinWidth(50), MinHeight(40))
			Box(b, MinWidth(30), MinHeight(30))
		}))

		if err != nil {
			t.Error(err)
			return
		}

		box.Width(10)
		box.Height(10)
		// This is a displayStack, so only the wider child expands parent.
		assert.Equal(t, box.GetWidth(), 70.0)
		// assert.Equal(t, box.GetHeight(), 60.0)
	})

	t.Run("GetPath with depth", func(t *testing.T) {
		var one, two, three, four Displayable
		Box(NewBuilder(), Id("root"), Children(func(b Builder) {
			one, _ = Box(b, Id("one"), Children(func() {
				two, _ = Box(b, Id("two"), Children(func() {
					three, _ = Box(b, Id("three"))
				}))
				four, _ = Box(b, Id("four"))
			}))
		}))

		assert.Equal(t, one.GetPath(), "/root/one")
		assert.Equal(t, two.GetPath(), "/root/one/two")
		assert.Equal(t, three.GetPath(), "/root/one/two/three")
		assert.Equal(t, four.GetPath(), "/root/one/four")
	})

	t.Run("GetOffsetFor", func(t *testing.T) {
		t.Run("Root at 0,0", func(t *testing.T) {
			root, _ := Box(NewBuilder())
			xOffset := root.GetXOffset()
			yOffset := root.GetYOffset()
			assert.Equal(t, xOffset, 0)
			assert.Equal(t, yOffset, 0)
		})

		t.Run("Root at offset", func(t *testing.T) {
			root, _ := Box(NewBuilder(), X(10), Y(15))
			xOffset := root.GetXOffset()
			yOffset := root.GetYOffset()
			assert.Equal(t, xOffset, 10)
			assert.Equal(t, yOffset, 15)
		})

		t.Run("Child at double offset", func(t *testing.T) {
			var nestedChild Displayable
			root, _ := Box(NewBuilder(), Padding(10), Children(func(b Builder) {
				Box(b, Padding(15), Children(func() {
					nestedChild, _ = Box(b, Padding(10))
				}))
			}))
			root.Layout()

			xOffset := nestedChild.GetXOffset()
			yOffset := nestedChild.GetYOffset()
			assert.Equal(t, xOffset, 25)
			assert.Equal(t, yOffset, 25)
		})
	})

	t.Run("Padding", func(t *testing.T) {
		t.Run("Applying Padding spreads to all four sides", func(t *testing.T) {
			root, _ := TestComponent(NewBuilder(), Padding(10))

			assert.Equal(t, root.GetHorizontalPadding(), 20.0)
			assert.Equal(t, root.GetVerticalPadding(), 20.0)

			assert.Equal(t, root.GetPaddingBottom(), 10.0)
			assert.Equal(t, root.GetPaddingLeft(), 10.0)
			assert.Equal(t, root.GetPaddingRight(), 10.0)
			assert.Equal(t, root.GetPaddingTop(), 10.0)
		})

		t.Run("PaddingTop overrides Padding", func(t *testing.T) {
			root, _ := TestComponent(NewBuilder(), Padding(10), PaddingTop(5))
			assert.Equal(t, root.GetPaddingTop(), 5.0)
			assert.Equal(t, root.GetPaddingBottom(), 10.0)
			assert.Equal(t, root.GetPadding(), 10.0)
		})

		t.Run("PaddingTop overrides Padding regardless of order", func(t *testing.T) {
			root, _ := TestComponent(NewBuilder(), PaddingTop(5), Padding(10))
			assert.Equal(t, root.GetPaddingTop(), 5.0)
			assert.Equal(t, root.GetPaddingBottom(), 10.0)
			assert.Equal(t, root.GetPadding(), 10.0)
		})
	})

	t.Run("PrefWidth default value", func(t *testing.T) {
		one := NewComponent()
		assert.Equal(t, 0.0, one.GetPrefWidth())
	})

	t.Run("PrefWidth ComponentModel value", func(t *testing.T) {
		one, _ := TestComponent(NewBuilder(), PrefWidth(200))
		assert.Equal(t, 200.0, one.GetPrefWidth())
	})

	t.Run("AddChild", func(t *testing.T) {
		root := NewComponent()
		one := NewComponent()
		two := NewComponent()
		root.Width(200)
		assert.Equal(t, root.AddChild(one), 1)
		assert.Equal(t, root.AddChild(two), 2)

		assert.Equal(t, one.GetParent().GetId(), root.GetId())
		assert.Equal(t, two.GetParent().GetId(), root.GetId())

		if root.GetParent() != nil {
			t.Error("Expected root.GetParent() to be nil")
		}
	})

	t.Run("GetChildCount", func(t *testing.T) {
		var one, two, three Displayable
		root, _ := Box(NewBuilder(), Children(func(b Builder) {
			one, _ = Box(b, Children(func() {
				two, _ = Box(b)
				three, _ = Box(b)
			}))
		}))

		assert.Equal(t, root.GetChildCount(), 1)
		assert.Equal(t, root.GetChildAt(0), one)

		assert.Equal(t, one.GetChildCount(), 2)
		assert.Equal(t, one.GetChildAt(0), two)
		assert.Equal(t, one.GetChildAt(1), three)
	})

	t.Run("GetFilteredChildren", func(t *testing.T) {
		createTree := func() (Displayable, []Displayable) {
			var root, one, two, three, four Displayable
			root, _ = Box(NewBuilder(), Children(func(b Builder) {
				one, _ = Box(b, Id("a-t-one"))
				two, _ = Box(b, Id("a-t-two"))
				three, _ = Box(b, Id("b-t-three"))
				four, _ = Box(b, Id("b-t-four"))
			}))

			return root, []Displayable{one, two, three, four}
		}

		allKids := func(d Displayable) bool {
			return strings.Index(d.GetId(), "-t-") > -1
		}

		bKids := func(d Displayable) bool {
			return strings.Index(d.GetId(), "b-") > -1
		}

		t.Run("returns Empty slice", func(t *testing.T) {
			root := NewComponent()
			filtered := root.GetFilteredChildren(allKids)
			assert.Equal(t, len(filtered), 0)
		})

		t.Run("returns all matched children in simple match", func(t *testing.T) {
			root, _ := createTree()
			filtered := root.GetFilteredChildren(allKids)
			assert.Equal(t, len(filtered), 4)
		})

		t.Run("returns all matched children in harder match", func(t *testing.T) {
			root, _ := createTree()
			filtered := root.GetFilteredChildren(bKids)
			assert.Equal(t, len(filtered), 2)
			assert.Equal(t, filtered[0].GetId(), "b-t-three")
			assert.Equal(t, filtered[1].GetId(), "b-t-four")
		})
	})

	t.Run("GetChildren returns empty list", func(t *testing.T) {
		root := NewComponent()
		children := root.GetChildren()

		if children == nil {
			t.Error("GetChildren should not return nil")
		}

		assert.Equal(t, len(children), 0)
	})

	t.Run("GetChildren returns new list", func(t *testing.T) {
		root, _ := Box(NewBuilder(), Children(func(b Builder) {
			Box(b)
			Box(b)
			Box(b)
		}))

		children := root.GetChildren()
		assert.Equal(t, len(children), 3)
	})
}
