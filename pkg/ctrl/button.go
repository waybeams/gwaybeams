package ctrl

import (
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/views"
)

// Button is a stub control pending implementation.
var Button = func(options ...spec.Option) spec.ReadWriter {
	defaults := []spec.Option{
		opts.SpecName("Button"),
		opts.LayoutType(spec.StackLayoutType),
		opts.IsFocusable(true),
		opts.IsMeasured(true),
		opts.Padding(5),
		opts.View(views.LabelView),

		// The following is terrible. This felt like a reasonable solution to introduce declarative
		// support for Spec states because it allowed us to generalize common functionality, like
		// cursor interactions. But this only works if we mutate and avoid replacing that spec
		// instance over time. This violates our Pull-based or functional approach to components and
		// will require sophisticated merging between existing and new component instances.
		// Purely functional, declarative specifications require all mutable state to be stored
		// externally from the instance.
		// Difficult to debug.
		opts.OnState("active", opts.BgColor(0xce3262ff)),
		opts.OnState("hovered", opts.BgColor(0x00acd7ff)),
		opts.OnState("pressed", opts.BgColor(0x5dc9e2ff)),
		opts.OnState("disabled", opts.BgColor(0xdbd9d6ff)),
		opts.On(events.Entered, opts.OptionsHandler(opts.SetState("hovered"))),
		opts.On(events.Exited, opts.OptionsHandler(opts.SetState("active"))),
		opts.On(events.Pressed, opts.OptionsHandler(opts.SetState("pressed"))),
		opts.On(events.Released, opts.OptionsHandler(opts.SetState("hovered"))),
	}
	button := &LabelSpec{}
	spec.ApplyAll(button, defaults, options)
	return button
}
