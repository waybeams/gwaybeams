package builder

import (
	"display"
	"errors"
)

type Builder interface {
	Build(factory ComponentComposer) (root display.Displayable, err error)
	GetFrameRate() int
	GetHeight() int
	GetSize() (width int, height int)
	GetSurfaceType() SurfaceTypeName
	GetTitle() string
	GetWidth() int
	GetWindowHint(hintName GlfwWindowHint) interface{}
	GetWindowHints() []*windowHint
	Push(d display.Displayable)
}

type builder struct {
	surfaceTypeName SurfaceTypeName
	frameRate       int
	windowHints     []*windowHint
	width           int
	height          int
	title           string
	root            display.Displayable
	stack           display.Stack // TODO: Move THIS stack def into builder package
	lastError       error
}

func (b *builder) applyDefaults() {
	b.frameRate = DefaultFrameRate
	b.width = DefaultWidth
	b.height = DefaultHeight
	b.title = DefaultTitle
	b.applyDefaultWindowHints()
	b.stack = display.NewStack()
}

func (b *builder) applyDefaultWindowHints() {
	b.windowHints = []*windowHint{
		&windowHint{name: AutoIconify, value: false},
		&windowHint{name: Decorated, value: false},
		&windowHint{name: Floating, value: true},
		&windowHint{name: Focused, value: true},
		&windowHint{name: Iconified, value: false},
		&windowHint{name: Maximized, value: false},
		&windowHint{name: Resizable, value: true},
		&windowHint{name: Visible, value: true},
	}
}

func (b *builder) removeWindowHint(hintName GlfwWindowHint) {
	hints := b.windowHints
	for i := 0; i < len(hints); i++ {
		if hints[i].name == hintName {
			b.windowHints = append(hints[:i], hints[i+1:]...)
			return
		}
	}
}

func (b *builder) createSurface() display.Surface {
	// create new surface here
	return nil
}

func (b *builder) Push(d display.Displayable) {
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

	// Push the element onto the stack
	b.stack.Push(d)

	// Render the element children by calling it's compose function
	decl := d.GetDeclaration()
	if decl.Compose != nil {
		decl.Compose()
	} else if decl.ComposeWithUpdate != nil {
		panic("Not yet implemented")
	}

	// Pop the element off the stack
	b.stack.Pop()
}

func (b *builder) Build(factory ComponentComposer) (root display.Displayable, err error) {
	factory(b)

	if b.lastError != nil {
		return nil, b.lastError
	}

	return b.root, nil
}

func (b *builder) GetSurfaceType() SurfaceTypeName {
	return b.surfaceTypeName
}

func (b *builder) GetFrameRate() int {
	return b.frameRate
}

func (b *builder) GetWindowHint(hintName GlfwWindowHint) interface{} {
	for _, hint := range b.windowHints {
		if hint.name == hintName {
			return hint.value
		}
	}
	return nil
}

func (b *builder) GetWindowHints() []*windowHint {
	return b.windowHints
}

func (b *builder) GetWidth() int {
	return b.width
}

func (b *builder) GetHeight() int {
	return b.height
}

func (b *builder) GetSize() (width, height int) {
	return b.width, b.height
}

func (b *builder) GetTitle() string {
	return b.title
}

// Create a new builder instance with the provided options.
// This pattern was discovered by Rob Pike and published here:
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
// It was also supported by Dave Cheney here:
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
// and here:
// https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions
// I'm exploring the idea and finding it to be pretty compelling, especially for what
// we'd like to consider "immutable" values.
func New(args ...BuilderOption) (Builder, error) {
	b := &builder{}
	b.applyDefaults()

	for _, arg := range args {
		err := arg(b)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}
