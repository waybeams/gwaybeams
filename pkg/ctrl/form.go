package ctrl

import (
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/spec"
)

var Form = func(options ...spec.Option) spec.ReadWriter {
	f := spec.New()
	f.SetSpecName("Form")
	f.SetLayoutType(spec.VerticalFlowLayoutType)
	f.On(events.EnterKeyReleased, events.BubbleAs(events.Submitted))
	spec.Apply(f, options...)
	return f
}
