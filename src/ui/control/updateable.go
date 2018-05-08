package control

import "ui"

// TODO(lbayes): These should not clutter up the public interface of controls!
// Invalid children should be managed outside of the control tree.

func (c *Control) SetUpdateableChildren(types ui.ChildrenTypeMap) {
	c.updateableChildrenMap = types
}

func (c *Control) UpdateableChildren() ui.ChildrenTypeMap {
	return c.updateableChildrenMap
}
