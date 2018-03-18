package display

// Postorder visitor will call provided handler with each node in the provided
// tree, beginning with the left most deepest root and working up to the root
// node.
func PostOrderVisit(node Displayable, onNode func(result Displayable)) {
	stack := NewStack()
	var visitChildren func(parent Displayable)
	var visitNode func(node Displayable)

	visitChildren = func(parent Displayable) {
		stack.Push(node)
		for i := 0; i < parent.ChildCount(); i++ {
			visitNode(parent.ChildAt(i))
		}
	}

	visitNode = func(node Displayable) {
		if node.ChildCount() > 0 {
			stack.Push(node)
			visitChildren(node)
			stack.Pop()
		}
		onNode(node)
	}

	visitNode(node)
}
