package fakes

import (
	"opts"
	"spec"
)

type FakeSpec struct {
	spec.Spec

	placeholder string
}

func (f *FakeSpec) Placeholder() string {
	text := f.Text()
	if text == "" {
		return f.placeholder
	}
	return text
}

func Placeholder(value string) spec.Option {
	return func(w spec.ReadWriter) {
		f := w.(*FakeSpec)
		f.placeholder = value
	}
}

func Fake(options ...spec.Option) *FakeSpec {
	defaults := []spec.Option{
		opts.SpecName("FakeControl"),
	}
	instance := &FakeSpec{}
	spec.ApplyAll(instance, defaults, options)
	return instance
}

type FakeContainerSpec struct {
	spec.Spec
}

func FakeContainer(options ...spec.Option) *FakeContainerSpec {
	defaults := []spec.Option{
		opts.Child(Fake(opts.Key("one"))),
		opts.Child(Fake(opts.Key("two"))),
		opts.Child(Fake(opts.Key("three"))),
	}
	options = append(defaults, options...)

	instance := &FakeContainerSpec{}
	spec.Apply(instance, options...)
	return instance
}

/*
func NewBuilder(root spec.ReadWriter) *builder.Builder {
	b := builder.New(root)
	b.Context().AddFont("Roboto", "../../testdata/Roboto-Regular.ttf")
	return b
}
*/
