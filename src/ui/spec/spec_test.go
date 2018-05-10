package spec_test

import (
	"assert"
	"testing"
	"ui/spec"
)

type fakeSpec struct {
	spec.ControlSpec
	placeholder string
}

func (f *fakeSpec) Placeholder() string {
	name := f.Name()
	if name == "" {
		return f.placeholder
	}
	return name
}

func Placeholder(value string) spec.Option {
	return func(w spec.ReadWriter) {
		f := w.(*fakeSpec)
		f.placeholder = value
	}
}

func fake(options ...spec.Option) *fakeSpec {
	defaults := []spec.Option{
		spec.Width(30),
		spec.Height(25),
	}
	instance := &fakeSpec{}
	spec.ApplyAll(instance, defaults, options)
	return instance
}

type fakeContainerSpec struct {
	spec.ControlSpec
}

func fakeContainer(options ...spec.Option) *fakeContainerSpec {
	defaults := []spec.Option{
		spec.Child(fake(spec.Key("one"))),
		spec.Child(fake(spec.Key("two"))),
		spec.Child(fake(spec.Key("three"))),
	}
	options = append(defaults, options...)

	instance := &fakeContainerSpec{}
	spec.Apply(instance, options...)
	return instance
}

func TestSpec(t *testing.T) {
	t.Run("Apply", func(t *testing.T) {
		instance := spec.Apply(&fakeSpec{},
			Placeholder("abcd"),
			spec.Width(20),
			spec.Height(30)).(*fakeSpec)

		assert.Equal(t, instance.Placeholder(), "abcd")
		assert.Equal(t, instance.Width(), 20)
		assert.Equal(t, instance.Height(), 30)
	})

	t.Run("ApplyAll", func(t *testing.T) {
		defaults := []spec.Option{spec.Width(100)}
		options := []spec.Option{spec.Height(110)}
		instance := spec.ApplyAll(&fakeSpec{}, defaults, options)

		assert.Equal(t, instance.Width(), 100)
		assert.Equal(t, instance.Height(), 110)
	})

	t.Run("Adds Child nodes", func(t *testing.T) {
		root := fake(spec.Key("root"),
			spec.Child(fake(spec.Key("abcd"), spec.Width(40))),
			spec.Child(fake(spec.Key("efgh"), spec.Width(45),
				spec.Child(fake(spec.Key("ijkl")))),
			))

		assert.Equal(t, root.ChildCount(), 2)
		assert.Equal(t, root.ChildAt(0).Key(), "abcd")
		assert.Equal(t, root.ChildAt(1).ChildAt(0).Key(), "ijkl")
	})

	t.Run("Container type", func(t *testing.T) {
		root := fakeContainer(spec.Key("root"), spec.Width(50), spec.Height(55))
		assert.Equal(t, root.ChildCount(), 3)
		assert.Nil(t, root.Parent())

		// Child one
		one := root.ChildAt(0)
		assert.Equal(t, one.Key(), "one")
		assert.Equal(t, one.Parent().Key(), "root")

		// Child two
		two := root.ChildAt(1)
		assert.Equal(t, two.Key(), "two")
		assert.Equal(t, two.Parent().Key(), "root")

		// Child three
		three := root.ChildAt(2)
		assert.Equal(t, three.Key(), "three")
		assert.Equal(t, three.Parent().Key(), "root")
	})

	t.Run("Dynamic fields", func(t *testing.T) {
		root := func() spec.ReadWriter {
			key := "root"
			return fake(spec.Key(key),
				spec.Child(fake(spec.Key(key+"-child"))))
		}()
		assert.Equal(t, root.ChildCount(), 1)
		assert.Equal(t, root.ChildAt(0).Key(), "root-child")
	})

	t.Run("Bag", func(t *testing.T) {
		b := spec.Bag(spec.Width(30), spec.Height(40))
		node := fake(b)
		assert.Equal(t, node.Width(), 30)
		assert.Equal(t, node.Height(), 40)
	})
}
