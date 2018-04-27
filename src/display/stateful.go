package display

type Stateful interface {
	AddState(name string, options ...ComponentOption)
	ApplyCurrentState() error
	HasState(name string) bool
	SetState(name string)
	State() string
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
