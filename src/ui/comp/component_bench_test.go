package comp_test

import (
	"log"
	"surface"
	"testing"
	. "ui"
	"ui/context"
	. "ui/controls"
	. "ui/opts"
)

func nodeCount(d Displayable) int {
	count := 0
	PostOrderVisit(d, func(node Displayable) bool {
		count++
		return false
	})

	return count
}

func BenchmarkComponent(b *testing.B) {
	var createTree = func(c Context) Displayable {
		return VBox(c, Children(func() {
			for i := 0; i < 100; i++ {
				HBox(c, Children(func() {
					Box(c, Children(func() {
						Box(c, Children(func() {
							Box(c, Children(func() {
								Box(c)
								Box(c)
								Box(c)
								Box(c)
								Box(c)
								Box(c)
							}))
						}))
					}))
				}))
			}
		}))
	}

	b.Run("Basic instantiation", func(b *testing.B) {
		tree := createTree(context.New())
		count := nodeCount(tree)
		log.Printf("BENCHMARK WITH %v NODES", count)

		for i := 0; i < b.N; i++ {
			c := context.New()
			s := surface.NewFake()
			tree := createTree(c)
			tree.Draw(s)
		}
	})

	b.Run("Singular instantiation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b := Box(context.New(), Width(100), Height(100))
			b.Draw(surface.NewFake())
		}
	})
}
