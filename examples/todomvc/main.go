package main

import (
	"builder"
	ctrl "controls"
	"events"
	"fmt"
	"opts"
	"runtime"
	"spec"
	"surface/nano"
	"time"
)

func init() {
	runtime.LockOSThread()
}

type TodoItemModel struct {
	CreatedAt   time.Time
	CompletedAt time.Time
	Description string
	collection  *TodoAppModel
}

func (t *TodoItemModel) SetDescription(text string) {
	t.Description = text
}

func (t *TodoItemModel) Delete() {
	t.collection.DeleteItem(t)
}

func (t *TodoItemModel) ToggleCompleted() {
	if t.CompletedAt.IsZero() {
		t.CompletedAt = time.Now()
	} else {
		t.MarkActive()
	}
}

func (t *TodoItemModel) MarkActive() {
	t.CompletedAt = time.Time{}
}

const ShowAllItems = "ShowAllItems"
const ShowActiveItems = "ShowActiveItems"
const ShowCompletedItems = "ShowCompletedItems"

type TodoAppModel struct {
	CurrentListName string
	allItems        []*TodoItemModel
	enteredText     string
}

func (t *TodoAppModel) ClearCompleted() {
	t.allItems = t.ActiveItems()
}

func (t *TodoAppModel) DeleteItem(deletedItem *TodoItemModel) {
	result := []*TodoItemModel{}
	// Splice instead of rebuild!
	for _, item := range t.allItems {
		if item != deletedItem {
			result = append(result, item)
		}
	}
	t.allItems = result
}

func (t *TodoAppModel) EnteredText() string {
	return t.enteredText
}

func (t *TodoAppModel) UpdateEnteredText(str string) {
	t.enteredText = str
}

func (t *TodoAppModel) CreateItemFromEnteredText() {
	t.CreateItem(t.EnteredText())
	t.enteredText = ""
}

func (t *TodoAppModel) CreateItem(desc string) {
	t.allItems = append(t.allItems, &TodoItemModel{
		Description: desc,
		CreatedAt:   time.Now(),
		collection:  t,
	})
}

func (t *TodoAppModel) CurrentItems() []*TodoItemModel {
	switch t.CurrentListName {
	case ShowAllItems:
		return t.allItems
	case ShowCompletedItems:
		return t.CompletedItems()
	default:
		return t.ActiveItems()
	}
}

func (t *TodoAppModel) CompletedItems() []*TodoItemModel {
	result := []*TodoItemModel{}
	for _, item := range t.allItems {
		if !item.CompletedAt.IsZero() {
			result = append(result, item)
		}
	}
	return result
}

func (t *TodoAppModel) ActiveItems() []*TodoItemModel {
	result := []*TodoItemModel{}
	for _, item := range t.allItems {
		if item.CompletedAt.IsZero() {
			result = append(result, item)
		}
	}
	return result
}

func TodoItemSpec(model *TodoItemModel, index int) spec.ReadWriter {
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

func todoModelsToSpecs(items []*TodoItemModel) []spec.ReadWriter {
	result := []spec.ReadWriter{}
	for index, itemModel := range items {
		result = append(result, TodoItemSpec(itemModel, index))
	}
	return result
}

func CreateAppRenderer(model *TodoAppModel) func() spec.ReadWriter {
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
					opts.Text(model.EnteredText()),
					opts.Padding(18),
					opts.FontSize(36),
					opts.FlexWidth(1),
					opts.BgColor(0xccccccff),
					opts.BindStringPayloadTo(events.TextChanged, model.UpdateEnteredText),
					opts.OnEnterKey(events.Empty(model.CreateItemFromEnteredText)),
				)),
				opts.Child(ctrl.VBox(
					opts.Key("Todo Items"),
					opts.MinHeight(300),
					opts.FlexWidth(1),
					opts.BgColor(0xeeeeeeff),
					opts.Children(todoModelsToSpecs(model.CurrentItems())),
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
						opts.Text(fmt.Sprintf("%d items", len(model.CurrentItems()))),
						buttonStyle,
					)),
					opts.Child(ctrl.Button(
						opts.Text("All"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							model.CurrentListName = ShowAllItems
						}),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Active"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							model.CurrentListName = ShowActiveItems
						}),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Completed"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							model.CurrentListName = ShowCompletedItems
						}),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Clear Completed"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							model.ClearCompleted()
							fmt.Println("Clear Completed Clicked")
						}),
					)),
				)),
			)),
		)
	}
}

func main() {
	model := &TodoAppModel{}
	model.CreateItem("Item One")
	model.CreateItem("Item Two")
	model.CreateItem("Item Three")
	model.CreateItem("Item Four")
	model.CreateItem("Item Five")
	model.CreateItem("Item Six")

	// Create the Application specification.
	renderer := CreateAppRenderer(model)

	// Create and configure the NanoSurface.
	surface := nano.New(
		nano.Font("Roboto", "./third_party/fonts/Roboto/Roboto-Regular.ttf"),
		nano.Font("Roboto Light", "./third_party/fonts/Roboto/Roboto-Light.ttf"),
	)

	// Create and configure the Builder.
	build := builder.New(
		builder.Surface(surface),
		builder.Factory(renderer),
	)

	// Loop until exit.
	build.Listen()
}
