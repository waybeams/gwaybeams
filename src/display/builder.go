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
	childrenTypeMap ChildrenTypeMap
	clock           clock.Clock
	emitter         Emitter
	isDestroyed     bool
	lastError       error
	root            Displayable
	stack           Stack
}

func (b *BaseBuilder) getEmitter() Emitter {
	if b.emitter == nil {
		b.emitter = NewEmitter()
	}
	return b.emitter
}

func (b *BaseBuilder) getStack() Stack {
	if b.stack == nil {
		b.stack = NewDisplayStack()
	}
	return b.stack
}

func (b *BaseBuilder) OnEnterFrame(handler EventHandler) Unsubscriber {
	return b.getEmitter().On(events.EnterFrame, handler)
}

func (b *BaseBuilder) Clock() clock.Clock {
	if b.clock == nil {
		// Go get the clock from the provided Root component's builder
		b.clock = clock.New()
	}
	return b.clock
}

func (b *BaseBuilder) LastError() error {
	return b.lastError
}

func (b *BaseBuilder) Peek() Displayable {
	return b.getStack().Peek()
}

func (b *BaseBuilder) getExistingChild(d Displayable, parent Displayable) Displayable {
	updateableChildren := parent.updateableChildren()
	typeName := d.TypeName()
	kids := updateableChildren[typeName]
	var result Displayable
	if kids != nil {
		if len(kids) > 1 {
			result = kids[0]
			updateableChildren[typeName] = kids[1:]
		} else {
			result = kids[0]
			delete(updateableChildren, typeName)
		}
		parent.setUpdateableChildren(updateableChildren)
	}
	return result
}

func (b *BaseBuilder) Update(d Displayable) error {
	b.Push(d)
	// d.Layout()
	return b.lastError
}

func (b *BaseBuilder) Push(d Displayable, options ...ComponentOption) {
	stack := b.getStack()

	// Get the parent element if one exists
	parent := stack.Peek()

	if parent != nil {
		existingChild := b.getExistingChild(d, parent)
		if existingChild == nil {
			parent.AddChild(d)
		} else {
			d = existingChild
		}

		// Clear the composer function before triggering for what might be
		// a second time
		d.Composer(nil)
	} else {
		b.root = d
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
	if d.State() == "" && d.HasState("default") {
		d.SetState("default")
	}

	// Push the element onto the stack
	stack.Push(d)

	// Create the children type map for this parent
	d.setUpdateableChildren(b.createChildrenMapFor(d))

	// Process composition function to build children
	b.callComposeFunctionFor(d)

	// Remove the children that haven't been used.
	b.clearChildrenFromTypeMap(d)

	// Pop the element off the stack
	stack.Pop()

	if !stack.HasNext() {
		b.root.Layout()
	}
}

func (b *BaseBuilder) clearChildrenFromTypeMap(d Displayable) {
	types := d.updateableChildren()

	for _, entries := range types {
		for _, entry := range entries {
			entry.Parent().RemoveChild(entry)
		}
	}

	// Remove the children type mape for this parent
	b.childrenTypeMap = nil
}

func (b *BaseBuilder) createChildrenMapFor(d Displayable) ChildrenTypeMap {
	types := make(map[string][]Displayable)
	for _, child := range d.Children() {
		typeName := child.TypeName()
		entry := types[typeName]
		if entry == nil {
			types[typeName] = []Displayable{child}
		} else {
			types[typeName] = append(entry, child)
		}
	}
	return types
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

func (b *BaseBuilder) Destroy() {
	b.stack = nil
	b.lastError = nil
	b.root = nil
	b.isDestroyed = true
}

func (b *BaseBuilder) Listen() {
	var frameHandler = func() bool {
		root := b.root
		if root != nil {
			b.getEmitter().Emit(NewEvent(events.EnterFrame, root, nil))

			if root.ShouldRecompose() {
				root.RecomposeChildren()
			}
			// root.Layout()
		}

		return b.isDestroyed
	}
	clock.OnFrame(frameHandler, DefaultFrameRate, b.Clock())
}

// NewBuilder returns a clean builder instance.
func NewBuilder() *BaseBuilder {
	return &BaseBuilder{}
}

// NewBuilderUsing runs with provided clock instead of the real one.
// Mainly used by tests that need to provide a fake clock.
func NewBuilderUsing(clock clock.Clock) *BaseBuilder {
	return &BaseBuilder{clock: clock}
}
