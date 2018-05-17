package ui_test

import (
	"github.com/waybeams/assert"
	"testing"
	. "ui"
	"ui/control"
)

func createTree() (root, one, two, three, four, five Displayable) {
	root = control.New()
	one = control.New()
	two = control.New()
	three = control.New()
	four = control.New()
	five = control.New()

	two.Model().ID = "two"

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
	return root, one, two, three, four, five
}

func TestVisitor(t *testing.T) {
	t.Run("Empty node", func(t *testing.T) {
		root := control.New()
		wasCalled := false
		PostOrderVisit(root, func(node Displayable) bool {
			wasCalled = true
			return false
		})
		assert.True(t, wasCalled)
	})

	t.Run("Default state", func(t *testing.T) {
		root, one, two, three, four, five := createTree()
		visited := []Displayable{}

		PostOrderVisit(root, func(node Displayable) bool {
			visited = append(visited, node)
			return false
		})

		assert.Equal(t, len(visited), 6)
		assert.Equal(t, visited[0], four)
		assert.Equal(t, visited[1], five)
		assert.Equal(t, visited[2], two)
		assert.Equal(t, visited[3], three)
		assert.Equal(t, visited[4], one)
		assert.Equal(t, visited[5], root)
	})

	t.Run("PostOrder bails", func(t *testing.T) {
		root, _, _, _, _, _ := createTree()
		var result Displayable
		var visited = []Displayable{}

		PostOrderVisit(root, func(node Displayable) bool {
			visited = append(visited, node)
			if node.ID() == "two" {
				result = node
				return true
			}
			return false
		})
		assert.Equal(t, result.ID(), "two")
		assert.Equal(t, len(visited), 3)
	})

	t.Run("PreOrderVisit", func(t *testing.T) {
		root, one, two, three, four, five := createTree()
		visited := []Displayable{}

		PreOrderVisit(root, func(node Displayable) bool {
			visited = append(visited, node)
			return false
		})

		assert.Equal(t, len(visited), 6)
		assert.Equal(t, visited[0], root)
		assert.Equal(t, visited[1], one, "one")
		assert.Equal(t, visited[2], two, "two")
		assert.Equal(t, visited[3], four, "four")
		assert.Equal(t, visited[4], five, "five")
		assert.Equal(t, visited[5], three, "three")
	})

	t.Run("PreOrder bails", func(t *testing.T) {
		root, _, _, _, _, _ := createTree()
		var result Displayable
		var visited = []Displayable{}

		PreOrderVisit(root, func(node Displayable) bool {
			visited = append(visited, node)
			if node.ID() == "two" {
				result = node
				return true
			}
			return false
		})
		assert.Equal(t, result.ID(), "two")
		assert.Equal(t, len(visited), 3)
	})
}
