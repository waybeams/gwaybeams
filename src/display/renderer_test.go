package display

import "log"

import (
	"testing"
)

func TestFactory(t *testing.T) {
	surface := &FakeSurface{}

	t.Run("CreateRender", func(t *testing.T) {
		renderer := CreateRenderer(surface, func(s Surface) {
			log.Println("HELLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL")
			Box(s, &Opts{Width: 100, Height: 100, StyleName: "abcd"})
			log.Printf("INSIDE ROOT: %v", s.GetRoot())
		})

		renderer.Render()

		box := renderer.GetRoot()
		log.Printf("OUTSIDE ROOT %v:", box)
		if box == nil {
			t.Error("Expected renderer.GetRoot() to return a valid box")
		}
	})
}
