package display

import (
	"errors"
)

// BuilderOption is a configuration option for Builders.
type BuilderOption func(b Builder) error

// ComponentComposer is a composition function that components send to the
// Children() option when composing children using an anonymous function.
type ComponentComposer func(b Builder)

// Builder is a basic wrapper around a stack that enables component
// composition.
type Builder interface {
	Push(d Displayable, options ...ComponentOption)
	Peek() Displayable
	GetInvalidNodes() []Displayable
	InvalidateChild(d Displayable)
	Destroy()
}

type builder struct {
	root       Displayable
	window     Window
	stack      Stack
	lastError  error
	dirtyNodes []Displayable
}

func (b *builder) Destroy() {
	b.stack = nil
	b.lastError = nil
	b.dirtyNodes = nil
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

func (b *builder) pruneChildNodes(nodes []Displayable) []Displayable {
	results := []Displayable{}
	for nIndex, node := range nodes {
		ancestorFound := false
		for aIndex, possibleAncestor := range nodes {
			if aIndex != nIndex && node.GetIsContainedBy(possibleAncestor) {
				ancestorFound = true
				break
			}
		}
		if !ancestorFound {
			results = append(results, node)
		}
	}

	return results
}

func (b *builder) GetInvalidNodes() []Displayable {
	return b.pruneChildNodes(b.dirtyNodes)
}

func (b *builder) InvalidateChild(d Displayable) {
	b.dirtyNodes = append(b.dirtyNodes, d)
}

func (b *builder) callComposeFunctionFor(d Displayable) (err error) {
	composeEmpty := d.GetComposeEmpty()
	if composeEmpty != nil {
		composeEmpty()
		return nil
	}
	composeWithBuilder := d.GetComposeWithBuilder()
	if composeWithBuilder != nil {
		composeWithBuilder(b)
		return nil
	}
	composeWithComponent := d.GetComposeWithComponent()
	if composeWithComponent != nil {
		composeWithComponent(d)
		return nil
	}
	composeWithBuilderAndComponent := d.GetComposeWithBuilderAndComponent()
	if composeWithBuilderAndComponent != nil {
		composeWithBuilderAndComponent(b, d)
		return nil
	}

	return errors.New("Expected compose function not found")
}

// Push accepts a new Displayable to place on the stack and processes the
// optional component composition function if one was provided.
func (b *builder) Push(d Displayable, options ...ComponentOption) {
	stack := b.getStack()

	// Get the parent element if one exists
	parent := stack.Peek()

	if parent == nil {
		b.root = d
	} else {
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
	composeError := b.callComposeFunctionFor(d)
	if composeError != nil && b.lastError == nil {
		b.lastError = composeError
	}

	// Pop the element off the stack
	stack.Pop()
}

// NewBuilder returns a clean builder instance.
func NewBuilder() Builder {
	return &builder{}
}
