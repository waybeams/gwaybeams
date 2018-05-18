package spec_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/fakes"
	"github.com/waybeams/waybeams/pkg/opts"
	"reflect"
	"github.com/waybeams/waybeams/pkg/spec"
	"testing"
)

func TestComposable(t *testing.T) {
	t.Run("Is instantiable", func(t *testing.T) {
		box := &spec.Spec{}
		assert.NotNil(box)
	})

	t.Run("Accepts key", func(t *testing.T) {
		ctrl := spec.Apply(&spec.Spec{}, opts.Key("abcd"))
		assert.Equal(ctrl.Key(), "abcd")
	})

	t.Run("Adds Child nodes", func(t *testing.T) {
		root := fakes.Fake(opts.Key("root"),
			opts.Child(fakes.Fake(opts.Key("abcd"), opts.Width(40))),
			opts.Child(fakes.Fake(opts.Key("efgh"), opts.Width(45),
				opts.Child(fakes.Fake(opts.Key("ijkl")))),
			))

		assert.Equal(root.ChildCount(), 2)
		assert.Equal(root.ChildAt(0).Key(), "abcd")
		assert.Equal(root.ChildAt(1).ChildAt(0).Key(), "ijkl")
	})

	t.Run("Container type", func(t *testing.T) {
		root := fakes.FakeContainer(opts.Key("root"), opts.Width(50), opts.Height(55))
		assert.Equal(root.ChildCount(), 3)
		assert.Nil(root.Parent())

		// Child one
		one := root.ChildAt(0)
		assert.Equal(one.Key(), "one")
		assert.Equal(one.Parent().Key(), "root")

		// Child two
		two := root.ChildAt(1)
		assert.Equal(two.Key(), "two")
		assert.Equal(two.Parent().Key(), "root")

		// Child three
		three := root.ChildAt(2)
		assert.Equal(three.Key(), "three")
		assert.Equal(three.Parent().Key(), "root")
	})

	t.Run("ChildCount", func(t *testing.T) {
		root := fakes.Fake(opts.Key("root"),
			opts.Child(fakes.Fake(
				opts.Key("one"),
				opts.Child(fakes.Fake(opts.Key("two"))),
				opts.Child(fakes.Fake(opts.Key("three"))),
			)),
		)

		one := root.ChildAt(0)
		two := one.ChildAt(0)
		three := one.ChildAt(1)

		assert.Equal(root.ChildCount(), 1)
		assert.Equal(root.ChildAt(0), one)

		assert.Equal(one.ChildCount(), 2)
		assert.Equal(one.ChildAt(0), two)
		assert.Equal(one.ChildAt(1), three)
	})

	t.Run("Children() returns empty list", func(t *testing.T) {
		ctrl := fakes.Fake()
		assert.Equal(len(ctrl.Children()), 0)
	})

	t.Run("ChildCount() and Children() agree", func(t *testing.T) {
		root := fakes.Fake(
			opts.Child(fakes.Fake(opts.Key("one"))),
			opts.Child(fakes.Fake(opts.Key("two"))),
			opts.Child(fakes.Fake(opts.Key("three"))),
		)

		assert.Equal(len(root.Children()), 3)
		assert.Equal(root.ChildCount(), 3)
	})

	t.Run("FindByKey", func(t *testing.T) {

		t.Run("returns current instance", func(t *testing.T) {
			root := fakes.Fake(opts.Key("abcd"))
			result := spec.FirstByKey(root, "abcd")
			assert.Equal(root, result)
		})

		t.Run("returns nil for no result", func(t *testing.T) {
			root := fakes.Fake(opts.Key("abcd"))
			result := spec.FirstByKey(root, "no-match")
			assert.Nil(result)
		})

		t.Run("returns nested instance", func(t *testing.T) {
			root := fakes.Fake(opts.Key("root"),
				opts.Child(fakes.Fake(opts.Key("one"),
					opts.Child(fakes.Fake(opts.Key("two"),
						opts.Child(fakes.Fake(opts.Key("three"))),
					)),
				)),
			)

			three := spec.FirstByKey(root, "three")
			assert.Equal(three.Key(), "three")
		})

		t.Run("Root() returns concrete type", func(t *testing.T) {
			tree := fakes.Fake(opts.Key("root"),
				opts.Child(fakes.Fake(opts.Key("one"),
					opts.Child(fakes.Fake(opts.Key("two"),
						opts.Child(fakes.Fake(opts.Key("three"))),
					)),
				)),
			)

			result := spec.Root(tree)
			// CRITICAL detail regarding Go's embed vs inheritance. This is why we need Root()
			// to be on the module and not the struct. The base struct will only return the
			// base struct and never the embedding struct.
			assert.Equal(reflect.ValueOf(result).Type().String(), "*fakes.FakeSpec")
		})
	})

	t.Run("Path", func(t *testing.T) {
		t.Run("root", func(t *testing.T) {
			root := fakes.Fake(opts.Key("root"))
			assert.Equal(spec.Path(root), "/root")
		})

		t.Run("uses SpecName if Key is not present", func(t *testing.T) {
			root := fakes.Fake()
			assert.Equal(spec.Path(root), "/FakeControl")
		})

	})

	t.Run("defaults to SpecName and parent index", func(t *testing.T) {
		root := fakes.Fake(opts.Key("root"),
			opts.Child(fakes.Fake()),
			opts.Child(fakes.Fake()),
			opts.Child(fakes.Fake()),
		)

		kids := root.Children()
		assert.Equal(spec.Path(kids[0]), "/root/FakeControl0")
		assert.Equal(spec.Path(kids[1]), "/root/FakeControl1")
		assert.Equal(spec.Path(kids[2]), "/root/FakeControl2")
	})

	t.Run("with depth", func(t *testing.T) {
		root := fakes.Fake(opts.Key("root"),
			opts.Child(fakes.Fake(opts.Key("one"),
				opts.Child(fakes.Fake(opts.Key("two"),
					opts.Child(fakes.Fake(opts.Key("three"))),
				)),
				opts.Child(fakes.Fake(opts.Key("four"))),
			)),
		)

		one := spec.FirstByKey(root, "one")
		two := spec.FirstByKey(root, "two")
		three := spec.FirstByKey(root, "three")
		four := spec.FirstByKey(root, "four")

		assert.Equal(spec.Path(one), "/root/one")
		assert.Equal(spec.Path(two), "/root/one/two")
		assert.Equal(spec.Path(three), "/root/one/two/three")
		assert.Equal(spec.Path(four), "/root/one/four")
	})

	/*
		t.Run("Events bubble on FakeSpec", func(t *testing.T) {
			var root, one, two, three, four Displayable
			var received []events.Event
			var receivers []Displayable
			var getHandlerFor = func(d Displayable) events.EventHandler {
				return func(e events.Event) {
					receivers = append(receivers, d)
					received = append(received, e)
				}
			}

			root = Box(context.New(), ID("root"), Children(func(c Context) {
				one = Box(c, ID("one"), Children(func() {
					two = Box(c, ID("two"), Children(func() {
						three = Box(c, ID("three"))
					}))
				}))
				four = Box(c, ID("four"))
			}))

			root.On("fake-event", getHandlerFor(root))
			one.On("fake-event", getHandlerFor(one))
			two.On("fake-event", getHandlerFor(two))
			three.On("fake-event", getHandlerFor(three))
			four.On("fake-event", getHandlerFor(four))

			three.Bubble(events.New("fake-event", three, nil))
			four.Emit(events.New("fake-event", nil, nil))

			assert.Equal(len(received), 5)
			assert.Equal(receivers[0].Path(), "/root/one/two/three")
			assert.Equal(receivers[1].Path(), "/root/one/two")
			assert.Equal(receivers[2].Path(), "/root/one")
			assert.Equal(receivers[3].Path(), "/root")
			assert.Equal(receivers[4].Path(), "/root/four")
		})

		t.Run("Events can be cancelled", func(t *testing.T) {
			secondCalled := false

			instance := events.NewEmitter()

			instance.On("fake-event", func(e events.Event) {
				e.Cancel()
			})
			instance.On("fake-event", func(e events.Event) {
				secondCalled = true
			})
			instance.Emit(events.New("fake-event", nil, nil))
			assert.False(secondCalled, "Expected Cancel to stop event")
		})
	*/
}
