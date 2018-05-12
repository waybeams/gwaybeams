package control

import (
	"fmt"
	"ui"
)

func (c *Control) AddChild(child ui.Displayable) int {
	c.children = append(c.Children(), child)
	child.SetParent(c)
	return len(c.children)
}

func (c *Control) Context() ui.Context {
	if c.parent != nil {
		return c.parent.Context()
	}
	return c.context
}

func (c *Control) ChildAt(index int) ui.Displayable {
	return c.Children()[index]
}

func (c *Control) ChildCount() int {
	return len(c.Children())
}

func (c *Control) Children() []ui.Displayable {
	if c.children == nil {
		c.children = make([]ui.Displayable, 0)
	}
	return c.children
}

func (c *Control) SetComposer(composer interface{}) {
	// Clear all/any existing Compose functions
	c.composeEmpty = nil
	c.composeWithContext = nil
	c.composeWithControl = nil
	c.composeWithContextAndControl = nil

	switch composer.(type) {
	case func():
		c.composeEmpty = composer.(func())
	case func(ui.Displayable):
		c.composeWithControl = composer.(func(ui.Displayable))
	case func(ui.Context):
		c.composeWithContext = composer.(func(ui.Context))
	case func(ui.Context, ui.Displayable):
		c.composeWithContextAndControl = composer.(func(ui.Context, ui.Displayable))
	case nil:
		c.composeEmpty = nil
		c.composeWithControl = nil
		c.composeWithContext = nil
		c.composeWithContextAndControl = nil
	default:
		panic("Spec.Composer() called with unexpected signature")
	}
}

func (c *Control) FindControlById(id string) ui.Displayable {
	if id == c.ID() {
		return c
	}
	for _, child := range c.Children() {
		result := child.FindControlById(id)
		if result != nil {
			return result
		}
	}
	return nil
}

func (c *Control) FirstChild() ui.Displayable {
	return c.ChildAt(0)
}

func (c *Control) GetComposeEmpty() func() {
	return c.composeEmpty
}

func (c *Control) GetComposeWithContext() func(ui.Context) {
	return c.composeWithContext
}

func (c *Control) GetComposeWithControl() func(ui.Displayable) {
	return c.composeWithControl
}

func (c *Control) GetComposeWithContextAndControl() func(ui.Context, ui.Displayable) {
	return c.composeWithContextAndControl
}

func (c *Control) GetFilteredChildren(filter ui.DisplayableFilter) []ui.Displayable {
	result := []ui.Displayable{}
	kids := c.Children()
	for _, child := range kids {
		if filter(child) {
			result = append(result, child)
		}
	}
	return result
}

func (c *Control) ID() string {
	return c.Model().ID
}

func (c *Control) IsContainedBy(node ui.Displayable) bool {
	current := c.Parent()
	for current != nil {
		if current == node {
			return true
		}
		current = current.Parent()
	}

	return false
}

func (c *Control) Key() string {
	key := c.Model().Key

	if key != "" {
		return key
	}

	return c.ID()
}

func (c *Control) LastChild() ui.Displayable {
	return c.ChildAt(c.ChildCount() - 1)
}

func (c *Control) Path() string {
	parent := c.Parent()
	localPath := "/" + c.pathPart()

	if parent != nil {
		return parent.Path() + localPath
	}
	return localPath
}

func (c *Control) pathPart() string {
	// Try ID first
	id := c.ID()
	if id != "" {
		return c.ID()
	}

	// Empty ID, try Key
	key := c.Key()
	if key != "" {
		return c.Key()
	}

	parent := c.Parent()
	if parent != nil {
		siblings := parent.Children()
		for index, child := range siblings {
			if child == c {
				return fmt.Sprintf("%v%v", c.TypeName(), index)
			}
		}
	}

	// Empty ID and Key, and Parent just use TypeName
	return c.TypeName()
}

// QuerySelector scans the tree from the current node forward and returns
// the first node that matches the provided selector.
func (c *Control) QuerySelector(selector string) ui.Displayable {
	var result ui.Displayable
	ui.PreOrderVisit(c, func(d ui.Displayable) bool {
		if ui.QuerySelectorMatches(selector, d) {
			result = d
			return true
		}
		return false
	})

	return result
}

// QuerySelectorAll scans the tree from the current node forward and returns
// all of the nodes that match the provided selector.
func (c *Control) QuerySelectorAll(selector string) []ui.Displayable {
	var result = []ui.Displayable{}
	ui.PreOrderVisit(c, func(d ui.Displayable) bool {
		if ui.QuerySelectorMatches(selector, d) {
			result = append(result, d)
		}
		return false
	})
	return result
}

func (c *Control) Parent() ui.Displayable {
	return c.parent
}

func (c *Control) RemoveAllChildren() {
	children := c.Children()
	c.children = make([]ui.Displayable, 0)
	for _, child := range children {
		c.onChildRemoved(child)
	}
}

func (c *Control) RemoveChild(toRemove ui.Displayable) int {
	children := c.Children()
	for index, child := range children {
		if child == toRemove {
			c.children = append(children[:index], children[index+1:]...)
			c.onChildRemoved(child)
			return index
		}
	}

	return -1
}

func (c *Control) onChildRemoved(child ui.Displayable) {
	child.SetParent(nil)
}

func (c *Control) RecomposeChildren() []ui.Displayable {
	// TODO(lbayes): Rename to something less confusing
	nodes := c.InvalidNodes()
	b := c.Context()
	for _, node := range nodes {
		b.Builder().Update(node)
	}
	c.dirtyNodes = []ui.Displayable{}
	return nodes
}

// Root returns a outermost Displayable in the current tree.
func (c *Control) Root() ui.Displayable {
	parent := c.Parent()
	if parent != nil {
		return parent.Root()
	}
	return c
}

func (c *Control) SetParent(parent ui.Displayable) {
	c.parent = parent
}

func (c *Control) ShouldRecompose() bool {
	return len(c.dirtyNodes) > 0
}

func (c *Control) SetTraitNames(names ...string) {
	c.Model().TraitNames = names
}

func (c *Control) SetTypeName(name string) {
	c.Model().TypeName = name
}

func (c *Control) SetContext(context ui.Context) {
	c.context = context
}

func (c *Control) TraitNames() []string {
	return c.Model().TraitNames
}

func (c *Control) TypeName() string {
	return c.Model().TypeName
}
