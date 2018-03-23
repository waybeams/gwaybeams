package display

import (
	"testing"
)

func TestGlfwBuilder(t *testing.T) {

	t.Run("With defaults", func(t *testing.T) {
		builder := NewGlfwBuilder()
		if builder.GetFrameRate() != DefaultFrameRate {
			t.Errorf("Unexpected default frame rate: %d", builder.GetFrameRate())
		}

		if builder.GetWindowWidth() != DefaultWindowWidth {
			t.Errorf("Unexpected default size %d", builder.GetWindowWidth())
		}
		if builder.GetWindowHeight() != DefaultWindowHeight {
			t.Errorf("Unexpected default size %d", builder.GetWindowHeight())
		}

		if builder.GetWindowHint(Resizable) != true {
			t.Errorf("Unexpected default window to be Resizable")
		}
		if builder.GetWindowHint(Focused) != true {
			t.Errorf("Unexpected default window to be Focused")
		}
		if builder.GetWindowHint(Visible) != true {
			t.Errorf("Unexpected default window to be Visible")
		}
		if builder.GetWindowHint(Floating) != true {
			t.Errorf("Unexpected default window to be Floating")
		}
		if builder.GetWindowHint(Decorated) != false {
			t.Errorf("Unexpected default window to not be Decorated")
		}
		if builder.GetWindowTitle() != "Default Title" {
			t.Errorf("Unexpected default title %s", builder.GetWindowTitle())
		}
	})

	t.Run("With multiple options", func(t *testing.T) {
		builder := NewGlfwBuilder(WindowSize(120, 240), FrameRate(24))

		width, _ := builder.GetWindowSize()
		if width != 120 {
			t.Error("Expected Width")
		}
		if builder.GetFrameRate() != 24 {
			t.Errorf("Expected configured FrameRate, but found %d", builder.GetFrameRate())
		}
	})

	t.Run("Accepts FrameRate", func(t *testing.T) {
		builder := NewGlfwBuilder(FrameRate(12))
		if builder.GetFrameRate() != 12 {
			t.Errorf("Expected configured FrameRate, but found %d", builder.GetFrameRate())
		}
	})

	t.Run("Accepts Size", func(t *testing.T) {
		builder := NewGlfwBuilder(WindowSize(800, 600))
		width, height := builder.GetWindowSize()
		if width != 800 {
			t.Errorf("Expected configured Width, but found %d", width)
		}
		if height != 600 {
			t.Errorf("Expected configured Height, but found %d", height)
		}
	})

	t.Run("Accepts WindowHints", func(t *testing.T) {
		builder := NewGlfwBuilder(WindowHint(Floating, true), WindowHint(Resizable, false))
		if builder.GetWindowHint(Floating) != true {
			t.Errorf("Expected WindowHint to be floating")
		}
		if builder.GetWindowHint(Resizable) != false {
			t.Errorf("Expected WindowHint to be resizable")
		}
	})
}
