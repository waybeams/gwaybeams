package display

import (
	"events"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type InputController interface {
	Update()
}

type GestureSource interface {
	GetCursorPos() (xpos, ypos float64)
	SetCursorByName(name glfw.StandardCursor)
}

type GlfwInput struct {
	lastMoveTarget Displayable
	root           Displayable
	source         GestureSource
	lastXpos       float64
	lastYpos       float64
}

// Update should be called on every frame and will collect any pending
// changes from the configured glfw.Window object and then bubble as events
// into the appropriate nodes of the tree.
func (i *GlfwInput) Update() {
	i.UpdateCursorPos()
}

func (i *GlfwInput) UpdateCursorPos() {
	xpos, ypos := i.source.GetCursorPos()
	if i.lastXpos == xpos && i.lastYpos == ypos {
		return
	}
	i.lastXpos = xpos
	i.lastYpos = ypos

	target := CoordToComponent(i.root, xpos, ypos)
	lastTarget := i.lastMoveTarget

	if lastTarget != target {
		if lastTarget != nil {
			lastTarget.Bubble(NewEvent(events.Exited, lastTarget, nil))
		}

		if target.IsFocusable() {
			cursorName := glfw.HandCursor
			if target.IsText() {
				cursorName = glfw.IBeamCursor
			}
			i.source.SetCursorByName(cursorName)

			target.Bubble(NewEvent(events.Entered, target, nil))
		} else {
			i.source.SetCursorByName(glfw.ArrowCursor)
		}
	}

	if target != nil {
		target.Bubble(NewEvent(events.Moved, target, nil))
	}
	i.lastMoveTarget = target
}

func NewGlfwInput(root Displayable, win GestureSource) *GlfwInput {
	return &GlfwInput{root: root, source: win}
}
