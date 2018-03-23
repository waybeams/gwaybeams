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
	stack     DisplayStack // TODO: Move THIS displayStack def into builder package
	lastError error
}

func (b *builder) Push(d Displayable) {
	if b.stack == nil {
		b.stack = NewDisplayStack()
	}

	if b.root == nil {
		b.root = d
	}

	// Get the parent element if one exists
	parent := b.stack.Peek()

	if parent == nil {
		if b.root != d {
			// It looks like we have a second root definition in the outer factory function
			b.lastError = errors.New("Component factory function should only have a single root node")
			return
		}
	} else {
		parent.AddChild(d)
	}

	// Push the element onto the displayStack
	b.stack.Push(d)

	// Render the element children by calling it's compose function
	decl := d.GetDeclaration()
	if decl.Compose != nil {
		decl.Compose()
	} else if decl.ComposeWithUpdate != nil {
		panic("Not yet implemented")
	}

	// Pop the element off the displayStack
	b.stack.Pop()
}

func (b *builder) Build(factory ComponentComposer) (root Displayable, err error) {
	// We may have a configuration error that was stored for later. If so, stop
	// and return it now.
	if b.lastError != nil {
		return nil, err
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
