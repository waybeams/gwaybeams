package component

import "ui"

// TODO(lbayes): These should not clutter up the public interface of Components!
// Invalid children should be managed outside of the component tree.

func (c *Component) SetUpdateableChildren(types ui.ChildrenTypeMap) {
	c.updateableChildrenMap = types
}

func (c *Component) UpdateableChildren() ui.ChildrenTypeMap {
	return c.updateableChildrenMap
}
