package display

// The PostOrder Visitor should return the deepest, left-most node
// first, followed by all siblings and then walk up the tree (and
// back down any other paths) until the root node is returned.
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
func PostOrderVisit(node Displayable, onNode func(result Displayable)) {
	stack := NewStack()
	var visitChildren func(parent Displayable)
	var visitNode func(node Displayable)

	visitChildren = func(parent Displayable) {
		stack.Push(node)
		for i := 0; i < parent.GetChildCount(); i++ {
			visitNode(parent.GetChildAt(i))
		}
	}

	visitNode = func(node Displayable) {
		if node.GetChildCount() > 0 {
			stack.Push(node)
			visitChildren(node)
			stack.Pop()
		}
		onNode(node)
	}

	visitNode(node)
}
