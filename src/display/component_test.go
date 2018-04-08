package display

import (
	"assert"
	"strings"
	"testing"
)

func TestBaseComponent(t *testing.T) {
	t.Run("Generated ID", func(t *testing.T) {
		root := NewComponent()
		assert.Equal(t, len(root.ID()), 20, "ID")
	})

	t.Run("Default Size", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		assert.Equal(t, box.FixedWidth(), -1, "FixedWidth")
		assert.Equal(t, box.FixedHeight(), -1, "FixedHeight")
	})

	t.Run("Default Size after Layout", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		box.Layout()
		if box.Width() != 0 {
			t.Errorf("Expected width to be 0 but was %v", box.Width())
		}
		if box.Height() != 0 {
			t.Errorf("Expected height to be 0 but was %v", box.HAlign())

		}
	})

	t.Run("Provided ID", func(t *testing.T) {
		root, _ := Box(NewBuilder(), ID("root"))
		assert.Equal(t, root.ID(), "root")
	})

	t.Run("GetPath for root", func(t *testing.T) {
		root, _ := Box(NewBuilder(), ID("root"))
		assert.Equal(t, root.Path(), "/root")
	})

	t.Run("GetLayoutType default value", func(t *testing.T) {
		root, _ := Box(NewBuilder())
		if root.LayoutType() != StackLayoutType {
			t.Errorf("Expected %v but got %v", StackLayoutType, root.LayoutType())
		}
	})

	t.Run("MinHeight becomes unset Height", func(t *testing.T) {
		box, _ := Box(NewBuilder(), MinHeight(20))
		assert.Equal(t, box.Height(), 20.0)
	})

	t.Run("MinWidth becomes unset Width", func(t *testing.T) {
		box, _ := Box(NewBuilder(), MinWidth(20))
		assert.Equal(t, box.Width(), 20.0)
	})

	t.Run("MinHeight replaces existing Height", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		box.SetHeight(10)
		box.SetMinHeight(20)
		assert.Equal(t, box.Height(), 20.0)
	})

	t.Run("MinWidth replaces existing Width", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		box.SetWidth(10)
		box.SetMinWidth(20)
		assert.Equal(t, box.Width(), 20.0)
	})

	t.Run("MaxWidth constrained Width", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Width(50), MaxWidth(40))
		assert.Equal(t, box.Width(), 40.0)
	})

	t.Run("MaxHeight constrained Height", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Height(51), MaxHeight(41))
		assert.Equal(t, box.Height(), 41.0)
	})

	t.Run("MinWidth might expand actual", func(t *testing.T) {
		box, _ := Box(NewBuilder(), Width(10), Height(11), MinWidth(20), MinHeight(21))

		assert.Equal(t, box.Width(), 20.0)
		assert.Equal(t, box.Height(), 21.0)
	})

	t.Run("WidthInBounds", func(t *testing.T) {
		box, _ := Box(NewBuilder(), MinWidth(10), MaxWidth(20), Width(15))
		box.SetWidth(21)
		assert.Equal(t, box.Width(), 20.0)
		box.SetWidth(9)
		assert.Equal(t, box.Width(), 10.0)
		box.SetWidth(16)
		assert.Equal(t, box.Width(), 16.0)
	})

	t.Run("Padding", func(t *testing.T) {

		t.Run("DefaultPadding", func(t *testing.T) {
			box, err := Box(NewBuilder())
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, box.Padding(), -1, "Default Padding")
			assert.Equal(t, box.PaddingBottom(), -1, "Default PaddingBottom")
			assert.Equal(t, box.PaddingTop(), -1, "Default PaddingTop")
			assert.Equal(t, box.PaddingLeft(), -1, "Default PaddingLeft")
			assert.Equal(t, box.PaddingRight(), -1, "Default PaddingRight")

			assert.Equal(t, box.MinWidth(), -1, "GetMinWidth")
			assert.Equal(t, box.MinHeight(), -1, "GetMinWidth")

			assert.Equal(t, box.Width(), 0, "Width")
		})

		t.Run("Override side padding", func(t *testing.T) {
			box, err := Box(NewBuilder(), Padding(10))
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, box.Padding(), 10, "Default Padding")
			assert.Equal(t, box.PaddingBottom(), 10, "Default PaddingBottom")
			assert.Equal(t, box.PaddingTop(), 10, "Default PaddingTop")
			assert.Equal(t, box.PaddingLeft(), 10, "Default PaddingLeft")
			assert.Equal(t, box.PaddingRight(), 10, "Default PaddingRight")
		})

		t.Run("Interacts with GetMinWidth()", func(t *testing.T) {
			box, err := Box(NewBuilder(), Padding(10))
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, box.MinWidth(), 20, "GetMinWidth")
			assert.Equal(t, box.MinHeight(), 20, "GetMinWidth")
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

		box.SetWidth(10)
		box.SetHeight(10)
		// This is a stack, so only the wider child expands parent.
		assert.Equal(t, box.Width(), 70.0)
		// assert.Equal(t, box.GetHeight(), 60.0)
	})

	t.Run("GetPath with depth", func(t *testing.T) {
		var one, two, three, four Displayable
		Box(NewBuilder(), ID("root"), Children(func(b Builder) {
			one, _ = Box(b, ID("one"), Children(func() {
				two, _ = Box(b, ID("two"), Children(func() {
					three, _ = Box(b, ID("three"))
				}))
				four, _ = Box(b, ID("four"))
			}))
		}))

		assert.Equal(t, one.Path(), "/root/one")
		assert.Equal(t, two.Path(), "/root/one/two")
		assert.Equal(t, three.Path(), "/root/one/two/three")
		assert.Equal(t, four.Path(), "/root/one/four")
	})

	t.Run("GetOffsetFor", func(t *testing.T) {
		t.Run("Root at 0,0", func(t *testing.T) {
			root, _ := Box(NewBuilder())
			xOffset := root.XOffset()
			yOffset := root.YOffset()
			assert.Equal(t, xOffset, 0)
			assert.Equal(t, yOffset, 0)
		})

		t.Run("Root at offset", func(t *testing.T) {
			root, _ := Box(NewBuilder(), X(10), Y(15))
			xOffset := root.XOffset()
			yOffset := root.YOffset()
			assert.Equal(t, xOffset, 10)
			assert.Equal(t, yOffset, 15)
		})

		t.Run("Child receives offset for padding", func(t *testing.T) {
			var root, child Displayable
			root, _ = Box(NewBuilder(), Padding(10), Width(100), Height(100), Children(func(b Builder) {
				child, _ = Box(b, FlexWidth(1), FlexHeight(1))
			}))

			root.Layout()
			assert.Equal(t, root.XOffset(), 0)
			assert.Equal(t, child.XOffset(), 10)
		})

		t.Run("Child at double offset", func(t *testing.T) {
			var nestedChild Displayable
			root, _ := Box(NewBuilder(), Padding(10), Children(func(b Builder) {
				Box(b, Padding(15), Children(func() {
					nestedChild, _ = Box(b, Padding(10))
				}))
			}))
			root.Layout()

			xOffset := nestedChild.XOffset()
			yOffset := nestedChild.YOffset()
			assert.Equal(t, xOffset, 25)
			assert.Equal(t, yOffset, 25)
		})
	})

	t.Run("Padding", func(t *testing.T) {
		t.Run("Applying Padding spreads to all four sides", func(t *testing.T) {
			root, _ := TestComponent(NewBuilder(), Padding(10))

			assert.Equal(t, root.HorizontalPadding(), 20.0)
			assert.Equal(t, root.VerticalPadding(), 20.0)

			assert.Equal(t, root.PaddingBottom(), 10.0)
			assert.Equal(t, root.PaddingLeft(), 10.0)
			assert.Equal(t, root.PaddingRight(), 10.0)
			assert.Equal(t, root.PaddingTop(), 10.0)
		})

		t.Run("PaddingTop overrides Padding", func(t *testing.T) {
			root, _ := TestComponent(NewBuilder(), Padding(10), PaddingTop(5))
			assert.Equal(t, root.PaddingTop(), 5.0)
			assert.Equal(t, root.PaddingBottom(), 10.0)
			assert.Equal(t, root.Padding(), 10.0)
		})

		t.Run("PaddingTop overrides Padding regardless of order", func(t *testing.T) {
			root, _ := TestComponent(NewBuilder(), PaddingTop(5), Padding(10))
			assert.Equal(t, root.PaddingTop(), 5.0)
			assert.Equal(t, root.PaddingBottom(), 10.0)
			assert.Equal(t, root.Padding(), 10.0)
		})
	})

	t.Run("PrefWidth default value", func(t *testing.T) {
		one := NewComponent()
		assert.Equal(t, 0.0, one.PrefWidth())
	})

	t.Run("PrefWidth ComponentModel value", func(t *testing.T) {
		one, _ := TestComponent(NewBuilder(), PrefWidth(200))
		assert.Equal(t, 200.0, one.PrefWidth())
	})

	t.Run("AddChild", func(t *testing.T) {
		root := NewComponent()
		one := NewComponent()
		two := NewComponent()
		root.SetWidth(200)
		assert.Equal(t, root.AddChild(one), 1)
		assert.Equal(t, root.AddChild(two), 2)

		assert.Equal(t, one.Parent().ID(), root.ID())
		assert.Equal(t, two.Parent().ID(), root.ID())

		if root.Parent() != nil {
			t.Error("Expected root.Parent() to be nil")
		}
	})

	t.Run("ChildCount", func(t *testing.T) {
		var one, two, three Displayable
		root, _ := Box(NewBuilder(), Children(func(b Builder) {
			one, _ = Box(b, Children(func() {
				two, _ = Box(b)
				three, _ = Box(b)
			}))
		}))

		assert.Equal(t, root.ChildCount(), 1)
		assert.Equal(t, root.ChildAt(0), one)

		assert.Equal(t, one.ChildCount(), 2)
		assert.Equal(t, one.ChildAt(0), two)
		assert.Equal(t, one.ChildAt(1), three)
	})

	t.Run("GetFilteredChildren", func(t *testing.T) {
		createTree := func() (Displayable, []Displayable) {
			var root, one, two, three, four Displayable
			root, _ = Box(NewBuilder(), Children(func(b Builder) {
				one, _ = Box(b, ID("a-t-one"))
				two, _ = Box(b, ID("a-t-two"))
				three, _ = Box(b, ID("b-t-three"))
				four, _ = Box(b, ID("b-t-four"))
			}))

			return root, []Displayable{one, two, three, four}
		}

		allKids := func(d Displayable) bool {
			return strings.Index(d.ID(), "-t-") > -1
		}

		bKids := func(d Displayable) bool {
			return strings.Index(d.ID(), "b-") > -1
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
			assert.Equal(t, filtered[0].ID(), "b-t-three")
			assert.Equal(t, filtered[1].ID(), "b-t-four")
		})
	})

	t.Run("GetChildren returns empty list", func(t *testing.T) {
		root := NewComponent()
		children := root.Children()

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

		children := root.Children()
		assert.Equal(t, len(children), 3)
	})

	t.Run("GetFontFace", func(t *testing.T) {
		root, _ := Box(NewBuilder())
		assert.Equal(t, root.FontFace(), "sans")
		assert.Equal(t, root.FontSize(), 12)
		assert.Equal(t, root.BgColor(), 0x999999ff, "BgColor")
		assert.Equal(t, root.StrokeColor(), 0x333333ff, "StrokeColor")
	})

	t.Run("GetBuilder", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		if box.Builder() == nil {
			t.Error("Component factory should assign the builder")
		}
	})

	t.Run("Root returns deeply nested root component", func(t *testing.T) {
		var descendant Displayable
		root, _ := Box(NewBuilder(), ID("root"), Children(func(b Builder) {
			Box(b, ID("one"), Children((func() {
				Box(b, ID("two"), Children(func() {
					Box(b, ID("three"), Children(func() {
						Box(b, ID("four"), Children(func() {
							Box(b, ID("five"), Children(func() {
								descendant, _ = Box(b, ID("child"))
							}))
						}))
					}))
				}))
			})))
		}))
		assert.Equal(t, root.ID(), descendant.Root().ID())
	})

	t.Run("IsContainedBy", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			one := NewComponent()
			two := NewComponent()
			if one.IsContainedBy(two) {
				t.Error("Unrelated nodes are not ancestors")
			}
		})

		t.Run("False for same component", func(t *testing.T) {
			one := NewComponent()
			if one.IsContainedBy(one) {
				t.Error("A component should not be contained by itself")
			}
		})

		t.Run("Child is true", func(t *testing.T) {
			one := NewComponent()
			two := NewComponent()
			one.AddChild(two)
			if !two.IsContainedBy(one) {
				t.Error("One should be an ancestor of two")
			}
			if one.IsContainedBy(two) {
				t.Error("Two is not an ancestor of one")
			}
		})

		t.Run("Deep descendants too", func(t *testing.T) {
			one := NewComponent()
			two := NewComponent()
			three := NewComponent()
			four := NewComponent()
			five := NewComponent()

			one.AddChild(two)
			two.AddChild(three)
			three.AddChild(four)
			four.AddChild(five)

			if !five.IsContainedBy(one) {
				t.Error("Five should be contained by one")
			}
			if !five.IsContainedBy(two) {
				t.Error("Five should be contained by two")
			}
			if !five.IsContainedBy(three) {
				t.Error("Five should be contained by three")
			}
			if !five.IsContainedBy(four) {
				t.Error("Five should be contained by four")
			}
		})

		t.Run("Prunes nested invalidations", func(t *testing.T) {
			var one, two, three Displayable
			root, _ := Box(NewBuilder(), ID("root"), Children(func(b Builder) {
				one, _ = Box(b, ID("one"), Children(func() {
					two, _ = Box(b, ID("two"), Children(func() {
						three, _ = Box(b, ID("three"))
					}))
				}))
			}))

			three.Invalidate()
			two.Invalidate()
			one.Invalidate()

			invalidNodes := root.InvalidNodes()
			assert.Equal(t, len(invalidNodes), 1)
			assert.Equal(t, invalidNodes[0].ID(), "one")
		})

		t.Run("RemoveAllChildren", func(t *testing.T) {
			var one, two, three Displayable
			root, _ := Box(NewBuilder(), Children(func(b Builder) {
				one, _ = Box(b)
				two, _ = Box(b)
				three, _ = Box(b)
			}))

			assert.Equal(t, root.ChildCount(), 3)
			root.RemoveAllChildren()
			assert.Equal(t, root.ChildCount(), 0)
			assert.Nil(t, one.Parent())
			assert.Nil(t, two.Parent())
			assert.Nil(t, three.Parent())
		})

		t.Run("Invalidated siblings are sorted fifo", func(t *testing.T) {
			var one, two, three Displayable
			root, _ := Box(NewBuilder(), ID("root"), Children(func(b Builder) {
				one, _ = Box(b, ID("one"), Children(func() {
					three, _ = Box(b, ID("three"))
				}))
				two, _ = Box(b, ID("two"))
			}))

			three.Invalidate()
			two.Invalidate()
			one.Invalidate()

			nodes := root.InvalidNodes()
			assert.Equal(t, len(nodes), 2, "Expected two")
			assert.Equal(t, nodes[0].ID(), "two")
			assert.Equal(t, nodes[1].ID(), "one")
		})

		t.Run("GetComponentByID", func(t *testing.T) {
			var aye, bee, cee, dee, eee Displayable

			var setUp = func() {
				aye, _ = Box(NewBuilder(), ID("aye"), Children(func(b Builder) {
					bee, _ = Box(b, ID("bee"), Children(func() {
						dee, _ = Box(b, ID("dee"))
						eee, _ = Box(b, ID("eee"))
					}))
					cee, _ = Box(b, ID("cee"))
				}))
			}

			t.Run("Matching returned", func(t *testing.T) {
				setUp()
				result := aye.FindComponentByID("aye")
				assert.NotNil(t, result)
				assert.Equal(t, result.ID(), "aye")
			})

			t.Run("First child returned", func(t *testing.T) {
				setUp()
				result := aye.FindComponentByID("bee")
				assert.NotNil(t, result)
				assert.Equal(t, result.ID(), "bee")
			})

			t.Run("Deep child returned", func(t *testing.T) {
				setUp()
				result := aye.FindComponentByID("eee")
				assert.NotNil(t, result)
				assert.Equal(t, result.ID(), "eee")
			})
		})
	})

	t.Run("Render Node", func(t *testing.T) {
		textValue := "abcd"

		var one, two, three Displayable
		var rootClosureCallCount, oneClosureCallCount int

		root, _ := Box(NewBuilder(), ID("root"), Children(func(b Builder) {
			rootClosureCallCount++
			one, _ = Box(b, ID("one"), Children(func(b Builder) {
				oneClosureCallCount++
				two, _ = Box(b, ID("two"), Text(textValue))
				three, _ = Box(b, ID("three"), Text("wxyz"))
			}))
		}))
		assert.Equal(t, rootClosureCallCount, 1)
		assert.Equal(t, oneClosureCallCount, 1)
		assert.NotNil(t, root)
		assert.Equal(t, two.Text(), "abcd")
		assert.Equal(t, three.Text(), "wxyz")

		firstInstanceOfTwo := two
		// Update a derived value
		textValue = "efgh"
		// Invalidate a nested child
		one.Invalidate()
		// Run validation from Root
		dirtyNodes := root.Validate()

		if firstInstanceOfTwo == two {
			t.Error("Expected the inner component to be re-instantiated")
		}

		assert.Equal(t, len(dirtyNodes), 1)
		assert.Equal(t, rootClosureCallCount, 1, "Root closure should NOT have been called again")
		assert.Equal(t, oneClosureCallCount, 2, "inner closure should have run twice")
		assert.Equal(t, one.ChildCount(), 2, "Children are rebuilt")
		assert.Equal(t, two.Text(), "efgh")
		assert.Equal(t, three.Text(), "wxyz")
	})
}
