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

func AppRenderer(appModel *model.App) func() spec.ReadWriter {
	boxStyle := opts.Bag(
		opts.BgColor(0xffffffff),
		opts.Padding(10),
		opts.Gutter(10),
	)

	headerText := opts.Bag(
		opts.FontColor(0xaf2f2f26),
		opts.FontFace("Roboto Light"),
		opts.FontSize(100),
	)

	mainStyle := opts.Bag(
		boxStyle,
		opts.FontColor(0x111111ff),
		opts.FontFace("Roboto"),
		opts.FontSize(24),
	)

	buttonStyle := opts.Bag(
		opts.BgColor(0xf8f8f8ff),
		// opts.StrokeColor(0x333333ff),
		// opts.StrokeSize(1),
	)

	return func() spec.ReadWriter {
		return ctrl.VBox(
			opts.Key("App"),
			mainStyle,
			opts.HAlign(spec.AlignCenter),
			opts.Child(ctrl.VBox(
				boxStyle,
				opts.Key("Body"),
				opts.FlexWidth(1),
				opts.FlexHeight(1),
				opts.Gutter(1),
				opts.MaxWidth(500),
				opts.MinWidth(350),
				opts.HAlign(spec.AlignCenter),

				opts.Child(ctrl.Label(
					headerText,
					opts.Text("TODO"),
				)),
				opts.Child(ctrl.TextInput(
					opts.Key("NewItemInput"),
					opts.Text(appModel.EnteredText()),
					opts.Padding(18),
					opts.FontSize(36),
					opts.FlexWidth(1),
					opts.BgColor(0xccccccff),
					opts.BindStringPayloadTo(events.TextChanged, appModel.UpdateEnteredText),
					opts.OnEnterKey(events.Empty(appModel.CreateItemFromEnteredText)),
				)),
				opts.Child(ctrl.VBox(
					opts.Key("Todo Items"),
					opts.MinHeight(300),
					opts.FlexWidth(1),
					opts.BgColor(0xeeeeeeff),
					opts.Children(todoModelsToSpecs(appModel.CurrentItems())),
				)),
				opts.Child(ctrl.HBox(
					boxStyle,
					opts.Key("Footer"),
					opts.FlexWidth(1),
					opts.FontColor(0xccccccff),
					opts.FontFace("Roboto"),
					opts.FontSize(18),
					opts.HAlign(spec.AlignCenter),
					opts.Padding(5),

					opts.Child(ctrl.Label(
						opts.Text(fmt.Sprintf("%d items", len(appModel.CurrentItems()))),
						buttonStyle,
					)),
					opts.Child(ctrl.Button(
						opts.Text("All"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							appModel.CurrentListName = model.ShowAllItems
						}),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Active"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							appModel.CurrentListName = model.ShowActiveItems
						}),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Completed"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							appModel.CurrentListName = model.ShowCompletedItems
						}),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Clear Completed"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							appModel.ClearCompleted()
							fmt.Println("Clear Completed Clicked")
						}),
					)),
				)),
			)),
		)
	}
}
