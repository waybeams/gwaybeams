package display

import (
	"errors"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type GlfwWindowResizeHandler func(width, height int)

const DefaultFrameRate = 63
const DefaultWindowWidth = 800
const DefaultWindowHeight = 600
const DefaultWindowTitle = "Default Title"

// GlfwWindowComponent is used an abstract composition class for client
// surface implementations that use GLFW window support (e.g., Cairo,
// NanoVG and possibly Skia).
type GlfwWindowComponent struct {
	ApplicationComponent

	nativeWindow *glfw.Window
}

func (g *GlfwWindowComponent) getNativeWindow() *glfw.Window {
	return g.nativeWindow
}

func (g *GlfwWindowComponent) OnWindowResize(handler GlfwWindowResizeHandler) {
	g.getNativeWindow().SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		handler(width, height)
	})
}

func (g *GlfwWindowComponent) initGlfw() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	width, height := g.Width(), g.Height()

	glfw.WindowHint(glfw.Floating, 1)
	glfw.WindowHint(glfw.Focused, 1)
	glfw.WindowHint(glfw.Resizable, 1)
	glfw.WindowHint(glfw.Visible, 1)

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	// glfw.WindowHint(glfw.DoubleBuffer, 1)
	win, err := glfw.CreateWindow(int(width), int(height), g.Title(), nil, nil)

	if err != nil {
		panic(err)
	}

	win.MakeContextCurrent()
	// glfw.SwapInterval(0)

	g.nativeWindow = win
}

func (g *GlfwWindowComponent) initGl() {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	width, height := g.Width(), g.Height()
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (g *GlfwWindowComponent) PollEvents() []Event {
	// TODO(lbayes): Find user input and send signals through tree
	glfw.PollEvents()
	return nil
}

func (g *GlfwWindowComponent) UpdateCursor() {
	// x, y := g.nativeWindow.GetCursorPos()
}

func (g *GlfwWindowComponent) OnClose() {
	glfw.Terminate()
}

func (g *GlfwWindowComponent) LayoutGl() {
	// log.Println("GlLayout with:", g.GetWidth(), g.GetHeight())
	gl.Viewport(0, 0, int32(g.Width()), int32(g.Height()))
}

func (g *GlfwWindowComponent) ClearGl() {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	gl.ClearColor(0, 0, 0, 0)
}

func (g *GlfwWindowComponent) EnableGlDepthTest() {
	gl.Enable(gl.DEPTH_TEST)
}

func (g *GlfwWindowComponent) SwapWindowBuffers() {
	gl.Enable(gl.DEPTH_TEST)
	g.getNativeWindow().SwapBuffers()
}

func GlfwFrameRate(value int) ComponentOption {
	return func(d Displayable) error {
		win := d.(*GlfwWindowComponent)
		if win == nil {
			return errors.New("Can only set FrameRate on GlfwWindowComponent")
		}
		win.frameRate = value
		return nil
	}
}
