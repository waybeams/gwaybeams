package main

import (
	"assert"
	"display"
	. "example"
	"fmt"
	"testing"
)

// Hypothetical display component
func Render() {
	Window(&Opts{FlexWidth: 80, FlexHeight: 60}, func() {
		VBox(func() {
			Header()
			HBox(func() {
				LeftNav()
				AppBody()
				Box()
			})
			Footer()
		})
	})
}

type RenderContext interface {
	Push(instance display.Displayable)
}

type FakeContext struct {
}

func (c *FakeContext) Push(instance display.Displayable) {
	fmt.Println("FakeContext.Push Called!")
}

type creationFunction func() display.Displayable

// General function that can bind a concrete RenderContext to component
// factory functions so that components can be instantiated with a minimal
// amount of duplicate boilerplate and ceremony.
func CreateRenderer(context RenderContext, creator creationFunction) (wrapper func(args ...interface{}), err error) {
	render := func(args ...interface{}) {
		fmt.Println("Creating new Component")
		// TODO(lbayes): Figure out the shape of the arguments, process them into:
		// Optional Opts object in position 0
		// Children renderer function in position 0 or 1 (depending on presence of Opts object)
		instance := creator()
		context.Push(instance)
	}

	return render, nil
}

type Foo struct{}

// Tests that validate or invalidate my assumptions about the Go Language itself
// These tests should eventually be removed, once comprehension is more
// complete.
func TestLang(t *testing.T) {
	t.Run("Passing a Struct constructor to a function", func(t *testing.T) {
		// Was unable to find a Type reference to annotate a function like the "new"
		// function. Looking at the docs for builtin.new, the type of the argument
		// is Type. But this seems to be unavailable in userspace. Not sure how to
		// provide this functionality. Was considering building a component registry
		// where library and user definitions could be aggregated and made available
		// from a single call site.
		instance := new(Foo)
		assert.NotNil(instance)
	})

	t.Run("Import definitions to root scope", func(t *testing.T) {
		Render()
	})

	t.Run("Wrap pseudo-constructor functions with factory functionality", func(t *testing.T) {
		context := &FakeContext{}

		factory, _ := CreateRenderer(context, display.NewBox)
		assert.NotNil(factory)

		factory(nil, func() {
			fmt.Println("Inside factory children")
		})
	})

	t.Run("Can I pass an int to a float field?", func(t *testing.T) {
		var foo = func(value float64) {}
		// Thankfully, this does throw a compile error.
		// myValue := 42
		// foo(myValue)

		// Interestingly, this does not:
		foo(42)
		// Of course, this also does not throw:
		foo(0.42)
	})
}
