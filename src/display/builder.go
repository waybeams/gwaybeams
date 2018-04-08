package display

// BuilderOption is a configuration option for Builders.
type BuilderOption func(b Builder) error

// ComponentComposer is a composition function that components send to the
// Children() option when composing children using an anonymous function.
type ComponentComposer func(b Builder)

// Builder is a transient, short-lived helper that allow us to use a natural Go
// syntax to declare component composition.
// The builder should fall out of scope once the component tree is created.
type Builder interface {
	UpdateChildren(d Displayable) error
	Push(d Displayable, options ...ComponentOption)
	Peek() Displayable
	Destroy()
	LastError() error
}

type builder struct {
	stack     Stack
	lastError error
}

func (b *builder) LastError() error {
	return b.lastError
}

func (b *builder) Destroy() {
	b.stack = nil
	b.lastError = nil
}

func (b *builder) getStack() Stack {
	if b.stack == nil {
		b.stack = NewDisplayStack()
	}
	return b.stack
}

// Current returns the current entry in the Builder stack.
// This method only works while the component declarations are being processed.
func (b *builder) Peek() Displayable {
	return b.getStack().Peek()
}

func (b *builder) callComposeFunctionFor(d Displayable) {
	composeEmpty := d.GetComposeEmpty()
	if composeEmpty != nil {
		composeEmpty()
		return
	}
	composeWithBuilder := d.GetComposeWithBuilder()
	if composeWithBuilder != nil {
		composeWithBuilder(b)
		return
	}
	composeWithComponent := d.GetComposeWithComponent()
	if composeWithComponent != nil {
		composeWithComponent(d)
		return
	}
	composeWithBuilderAndComponent := d.GetComposeWithBuilderAndComponent()
	if composeWithBuilderAndComponent != nil {
		composeWithBuilderAndComponent(b, d)
		return
	}
}

// Update will re-render the provided component's children
func (b *builder) UpdateChildren(d Displayable) error {
	// NOTE: Brute force update here. Long term, look into creating the
	// secondary tree and diffing it against the existing tree, only
	// applying deltas where necessary.
	d.RemoveAllChildren()
	b.Push(d)
	return b.lastError
}

// Push accepts a new Displayable to place on the stack and processes the
// optional component composition function if one was provided.
func (b *builder) Push(d Displayable, options ...ComponentOption) {
	stack := b.getStack()

	// Get the parent element if one exists
	parent := stack.Peek()

	if parent != nil {
		parent.AddChild(d)
	}

	// Push the element onto the stack
	stack.Push(d)

	// One of these options might be a Children(func()), which will recurse
	// back into this Push function.
	for _, option := range options {
		err := option(d)
		if err != nil {
			// If an option error is found, bail with it, for now.
			b.lastError = err
			return
		}
	}

	// Process composition function to build children
	b.callComposeFunctionFor(d)

	// Pop the element off the stack
	stack.Pop()
}

// NewBuilder returns a clean builder instance.
func NewBuilder() Builder {
	return &builder{}
}
