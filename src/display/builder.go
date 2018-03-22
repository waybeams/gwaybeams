package display

// import "errors"

type SurfaceType int

const (
	CairoSurfaceType = iota
	ImageSurfaceType
	FakeSurfaceType
)

type GlfwWindowHint int

const (
	AutoIconify = iota
	Decorated
	Floating
	Focused
	Iconified
	Maximized
	Resizable
	Visible
)

/*
// TODO(lbayes) Map local hints to glfw library hints
const (
	Focused     GlfwWindowHint = C.GLFW_FOCUSED      // Specifies whether the window will be given input focus when created. This hint is ignored for full screen and initially hidden windows.
	Iconified   Hint = C.GLFW_ICONIFIED    // Specifies whether the window will be minimized.
	Maximized   Hint = C.GLFW_MAXIMIZED    // Specifies whether the window is maximized.
	Visible     Hint = C.GLFW_VISIBLE      // Specifies whether the window will be initially visible.
	Resizable   Hint = C.GLFW_RESIZABLE    // Specifies whether the window will be resizable by the user.
	Decorated   Hint = C.GLFW_DECORATED    // Specifies whether the window will have window decorations such as a border, a close widget, etc.
	Floating    Hint = C.GLFW_FLOATING     // Specifies whether the window will be always-on-top.
	AutoIconify Hint = C.GLFW_AUTO_ICONIFY // Specifies whether fullscreen windows automatically iconify (and restore the previous video mode) on focus loss.
)

*/

type Builder2 interface {
	// Provide a named surface type (CairoSurface, ImageSurface, FakeSurface)
	WithSurfaceType(t SurfaceType) Builder2
	// Provide a concrete surface instance.
	WithSurface(t Surface) Builder2
	// Provide the desired frame rate (in frames per second)
	WithFrameRate(fps int) Builder2
	// Provide zero or more Window Hints to Glfw.
	WithHint(name GlfwWindowHint, value interface{}) Builder2
	// Provide width and height to the native window.
	WithSize(width int, height int) Builder2
	// Provide title to the native window.
	WithTitle(title string) Builder2
	// Execute the configured builder and receive a surface that can be used to
	//draw visual components.
	Build(handler func(s Surface)) error

	GetNativeWindowSize() (width int, height int)
	PollEvents() []interface{}
}

type builder2 struct {
	windowHints map[GlfwWindowHint]interface{}
}

// Provide a named surface type (CairoSurface, ImageSurface, FakeSurface)
func (b *builder2) WithSurfaceType(t SurfaceType) Builder2 {
	return b
}

// Provide a concrete surface instance.
func (b *builder2) WithSurface(t Surface) Builder2 {
	return b
}

// Provide the desired frame rate (in frames per second)
func (b *builder2) WithFrameRate(fps int) Builder2 {
	return b
}

// Provide zero or more Window Hints to Glfw.
func (b *builder2) WithHint(name GlfwWindowHint, value interface{}) Builder2 {
	return b
}

// Provide width and height to the native window.
func (b *builder2) WithSize(width int, height int) Builder2 {
	return b
}

// Provide title to the native window.
func (b *builder2) WithTitle(title string) Builder2 {
	return b
}

// Execute the configured builder and receive a surface that can be used to
//draw visual components.
func (b *builder2) Build(handler func(s Surface)) error {
	return nil
}

func (b *builder2) GetNativeWindowSize() (width int, height int) {
	return 0, 0
}

func (b *builder2) PollEvents() []interface{} {
	return nil
}

func NewBuilder2() Builder2 {
	return &builder2{}
}

//----------------------------------------------------------------------

type Builder interface {
	Build()
	BuildWithRoot(d Displayable)
	GetRoot() Displayable
}

// Factory that operates over semantic sugar that we use to describe the
// displayable hierarchy.
type builder struct {
	*SurfaceDelegate
	stack        Stack
	root         Displayable
	buildHandler func(Surface)
}

func (b *builder) getStack() Stack {
	if b.stack == nil {
		b.stack = NewStack()
	}
	return b.stack
}

func (b *builder) GetRoot() Displayable {
	return b.root
}

func (b *builder) Push(d Displayable) error {
	if b.root == nil {
		b.root = d
	}

	s := b.getStack()
	parent := s.Peek()

	if parent != nil {
		parent.AddChild(d)
	}

	err := s.Push(d)
	if err != nil {
		return err
	}

	decl := d.GetDeclaration()
	if decl.Compose != nil {
		decl.Compose()
	} else if decl.ComposeWithUpdate != nil {
		panic("Not yet implemented")
	}

	//if b.root == d {
	// d.builder(r)
	// } else {
	s.Pop()
	// }

	return nil
}

func (b *builder) Reset() {
	b.stack = NewStack()
	b.root = nil
}

func (b *builder) BuildWithRoot(root Displayable) {
	b.Reset()
	b.Push(root)
	b.buildHandler(b)
	b.root.Render(b)
}

func (b *builder) Build() {
	b.Reset()
	b.buildHandler(b)
	b.root.Render(b)
}

func CreateBuilder(s Surface, buildHandler func(s Surface)) Builder {
	return &builder{SurfaceDelegate: NewSurfaceDelegate(s), buildHandler: buildHandler}
}
