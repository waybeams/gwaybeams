package main

import (
	"assert"
	"display"
	"testing"
)

func TestBoxesMain(t *testing.T) {

	t.Run("", func(t *testing.T) {
		// Use the (default) headless builder to test the application composition
		root, _ := display.NewBuilder().Build(Composer)
		root.Render()
		// NOTE(lbayes): This should be 1, but is not yet implemented
		assert.Equal(root.GetChildCount(), 2)
	})
}
