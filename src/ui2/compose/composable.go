package compose

type Composite struct {
	Children []*Composite
	Parent   *Composite
}

func New() *Composite {
	return &Composite{}
}

func AddChild(parent *Composite, child *Composite) int {
	if parent == child {
		panic("Cannot add Composite to self")
	}
	parent.Children = append(parent.Children, child)
	child.Parent = parent
	return -1
}

func ChildAt(parent *Composite, index int) *Composite {
	return parent.Children[index]
}

func Root(node *Composite) *Composite {
	parent := node.Parent
	if parent != nil {
		return Root(parent)
	}
	return node
}
