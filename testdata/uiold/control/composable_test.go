package control

import (
	"github.com/waybeams/assert"
	"strings"
	"testing"
	. "ui"
	"uiold/context"
	"ui/control"
	. "ui/controls"
	. "uiold/opts"
)

func TestComposable(t *testing.T) {
	t.Run("ID can be empty", func(t *testing.T) {
		root := control.New()
		assert.Equal(root.ID(), "")
	})

	t.Run("Builder() assigns Context", func(t *testing.T) {
		box := Box(context.New())
		assert.NotNil(box.Context())
	})

	t.Run("Applied Key", func(t *testing.T) {
		root := Box(context.New(), Key("abcd"))
		assert.Equal(root.Key(), "abcd")
	})

	t.Run("Key can be empty", func(t *testing.T) {
		root := control.New()
		assert.Equal(root.Key(), "")
	})

	t.Run("Empty key will defer to ID if present", func(t *testing.T) {
		root := Box(context.New(), ID("abcd"))
		assert.Equal(root.Key(), "abcd")
	})

	t.Run("Provided ID", func(t *testing.T) {
		root := Box(context.New(), ID("root"))
		assert.Equal(root.ID(), "root")
	})

	t.Run("AddChild", func(t *testing.T) {
		root := control.New()
		one := control.New()
		two := control.New()
		root.SetWidth(200)
		assert.Equal(root.AddChild(one), 1)
		assert.Equal(root.AddChild(two), 2)

		assert.Equal(one.Parent().ID(), root.ID())
		assert.Equal(two.Parent().ID(), root.ID())

		if root.Parent() != nil {
			t.Error("Expected root.Parent() to be nil")
		}
	})

	t.Run("ChildCount", func(t *testing.T) {
		var one, two, three Displayable
		root := Box(context.New(), Children(func(c Context) {
			one = Box(c, Children(func() {
				two = Box(c)
				three = Box(c)
			}))
		}))

		assert.Equal(root.ChildCount(), 1)
		assert.Equal(root.ChildAt(0), one)

		assert.Equal(one.ChildCount(), 2)
		assert.Equal(one.ChildAt(0), two)
		assert.Equal(one.ChildAt(1), three)
	})

	t.Run("GetFilteredChildren", func(t *testing.T) {
		createTree := func() (Displayable, []Displayable) {
			var root, one, two, three, four Displayable
			root = Box(context.New(), Children(func(c Context) {
				one = Box(c, ID("a-t-one"))
				two = Box(c, ID("a-t-two"))
				three = Box(c, ID("b-t-three"))
				four = Box(c, ID("b-t-four"))
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
			root := control.New()
			filtered := root.GetFilteredChildren(allKids)
			assert.Equal(len(filtered), 0)
		})

		t.Run("returns all matched children in simple match", func(t *testing.T) {
			root, _ := createTree()
			filtered := root.GetFilteredChildren(allKids)
			assert.Equal(len(filtered), 4)
		})

		t.Run("returns all matched children in harder match", func(t *testing.T) {
			root, _ := createTree()
			filtered := root.GetFilteredChildren(bKids)
			assert.Equal(len(filtered), 2)
			assert.Equal(filtered[0].ID(), "b-t-three")
			assert.Equal(filtered[1].ID(), "b-t-four")
		})
	})

	t.Run("GetChildren returns empty list", func(t *testing.T) {
		root := control.New()
		children := root.Children()

		if children == nil {
			t.Error("GetChildren should not return nil")
		}

		assert.Equal(len(children), 0)
	})

	t.Run("GetChildren returns new list", func(t *testing.T) {
		root := Box(context.New(), Children(func(c Context) {
			Box(c)
			Box(c)
			Box(c)
		}))

		children := root.Children()
		assert.Equal(len(children), 3)
	})

	t.Run("Empty", func(t *testing.T) {
		one := control.New()
		two := control.New()
		if one.IsContainedBy(two) {
			t.Error("Unrelated nodes are not ancestors")
		}
	})

	t.Run("False for same control", func(t *testing.T) {
		one := control.New()
		if one.IsContainedBy(one) {
			t.Error("A control should not be contained by itself")
		}
	})

	t.Run("Child is true", func(t *testing.T) {
		one := control.New()
		two := control.New()
		one.AddChild(two)
		if !two.IsContainedBy(one) {
			t.Error("One should be an ancestor of two")
		}
		if one.IsContainedBy(two) {
			t.Error("Two is not an ancestor of one")
		}
	})

	t.Run("Deep descendants too", func(t *testing.T) {
		one := control.New()
		two := control.New()
		three := control.New()
		four := control.New()
		five := control.New()

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
		root := Box(context.New(), ID("root"), Children(func(c Context) {
			one = Box(c, ID("one"), Children(func() {
				two = Box(c, ID("two"), Children(func() {
					three = Box(c, ID("three"))
				}))
			}))
		}))

		three.InvalidateChildren()
		two.InvalidateChildren()
		one.InvalidateChildren()

		invalidNodes := root.InvalidNodes()
		assert.Equal(len(invalidNodes), 1)
		assert.Equal(invalidNodes[0].ID(), "one")
	})

	t.Run("InvalidateChildrenFor always goes to root", func(t *testing.T) {
		root := Box(context.New(), Children(func(c Context) {
			Box(c, Children(func() {
				Box(c, Children(func() {
					Box(c, ID("abcd"))
				}))
			}))
		}))

		child := root.FindControlById("abcd")
		child.InvalidateChildrenFor(child.Parent())
		assert.Equal(len(root.InvalidNodes()), 1)
	})

	t.Run("RemoveChild", func(t *testing.T) {
		var one, two, three Displayable
		root := Box(context.New(), Children(func(c Context) {
			one = Box(c)
			two = Box(c)
			three = Box(c)
		}))
		removedFromIndex := root.RemoveChild(two)
		assert.Equal(removedFromIndex, 1)

		removedFromIndex = root.RemoveChild(two)
		assert.Equal(removedFromIndex, -1, "Already removed, not found")
	})

	t.Run("RemoveAllChildren", func(t *testing.T) {
		var one, two, three Displayable
		root := Box(context.New(), Children(func(c Context) {
			one = Box(c)
			two = Box(c)
			three = Box(c)
		}))

		assert.Equal(root.ChildCount(), 3)
		root.RemoveAllChildren()
		assert.Equal(root.ChildCount(), 0)
		assert.Nil(one.Parent())
		assert.Nil(two.Parent())
		assert.Nil(three.Parent())
	})

	t.Run("Invalidated siblings are sorted fifo", func(t *testing.T) {
		var one, two, three Displayable
		root := Box(context.New(), ID("root"), Children(func(c Context) {
			one = Box(c, ID("one"), Children(func() {
				three = Box(c, ID("three"))
			}))
			two = Box(c, ID("two"))
		}))

		three.InvalidateChildren()
		two.InvalidateChildren()
		one.InvalidateChildren()

		nodes := root.InvalidNodes()
		assert.Equal(len(nodes), 2, "Expected two")
		assert.Equal(nodes[0].ID(), "two")
		assert.Equal(nodes[1].ID(), "one")
	})

	t.Run("GetControlByID", func(t *testing.T) {
		var aye, bee, cee, dee, eee Displayable

		var setUp = func() {
			aye = Box(context.New(), ID("aye"), Children(func(c Context) {
				bee = Box(c, ID("bee"), Children(func() {
					dee = Box(c, ID("dee"))
					eee = Box(c, ID("eee"))
				}))
				cee = Box(c, ID("cee"))
			}))
		}

		t.Run("Matching returned", func(t *testing.T) {
			setUp()
			result := aye.FindControlById("aye")
			assert.NotNil(result)
			assert.Equal(result.ID(), "aye")
		})

		t.Run("First child returned", func(t *testing.T) {
			setUp()
			result := aye.FindControlById("bee")
			assert.NotNil(result)
			assert.Equal(result.ID(), "bee")
		})

		t.Run("Deep child returned", func(t *testing.T) {
			setUp()
			result := aye.FindControlById("eee")
			assert.NotNil(result)
			assert.Equal(result.ID(), "eee")
		})
	})

	t.Run("SelectControls", func(t *testing.T) {
		t.Run("By Type", func(t *testing.T) {
			root := Box(context.New(), Children(func(c Context) {
				HBox(c)
			}))

			assert.NotNil(root.QuerySelector("HBox"))
		})

		t.Run("By TraitName", func(t *testing.T) {
			root := Box(context.New(), Children(func(c Context) {
				Box(c, TraitNames("abcd"))
				Box(c, TraitNames("efgh"))
			}))

			assert.NotNil(root.QuerySelector(".efgh"))
		})
	})

	t.Run("Root returns deeply nested root control", func(t *testing.T) {
		var descendant Displayable
		root := Box(context.New(), ID("root"), Children(func(c Context) {
			Box(c, ID("one"), Children((func() {
				Box(c, ID("two"), Children(func() {
					Box(c, ID("three"), Children(func() {
						Box(c, ID("four"), Children(func() {
							Box(c, ID("five"), Children(func() {
								descendant = Box(c, ID("child"))
							}))
						}))
					}))
				}))
			})))
		}))
		assert.Equal(root.ID(), descendant.Root().ID())
	})

	t.Run("Root gets Builder reference", func(t *testing.T) {
		var root, child Displayable

		root = Box(context.New(), Children(func(c Context) {
			Box(c, Children(func() {
				child = Box(c)
			}))
		}))

		assert.NotNil(root.Context())
		assert.NotNil(child.Context())
	})

	t.Run("Path", func(t *testing.T) {
		t.Run("root", func(t *testing.T) {
			root := Box(context.New(), ID("root"))
			assert.Equal(root.Path(), "/root")
		})

		t.Run("uses Key if ID is empty", func(t *testing.T) {
			root := Box(context.New(), Key("abcd"))
			assert.Equal(root.Path(), "/abcd")
		})

		t.Run("uses type if neither Key nor Id are present", func(t *testing.T) {
			root := Box(context.New())
			assert.Equal(root.Path(), "/Box")
		})

		t.Run("defaults to TypeName and parent index", func(t *testing.T) {
			root := VBox(context.New(), Children(func(c Context) {
				Box(c)
				Box(c)
				HBox(c)
			}))

			kids := root.Children()
			assert.Equal(kids[0].Path(), "/VBox/Box0")
			assert.Equal(kids[1].Path(), "/VBox/Box1")
			assert.Equal(kids[2].Path(), "/VBox/HBox2")
		})

		t.Run("with depth", func(t *testing.T) {
			var one, two, three, four Displayable
			Box(context.New(), ID("root"), Children(func(c Context) {
				one = Box(c, ID("one"), Children(func() {
					two = Box(c, ID("two"), Children(func() {
						three = Box(c, ID("three"))
					}))
					four = Box(c, ID("four"))
				}))
			}))

			assert.Equal(one.Path(), "/root/one")
			assert.Equal(two.Path(), "/root/one/two")
			assert.Equal(three.Path(), "/root/one/two/three")
			assert.Equal(four.Path(), "/root/one/four")
		})
	})
}
