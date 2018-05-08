package main

import (
	"testing"
	"ui/context"
)

func TestBoxesMain(t *testing.T) {
	t.Run("Main application configuration", func(t *testing.T) {
		win := createWindow(
			context.Font("Roboto", "../../third_party/fonts/Roboto/Roboto-Regular.ttf"),
		)
		one := win.ChildAt(0)

		if win == nil {
			t.Error("Expected win to be returned from createWindow")
		}
		if win.ChildCount() < 1 {
			t.Errorf("Expected at least 1 child on window, but got %d", win.ChildCount())
		}
		if one == nil {
			t.Errorf("Expected at least one child")
		}
	})
}
