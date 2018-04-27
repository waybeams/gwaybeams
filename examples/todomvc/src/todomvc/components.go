package todomvc

import (
	"clock"
	. "display"
)

func CreateTraits(b Builder, model *TodoAppModel) {
	// Make font loading into a Trait or Node?

	Trait(b, ".header-h1",
		FlexWidth(1),
		FlexHeight(1),
		FontFace("Roboto Light"),
		FontSize(100),
		HAlign(AlignCenter),
		FontColor(0xaf2f2f26)) // color.RGBA{175, 47, 47, 0.15}

	Trait(b, ".new-todo",
		Width(400),
		Height(48),
		Padding(16),
		// BoxShadow("inset", 0, -2, 1, 0xffffff0f))
	)

	Trait(b, ".destroy",
		Visible(false),
	)

	Trait(b, "#footer",
		Visible(len(model.AllItems()) > 0))
}

// This is the component definition for a new entry.
var Todo = NewComponentFactory("Todo", NewComponent, Children(func(b Builder, d Displayable) {
	// model := d.Data().(*TodoItemModel)

	HBox(b, Children(func() {
		// Checkbox(b, Selected(!model.CompletedAt.IsZero()), Text(model.Text))
		Button(b, TraitNames("destroy"), Text("X"))
	}))
}))

func createTodoHandler(model *TodoAppModel) EventHandler {
	return func(e Event) {
		t := e.Target().(*TextInputComponent)
		model.PushItem(t.Text())
		t.SetText("")
	}
}

func Create(c clock.Clock, model *TodoAppModel) (Displayable, error) {
	return NanoWindow(NewBuilderUsing(c), HAlign(AlignCenter), Children(func(b Builder) {
		// Create all of the application traits
		CreateTraits(b, model)

		VBox(b, Width(550), FlexHeight(1), PaddingTop(100), Children(func() {
			Label(b, TraitNames("header-h1"), Height(100), FlexWidth(1), Text("todos"))
			TextInput(b, TraitNames("new-todo"), Height(100), FlexWidth(1), OnEnter(createTodoHandler(model)), Placeholder("What needs to be done?"))
			Box(b, TraitNames("main"), FlexWidth(1), Children(func() {
				for _, entry := range model.AllItems() {
					Todo(b, Data(entry))
				}
			}))
			HBox(b, TraitNames("footer"), Height(65), FlexWidth(1), Children(func() {
				Label(b, TraitNames("remaining-label"), Text(model.PendingLabel()))
				Spacer(b, FlexWidth(1))
				RadioGroup(b, Children(func() {
					// filterSelection := model.FilterSelection()
					// Following is a good example where a Composition function will run over and over, and
					// this execution will always result in the expected output without relying on
					// accumulated, hidden component state.
					// ToggleButton(b, Selected(filterSelection == "All"), Text("All"))
					// ToggleButton(b, Selected(filterSelection == "Active"), Text("Active"))
					// ToggleButton(b, Selected(filterSelection == "Completed"), Text("Completed"))
				}))
				Spacer(b, FlexWidth(1))
				Button(b, TraitNames("clear-completed-button"), Text("Clear completed"))
			}))
		}))
	}))
}
