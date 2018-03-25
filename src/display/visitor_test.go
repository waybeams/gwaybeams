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

		assert.Equal(t, len(visited), 6)
		assert.Equal(t, visited[0], four)
		assert.Equal(t, visited[1], five)
		assert.Equal(t, visited[2], two)
		assert.Equal(t, visited[3], three)
		assert.Equal(t, visited[4], one)
		assert.Equal(t, visited[5], root)
	})
}
