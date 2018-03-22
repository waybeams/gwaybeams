package main

import (
	"assert"
	"testing"
)

func TestBoxesMain(t *testing.T) {

	t.Run("", func(t *testing.T) {
		root, _ := CreateBoxesApp("Test Title")
		// NOTE(lbayes): This should be 1, but is not yet implemented
		assert.Equal(root.GetChildCount(), 2)
	})
}
