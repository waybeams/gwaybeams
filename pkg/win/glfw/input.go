package glfw

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/spec"
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
			g.bubbleOn(lastTarget, events.New(events.Exited, lastTarget, nil))
		}

		if target.IsFocusable() {
			cursorName := glfw.HandCursor
			if target.IsText() || target.IsTextInput() {
				cursorName = glfw.IBeamCursor
			}
			g.source.SetCursorByName(cursorName)

			g.bubbleOn(target, events.New(events.Entered, target, nil))
		} else {
			g.source.SetCursorByName(glfw.ArrowCursor)
		}
	}

	if target != nil {
		g.bubbleOn(target, events.New(events.Moved, target, nil))
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
			g.bubbleOn(lastMoveTarget, events.New(events.Pressed, lastMoveTarget, payload))
		} else if action == glfw.Release {
			g.bubbleOn(lastMoveTarget, events.New(events.Released, lastMoveTarget, payload))
			g.bubbleOn(lastMoveTarget, events.New(events.Clicked, lastMoveTarget, payload))
		}
	} else {
		g.focusSpec(nil)
	}
}

func (g *GlfwInput) focusSpec(s spec.ReadWriter) {
	lastFocused := s.FocusedSpec()
	if lastFocused != nil && lastFocused != s {
		lastFocused.SetFocusedSpec(nil)
		g.bubbleOn(lastFocused, events.New(events.Blurred, lastFocused, s))
		g.lastFocused = nil
	}
	if s != nil {
		s.SetFocusedSpec(s)
		g.bubbleOn(s, events.New(events.Focused, s, lastFocused))
		g.lastFocused = s
	}

}

func (g *GlfwInput) onCharHandler(char rune) {
	if g.lastRoot == nil {
		return
	}
	focused := g.lastFocused
	if focused != nil && focused.IsTextInput() {
		g.bubbleOn(focused, events.New(events.CharEntered, focused, string(char)))
	}
}

func (g *GlfwInput) onKeyHandler(key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if g.lastRoot == nil {
		return
	}
	focused := g.lastFocused
	if focused != nil && focused.IsTextInput() {
		g.bubbleOn(focused, events.New(events.KeyEntered, focused, key))
		if key == glfw.KeyEnter && action == glfw.Release {
			g.bubbleOn(focused, events.New(events.EnterKeyReleased, focused, key))
		}
	}
}

func (g *GlfwInput) bubbleOn(s spec.ReadWriter, event events.Event) {
	s.Bubble(event)
	// Also Emit an Invalidated event on the root node, but include the node
	// that triggered it.
	g.lastRoot.Emit(events.New(events.Invalidated, s, nil))
}

func NewGlfwInput(win spec.GestureSource) *GlfwInput {
	instance := &GlfwInput{source: win}
	win.SetCharCallback(instance.onCharHandler)
	win.SetKeyCallback(instance.onKeyHandler)
	win.SetMouseButtonCallback(instance.onMouseButtonHandler)
	return instance
}
