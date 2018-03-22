package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Run("Instantiated", func(t *testing.T) {
		builder, _ := NewBuilder()
		if builder == nil {
			t.Error("Expected builder instance")
		}
	})

	t.Run("With defaults", func(t *testing.T) {
		builder, _ := NewBuilder()
		if builder.GetFrameRate() != DefaultFrameRate {
			t.Errorf("Unexpected default frame rate: %d", builder.GetFrameRate())
		}

		if builder.GetWidth() != DefaultWidth {
			t.Errorf("Unexpected default size %d", builder.GetWidth())
		}
		if builder.GetHeight() != DefaultHeight {
			t.Errorf("Unexpected default size %d", builder.GetHeight())
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
		if builder.GetTitle() != "Default Title" {
			t.Errorf("Unexpected default title %s", builder.GetTitle())
		}
	})

	t.Run("With multiple options", func(t *testing.T) {
		builder, _ := NewBuilder(Surface(ImageSurface), FrameRate(24))

		if builder.GetSurfaceType() != ImageSurface {
			t.Error("Expected FakeSurface")
		}
		if builder.GetFrameRate() != 24 {
			t.Errorf("Expected configured FrameRate, but found %d", builder.GetFrameRate())
		}
	})

	t.Run("Accepts SurfaceType", func(t *testing.T) {
		builder, _ := NewBuilder(Surface(FakeSurface))
		if builder.GetSurfaceType() != FakeSurface {
			t.Error("Expected FakeSurface")
		}
	})

	t.Run("Accepts FrameRate", func(t *testing.T) {
		builder, _ := NewBuilder(FrameRate(12))
		if builder.GetFrameRate() != 12 {
			t.Errorf("Expected configured FrameRate, but found %d", builder.GetFrameRate())
		}
	})

	t.Run("Accepts Size", func(t *testing.T) {
		builder, _ := NewBuilder(Size(800, 600))
		width, height := builder.GetSize()
		if width != 800 {
			t.Errorf("Expected configured Width, but found %d", width)
		}
		if height != 600 {
			t.Errorf("Expected configured Height, but found %d", height)
		}
	})

	t.Run("Accepts WindowHints", func(t *testing.T) {
		builder, _ := NewBuilder(WindowHint(Floating, true), WindowHint(Resizable, false))
		if builder.GetWindowHint(Floating) != true {
			t.Errorf("Expected WindowHint to be floating")
		}
		if builder.GetWindowHint(Resizable) != false {
			t.Errorf("Expected WindowHint to be resizable")
		}
	})
}
