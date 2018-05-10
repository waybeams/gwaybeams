package spec_test

import (
	"assert"
	"fakes"
	"opts"
	"spec"
	"testing"
)

func TestSpec(t *testing.T) {
	t.Run("Apply", func(t *testing.T) {
		instance := spec.Apply(&fakes.FakeSpec{},
			fakes.Placeholder("abcd"),
			opts.Width(20),
			opts.Height(30)).(*fakes.FakeSpec)

		assert.Equal(t, instance.Placeholder(), "abcd")
		assert.Equal(t, instance.Width(), 20)
		assert.Equal(t, instance.Height(), 30)
	})

	t.Run("ApplyAll", func(t *testing.T) {
		defaults := []spec.Option{opts.Width(100)}
		options := []spec.Option{opts.Height(110)}
		instance := spec.ApplyAll(&fakes.FakeSpec{}, defaults, options)

		assert.Equal(t, instance.Width(), 100)
		assert.Equal(t, instance.Height(), 110)
	})

	t.Run("Dynamic fields", func(t *testing.T) {
		root := func() spec.ReadWriter {
			key := "root"
			return fakes.Fake(opts.Key(key),
				opts.Child(fakes.Fake(opts.Key(key+"-child"))))
		}()
		assert.Equal(t, root.ChildCount(), 1)
		assert.Equal(t, root.ChildAt(0).Key(), "root-child")
	})

	t.Run("Bag", func(t *testing.T) {
		b := opts.Bag(opts.Width(30), opts.Height(40))
		node := fakes.Fake(b)
		assert.Equal(t, node.Width(), 30)
		assert.Equal(t, node.Height(), 40)
	})

	t.Run("Contains", func(t *testing.T) {
		t.Run("is false when unrelated", func(t *testing.T) {
			one := fakes.Fake()
			two := fakes.Fake()
			assert.False(t, spec.Contains(one, two))
		})

		t.Run("is true when descended", func(t *testing.T) {
			root := fakes.Fake(
				opts.Key("root"),
				opts.Child(fakes.Fake(opts.Key("child"))),
			)

			child := root.ChildAt(0)
			assert.True(t, spec.Contains(root, child))
			assert.False(t, spec.Contains(child, root))
		})

		t.Run("false for same control", func(t *testing.T) {
			root := fakes.Fake()
			assert.False(t, spec.Contains(root, root))
		})

		t.Run("deep descendants too", func(t *testing.T) {
			one := fakes.Fake(opts.Key("one"),
				opts.Child(fakes.Fake(opts.Key("two"),
					opts.Child(fakes.Fake(opts.Key("three"),
						opts.Child(fakes.Fake(opts.Key("four"),
							opts.Child(fakes.Fake(opts.Key("five"))),
						)),
					)),
				)),
			)

			two := spec.FirstByKey(one, "two")
			three := spec.FirstByKey(one, "three")
			four := spec.FirstByKey(one, "four")
			five := spec.FirstByKey(one, "five")

			assert.False(t, spec.Contains(five, one))
			assert.False(t, spec.Contains(five, two))
			assert.False(t, spec.Contains(five, three))
			assert.False(t, spec.Contains(five, four))
			assert.False(t, spec.Contains(five, five))

			assert.True(t, spec.Contains(one, five))
			assert.True(t, spec.Contains(two, five))
			assert.True(t, spec.Contains(three, five))
			assert.True(t, spec.Contains(four, five))
		})
	})
}
