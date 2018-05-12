package spec

type ComposableReader interface {
	ChildAt(index int) ReadWriter
	ChildCount() int
	Children() []ReadWriter
	Key() string
	Parent() ReadWriter
	SpecName() string
}

type ComposableWriter interface {
	SetChildren(children []ReadWriter)
	SetKey(key string)
	SetParent(parent ReadWriter)
	SetSpecName(name string)
}

type ComposableReadWriter interface {
	ComposableReader
	ComposableWriter
}

// GetFilteredChildren(DisplayableFilter) []Displayable
// IsContainedBy(d Displayable) bool
// QuerySelector(selector string) Displayable
// QuerySelectorAll(selector string) []Displayable

func (c *Spec) ChildAt(index int) ReadWriter {
	return c.children[index]
}

func (c *Spec) ChildCount() int {
	return len(c.children)
}

func (c *Spec) Children() []ReadWriter {
	return c.children
}

func (c *Spec) Composer() interface{} {
	return c.composer
}

func (c *Spec) Key() string {
	return c.key
}

func (c *Spec) SetSpecName(name string) {
	c.specName = name
}

func (c *Spec) SpecName() string {
	return c.specName
}

func (c *Spec) Parent() ReadWriter {
	return c.parent
}

func (c *Spec) SetChildren(children []ReadWriter) {
	c.children = children
}

func (c *Spec) SetKey(key string) {
	c.key = key
}

func (c *Spec) SetParent(parent ReadWriter) {
	c.parent = parent
}
