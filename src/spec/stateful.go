package spec

type StatefulReader interface {
	HasState(name string) bool
	OnState(name string, options ...Option)
	OptionsForState(stateName string) []Option
	State() string
}

type StatefulWriter interface {
	ApplyCurrentState()
	SetState(name string)
}

type StatefulReadWriter interface {
	StatefulReader
	StatefulWriter
}

func (c *Spec) getStates() map[string][]Option {
	if c.states == nil {
		c.states = make(map[string][]Option)
	}
	return c.states
}

func (c *Spec) OnState(name string, options ...Option) {
	if len(c.getStates()) == 0 {
		c.currentState = name
	}
	c.getStates()[name] = options
}

func (c *Spec) SetState(name string) {
	c.currentState = name
}

func (c *Spec) ApplyCurrentState() {
	options := c.OptionsForState(c.State())
	for _, option := range options {
		option(c)
	}
}

func (c *Spec) OptionsForState(stateName string) []Option {
	// TODO(lbayes): This exposes a risk of double-subscribing event handlers.
	// We cannot currently call UnsubAll() in here, because valid handlers may
	// have been added when the rest of the props were applied.
	if c.HasState(stateName) {
		return c.getStates()[stateName]
	}

	return nil
}

func (c *Spec) HasState(name string) bool {
	_, ok := c.getStates()[name]
	return ok
}

func (c *Spec) State() string {
	return c.currentState
}
