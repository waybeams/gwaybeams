package display

import (
	"errors"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/cairo/cairogl"
	"time"
)

const DefaultFrameRate = 12
const DefaultWindowWidth = 1024
const DefaultWindowHeight = 768
const DefaultWindowTitle = "Default Title"

type GlfwWindowComponent struct {
	Component

	cairoSurface *cairogl.Surface
	frameRate    int
	height       int
	nativeWindow *glfw.Window
	surface      Surface
	width        int
}

func (g *GlfwWindowComponent) updateSize(width, height int) {
	g.Width(float64(width))
	g.Height(float64(height))

	// Pull them from the component in order to respect layout constraints.
	g.cairoSurface.Update(int(g.GetWidth()), int(g.GetHeight()))
	// enqueue a render request
	g.LayoutDrawAndPaint()
}

func (g *GlfwWindowComponent) createSurface() Surface {
	// Create the Epiphyte -> Cairo Surface Adapter
	return NewCairoSurface(g.cairoSurface.Context())
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
	win.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		g.updateSize(width, height)
	})

	g.width, g.height = win.GetFramebufferSize()
	g.nativeWindow = win
}

func (g *GlfwWindowComponent) initGl() {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	width, height := g.GetWidth(), g.GetHeight()
	gl.Viewport(0, 0, int32(width), int32(height))
	g.cairoSurface = cairogl.NewSurface(int(width), int(height))
}

func (g *GlfwWindowComponent) initSurface() {
	// Create the Epiphyte -> Cairo Surface Adapter
	g.surface = NewCairoSurface(g.cairoSurface.Context())
}

func (g *GlfwWindowComponent) ProcessUserInput() {
	// TODO(lbayes): Find user input and send signals through tree
	glfw.PollEvents()
}

func (g *GlfwWindowComponent) OnClose() {
	g.cairoSurface.Destroy()
	glfw.Terminate()
}

func (g *GlfwWindowComponent) Loop() {
	g.initGlfw()
	g.initGl()
	g.surface = g.createSurface()

	// Clean up GL and GLFW entities before closing
	defer g.OnClose()
	for {
		t := time.Now()

		if g.nativeWindow.ShouldClose() {
			g.OnClose()
			return
		}

		g.ProcessUserInput()
		g.LayoutDrawAndPaint()

		// Wait for whatever amount of time remains between how long we just spent,
		// and when the next frame (at fps) should be.
		waitDuration := time.Second/time.Duration(g.GetFrameRate()) - time.Since(t)
		time.Sleep(waitDuration)
	}
}

func (g *GlfwWindowComponent) GlLayout() {
	g.Layout()
	gl.Viewport(0, 0, int32(g.GetWidth()), int32(g.GetHeight()))
}

func (g *GlfwWindowComponent) GlDraw() {
	g.Draw(g.surface)
}

func (g *GlfwWindowComponent) GlPaint() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(1, 1, 1, 1)
	g.cairoSurface.Draw()
	g.nativeWindow.SwapBuffers()
}

func (g *GlfwWindowComponent) LayoutDrawAndPaint() {
	g.GlLayout()
	g.GlDraw()
	g.GlPaint()
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
