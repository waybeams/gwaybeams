package display

import "events"

type InputController interface {
	Update()
}

type GestureWindow interface {
	GetCursorPos() (xpos, ypos float64)
}

type GlfwInput struct {
	lastMoveTarget Displayable
	root           Displayable
	window         GestureWindow
}

// Update should be called on every frame and will collect any pending
// changes from the configured glfw.Window object and then bubble as events
// into the appropriate nodes of the tree.
func (i *GlfwInput) Update() {
	i.UpdateCursorPos()
}

func (i *GlfwInput) UpdateCursorPos() {
	xpos, ypos := i.window.GetCursorPos()
	target := CursorPick(i.root, xpos, ypos)
	lastTarget := i.lastMoveTarget

	if lastTarget != target {
		if lastTarget != nil {
			lastTarget.Bubble(NewEvent(events.Exited, lastTarget, nil))
		}
		target.Bubble(NewEvent(events.Entered, target, nil))
	}
	i.lastMoveTarget = target

}

func NewGlfwInput(root Displayable, win GestureWindow) *GlfwInput {
	return &GlfwInput{root: root, window: win}
}
