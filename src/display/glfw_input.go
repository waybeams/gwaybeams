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
	SetCharCallback(callback CharCallback) Unsubscriber
	SetMouseButtonCallback(callback MouseButtonCallback) Unsubscriber
}

type MouseEventPayload struct {
	Button   glfw.MouseButton
	Action   glfw.Action
	Modifier glfw.ModifierKey
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

func (i *GlfwInput) onMouseButtonHandler(button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	lastMoveTarget := i.lastMoveTarget
	if button == glfw.MouseButton1 && lastMoveTarget != nil && lastMoveTarget.IsFocusable() {
		payload := &MouseEventPayload{
			Button:   button,
			Action:   action,
			Modifier: mod,
		}

		if action == glfw.Press {
			lastMoveTarget.Focus()
			lastMoveTarget.Bubble(NewEvent(events.Pressed, lastMoveTarget, payload))
		} else if action == glfw.Release {
			lastMoveTarget.Bubble(NewEvent(events.Released, lastMoveTarget, payload))
			lastMoveTarget.Bubble(NewEvent(events.Clicked, lastMoveTarget, payload))
		}
	} else {
		if i.root.FocusedChild() != nil {
			i.root.FocusedChild().Blur()
		}
	}
}

func (i *GlfwInput) onCharHandler(char rune) {
	focused := i.root.FocusedChild()
	if focused != nil && focused.IsTextInput() {
		focused.Bubble(NewEvent(events.CharEntered, focused, char))
	}
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
			if target.IsText() || target.IsTextInput() {
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
	instance := &GlfwInput{root: root, source: win}
	win.SetCharCallback(instance.onCharHandler)
	win.SetMouseButtonCallback(instance.onMouseButtonHandler)
	return instance
}
