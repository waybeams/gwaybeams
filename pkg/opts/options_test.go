package opts_test

import (
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/fakes"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

func TestOptions(t *testing.T) {
	t.Run("Width", func(t *testing.T) {
		f := fakes.Fake(opts.Width(10))
		assert.Equal(f.Width(), 10)
	})

	t.Run("Height", func(t *testing.T) {
		f := fakes.Fake(opts.Height(10))
		assert.Equal(f.Height(), 10)
	})

	t.Run("BgColor", func(t *testing.T) {
		f := fakes.Fake(opts.BgColor(0xffcc00ff))
		assert.Equal(f.BgColor(), 0xffcc00ff)
	})

	t.Run("Child", func(t *testing.T) {
		root := fakes.Fake(
			opts.Key("root"),
			opts.Child(fakes.Fake(opts.Key("child"))))

		assert.Equal(root.ChildCount(), 1)
		assert.Equal(root.Children()[0].Key(), "child")
	})

	t.Run("Childf", func(t *testing.T) {
		root := fakes.Fake(
			opts.Key("root"),
			opts.Childf(func() spec.ReadWriter {
				return fakes.Fake(opts.Key("child-1"))
			}),
		)
		assert.Equal(root.ChildCount(), 1)
		assert.Equal(root.Children()[0].Key(), "child-1")
	})

	t.Run("Childf", func(t *testing.T) {
		t.Run("Update", func(t *testing.T) {
			var childWidth = 20.0

			root := fakes.Fake(
				opts.Key("root"),
				opts.Childf(func() spec.ReadWriter {
					return fakes.Fake(
						opts.Key("child-1"),
						opts.Width(childWidth),
					)
				}),
			)

			assert.Equal(root.ChildCount(), 1)
			assert.Equal(spec.FirstChild(root).Width(), 20)

			// Re-run the factory function with new (hoisted) childWidth value
			childWidth = 30.0
			factory := spec.FirstChild(root).Factory()
			opts.Childf(factory)(root)

			assert.Equal(root.ChildCount(), 1)
			assert.Equal(spec.FirstChild(root).Width(), 30)
		})
	})

	t.Run("Childrenf", func(t *testing.T) {
		var prefix = "child"

		var createTree = func() spec.ReadWriter {
			return fakes.Fake(
				opts.Childrenf(func() []spec.ReadWriter {
					return []spec.ReadWriter{
						fakes.Fake(opts.Key(prefix + "-1")),
						fakes.Fake(opts.Key(prefix + "-2")),
						fakes.Fake(opts.Key(prefix + "-3")),
					}
				}),
			)
		}

		t.Run("Simple", func(t *testing.T) {
			root := createTree()
			assert.Equal(root.ChildCount(), 3)
			assert.Equal(root.Children()[2].Key(), "child-3")
		})

		t.Run("Update", func(t *testing.T) {
			root := createTree()
			// Update the lexical variable for re-construction
			prefix = "abcd"

			// Configure a new Childrenf execution with the original factory function.
			opt := opts.Childrenf(root.ChildAt(0).SiblingsFactory())
			// Execute the Childrenf execution on the original root component.
			opt(root)

			assert.Equal(root.ChildCount(), 3)
			assert.Equal(root.Children()[2].Key(), "abcd-3")
		})
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

	t.Run("Padding", func(t *testing.T) {
		f := fakes.Fake(opts.Padding(10))
		assert.Equal(f.PaddingBottom(), 10)
		assert.Equal(f.PaddingLeft(), 10)
		assert.Equal(f.PaddingRight(), 10)
		assert.Equal(f.PaddingTop(), 10)
	})

	t.Run("PaddingBottom", func(t *testing.T) {
		f := fakes.Fake(opts.PaddingBottom(10))
		assert.Equal(f.PaddingBottom(), 10)
	})

	t.Run("PaddingLeft", func(t *testing.T) {
		f := fakes.Fake(opts.PaddingLeft(10))
		assert.Equal(f.PaddingLeft(), 10)
	})

	t.Run("PaddingRight", func(t *testing.T) {
		f := fakes.Fake(opts.PaddingRight(10))
		assert.Equal(f.PaddingRight(), 10)
	})

	t.Run("PaddingTop", func(t *testing.T) {
		f := fakes.Fake(opts.PaddingTop(10))
		assert.Equal(f.PaddingTop(), 10)
	})

	t.Run("PrefHeight", func(t *testing.T) {
		f := fakes.Fake(opts.PrefHeight(10))
		assert.Equal(f.PrefHeight(), 10)
	})

	t.Run("PrefWidth", func(t *testing.T) {
		f := fakes.Fake(opts.PrefWidth(10))
		assert.Equal(f.PrefWidth(), 10)
	})

	t.Run("Width", func(t *testing.T) {
		f := fakes.Fake(opts.Width(10))
		assert.Equal(f.Width(), 10)
	})

	t.Run("Height", func(t *testing.T) {
		f := fakes.Fake(opts.Height(10))
		assert.Equal(f.Height(), 10)
	})

	t.Run("Visible", func(t *testing.T) {
		f := fakes.Fake(opts.Visible(false))
		assert.False(f.Visible())
	})
}
