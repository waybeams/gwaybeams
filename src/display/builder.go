package display

import (
	"clock"
	"events"
)

// BuilderOption is a configuration option for Builders.
type BuilderOption func(b Builder) error

// ComponentComposer is a composition function that components send to the
// Children() option when composing children using an anonymous function.
type ComponentComposer func(b Builder)

// Builder is a transient, short-lived helper that allow us to use a natural Go
// syntax to declare component composition.
// The builder should fall out of scope once the component tree is created.
type Builder interface {
	Clock() clock.Clock
	Destroy()
	LastError() error
	Listen()
	OnEnterFrame(handler EventHandler) Unsubscriber
	Peek() Displayable
	Push(d Displayable, options ...ComponentOption)
	Update(d Displayable) error
}

type BaseBuilder struct {
	clock       clock.Clock
	isDestroyed bool
	root        Displayable
	stack       Stack
	lastError   error
	emitter     Emitter
}

func (b *BaseBuilder) getStack() Stack {
	if b.stack == nil {
		b.stack = NewDisplayStack()
	}
	return b.stack
}

func (b *BaseBuilder) getEmitter() Emitter {
	if b.emitter == nil {
		b.emitter = NewEmitter()
	}
	return b.emitter
}

func (b *BaseBuilder) Clock() clock.Clock {
	return b.clock
}

func (b *BaseBuilder) OnEnterFrame(handler EventHandler) Unsubscriber {
	return b.getEmitter().On(events.EnterFrame, handler)
}

func (b *BaseBuilder) LastError() error {
	return b.lastError
}

// Listen() is the overall application loop and anything that wants to operate
// in the environment should listen for these events.
func (b *BaseBuilder) Listen() {
	var frameHandler = func() bool {
		root := b.root
		if root != nil {
			b.getEmitter().Emit(NewEvent(events.EnterFrame, root, nil))

			if root.ShouldRecompose() {
				root.RecomposeChildren()
			}
			root.Layout()
		}

		return b.isDestroyed
	}
	clock.OnFrame(frameHandler, DefaultFrameRate, b.Clock())
}

func (b *BaseBuilder) Destroy() {
	b.stack = nil
	b.lastError = nil
	b.root = nil
	b.isDestroyed = true
}

// Current returns the current entry in the Builder stack.
// This method only works while the component declarations are being processed.
func (b *BaseBuilder) Peek() Displayable {
	return b.getStack().Peek()
}

func (b *BaseBuilder) callComposeFunctionFor(d Displayable) {
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
func (b *BaseBuilder) Update(d Displayable) error {
	// NOTE: Brute force update here. Long term, look into creating the
	// secondary tree and diffing it against the existing tree, only
	// applying deltas where necessary.
	b.Push(d)
	return b.lastError
}

// Push accepts a new Displayable to place on the stack and processes the
// optional component composition function if one was provided.
func (b *BaseBuilder) Push(d Displayable, options ...ComponentOption) {
	stack := b.getStack()

	// Get the parent element if one exists
	parent := stack.Peek()

	if parent != nil {
		parent.AddChild(d)
	} else {
		// We're looking at the Root element
		if b.root == nil {
			b.root = d
		}
		d.SetBuilder(b)
	}

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

	// Push the element onto the stack
	stack.Push(d)

	// Process composition function to build children
	b.callComposeFunctionFor(d)

	// Pop the element off the stack
	stack.Pop()
}

// NewBuilder returns a clean builder instance.
func NewBuilder() *Updater {
	return &Updater{}
}

// NewBuilderUsing runs with provided clock instead of the real one.
// Mainly used by tests that need to provide a fake clock.
func NewBuilderUsing(clock clock.Clock) *Updater {
	return &Updater{clock: clock}
}
