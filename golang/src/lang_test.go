package main

import (
	"assert"
	. "display"
	"testing"
)

// Hypothetical display component
func Render() Surface {
	surface := &FakeSurface{}
	return CreateRenderer(surface, func(s Surface) {
		// MyComponent(Layouts(), Styles(BgColor(0xfc0), Rectangle()))

		Window(s, func() {
			/*
				Styles(s, func() {
					For("Window", BgColor(0xfc0), StrokeSize(5), StrokeStyle(STROKE_DASH), StrokeColor(0xff0000))
					For("Header", BgColor(0xccc))
					For("Window.VBox", BgColor(0x0f0))
					For("AppBody", BgColor(0xfff))
					For("Foo", FontSize(10))
					For("Bar", FontWeight(Bold))
					For("Bar:hover", FontWeight(Italic, Bold))
				})
			*/

			// On("Window", Resize(update))

			/*
				VBox(s, &Opts{FlexWidth: 1, FlexHeight: 1}, func() {
					Header(s, &Opts{FlexWidth: 1, Height: 80})
					HBox(s, &Opts{FlexWidth: 1, FlexHeight: 1}, func() {
						// LeftNav(s, &Opts{Traits: "Foo:Bar", FlexWidth: 1, FlexHeight: 1})
						// AppBody(s, &Opts{Traits: []Trait{Foo, Bar, Baz}, FlexWidth: 4, FlexHeight: 1})
					})
					Footer(s, &Opts{FlexWidth: 1, Height: 60})
				})
			*/
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
