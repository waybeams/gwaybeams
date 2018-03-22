package display

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/cairo/cairogl"
	// "runtime"
	// "time"
)

type glfwBuilder struct {
	builder

	cairoSurface *cairogl.Surface
	nativeWindow *glfw.Window
	surface      Surface
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

// Create a new builder instance with the provided options.
// This pattern was discovered by Rob Pike and published here:
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
// It was also supported by Dave Cheney here:
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
// and here:
// https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions
// I'm exploring the idea and finding it to be pretty compelling, especially for what
// we'd like to consider "immutable" values.
func NewGlfwBuilder(args ...BuilderOption) Builder {
	b := &glfwBuilder{}
	b.applyDefaults()

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
