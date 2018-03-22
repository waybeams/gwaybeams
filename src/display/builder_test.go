package display

import (
	"assert"
	"testing"
)

func FakeSprite(b Builder, opts *Opts) Displayable {
	sprite := NewSpriteWithOpts(opts)
	b.Push(sprite)
	return sprite
}

func TestBuilder(t *testing.T) {
	t.Run("Instantiated", func(t *testing.T) {
		builder := NewBuilder()
		if builder == nil {
			t.Error("Expected builder instance")
		}
	})

	t.Run("With defaults", func(t *testing.T) {
		builder := NewBuilder()
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
		builder := NewBuilder(SurfaceType(ImageSurfaceType), FrameRate(24))

		if builder.GetSurfaceType() != ImageSurfaceType {
			t.Error("Expected FakeSurfaceType")
		}
		if builder.GetFrameRate() != 24 {
			t.Errorf("Expected configured FrameRate, but found %d", builder.GetFrameRate())
		}
	})

	t.Run("Accepts SurfaceTypeName", func(t *testing.T) {
		builder := NewBuilder(SurfaceType(FakeSurfaceType))
		if builder.GetSurfaceType() != FakeSurfaceType {
			t.Error("Expected FakeSurfaceType")
		}
	})

	t.Run("Accepts FrameRate", func(t *testing.T) {
		builder := NewBuilder(FrameRate(12))
		if builder.GetFrameRate() != 12 {
			t.Errorf("Expected configured FrameRate, but found %d", builder.GetFrameRate())
		}
	})

	t.Run("Accepts Size", func(t *testing.T) {
		builder := NewBuilder(WindowSize(800, 600))
		width, height := builder.GetWindowSize()
		if width != 800 {
			t.Errorf("Expected configured Width, but found %d", width)
		}
		if height != 600 {
			t.Errorf("Expected configured Height, but found %d", height)
		}
	})

	t.Run("Accepts WindowHints", func(t *testing.T) {
		builder := NewBuilder(WindowHint(Floating, true), WindowHint(Resizable, false))
		if builder.GetWindowHint(Floating) != true {
			t.Errorf("Expected WindowHint to be floating")
		}
		if builder.GetWindowHint(Resizable) != false {
			t.Errorf("Expected WindowHint to be resizable")
		}
	})

	t.Run("Returns error when more than one root node is provided", func(t *testing.T) {
		builder := NewBuilder()
		box, err := builder.Build(func(b Builder) {
			FakeSprite(b, &Opts{})
			FakeSprite(b, &Opts{})
		})
		if err == nil {
			t.Error("Expected an error from builder")
		}
		assert.ErrorMatch("single root node", err)

		if box != nil {
			t.Errorf("Expected nil result with error state")
		}
	})

	t.Run("Builds provided elements", func(t *testing.T) {
		builder := NewBuilder()
		sprite, _ := builder.Build(func(b Builder) {
			FakeSprite(b, &Opts{Width: 200, Height: 100})
		})
		if sprite == nil {
			t.Error("Expected root displayable to be returned")
		}
		if sprite.GetWidth() != 200.0 {
			t.Errorf("Expected sprite width to be set but was %f", sprite.GetWidth())
		}
		if sprite.GetHeight() != 100 {
			t.Errorf("Expected sprite height to be set but was %f", sprite.GetHeight())
		}
	})
}
