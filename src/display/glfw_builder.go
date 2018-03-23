package display

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/cairo/cairogl"
	// "runtime"
	// "time"
)

const DefaultFrameRate = 60
const DefaultWindowWidth = 1024
const DefaultWindowHeight = 768
const DefaultWindowTitle = "Default Title"

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

type GlfwBuilder interface {
	Builder
	GetWindowHint(hintName GlfwWindowHint) interface{}
	GetWindowHints() []*windowHint
	PushWindowHint(hintName GlfwWindowHint, value interface{})
	RemoveWindowHint(hintName GlfwWindowHint)
	FrameRate(fps int)
	GetFrameRate() int
	GetWindowHeight() int
	GetWindowSize() (width int, height int)
	GetWindowTitle() string
	GetWindowWidth() int
	WindowSize(width, height int)
	WindowTitle(title string)
}

type glfwBuilder struct {
	builder

	cairoSurface *cairogl.Surface
	frameRate    int
	height       int
	nativeWindow *glfw.Window
	surface      Surface
	width        int
	windowHints  []*windowHint
	windowTitle  string
}

func (g *glfwBuilder) applyGlfwDefaults() {
	g.frameRate = DefaultFrameRate
	g.width = DefaultWindowWidth
	g.height = DefaultWindowHeight
	g.windowTitle = DefaultWindowTitle

	g.windowHints = []*windowHint{
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

func (g *glfwBuilder) createSurface() {
	g.initGlfw()
	g.initGl()
}

func (g *glfwBuilder) initGlfw() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	width, height := g.GetWindowSize()

	glfw.WindowHint(glfw.Floating, 1)
	glfw.WindowHint(glfw.Focused, 1)
	glfw.WindowHint(glfw.Resizable, 1)
	glfw.WindowHint(glfw.Visible, 1)

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	win, err := glfw.CreateWindow(int(width), int(height), g.GetWindowTitle(), nil, nil)

	if err != nil {
		panic(err)
	}

	win.MakeContextCurrent()
	win.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		g.updateSize(width, height)
	})

	g.width, g.height = win.GetFramebufferSize()
	g.nativeWindow = win
}

// TODO(lbayes): Pretty sure these should move out to GlfwBuilder.
func (g *glfwBuilder) GetWindowHint(hintName GlfwWindowHint) interface{} {
	for _, hint := range g.windowHints {
		if hint.name == hintName {
			return hint.value
		}
	}
	return nil
}

func (g *glfwBuilder) PushWindowHint(hintName GlfwWindowHint, value interface{}) {
	g.RemoveWindowHint(hintName)

	wHint := &windowHint{
		name:  hintName,
		value: value,
	}

	g.windowHints = append(g.windowHints, wHint)
}

func (g *glfwBuilder) RemoveWindowHint(hintName GlfwWindowHint) {
	hints := g.windowHints
	for i := 0; i < len(hints); i++ {
		if hints[i].name == hintName {
			g.windowHints = append(hints[:i], hints[i+1:]...)
			return
		}
	}
}

func (g *glfwBuilder) GetWindowHints() []*windowHint {
	return g.windowHints
}
func (g *glfwBuilder) initGl() {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Viewport(0, 0, int32(g.width), int32(g.height))
	g.cairoSurface = cairogl.NewSurface(g.width, g.height)
}

func (g *glfwBuilder) initSurface() {
	// composeChildren := g.GetDeclaration().ComposeWithSurface
	// if composeChildren == nil {
	// panic("Application must be provided a function callback that receives a Surface as an argument")
	// }

	// Create the Epiphyte -> Cairo Surface Adapter
	g.surface = NewCairoSurface(g.cairoSurface.Context())
	// g.builder = CreateBuilder(g.surface, composeChildren)
}

func (g *glfwBuilder) updateSize(width int, height int) {
	g.width = width
	g.height = height
	g.cairoSurface.Update(width, height)
	// enqueue a render request
	// g.BuildAndRender()
}

func (g *glfwBuilder) FrameRate(fps int) {
	g.frameRate = fps
}

func (g *glfwBuilder) GetFrameRate() int {
	return g.frameRate
}

func (g *glfwBuilder) GetWindowWidth() int {
	return g.width
}

func (g *glfwBuilder) GetWindowHeight() int {
	return g.height
}

func (g *glfwBuilder) WindowSize(width, height int) {
	g.width = width
	g.height = height
}

func (g *glfwBuilder) GetWindowSize() (width, height int) {
	return g.width, g.height
}

func (g *glfwBuilder) GetWindowTitle() string {
	return g.windowTitle
}

func (g *glfwBuilder) WindowTitle(title string) {
	g.windowTitle = title
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
func NewGlfwBuilder(args ...GlfwBuilderOption) GlfwBuilder {
	g := &glfwBuilder{}
	g.applyGlfwDefaults()

	for _, arg := range args {
		// Options write directly to the builder struct, not via the interface
		err := arg(g)
		if err != nil {
			// Store any errors until Build is called. This allows us to chain
			// the calls and makes clients much more readable.
			g.lastError = err
			break
		}
	}

	return g
}
