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
	model := d.Data().(*TodoItemModel)

	HBox(b, Children(func() {
		// Checkbox(b, Checked(model.CompletedAt.IsZero()), Text(model.Text))
		Checkbox(b, Text(model.Text))
		Button(b, TraitNames("destroy"), Text("X"))
	}))

}))

func CreateTodoHandler(model *TodoAppModel) EventHandler {
	return func(e Event) {
		t := e.Target().(*TextInputComponent)
		model.PushItem(t.Text())
		t.SetText("")
	}
}

func Create(c clock.Clock, model *TodoAppModel) (Displayable, error) {
	return NanoWindow(NewBuilderUsing(c), Children(func(b Builder) {
		// Create all of the application traits
		CreateTraits(b, model)

		Box(b, TraitNames("header"), Children(func() {
			Label(b, TraitNames("header-h1"), Text("todos"))
			TextInput(b, TraitNames("new-todo"), OnEnter(CreateTodoHandler(model)), Placeholder("What needs to be done?"))
		}))
		Box(b, TraitNames("main"), Children(func() {
			for _, entry := range model.AllItems() {
				Todo(b, Data(entry))
			}
		}))
	}))
}
