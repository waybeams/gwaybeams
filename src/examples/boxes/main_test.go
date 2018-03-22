package main

import (
	"assert"
	"display"
	"testing"
)

func TestBoxesMain(t *testing.T) {

	t.Run("", func(t *testing.T) {
		surface := &display.FakeSurface{}
		root, _ := CreateBoxesApp()
		root.Render(surface)
		// NOTE(lbayes): This should be 1, but is not yet implemented
		assert.Equal(root.GetChildCount(), 0)
	})
}
