package display

// Component functions related to State management and handling.

const DefaultState = "default"

// CursorState is how a component responds to cursor movement.
// A cursor may be any pointing device including fingers.
type CursorState int

const (
	CursorActive = iota
	CursorHovered
	CursorPressed
	CursorDisabled
)

type Stateful interface {
	AddState(name string, options ...ComponentOption)
	ApplyCurrentState() error
	HasState(name string) bool
	SetState(name string)
	State() string

	// REMOVE THESE
	Selected() bool
	// REMOVE THESE
	SetCursorState(CursorState)
	// REMOVE THESE
	SetSelected(value bool)
}

func (c *Component) getStates() map[string][]ComponentOption {
	if c.states == nil {
		c.states = make(map[string][]ComponentOption)
	}
	return c.states
}

func (c *Component) AddState(name string, options ...ComponentOption) {
	if len(c.getStates()) == 0 {
		c.currentState = name
	}
	c.getStates()[name] = options
}

func (c *Component) SetState(name string) {
	c.currentState = name
	c.Invalidate()
}

// ApplyCurrentState is called from Builder.Push after a new component is
// instantiated.
func (c *Component) ApplyCurrentState() error {
	if !c.HasState(c.currentState) {
		return nil
	}
	options := c.getStates()[c.State()]
	for _, option := range options {
		err := option(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Component) HasState(name string) bool {
	_, ok := c.getStates()[name]
	return ok
}

func (c *Component) State() string {
	return c.currentState
}

func (c *Component) Selected() bool {
	return c.Model().Selected
}

func (c *Component) SetCursorState(state CursorState) {
	if c.cursorState != state {
		c.cursorState = state
		c.Invalidate()
	}
}

func (c *Component) CursorState() CursorState {
	return c.cursorState
}

func (c *Component) SetSelected(value bool) {
	c.Model().Selected = value
}
