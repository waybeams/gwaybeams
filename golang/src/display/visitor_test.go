package display

import (
	"assert"
	"testing"
)

func TestVisitor(t *testing.T) {
	t.Run("Empty node", func(t *testing.T) {
		root := NewSprite()
		PostOrderVisit(root, func(node Displayable) {})
	})

	t.Run("Default state", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		three := NewSprite()
		four := NewSprite()
		five := NewSprite()

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

		assert.Equal(len(visited), 6)
		assert.Equal(visited[0], four)
		assert.Equal(visited[1], five)
		assert.Equal(visited[2], two)
		assert.Equal(visited[3], three)
		assert.Equal(visited[4], one)
		assert.Equal(visited[5], root)
	})
}
