package glfw

import (
	"events"
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
	"spec"
)

type MouseEventPayload struct {
	Button   glfw.MouseButton
	Action   glfw.Action
	Modifier glfw.ModifierKey
}

type GlfwInput struct {
	lastMoveTarget spec.ReadWriter
	source         spec.GestureSource
	lastXpos       float64
	lastYpos       float64
	lastRoot       spec.ReadWriter
	lastFocused    spec.ReadWriter
}

// Update should be called on every frame and will collect any pending
// changes from the configured glfw.Window object and then bubble as events
// into the appropriate nodes of the tree.
func (g *GlfwInput) Update(root spec.ReadWriter) {
	g.lastRoot = root

	xpos, ypos := g.source.GetCursorPos()
	if g.lastXpos == xpos && g.lastYpos == ypos {
		return
	}
	g.lastXpos = xpos
	g.lastYpos = ypos

	target := spec.CoordToControl(root, xpos, ypos)
	lastTarget := g.lastMoveTarget

	if lastTarget != target {
		if lastTarget != nil {
			lastTarget.Bubble(events.New(events.Exited, lastTarget, nil))
		}

		if target.IsFocusable() {
			cursorName := glfw.HandCursor
			if target.IsText() || target.IsTextInput() {
				cursorName = glfw.IBeamCursor
			}
			g.source.SetCursorByName(cursorName)

			target.Bubble(events.New(events.Entered, target, nil))
		} else {
			g.source.SetCursorByName(glfw.ArrowCursor)
		}
	}

	if target != nil {
		target.Bubble(events.New(events.Moved, target, nil))
	}
	g.lastMoveTarget = target
}

func (g *GlfwInput) onMouseButtonHandler(button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if g.lastRoot == nil {
		return
	}

	lastMoveTarget := g.lastMoveTarget
	if button == glfw.MouseButton1 && lastMoveTarget != nil && lastMoveTarget.IsFocusable() {
		payload := &MouseEventPayload{
			Button:   button,
			Action:   action,
			Modifier: mod,
		}

		if action == glfw.Press {
			g.focusSpec(lastMoveTarget)
			lastMoveTarget.Bubble(events.New(events.Pressed, lastMoveTarget, payload))
		} else if action == glfw.Release {
			lastMoveTarget.Bubble(events.New(events.Released, lastMoveTarget, payload))
			lastMoveTarget.Bubble(events.New(events.Clicked, lastMoveTarget, payload))
		}
	} else {
		g.focusSpec(nil)
	}
}

func (g *GlfwInput) focusSpec(s spec.ReadWriter) {
	if g.lastFocused != nil && g.lastFocused != s {
		s.Blur()
		s.Bubble(events.New(events.Blurred, g.lastFocused, s))
		g.lastFocused = nil
	}
	if s != nil {
		s.Focus()
		s.Bubble(events.New(events.Focused, s, g.lastFocused))
		g.lastFocused = s
	}

}

func (g *GlfwInput) onCharHandler(char rune) {
	fmt.Println("onCharHandler with:", string(char))
	if g.lastRoot == nil {
		return
	}
	focused := g.lastFocused
	if focused != nil && focused.IsTextInput() {
		focused.Bubble(events.New(events.CharEntered, focused, char))
	}
}

func NewGlfwInput(win spec.GestureSource) *GlfwInput {
	instance := &GlfwInput{source: win}
	win.SetCharCallback(instance.onCharHandler)
	win.SetMouseButtonCallback(instance.onMouseButtonHandler)
	return instance
}
