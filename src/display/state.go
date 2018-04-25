package display

// Component functions related to State management and handling.

func (c *Component) getStates() map[string][]ComponentOption {
	if c.states == nil {
		c.states = make(map[string][]ComponentOption)
	}
	return c.states
}

func (c *Component) AddState(name string, options ...ComponentOption) {
	states := c.getStates()
	states[name] = options
}

func (c *Component) SetState(name string, payloads ...interface{}) {
	c.currentState = name

	options, ok := c.states[name]
	if !ok {
		return
	}
	for _, option := range options {
		option(c)
	}
	// Probably too aggressive?
	c.Invalidate()
}

func (c *Component) HasState(name string) bool {
	_, ok := c.getStates()[name]
	return ok
}

func (c *Component) State() string {
	return c.currentState
}

func AddState(name string, options ...ComponentOption) ComponentOption {
	return func(d Displayable) error {
		d.AddState(name, options...)
		return nil
	}
}

func SetState(name string, payloads ...interface{}) ComponentOption {
	return func(d Displayable) error {
		d.SetState(name, payloads...)
		return nil
	}
}
