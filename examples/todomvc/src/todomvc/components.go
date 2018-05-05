package todomvc

import (
	"clock"
	"component"
	. "controls"
	"ctx"
	"events"
	. "opts"
	. "ui"
)

/*
func CreateTraits(c Context, model *TodoAppModel) {
	// Make font loading into a Trait or Node?

	Trait(c, ".header-h1",
		FlexWidth(1),
		FlexHeight(1),
		FontFace("Roboto Light"),
		FontSize(100),
		HAlign(AlignCenter),
		FontColor(0xaf2f2f26)) // color.RGBA{175, 47, 47, 0.15}

	Trait(c, ".new-todo",
		Width(400),
		Height(48),
		Padding(16),
		// BoxShadow("inset", 0, -2, 1, 0xffffff0f))
	)

	Trait(c, ".destroy",
		Visible(false),
	)

	Trait(c, "#footer",
		Visible(len(model.AllItems()) > 0))
}
*/

// This is the component definition for a new entry.
var Todo = component.Define("Todo", component.New, Children(func(c Context, d Displayable) {
	// model := d.Data().(*TodoItemModel)

	HBox(c, Children(func() {
		// Checkbox(c, Selected(!model.CompletedAt.IsZero()), Text(model.Text))
		Button(c, Visible(false), Text("X"))
	}))
}))

func createTodoHandler(model *TodoAppModel) events.EventHandler {
	return func(e events.Event) {
		t := e.Target().(Displayable)
		model.PushItem(t.Text())
		t.SetText("")
	}
}

func Create(c clock.Clock, model *TodoAppModel) Displayable {
	return NanoWindow(ctx.New(ctx.Clock(c)), HAlign(AlignCenter), Children(func(c Context) {
		// Create all of the application traits
		// CreateTraits(c, model)
		HeaderBag := Bag(
			FlexWidth(1),
			FlexHeight(1),
			FontFace("Roboto Light"),
			FontSize(100),
			HAlign(AlignCenter),
			FontColor(0xaf2f2f26),
		) // color.RGBA{175, 47, 47, 0.15}

		NewTodoBag := Bag(
			Width(400),
			Height(48),
			Padding(16),
			// BoxShadow("inset", 0, -2, 1, 0xffffff0f))
		)

		FooterBag := Bag(
			Visible(len(model.AllItems()) > 0),
		)

		VBox(c, Width(550), FlexHeight(1), PaddingTop(100), Children(func() {
			Label(c, HeaderBag, Height(100), FlexWidth(1), Text("todos"))
			TextInput(c, NewTodoBag, Height(100), FlexWidth(1), OnEnter(createTodoHandler(model)), Placeholder("What needs to be done?"))
			Box(c, TraitNames("main"), FlexWidth(1), Children(func() {
				for _, entry := range model.AllItems() {
					Todo(c, Data(entry))
				}
			}))
			HBox(c, FooterBag, Height(65), FlexWidth(1), Children(func() {
				Label(c, TraitNames("remaining-label"), Text(model.PendingLabel()))
				Spacer(c, FlexWidth(1))
				Box(c, Children(func() {
					// filterSelection := model.FilterSelection()
					// Following is a good example where a Composition function will run over and over, and
					// this execution will always result in the expected output without relying on
					// accumulated, hidden component state.
					// ToggleButton(c, Selected(filterSelection == "All"), Text("All"))
					// ToggleButton(c, Selected(filterSelection == "Active"), Text("Active"))
					// ToggleButton(c, Selected(filterSelection == "Completed"), Text("Completed"))
				}))
				Spacer(c, FlexWidth(1))
				Button(c, TraitNames("clear-completed-button"), Text("Clear completed"))
			}))
		}))
	}))
}
