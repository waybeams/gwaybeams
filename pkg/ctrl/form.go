package ctrl

import (
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/spec"
)

var Form = func(options ...spec.Option) spec.ReadWriter {
	f := spec.New()
	f.SetSpecName("Form")
	f.SetLayoutType(spec.VerticalFlowLayoutType)
	f.On(events.EnterKeyReleased, func(e events.Event) {
		values := make(map[string]interface{})
		kids := f.Children()
		for i := 0; i < len(kids); i++ {
			values[kids[i].Key()] = kids[i].Value()
		}
		f.Bubble(events.New(events.Submitted, f, values))
	})
	spec.Apply(f, options...)
	return f
}
