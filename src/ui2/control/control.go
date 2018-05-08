package control

import "ui2/compose"

type Control struct {
	compose *compose.Composite
}

func New() *Control {
	return &Control{}
}

func AddChild(parent *Control, child *Control) int {
	return compose.AddChild(parent.compose, child.compose)
}

func ChildAt(parent *Control, index int) *Control {
	// Incorrect return type. How do I get from a Composite to a Control reference?
	// return compose.ChildAt(parent.compose, index)
	return nil
}
