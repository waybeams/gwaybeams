package display

import (
	"log"
	"testing"
)

func nodeCount(d Displayable) int {
	count := 0
	PostOrderVisit(d, func(node Displayable) {
		count++
	})

	return count
}

func BenchmarkComponent(b *testing.B) {
	var createTree = func(b Builder) (Displayable, error) {
		return VBox(b, Children(func() {
			for i := 0; i < 100; i++ {
				HBox(b, Children(func() {
					Box(b, Children(func() {
						Box(b, Children(func() {
							Box(b, Children(func() {
								Box(b)
								Box(b)
								Box(b)
								Box(b)
								Box(b)
								Box(b)
							}))
						}))
					}))
				}))
			}
		}))
	}

	b.Run("Basic instantiation", func(b *testing.B) {
		tree, _ := createTree(NewBuilder())
		count := nodeCount(tree)
		log.Printf("BENCHMARK WITH %v NODES", count)

		for i := 0; i < b.N; i++ {
			builder := NewBuilder()
			surface := NewFakeSurface()
			tree, _ := createTree(builder)
			tree.Layout()
			tree.Draw(surface)
		}
	})
}
