package ctrl

import (
	"github.com/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

const NewItemInput = "New Item Input"

func ItemCreate(model *model.App, options ...spec.Option) spec.ReadWriter {
	input := ctrl.TextInput(
		ctrl.Placeholder("Description"),
		opts.BgColor(0xecececff),
		opts.FlexWidth(1),
		opts.FontSize(36),
		opts.Key(NewItemInput),
		opts.On(events.TextChanged, events.StringPayload(model.UpdateEnteredText)),
		opts.Padding(18),
		opts.Text(model.EnteredText()),
	)

	return ctrl.Form(
		opts.FlexWidth(1),
		opts.Child(input),
		opts.On(events.Submitted, func(e events.Event) {
			m := e.Payload().(map[string]interface{})
			model.CreateItem(m[NewItemInput].(string))
			// TODO(lbayes): FOLLOWING SHOULD NOT BE NECESSARY!
			input.SetText("")
		}),
		// Allow option overrides
		opts.Bag(options...),
	)
}
