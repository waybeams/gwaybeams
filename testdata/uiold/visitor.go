package ui

type VisitorHandler func(result Displayable) bool

// PostOrderVisit should return the deepest, left-most node first, followed
// by all siblings and then walk up the tree (and back down any other paths)
// until the root node is returned.
//
// For this example tree, this means an order of:
// [four, five, two, three, one, root]
//
//           [root]
//             |
//             |
//           [one]
//            / \
//           /   \
//        [two] [three]
//         / \
//        /   \
//    [four] [five]
//
func PostOrderVisit(node Displayable, onNode VisitorHandler) {
	stack := NewStack()
	var visitChildren func(parent Displayable) bool
	var visitNode func(node Displayable) bool

	visitChildren = func(parent Displayable) bool {
		stack.Push(node)
		for i := 0; i < parent.ChildCount(); i++ {
			if visitNode(parent.ChildAt(i)) {
				return true
			}
		}
		return false
	}

	visitNode = func(node Displayable) bool {
		if node.ChildCount() > 0 {
			stack.Push(node)
			if visitChildren(node) {
				return true
			}
			stack.Pop()
		}
		return onNode(node)
	}

	visitNode(node)
}

// PreOrderVisit should call onNode as it passes through each node in the
// provided tree.
func PreOrderVisit(node Displayable, onNode VisitorHandler) {
	stack := NewStack()
	var visitChildren func(parent Displayable) bool
	var visitNode func(node Displayable) bool

	visitChildren = func(parent Displayable) bool {
		stack.Push(node)
		for i := 0; i < parent.ChildCount(); i++ {
			if visitNode(parent.ChildAt(i)) {
				return true
			}
		}
		return false
	}

	visitNode = func(node Displayable) bool {
		if onNode(node) {
			return true
		}
		if node.ChildCount() > 0 {
			stack.Push(node)
			if visitChildren(node) {
				return true
			}
			stack.Pop()
		}
		return false
	}

	visitNode(node)
}
