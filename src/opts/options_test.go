package opts_test

import (
	"github.com/waybeams/assert"
	"fakes"
	"opts"
	"testing"
)

func TestOptions(t *testing.T) {
	t.Run("ActualWidth", func(t *testing.T) {
		f := fakes.Fake(opts.ActualWidth(10))
		assert.Equal(f.ActualWidth(), 10)
	})

	t.Run("ActualHeight", func(t *testing.T) {
		f := fakes.Fake(opts.ActualHeight(10))
		assert.Equal(f.ActualHeight(), 10)
	})

	t.Run("BgColor", func(t *testing.T) {
		f := fakes.Fake(opts.BgColor(0xffcc00ff))
		assert.Equal(f.BgColor(), 0xffcc00ff)
	})

	t.Run("ExcludeFromlayout(true)", func(t *testing.T) {
		f := fakes.Fake(opts.ExcludeFromLayout(true))
		assert.Equal(f.ExcludeFromLayout(), true)
	})

	t.Run("ExcludeFromlayout(false)", func(t *testing.T) {
		f := fakes.Fake(opts.ExcludeFromLayout(false))
		assert.Equal(f.ExcludeFromLayout(), false)
	})

	t.Run("FlexHeight", func(t *testing.T) {
		f := fakes.Fake(opts.FlexHeight(2))
		assert.Equal(f.FlexHeight(), 2)
	})

	t.Run("FlexWidth", func(t *testing.T) {
		f := fakes.Fake(opts.FlexWidth(2))
		assert.Equal(f.FlexWidth(), 2)
	})

	t.Run("FontColor", func(t *testing.T) {
		f := fakes.Fake(opts.FontColor(0xffcc00ff))
		assert.Equal(f.FontColor(), 0xffcc00ff)
	})

	t.Run("FontFace", func(t *testing.T) {
		f := fakes.Fake(opts.FontFace("abcd"))
		assert.Equal(f.FontFace(), "abcd")
	})

	t.Run("FontSize", func(t *testing.T) {
		f := fakes.Fake(opts.FontSize(23))
		assert.Equal(f.FontSize(), 23)
	})

	t.Run("Gutter", func(t *testing.T) {
		f := fakes.Fake(opts.Gutter(10))
		assert.Equal(f.Gutter(), 10)
	})

	t.Run("IsFocusable", func(t *testing.T) {
		f := fakes.Fake(opts.IsFocusable(true))
		assert.Equal(f.IsFocusable(), true)
	})

	t.Run("IsMeasured", func(t *testing.T) {
		f := fakes.Fake(opts.IsMeasured(true))
		assert.Equal(f.IsMeasured(), true)
	})

	t.Run("Width", func(t *testing.T) {
		f := fakes.Fake(opts.Width(10))
		assert.Equal(f.Width(), 10)
	})

	t.Run("Height", func(t *testing.T) {
		f := fakes.Fake(opts.Height(10))
		assert.Equal(f.Height(), 10)
	})
}
