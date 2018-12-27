package ctrl

import (
	"strconv"

	"github.com/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

const ActiveButton = "Active Button"
const AllButton = "All Button"
const ClearCompletedButton = "Clear Completed Button"
const CompletedButton = "Completed Button"

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
			opts.Text(strconv.Itoa(len(appModel.CurrentItems()))+" items"),
			styles.Button,
		)),
		opts.Child(ctrl.Button(
			opts.Key(AllButton),
			opts.Text("All"),
			styles.Button,
			styles.SelectedFilter(appModel.Showing() == model.AllItems),
			opts.OnClick(events.EmptyHandler(appModel.ShowAllItems)),
		)),
		opts.Child(ctrl.Button(
			opts.Key(ActiveButton),
			opts.Text("Active"),
			opts.IsDisabled(len(appModel.ActiveItems()) == 0),
			styles.Button,
			styles.SelectedFilter(appModel.Showing() == model.ActiveItems),
			opts.OnClick(events.EmptyHandler(appModel.ShowActiveItems)),
		)),
		opts.Child(ctrl.Button(
			opts.Key(CompletedButton),
			opts.Text("Completed"),
			opts.IsDisabled(len(appModel.CompletedItems()) == 0),
			styles.Button,
			styles.SelectedFilter(appModel.Showing() == model.CompletedItems),
			opts.OnClick(events.EmptyHandler(appModel.ShowCompletedItems)),
		)),
		opts.Child(ctrl.Button(
			opts.Key(ClearCompletedButton),
			opts.Text("Clear Completed"),
			opts.IsDisabled(len(appModel.CompletedItems()) == 0),
			styles.Button,
			opts.OnClick(func(e events.Event) {
				appModel.ClearCompleted()
			}),
		)),
	)
}
