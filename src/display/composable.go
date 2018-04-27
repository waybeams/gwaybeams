package display

import (
	"errors"
	"fmt"
)

// DisplayableFilter is a function that accepts a Displayable and returns
// true if it should be included.
type DisplayableFilter = func(Displayable) bool

// Composable is a set of methods that are used for composition and tree
// traversal.
type Composable interface {
	AddChild(child Displayable) int
	Builder() Builder
	ChildAt(index int) Displayable
	ChildCount() int
	Children() []Displayable
	Composer(composeFunc interface{}) error
	FindComponentByID(id string) Displayable
	FirstChild() Displayable
	GetComposeEmpty() func()
	GetComposeWithBuilder() func(Builder)
	GetComposeWithBuilderAndComponent() func(Builder, Displayable)
	GetComposeWithComponent() func(Displayable)
	GetFilteredChildren(DisplayableFilter) []Displayable
	// ID should be a tree-unique identifier and should not change
	// for a given component reference at any point in time.
	// Uniqueness constraints are not enforced at this time, but if duplicate
	// IDs are used, selectors, rendering and other features might fail in
	// unanticipated ways. We will generally default to using the first found
	// match.
	ID() string
	IsContainedBy(d Displayable) bool
	// Key should be unique for all children of a given parent and will be used
	// to determine whether an instance created by re-running the compose
	// function should replace or reuse an existing child.
	Key() string
	LastChild() Displayable
	Parent() Displayable
	// Path returns a slash-delimited path string that is the canonical
	// location for the given component.
	Path() string
	QuerySelector(selector string) Displayable
	QuerySelectorAll(selector string) []Displayable
	RecomposeChildren() []Displayable
	RemoveAllChildren()
	RemoveChild(child Displayable) int
	Root() Displayable
	SetBuilder(b Builder)
	SetParent(parent Displayable)
	SetTraitNames(name ...string)
	SetTypeName(name string)
	ShouldRecompose() bool
	TraitNames() []string
	TypeName() string
}

func (c *Component) AddChild(child Displayable) int {
	c.children = append(c.Children(), child)
	child.SetParent(c)
	return len(c.children)
}

func (c *Component) Builder() Builder {
	if c.parent != nil {
		return c.parent.Builder()
	}
	return c.builder
}

func (c *Component) ChildAt(index int) Displayable {
	return c.Children()[index]
}

func (c *Component) ChildCount() int {
	return len(c.Children())
}

func (c *Component) Children() []Displayable {
	if c.children == nil {
		c.children = make([]Displayable, 0)
	}
	return c.children
}

func (c *Component) Composer(composer interface{}) error {
	// Clear all/any existing Compose functions
	c.composeEmpty = nil
	c.composeWithBuilder = nil
	c.composeWithComponent = nil
	c.composeWithBuilderAndComponent = nil

	switch composer.(type) {
	case func():
		c.composeEmpty = composer.(func())
	case func(Builder):
		c.composeWithBuilder = composer.(func(Builder))
	case func(Displayable):
		c.composeWithComponent = composer.(func(Displayable))
	case func(Builder, Displayable):
		c.composeWithBuilderAndComponent = composer.(func(Builder, Displayable))
	case nil:
		c.composeEmpty = nil
		c.composeWithBuilder = nil
		c.composeWithComponent = nil
		c.composeWithBuilderAndComponent = nil
	default:
		return errors.New("Component.Composer() called with unexpected signature")
	}
	return nil
}

func (c *Component) FindComponentByID(id string) Displayable {
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

func (c *Component) FirstChild() Displayable {
	return c.ChildAt(0)
}

func (c *Component) GetComposeEmpty() func() {
	return c.composeEmpty
}

func (c *Component) GetComposeWithBuilder() func(Builder) {
	return c.composeWithBuilder
}

func (c *Component) GetComposeWithComponent() func(Displayable) {
	return c.composeWithComponent
}

func (c *Component) GetComposeWithBuilderAndComponent() func(Builder, Displayable) {
	return c.composeWithBuilderAndComponent
}

func (c *Component) GetFilteredChildren(filter DisplayableFilter) []Displayable {
	result := []Displayable{}
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

func (c *Component) IsContainedBy(node Displayable) bool {
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

func (c *Component) LastChild() Displayable {
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
func (c *Component) QuerySelector(selector string) Displayable {
	var result Displayable
	PreOrderVisit(c, func(d Displayable) bool {
		if QuerySelectorMatches(selector, d) {
			result = d
			return true
		}
		return false
	})

	return result
}

// QuerySelectorAll scans the tree from the current node forward and returns
// all of the nodes that match the provided selector.
func (c *Component) QuerySelectorAll(selector string) []Displayable {
	var result = []Displayable{}
	PreOrderVisit(c, func(d Displayable) bool {
		if QuerySelectorMatches(selector, d) {
			result = append(result, d)
		}
		return false
	})
	return result
}

func (c *Component) Parent() Displayable {
	return c.parent
}

func (c *Component) RemoveAllChildren() {
	children := c.Children()
	c.children = make([]Displayable, 0)
	for _, child := range children {
		c.onChildRemoved(child)
	}
}

func (c *Component) RemoveChild(toRemove Displayable) int {
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

func (c *Component) onChildRemoved(child Displayable) {
	child.SetParent(nil)
}

func (c *Component) RecomposeChildren() []Displayable {
	// TODO(lbayes): Rename to something less confusing
	nodes := c.InvalidNodes()
	b := c.Builder()
	for _, node := range nodes {
		err := b.Update(node)
		if err != nil {
			panic(err)
		}
	}
	// TODO(lbayes): Ensure this happens, even if there's a panic before!
	c.dirtyNodes = []Displayable{}
	return nodes
}

// Root returns a outermost Displayable in the current tree.
func (c *Component) Root() Displayable {
	parent := c.Parent()
	if parent != nil {
		return parent.Root()
	}
	return c
}

func (c *Component) SetParent(parent Displayable) {
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

func (c *Component) SetBuilder(b Builder) {
	// NOTE(lbayes): This method is called on temporary components
	// that are never going to be added to a valid tree. Be sure
	// we do not mutate the state of the Builder for any reason
	// from this call.
	c.builder = b
}

func (c *Component) TraitNames() []string {
	return c.Model().TraitNames
}

func (c *Component) TypeName() string {
	return c.Model().TypeName
}
