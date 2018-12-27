package ctrl

import (
	"github.com/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

func ItemSpec(model *model.Item, index int) spec.ReadWriter {
	var completedLabel string = "[  ]"
	if !model.CompletedAt.IsZero() {
		completedLabel = "[X]"
	}
	var bgColor uint = 0xdededeff
	if !model.CompletedAt.IsZero() {
		bgColor = 0x9e9e9eff
	}
	return ctrl.HBox(
		opts.Key("item-"+string(index)),
		opts.BgColor(bgColor),
		opts.StrokeColor(0x333333ff),
		opts.StrokeSize(1),
		opts.FlexWidth(1),
		opts.Child(ctrl.Button(
			opts.Key("btn"),
			opts.Text(completedLabel),
			opts.OnClick(events.EmptyHandler(model.ToggleCompleted)),
		)),
		opts.Child(ctrl.Label(
			opts.BgColor(0),
			opts.FlexWidth(1),
			opts.Key("desc"),
			opts.StrokeColor(0),
			opts.StrokeSize(0),
			opts.Text(model.Description),
		)),
		opts.Child(ctrl.Button(
			opts.Key("del"),
			opts.Text("X"),
			opts.OnClick(events.EmptyHandler(model.Delete)),
		)),
	)
}
