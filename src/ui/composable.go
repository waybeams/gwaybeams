package ui

// DisplayableFilter is a function that accepts a Displayable and returns
// true if it should be included.
type DisplayableFilter = func(Displayable) bool

// Composable is a set of methods that are used for composition and tree
// traversal.
type Composable interface {
	AddChild(child Displayable) int
	Context() Context
	ChildAt(index int) Displayable
	ChildCount() int
	Children() []Displayable
	Composer(composeFunc interface{}) error
	FindComponentByID(id string) Displayable
	FirstChild() Displayable
	GetComposeEmpty() func()
	GetComposeWithContext() func(Context)
	GetComposeWithContextAndComponent() func(Context, Displayable)
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
	SetContext(c Context)
	SetParent(parent Displayable)
	SetTraitNames(name ...string)
	SetTypeName(name string)
	ShouldRecompose() bool
	TraitNames() []string
	TypeName() string
}
