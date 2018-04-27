package display

type ChildrenTypeMap map[string][]Displayable

type Updateable interface {
	setUpdateableChildren(types ChildrenTypeMap)
	updateableChildren() ChildrenTypeMap
}

func (c *Component) setUpdateableChildren(types ChildrenTypeMap) {
	c.updateableChildrenMap = types
}

func (c *Component) updateableChildren() ChildrenTypeMap {
	return c.updateableChildrenMap
}
