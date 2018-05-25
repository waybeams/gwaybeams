package spec_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/fakes"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
	"testing"
)

func TestSpec(t *testing.T) {
	t.Run("Apply", func(t *testing.T) {
		instance := spec.Apply(&fakes.FakeSpec{},
			fakes.Placeholder("abcd"),
			opts.Width(20),
			opts.Height(30)).(*fakes.FakeSpec)

		assert.Equal(instance.Placeholder(), "abcd")
		assert.Equal(instance.Width(), 20)
		assert.Equal(instance.Height(), 30)
	})

	t.Run("Apply with empty entry", func(t *testing.T) {
		instance := spec.Apply(&fakes.FakeSpec{},
			opts.Width(20),
			opts.Height(30),
			opts.Empty(),
		)
		assert.Equal(instance.Width(), 20)
		assert.Equal(instance.Height(), 30)
	})

	t.Run("Apply fails with nil entry", func(t *testing.T) {
		// Call with opts.Empty() instead.
		assert.Panic("runtime error: invalid memory address or nil pointer dereference", func() {
			spec.Apply(&fakes.FakeSpec{},
				opts.Width(34),
				nil,
			)
			assert.True(false, "Execution should not reach this line")
		})
	})

	t.Run("ApplyAll", func(t *testing.T) {
		defaults := []spec.Option{opts.Width(100)}
		options := []spec.Option{opts.Height(110)}
		instance := spec.ApplyAll(&fakes.FakeSpec{}, defaults, options)

		assert.Equal(instance.Width(), 100)
		assert.Equal(instance.Height(), 110)
	})

	t.Run("Dynamic fields", func(t *testing.T) {
		root := func() spec.ReadWriter {
			key := "root"
			return fakes.Fake(opts.Key(key),
				opts.Child(fakes.Fake(opts.Key(key+"-child"))))
		}()
		assert.Equal(root.ChildCount(), 1)
		assert.Equal(root.ChildAt(0).Key(), "root-child")
	})

	t.Run("Bag", func(t *testing.T) {
		b := opts.Bag(opts.Width(30), opts.Height(40))
		node := fakes.Fake(b)
		assert.Equal(node.Width(), 30)
		assert.Equal(node.Height(), 40)
	})

	t.Run("Contains", func(t *testing.T) {
		t.Run("is false when unrelated", func(t *testing.T) {
			one := fakes.Fake()
			two := fakes.Fake()
			assert.False(spec.Contains(one, two))
		})

		t.Run("is true when descended", func(t *testing.T) {
			root := fakes.Fake(
				opts.Key("root"),
				opts.Child(fakes.Fake(opts.Key("child"))),
			)

			child := root.ChildAt(0)
			assert.True(spec.Contains(root, child))
			assert.False(spec.Contains(child, root))
		})

		t.Run("false for same control", func(t *testing.T) {
			root := fakes.Fake()
			assert.False(spec.Contains(root, root))
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

			assert.False(spec.Contains(five, one))
			assert.False(spec.Contains(five, two))
			assert.False(spec.Contains(five, three))
			assert.False(spec.Contains(five, four))
			assert.False(spec.Contains(five, five))

			assert.True(spec.Contains(one, five))
			assert.True(spec.Contains(two, five))
			assert.True(spec.Contains(three, five))
			assert.True(spec.Contains(four, five))
		})
	})

	t.Run("Invalidate", func(t *testing.T) {
		received := []events.Event{}
		root := fakes.Fake(
			opts.Child(fakes.Fake(
				opts.Child(fakes.Fake(opts.Key("Nested Child"))))),
		)
		root.On(events.Invalidated, func(e events.Event) {
			received = append(received, e)
		})

		child := spec.FirstByKey(root, "Nested Child")
		child.Invalidate()
		child.Invalidate()
		child.Invalidate()

		assert.Equal(len(received), 3)
		assert.Nil(received[0].Target(), "Expected nil target because Go embed != inherit")
		assert.Nil(received[0].Payload())
	})
}
