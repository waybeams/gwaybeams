package compose

type Composite struct {
	children []*Composite
	parent   *Composite
}

func New() *Composite {
	return &Composite{}
}

func AddChild(parent *Composite, child *Composite) int {
	if parent == child {
		panic("Cannot add Composite to self")
	}
	parent.children = append(parent.children, child)
	child.parent = parent
	return -1
}

func ChildAt(parent *Composite, index int) *Composite {
	return parent.children[index]
}

func ChildCount(node *Composite) int {
	return len(node.children)
}

func Children(node *Composite) []*Composite {
	return node.children
}

func Parent(node *Composite) *Composite {
	return node.parent
}

func Root(node *Composite) *Composite {
	parent := node.parent
	if parent != nil {
		return Root(parent)
	}
	return node
}
