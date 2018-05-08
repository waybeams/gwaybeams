package control

import "ui"

func (c *Control) getStates() map[string][]ui.Option {
	if c.states == nil {
		c.states = make(map[string][]ui.Option)
	}
	return c.states
}

func (c *Control) OnState(name string, options ...ui.Option) {
	if len(c.getStates()) == 0 {
		c.currentState = name
	}
	c.getStates()[name] = options
}

func (c *Control) SetState(name string) {
	c.currentState = name
	c.Invalidate()
}

// ApplyCurrentState is called from Builder.Push after a new control is
// instantiated.
func (c *Control) ApplyCurrentState() {
	options := c.OptionsForState(c.State())
	for _, option := range options {
		option(c)
	}
}

func (c *Control) OptionsForState(stateName string) []ui.Option {
	// TODO(lbayes): This exposes a risk of double-subscribing event handlers.
	// We cannot currently call UnsubAll() in here, because valid handlers may
	// have been added when the rest of the props were applied.
	if c.HasState(stateName) {
		return c.getStates()[stateName]
	}

	return nil
}

func (c *Control) HasState(name string) bool {
	_, ok := c.getStates()[name]
	return ok
}

func (c *Control) State() string {
	return c.currentState
}
