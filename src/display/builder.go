package display

import (
	"errors"
)

type ComponentComposer func(b Builder)

type Builder interface {
	Build(factory ComponentComposer) (root Displayable, err error)
	Push(d Displayable)
}

type builder struct {
	root      Displayable
	stack     DisplayStack
	lastError error
}

func (b *builder) getStack() DisplayStack {
	if b.stack == nil {
		b.stack = NewDisplayStack()
	}
	return b.stack
}

func (b *builder) Current() Displayable {
	return b.getStack().Peek()
}

func (b *builder) Push(d Displayable) {
	if b.root == nil {
		b.root = d
	}

	stack := b.getStack()

	// Get the parent element if one exists
	parent := stack.Peek()

	if parent == nil {
		if b.root != d {
			// It looks like we have a second root definition in the outer factory
			// function
			b.lastError = errors.New("Box factory function should only have a single root node")
			return
		}
	} else {
		parent.AddChild(d)
	}

	// Push the element onto the displayStack
	stack.Push(d)

	// Create the element's children by calling the associated Children(compose) function
	decl := d.GetDeclaration()
	if decl.Compose != nil {
		decl.Compose()
	} else if decl.ComposeWithBuilder != nil {
		decl.ComposeWithBuilder(b)
	} else if decl.ComposeWithUpdate != nil {
		panic("Not yet implemented")
	}

	// Pop the element off the displayStack
	stack.Pop()
}

// This method should be deprecated, clients should use the Component factory functions directly
// instead.
func (b *builder) Build(factory ComponentComposer) (Displayable, error) {
	// We may have a configuration error that was stored for later. If so, stop
	// and return it now.
	if b.lastError != nil {
		return nil, b.lastError
	}

	factory(b)

	if b.lastError != nil {
		return nil, b.lastError
	}

	return b.root, nil
}

func NewBuilder() Builder {
	return &builder{}
}
