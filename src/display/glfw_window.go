package display

import (
	"errors"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type GlfwWindowResizeHandler func(width, height int)

const DefaultFrameRate = 60
const DefaultWindowWidth = 1024
const DefaultWindowHeight = 768
const DefaultWindowTitle = "Default Title"

type GlfwWindowComponent struct {
	Component

	frameRate    int
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

	width, height := g.GetWidth(), g.GetHeight()

	glfw.WindowHint(glfw.Floating, 1)
	glfw.WindowHint(glfw.Focused, 1)
	glfw.WindowHint(glfw.Resizable, 1)
	glfw.WindowHint(glfw.Visible, 1)

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	win, err := glfw.CreateWindow(int(width), int(height), g.GetTitle(), nil, nil)

	if err != nil {
		panic(err)
	}

	win.MakeContextCurrent()
	glfw.SwapInterval(0)

	g.nativeWindow = win
}

func (g *GlfwWindowComponent) initGl() {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	width, height := g.GetWidth(), g.GetHeight()
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (g *GlfwWindowComponent) PollEvents() {
	// TODO(lbayes): Find user input and send signals through tree
	glfw.PollEvents()
}

func (g *GlfwWindowComponent) OnClose() {
	glfw.Terminate()
}

func (g *GlfwWindowComponent) GlLayout() {
	// log.Println("GlLayout with:", g.GetWidth(), g.GetHeight())
	gl.Viewport(0, 0, int32(g.GetWidth()), int32(g.GetHeight()))

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)
}

func (g *GlfwWindowComponent) GlClear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(1, 1, 1, 1)
}

func (g *GlfwWindowComponent) SwapWindowBuffers() {
	gl.Enable(gl.DEPTH_TEST)
	g.getNativeWindow().SwapBuffers()
}

func (g *GlfwWindowComponent) GetFrameRate() int {
	return g.frameRate
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

func NewGlfwWindow() Displayable {
	win := &GlfwWindowComponent{}
	win.frameRate = DefaultFrameRate
	win.Title(DefaultWindowTitle)
	return win
}

// Debating whether this belongs in this file, or if they should all be
// defined in component_factory.go, or maybe someplace else?
// This is the hook that is used within the Builder context.
var GlfwWindow = NewComponentFactory(NewGlfwWindow, LayoutType(VerticalFlowLayoutType))
