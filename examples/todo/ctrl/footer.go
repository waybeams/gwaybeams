package ctrl

import (
	"fmt"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

func Footer(appModel *model.App, styles *Styles) spec.ReadWriter {
	return ctrl.HBox(
		styles.Box,
		opts.Key("Footer"),
		opts.FlexWidth(1),
		opts.FontColor(0xccccccff),
		opts.FontFace("Roboto"),
		opts.FontSize(18),
		opts.HAlign(spec.AlignCenter),
		opts.Padding(5),

		opts.Child(ctrl.Label(
			opts.Key("Item Count"),
			opts.Text(fmt.Sprintf("%d items", len(appModel.CurrentItems()))),
			styles.Button,
		)),
		opts.Child(ctrl.Button(
			opts.Key("All Button"),
			opts.Text("All"),
			styles.Button,
			styles.SelectedFilter(appModel.Showing() == model.AllItems),
			opts.OnClick(events.Empty(appModel.ShowAllItems)),
		)),
		opts.Child(ctrl.Button(
			opts.Key("Active Button"),
			opts.Text("Active"),
			styles.Button,
			styles.SelectedFilter(appModel.Showing() == model.ActiveItems),
			opts.OnClick(events.Empty(appModel.ShowActiveItems)),
		)),
		opts.Child(ctrl.Button(
			opts.Key("Completed Button"),
			opts.Text("Completed"),
			styles.Button,
			styles.SelectedFilter(appModel.Showing() == model.CompletedItems),
			opts.OnClick(events.Empty(appModel.ShowCompletedItems)),
		)),
		opts.Child(ctrl.Button(
			opts.Key("Clear Completed Button"),
			opts.Text("Clear Completed"),
			styles.Button,
			opts.OnClick(func(e events.Event) {
				appModel.ClearCompleted()
			}),
		)),
	)
}
