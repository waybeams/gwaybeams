package display

import (
	"events"
)

const ToggleUnselected = "unselected"
const ToggleSelected = "selected"

func UnselectedToggleView(s Surface, d Displayable) error {
	// TODO(lbayes): Draw toggle button UNSELECTED state here
	return nil
}

func SelectedToggleView(s Surface, d Displayable) error {
	// TODO(lbayes): Draw toggle button SELECTED state here
	return nil
}

func setSelectedHandler(d Displayable) EventHandler {
	return toggleSelectedHandler(d, ToggleSelected)
}

func setUnSelectedHandler(d Displayable) EventHandler {
	return toggleSelectedHandler(d, ToggleUnselected)
}

func toggleSelectedHandler(d Displayable, state string) EventHandler {
	return func(e Event) {
		d.SetState(state)
		e.Cancel()
		// Refire the event from here:
		d.Bubble(NewEvent(events.Clicked, d, nil))
	}
}

var ToggleButton = NewComponentFactory("ToggleButton", NewComponent,
	OnState(ToggleUnselected, Children(func(b Builder, d Displayable) {
		Button(b, OnClick(setSelectedHandler(d)), View(UnselectedToggleView))
	})),
	OnState(ToggleSelected, Children(func(b Builder, d Displayable) {
		Button(b, OnClick(setUnSelectedHandler(d)), View(SelectedToggleView))
	})),
)
