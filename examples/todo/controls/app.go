package controls

import (
	"fmt"
	"github.com/waybeams/waybeams/examples/todo/model"
	ctrl "github.com/waybeams/waybeams/pkg/controls"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

func todoModelsToSpecs(items []*model.Item) []spec.ReadWriter {
	result := []spec.ReadWriter{}
	for index, itemModel := range items {
		result = append(result, ItemSpec(itemModel, index))
	}
	return result
}

// AppRenderer returns a function that, when called, will create a tree
// of specifications that describe the current state of the provided model.
func AppRenderer(appModel *model.App) func() spec.ReadWriter {
	styles := CreateStyles()

	return func() spec.ReadWriter {
		return ctrl.VBox(
			opts.Key("App"),
			styles.Box,
			opts.FontColor(0x111111ff),
			opts.FontFace("Roboto"),
			opts.FontSize(24),
			opts.HAlign(spec.AlignCenter),
			opts.Child(ctrl.VBox(
				styles.Box,
				opts.Key("Body"),
				opts.FlexWidth(1),
				opts.FlexHeight(1),
				opts.Gutter(1),
				opts.MaxWidth(500),
				opts.MinWidth(350),
				opts.HAlign(spec.AlignCenter),

				opts.Child(ctrl.Label(
					opts.FontColor(0xaf2f2f26),
					opts.FontFace("Roboto Light"),
					opts.FontSize(100),
					opts.Text("TODO"),
				)),
				opts.Child(ctrl.TextInput(
					ctrl.Placeholder("Description"),
					opts.BgColor(0xecececff),
					opts.BindStringPayloadTo(events.TextChanged, appModel.UpdateEnteredText),
					opts.FlexWidth(1),
					opts.FontSize(36),
					opts.Key("NewItemInput"),
					opts.OnEnterKey(events.Empty(appModel.CreateItemFromEnteredText)),
					opts.Padding(18),
					opts.Text(appModel.EnteredText()),
				)),
				opts.Child(ctrl.VBox(
					opts.Key("Todo Items"),
					opts.MinHeight(300),
					opts.FlexWidth(1),
					opts.BgColor(0xeeeeeeff),
					opts.Childrenf(func() []spec.ReadWriter {
						result := []spec.ReadWriter{}
						for index, itemModel := range appModel.CurrentItems() {
							result = append(result, ItemSpec(itemModel, index))
						}
						return result
					}),
				)),
				opts.Child(ctrl.HBox(
					styles.Box,
					opts.Key("Footer"),
					opts.FlexWidth(1),
					opts.FontColor(0xccccccff),
					opts.FontFace("Roboto"),
					opts.FontSize(18),
					opts.HAlign(spec.AlignCenter),
					opts.Padding(5),

					opts.Child(ctrl.Label(
						opts.Text(fmt.Sprintf("%d items", len(appModel.CurrentItems()))),
						styles.Button,
					)),
					opts.Child(ctrl.Button(
						opts.Text("All"),
						styles.Button,
						styles.SelectedFilter(appModel.Showing() == model.AllItems),
						opts.OnClick(events.Empty(appModel.ShowAllItems)),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Active"),
						styles.Button,
						styles.SelectedFilter(appModel.Showing() == model.ActiveItems),
						opts.OnClick(events.Empty(appModel.ShowActiveItems)),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Completed"),
						styles.Button,
						styles.SelectedFilter(appModel.Showing() == model.CompletedItems),
						opts.OnClick(events.Empty(appModel.ShowCompletedItems)),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Clear Completed"),
						styles.Button,
						opts.OnClick(func(e events.Event) {
							appModel.ClearCompleted()
						}),
					)),
				)),
			)),
		)
	}
}
