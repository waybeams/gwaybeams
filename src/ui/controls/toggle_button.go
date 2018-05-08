package controls

import (
	"ui/comp"
	"events"
	"ui/opts"
	"ui"
)

const ToggleUnselected = "unselected"
const ToggleSelected = "selected"

func UnselectedToggleView(s ui.Surface, d ui.Displayable) error {
	// TODO(lbayes): Draw toggle button UNSELECTED state here
	return nil
}

func SelectedToggleView(s ui.Surface, d ui.Displayable) error {
	// TODO(lbayes): Draw toggle button SELECTED state here
	return nil
}

func setSelectedHandler(d ui.Displayable) events.EventHandler {
	return toggleSelectedHandler(d, ToggleSelected)
}

func setUnSelectedHandler(d ui.Displayable) events.EventHandler {
	return toggleSelectedHandler(d, ToggleUnselected)
}

func toggleSelectedHandler(d ui.Displayable, state string) events.EventHandler {
	return func(e events.Event) {
		d.SetState(state)
		e.Cancel()
		// Refire the event from here:
		d.Bubble(events.New(events.Clicked, d, nil))
	}
}

var ToggleButton = comp.Define("ToggleButton", comp.New,
	opts.OnState(ToggleUnselected, opts.Children(func(c ui.Context, d ui.Displayable) {
		Button(c, opts.OnClick(setSelectedHandler(d)), opts.View(UnselectedToggleView))
	})),
	opts.OnState(ToggleSelected, opts.Children(func(c ui.Context, d ui.Displayable) {
		Button(c, opts.OnClick(setUnSelectedHandler(d)), opts.View(SelectedToggleView))
	})),
)
