package controls

import (
	"github.com/waybeams/waybeams/examples/todo/model"
	ctrl "github.com/waybeams/waybeams/pkg/controls"
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
			opts.Text(completedLabel),
			opts.OnClick(events.Empty(model.ToggleCompleted)),
		)),
		opts.Child(ctrl.Label(
			opts.Text(model.Description),
			opts.FlexWidth(1),
		)),
		opts.Child(ctrl.Button(
			opts.Text("X"),
			opts.OnClick(events.Empty(model.Delete)),
		)),
	)
}
