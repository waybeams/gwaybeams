package display

import (
	"errors"
)

type Builder interface {
	Build(factory ComponentComposer) (root Displayable, err error)
	FrameRate(fps int)
	GetFrameRate() int
	GetWindowHeight() int
	GetWindowHint(hintName GlfwWindowHint) interface{}
	GetWindowHints() []*windowHint
	GetWindowSize() (width int, height int)
	GetWindowTitle() string
	GetWindowWidth() int
	Push(d Displayable)
	PushWindowHint(hintName GlfwWindowHint, value interface{})
	RemoveWindowHint(hintName GlfwWindowHint)
	WindowSize(width, height int)
	WindowTitle(title string)
}

type builder struct {
	frameRate   int
	width       int
	height      int
	windowTitle string
	root        Displayable
	stack       DisplayStack // TODO: Move THIS displayStack def into builder package
	surface     Surface
	lastError   error
}

func (b *builder) applyBaseDefaults() {
	b.frameRate = DefaultFrameRate
	b.width = DefaultWindowWidth
	b.height = DefaultWindowHeight
	b.windowTitle = DefaultWindowTitle
	b.stack = NewDisplayStack()
}

func (b *builder) createSurface() Surface {
	// create new surface here
	return nil
}

func (b *builder) Push(d Displayable) {
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

func (b *builder) draw() {
	b.root.Render()
	b.root.Draw(b.surface)
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

	b.surface = b.createSurface()

	return b.root, nil
}

func (b *builder) FrameRate(fps int) {
	b.frameRate = fps
}

func (b *builder) GetFrameRate() int {
	return b.frameRate
}

func (b *builder) GetWindowWidth() int {
	return b.width
}

func (b *builder) GetWindowHeight() int {
	return b.height
}

func (b *builder) WindowSize(width, height int) {
	b.width = width
	b.height = height
}

func (b *builder) GetWindowSize() (width, height int) {
	return b.width, b.height
}

func (b *builder) GetWindowTitle() string {
	return b.windowTitle
}

func (b *builder) WindowTitle(title string) {
	b.windowTitle = title
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
func NewBuilder(args ...BuilderOption) Builder {
	b := &builder{}
	b.applyBaseDefaults()

	for _, arg := range args {
		// Options write directly to the builder struct, not via the interface
		err := arg(b)
		if err != nil {
			// Store any errors until Build is called. This allows us to chain
			// the calls and makes clients much more readable.
			b.lastError = err
			break
		}
	}

	return b
}
