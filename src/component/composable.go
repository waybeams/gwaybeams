package component

import (
	"errors"
	"fmt"
	"ui"
)

func (c *Component) AddChild(child ui.Displayable) int {
	c.children = append(c.Children(), child)
	child.SetParent(c)
	return len(c.children)
}

func (c *Component) Context() ui.Context {
	if c.parent != nil {
		return c.parent.Context()
	}
	return c.context
}

func (c *Component) ChildAt(index int) ui.Displayable {
	return c.Children()[index]
}

func (c *Component) ChildCount() int {
	return len(c.Children())
}

func (c *Component) Children() []ui.Displayable {
	if c.children == nil {
		c.children = make([]ui.Displayable, 0)
	}
	return c.children
}

// TODO(lbayes): Rename to SetComposer
func (c *Component) Composer(composer interface{}) error {
	// Clear all/any existing Compose functions
	c.composeEmpty = nil
	c.composeWithContext = nil
	c.composeWithComponent = nil
	c.composeWithContextAndComponent = nil

	switch composer.(type) {
	case func():
		c.composeEmpty = composer.(func())
	case func(ui.Displayable):
		c.composeWithComponent = composer.(func(ui.Displayable))
	case func(ui.Context):
		c.composeWithContext = composer.(func(ui.Context))
	case func(ui.Context, ui.Displayable):
		c.composeWithContextAndComponent = composer.(func(ui.Context, ui.Displayable))
	case nil:
		c.composeEmpty = nil
		c.composeWithComponent = nil
		c.composeWithContext = nil
		c.composeWithContextAndComponent = nil
	default:
		return errors.New("Component.Composer() called with unexpected signature")
	}
	return nil
}

func (c *Component) FindComponentByID(id string) ui.Displayable {
	if id == c.ID() {
		return c
	}
	for _, child := range c.Children() {
		result := child.FindComponentByID(id)
		if result != nil {
			return result
		}
	}
	return nil
}

func (c *Component) FirstChild() ui.Displayable {
	return c.ChildAt(0)
}

func (c *Component) GetComposeEmpty() func() {
	return c.composeEmpty
}

func (c *Component) GetComposeWithContext() func(ui.Context) {
	return c.composeWithContext
}

func (c *Component) GetComposeWithComponent() func(ui.Displayable) {
	return c.composeWithComponent
}

func (c *Component) GetComposeWithContextAndComponent() func(ui.Context, ui.Displayable) {
	return c.composeWithContextAndComponent
}

func (c *Component) GetFilteredChildren(filter ui.DisplayableFilter) []ui.Displayable {
	result := []ui.Displayable{}
	kids := c.Children()
	for _, child := range kids {
		if filter(child) {
			result = append(result, child)
		}
	}
	return result
}

func (c *Component) ID() string {
	return c.Model().ID
}

func (c *Component) IsContainedBy(node ui.Displayable) bool {
	current := c.Parent()
	for current != nil {
		if current == node {
			return true
		}
		current = current.Parent()
	}

	return false
}

func (c *Component) Key() string {
	key := c.Model().Key

	if key != "" {
		return key
	}

	return c.ID()
}

func (c *Component) LastChild() ui.Displayable {
	return c.ChildAt(c.ChildCount() - 1)
}

func (c *Component) Path() string {
	parent := c.Parent()
	localPath := "/" + c.pathPart()

	if parent != nil {
		return parent.Path() + localPath
	}
	return localPath
}

func (c *Component) pathPart() string {
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
func (c *Component) QuerySelector(selector string) ui.Displayable {
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
func (c *Component) QuerySelectorAll(selector string) []ui.Displayable {
	var result = []ui.Displayable{}
	ui.PreOrderVisit(c, func(d ui.Displayable) bool {
		if ui.QuerySelectorMatches(selector, d) {
			result = append(result, d)
		}
		return false
	})
	return result
}

func (c *Component) Parent() ui.Displayable {
	return c.parent
}

func (c *Component) RemoveAllChildren() {
	children := c.Children()
	c.children = make([]ui.Displayable, 0)
	for _, child := range children {
		c.onChildRemoved(child)
	}
}

func (c *Component) RemoveChild(toRemove ui.Displayable) int {
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

func (c *Component) onChildRemoved(child ui.Displayable) {
	child.SetParent(nil)
}

func (c *Component) RecomposeChildren() []ui.Displayable {
	// TODO(lbayes): Rename to something less confusing
	nodes := c.InvalidNodes()
	b := c.Context()
	for _, node := range nodes {
		err := b.Builder().Update(node)
		if err != nil {
			panic(err)
		}
	}
	// TODO(lbayes): Ensure this happens, even if there's a panic before!
	c.dirtyNodes = []ui.Displayable{}
	return nodes
}

// Root returns a outermost Displayable in the current tree.
func (c *Component) Root() ui.Displayable {
	parent := c.Parent()
	if parent != nil {
		return parent.Root()
	}
	return c
}

func (c *Component) SetParent(parent ui.Displayable) {
	c.parent = parent
}

func (c *Component) ShouldRecompose() bool {
	return len(c.dirtyNodes) > 0
}

func (c *Component) SetTraitNames(names ...string) {
	c.Model().TraitNames = names
}

func (c *Component) SetTypeName(name string) {
	c.Model().TypeName = name
}

func (c *Component) SetContext(ctx ui.Context) {
	c.context = ctx
}

func (c *Component) TraitNames() []string {
	return c.Model().TraitNames
}

func (c *Component) TypeName() string {
	return c.Model().TypeName
}
