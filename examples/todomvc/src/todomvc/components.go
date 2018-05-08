package todomvc

import (
	"events"
	. "ui"
	"ui/control"
	"ui/context"
	. "ui/controls"
	. "ui/opts"
)

// This is the control definition for a new entry.
var Todo = control.Define("Todo", control.New, Children(func(c Context, d Displayable) {
	// model := d.Data().(*TodoItemModel)

	HBox(c, Children(func() {
		// Checkbox(c, Selected(!model.CompletedAt.IsZero()), Text(model.Text))
		Button(c, Visible(false), Text("X"))
	}))
}))

func todoCreateHandler(model *TodoAppModel) events.EventHandler {
	return func(e events.Event) {
		t := e.Target().(Displayable)
		model.PushItem(t.Text())
		t.SetText("")
	}
}

func Create(model *TodoAppModel, opts ...context.Option) Displayable {
	return NanoWindow(context.New(opts...),
		BgColor(0xffffffff),
		Padding(10),
		FontFace("Roboto"),
		FontSize(24),
		FontColor(0x333333ff),
		HAlign(AlignCenter),
		Children(func(c Context) {
			VBox(c,
				Width(550),
				HAlign(AlignCenter),
				StrokeColor(0xff0000ff),
				FlexHeight(1),
				Gutter(20),
				Padding(20),
				Children(func() {
					Label(c,
						FontFace("Roboto-Light"),
						FontSize(100),
						Padding(10),
						FontColor(0xaf2f2f26),
						Text("todos"),
					)
					TextInput(c,
						BgColor(0xffffffff),
						StrokeColor(0xccccccff),
						StrokeSize(1),
						FontColor(0x222222ff),
						Height(50),
						FlexWidth(1),
						Padding(10),
						Placeholder("What needs to be done?"),
						OnEnterKey(todoCreateHandler(model)))
					/*
						Box(c, TraitNames("main"), FlexWidth(1), Children(func() {
							for _, entry := range model.AllItems() {
								Todo(c, Data(entry))
							}
						}))
							HBox(c, Visible(len(model.AllItems()) > 0), Height(65), FlexWidth(1), Children(func() {
								Label(c, TraitNames("remaining-label"), Text(model.PendingLabel()))
								Spacer(c, FlexWidth(1))
								Box(c, Children(func() {
									// filterSelection := model.FilterSelection()
									// Following is a good example where a Composition function will run over and over, and
									// this execution will always result in the expected output without relying on
									// accumulated, hidden control state.
									// ToggleButton(c, Selected(filterSelection == "All"), Text("All"))
									// ToggleButton(c, Selected(filterSelection == "Active"), Text("Active"))
									// ToggleButton(c, Selected(filterSelection == "Completed"), Text("Completed"))
								}))
								Spacer(c, FlexWidth(1))
								Button(c, TraitNames("clear-completed-button"), Text("Clear completed"))
							}))
					*/
				}))
		}))
}
