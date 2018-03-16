package main

import (
	"assert"
	. "example"
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
}
