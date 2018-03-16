package display

var lastId = 0

type Composable interface {
	Id() int
	Parent() Composable
	AddChild(child Composable) int
	Layout() error
	Render() error
	setParent(parent Composable)
}

// Concrete Sprite implementation
type sprite struct {
	children []Composable
	id       int
	parent   Composable
}

func (c *sprite) setParent(parent Composable) {
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

func (c *sprite) AddChild(child Composable) int {
	if c.children == nil {
		c.children = make([]Composable, 0)
	}

	c.children = append(c.children, child)
	child.setParent(c)
	return len(c.children)
}

func (c *sprite) Parent() Composable {
	return c.parent
}

func NewSprite() Composable {
	lastId++
	return &sprite{
		id: lastId,
	}
}
