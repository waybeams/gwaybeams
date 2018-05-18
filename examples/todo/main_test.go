package main

import (
	"github.com/waybeams/assert"
	"testing"
)

func TestTodoMain(t *testing.T) {

	t.Run("Default Surface", func(t *testing.T) {
		s := CreateSurface()
		font := s.Font("Roboto")
		if font == nil {
			t.Error("Expected Roboto")
		}
		font = s.Font("Roboto Light")
		if font == nil {
			t.Error("Expected Roboto Light")
		}
	})

	t.Run("App Model", func(t *testing.T) {
		m := CreateModel()
		if m == nil {
			t.Error("Expected model")
		}
		assert.Equal(len(m.CurrentItems()), 6)
	})
}
