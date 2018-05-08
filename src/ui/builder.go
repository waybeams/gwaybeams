package ui

import (
	"events"
)

// BuilderOption is a configuration option for Builders.
type BuilderOption func(b Builder) error

// ComponentComposer is a composition function that controls send to the
// Children() option when composing children using an anonymous function.
type ComponentComposer func(b Builder)

// Builder is a transient, short-lived helper that allow us to use a natural Go
// syntax to declare control composition.
// The builder should fall out of scope once the control tree is created.
type Builder interface {
	Destroy()
	Peek() Displayable
	Push(d Displayable, options ...Option)
	Update(d Displayable)
}

type BaseBuilder struct {
	childrenTypeMap ChildrenTypeMap
	isDestroyed     bool
	root            Displayable
	stack           Stack
}

func (b *BaseBuilder) getStack() Stack {
	if b.stack == nil {
		b.stack = NewStack()
	}
	return b.stack
}

func (b *BaseBuilder) Peek() Displayable {
	return b.getStack().Peek()
}

func (b *BaseBuilder) getExistingChild(d Displayable, parent Displayable) Displayable {
	updateableChildren := parent.UpdateableChildren()
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
		parent.SetUpdateableChildren(updateableChildren)
	}
	return result
}

func (b *BaseBuilder) Update(d Displayable) {
	b.Push(d)
	d.Layout()
}

func (b *BaseBuilder) Push(d Displayable, options ...Option) {
	firstRun := b.root == nil
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
		d.SetComposer(nil)
	} else if b.root == nil {
		b.root = d
	}

	d.UnsubAll()
	// Apply all options up to this point.
	b.applyOptions(d, options)
	// NOTE(lbayes): This MUST be done AFTER applying all other options, as
	// they may include an Option to SetState which allows this to work.
	b.applyOptions(d, d.OptionsForState(d.State()))

	// Trigger Configured lifecycle event.
	d.Emit(events.New(events.Configured, d, nil))

	// Push the element onto the stack
	stack.Push(d)

	// Create the children type map for this parent
	d.SetUpdateableChildren(b.createChildrenMapFor(d))

	// Process composition function to build children
	b.callComposeFunctionFor(d)

	// Remove the children that haven't been used.
	b.clearChildrenFromTypeMap(d)

	// Pop the element off the stack
	stack.Pop()

	if firstRun && !stack.HasNext() {
		b.root.Layout()
	}
}

func (b *BaseBuilder) applyOptions(d Displayable, options []Option) {
	// One of these options might be a Children(func()), which will recurse
	// back into this Push function.
	for _, option := range options {
		option(d)
	}
}

func (b *BaseBuilder) clearChildrenFromTypeMap(d Displayable) {
	types := d.UpdateableChildren()

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
	composeWithContext := d.GetComposeWithContext()
	if composeWithContext != nil {
		composeWithContext(b.root.Context())
		return
	}
	composeWithControl := d.GetComposeWithControl()
	if composeWithControl != nil {
		composeWithControl(d)
		return
	}
	composeWithContextAndControl := d.GetComposeWithContextAndControl()
	if composeWithContextAndControl != nil {
		composeWithContextAndControl(b.root.Context(), d)
		return
	}
}

func (b *BaseBuilder) Destroy() {
	b.stack = nil
	b.root = nil
	b.isDestroyed = true
}

// NewBuilder returns a clean builder instance.
func NewBuilder() *BaseBuilder {
	return &BaseBuilder{}
}
