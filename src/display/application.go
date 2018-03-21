package display

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/cairo/cairogl"
	"log"
	"runtime"
	"time"
)

const (
	DefaultWidth           = 1024
	DefaultHeight          = 768
	DefaultFramesPerSecond = 12
)

func init() {
	runtime.LockOSThread()
}

type application struct {
	vbox

	cairoSurface *cairogl.Surface
	nativeWindow *glfw.Window
	fps          int
	surface      Surface
	renderer     Renderer
}

func (a *application) GetFramesPerSecond() int {
	if a.fps == 0 {
		fps := a.GetDeclaration().Options.FramesPerSecond
		if fps > 0 {
			a.fps = fps
		} else {
			a.fps = DefaultFramesPerSecond
		}
	}
	return a.fps
}

func (a *application) GetSize() (float64, float64) {
	width := a.GetWidth()
	if width == 0 {
		width = DefaultWidth
		a.Width(width)
	}
	height := a.GetHeight()
	if height == 0 {
		height = DefaultHeight
		a.Height(height)
	}

	return width, height
}

func (a *application) initGlfw() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	width, height := a.GetSize()

	glfw.WindowHint(glfw.Floating, 1)
	glfw.WindowHint(glfw.Focused, 1)
	glfw.WindowHint(glfw.Resizable, 1)
	glfw.WindowHint(glfw.Visible, 1)

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	win, err := glfw.CreateWindow(int(width), int(height), a.GetTitle(), nil, nil)
	log.Println("TITLE:", a.GetTitle())

	if err != nil {
		panic(err)
	}

	win.MakeContextCurrent()
	a.nativeWindow = win
}

func (a *application) initGl() {
	ww, wh := a.nativeWindow.GetSize()
	width, height := a.nativeWindow.GetFramebufferSize()
	log.Printf("glfw: created window %dx%d (framebuffer: %dx%d)", ww, wh, width, height)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Viewport(0, 0, int32(width), int32(height))
	a.cairoSurface = cairogl.NewSurface(width, height)
}

func (a *application) initApplication() {
	a.nativeWindow.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		a.updateSize(width, height)
	})

	composeChildren := a.GetDeclaration().ComposeWithSurface

	if composeChildren == nil {
		panic("Application must be provided a function callback that receives a Surface as an argument")
	}

	// Create the Epiphyte -> Cairo Surface Adapter
	a.surface = NewCairoSurface(a.cairoSurface.Context())
	a.renderer = CreateRenderer(a.surface, composeChildren)
}

func (a *application) OnClose() {
	a.cairoSurface.Destroy()
	glfw.Terminate()
}

func (a *application) ProcessUserInput() {
	// TODO(lbayes): Find user input and send signals through tree
	glfw.PollEvents()
}

func (a *application) RenderAndDraw() {
	log.Println("Application.children:", a.GetChildCount())
	width := a.GetWidth()
	height := a.GetHeight()

	// Render application
	a.renderer.RenderWithRoot(a)

	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(1, 1, 1, 1)
	a.cairoSurface.Draw()
	a.nativeWindow.SwapBuffers()
	log.Println("render complete")
}

func (a *application) Loop() {
	// Clean up GL and GLFW entities before closing
	defer a.OnClose()
	for {
		t := time.Now()

		if a.nativeWindow.ShouldClose() {
			log.Println("SHOULD CLOSE FOUND!")
			a.cairoSurface.Destroy()
			glfw.Terminate()
			return
		}

		a.ProcessUserInput()
		a.RenderAndDraw()

		// Wait for whatever amount of time remains between how long we just spent,
		// and when the next frame (at fps) should be.
		waitDuration := time.Second/time.Duration(a.GetFramesPerSecond()) - time.Since(t)
		// log.Println("Render waiting:", waitDuration)
		time.Sleep(waitDuration)
	}
}

func (a *application) init() {
	log.Println("Application.init called")
	a.initGlfw()
	a.initGl()
	a.initApplication()
}

func (a *application) updateSize(width int, height int) {
	log.Println("updateSize called with", width, height)
	a.Width(float64(width))
	a.Height(float64(height))
	a.cairoSurface.Update(width, height)
	// enqueue a render request
	a.RenderAndDraw()
}

func Application(args ...interface{}) *application {
	decl, err := NewDeclaration(args)
	if err != nil {
		panic(err)
	}

	instance := &application{}
	instance.Declaration(decl)
	instance.init()
	return instance
}
