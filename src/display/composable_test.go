package display

import (
	"assert"
	"strings"
	"testing"
)

func TestComposable(t *testing.T) {
	t.Run("ID can be empty", func(t *testing.T) {
		root := NewComponent()
		assert.Equal(t, root.ID(), "")
	})

	t.Run("Builder() returns", func(t *testing.T) {
		box, _ := Box(NewBuilder())
		if box.Builder() == nil {
			t.Error("Component factory should assign the builder")
		}
	})

	t.Run("Applied Key", func(t *testing.T) {
		root, _ := Box(NewBuilder(), Key("abcd"))
		assert.Equal(t, root.Key(), "abcd")
	})

	t.Run("Key can be empty", func(t *testing.T) {
		root := NewComponent()
		assert.Equal(t, root.Key(), "")
	})

	t.Run("Empty key will defer to ID if present", func(t *testing.T) {
		root, _ := Box(NewBuilder(), ID("abcd"))
		assert.Equal(t, root.Key(), "abcd")
	})

	t.Run("Provided ID", func(t *testing.T) {
		root, _ := Box(NewBuilder(), ID("root"))
		assert.Equal(t, root.ID(), "root")
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

		three.InvalidateChildren()
		two.InvalidateChildren()
		one.InvalidateChildren()

		invalidNodes := root.InvalidNodes()
		assert.Equal(t, len(invalidNodes), 1)
		assert.Equal(t, invalidNodes[0].ID(), "one")
	})

	t.Run("InvalidateChildrenFor always goes to root", func(t *testing.T) {
		root, _ := Box(NewBuilder(), Children(func(b Builder) {
			Box(b, Children(func() {
				Box(b, Children(func() {
					Box(b, ID("abcd"))
				}))
			}))
		}))

		child := root.FindComponentByID("abcd")
		child.InvalidateChildrenFor(child.Parent())
		assert.Equal(t, len(root.InvalidNodes()), 1)
	})

	t.Run("RemoveChild", func(t *testing.T) {
		var one, two, three Displayable
		root, _ := Box(NewBuilder(), Children(func(b Builder) {
			one, _ = Box(b)
			two, _ = Box(b)
			three, _ = Box(b)
		}))
		removedFromIndex := root.RemoveChild(two)
		assert.Equal(t, removedFromIndex, 1)

		removedFromIndex = root.RemoveChild(two)
		assert.Equal(t, removedFromIndex, -1, "Already removed, not found")
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

		three.InvalidateChildren()
		two.InvalidateChildren()
		one.InvalidateChildren()

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

	t.Run("SelectComponents", func(t *testing.T) {
		t.Run("By Type", func(t *testing.T) {
			root, _ := Box(NewBuilder(), Children(func(b Builder) {
				HBox(b)
			}))

			assert.NotNil(t, root.QuerySelector("HBox"))
		})

		t.Run("By TraitName", func(t *testing.T) {
			root, _ := Box(NewBuilder(), Children(func(b Builder) {
				Box(b, TraitNames("abcd"))
				Box(b, TraitNames("efgh"))
			}))

			assert.NotNil(t, root.QuerySelector(".efgh"))
		})
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

	t.Run("Root gets Builder reference", func(t *testing.T) {
		var root, child Displayable

		root, _ = Box(NewBuilder(), Children(func(b Builder) {
			Box(b, Children(func() {
				child, _ = Box(b)
			}))
		}))

		assert.NotNil(t, root.Builder())
		assert.NotNil(t, child.Builder())
	})

	t.Run("Path", func(t *testing.T) {
		t.Run("root", func(t *testing.T) {
			root, _ := Box(NewBuilder(), ID("root"))
			assert.Equal(t, root.Path(), "/root")
		})

		t.Run("uses Key if ID is empty", func(t *testing.T) {
			root, _ := Box(NewBuilder(), Key("abcd"))
			assert.Equal(t, root.Path(), "/abcd")
		})

		t.Run("uses type if neither Key nor Id are present", func(t *testing.T) {
			root, _ := Box(NewBuilder())
			assert.Equal(t, root.Path(), "/Box")
		})

		t.Run("defaults to TypeName and parent index", func(t *testing.T) {
			root, _ := VBox(NewBuilder(), Children(func(b Builder) {
				Box(b)
				Box(b)
				HBox(b)
			}))

			kids := root.Children()
			assert.Equal(t, kids[0].Path(), "/VBox/Box0")
			assert.Equal(t, kids[1].Path(), "/VBox/Box1")
			assert.Equal(t, kids[2].Path(), "/VBox/HBox2")
		})

		t.Run("with depth", func(t *testing.T) {
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
	})
}
