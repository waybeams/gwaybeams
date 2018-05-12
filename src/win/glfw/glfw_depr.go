package glfw

/*
type GlfwWindowResizeHandler func(width, height int)

// GlfwWindowControl is used an abstract composition class for client
// surface implementations that use GLFW source support (e.g., Cairo,
// NanoVG and possibly Skia).
type GlfwWindowControl struct {
	spec.Spec

	nativeWindow *glfw.Window
}

type MouseButtonCallback func(button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey)
type CharCallback func(char rune)

func (g *GlfwWindowControl) SetMouseButtonCallback(callback MouseButtonCallback) events.Unsubscriber {
	// type MouseButtonCallback func(w *Window, button MouseButton, action Action, mod ModifierKey)
	g.nativeWindow.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		callback(button, action, mod)

	})
	// Unsubscribe the registered callback
	return func() bool {
		g.nativeWindow.SetMouseButtonCallback(nil)
		return true
	}
}

func (g *GlfwWindowControl) SetCharCallback(callback CharCallback) events.Unsubscriber {
	g.nativeWindow.SetCharCallback(func(w *glfw.Window, char rune) {
		callback(char)
	})
	// Unsubscribe the registered callback
	return func() bool {
		g.nativeWindow.SetCharCallback(nil)
		return true
	}
}

func (g *GlfwWindowControl) SetCursorByName(cursorName glfw.StandardCursor) {
	g.nativeWindow.SetCursor(glfw.CreateStandardCursor(cursorName))
}

func (g *GlfwWindowControl) GetCursorPos() (xpos, ypos float64) {
	return g.getNativeWindow().GetCursorPos()
}

func (g *GlfwWindowControl) getNativeWindow() *glfw.Window {
	return g.nativeWindow
}

func (g *GlfwWindowControl) OnWindowResize(handler GlfwWindowResizeHandler) {
	g.getNativeWindow().SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		handler(width, height)
	})
}

func (g *GlfwWindowControl) initGlfw() {
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

func (g *GlfwWindowControl) initGl() {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	width, height := g.Width(), g.Height()
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (g *GlfwWindowControl) PollEvents() []events.Event {
	// TODO(lbayes): Find user input and send signals through tree
	glfw.PollEvents()
	return nil
}

func (g *GlfwWindowControl) UpdateCursor() {
	// x, y := g.nativeWindow.GetCursorPos()
}

func (g *GlfwWindowControl) OnClose() {
	glfw.Terminate()
}

func (g *GlfwWindowControl) LayoutGl() {
	// log.Println("GlLayout with:", g.GetWidth(), g.GetHeight())
	gl.Viewport(0, 0, int32(g.Width()), int32(g.Height()))
}

func (g *GlfwWindowControl) ClearGl() {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	gl.ClearColor(0, 0, 0, 0)
}

func (g *GlfwWindowControl) EnableGlDepthTest() {
	gl.Enable(gl.DEPTH_TEST)
}

func (g *GlfwWindowControl) SwapWindowBuffers() {
	gl.Enable(gl.DEPTH_TEST)
	g.getNativeWindow().SwapBuffers()
}
*/
