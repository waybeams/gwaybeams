package display

import (
	"assert"
	"testing"
)

func TestVisitor(t *testing.T) {
	t.Run("Empty node", func(t *testing.T) {
		root := NewComponent()
		PostOrderVisit(root, func(node Displayable) {})
	})

	t.Run("Default state", func(t *testing.T) {
		root := NewComponent()
		one := NewComponent()
		two := NewComponent()
		three := NewComponent()
		four := NewComponent()
		five := NewComponent()

		root.AddChild(one)
		one.AddChild(two)
		one.AddChild(three)
		two.AddChild(four)
		two.AddChild(five)

		// Creating a structure as follows:
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
		// Expect an order like: [four, five, two, three, one, root]

		visited := []Displayable{}

		PostOrderVisit(root, func(node Displayable) {
			visited = append(visited, node)
		})

		assert.TEqual(t, len(visited), 6)
		assert.TEqual(t, visited[0], four)
		assert.TEqual(t, visited[1], five)
		assert.TEqual(t, visited[2], two)
		assert.TEqual(t, visited[3], three)
		assert.TEqual(t, visited[4], one)
		assert.TEqual(t, visited[5], root)
	})
}
