package display

var lastId = 0

type Displayable interface {
	Id() int
	Parent() Displayable
	AddChild(child Displayable) int
	Layout() error
	Render() error
	setParent(parent Displayable)
}

// Concrete Sprite implementation
type sprite struct {
	children []Displayable
	id       int
	parent   Displayable
}

func (c *sprite) setParent(parent Displayable) {
	c.parent = parent
}

func (c *sprite) Id() int {
	return c.id
}

func (c *sprite) Render() error {
	return nil
}

func (c *sprite) Layout() error {
	return nil
}

func (c *sprite) AddChild(child Displayable) int {
	if c.children == nil {
		c.children = make([]Displayable, 0)
	}

	c.children = append(c.children, child)
	child.setParent(c)
	return len(c.children)
}

func (c *sprite) Parent() Displayable {
	return c.parent
}

func NewSprite() Displayable {
	lastId++
	return &sprite{
		id: lastId,
	}
}
