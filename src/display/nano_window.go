package display

import (
	"github.com/shibukawa/nanovgo"
	"time"
)

type NanoWindowComponent struct {
	GlfwWindowComponent

	nanoContext *nanovgo.Context
	nanoSurface Surface
}

func (c *NanoWindowComponent) updateSize(width, height int) {
	if float64(width) != c.GetWidth() || float64(height) != c.GetHeight() {
		c.Width(float64(width))
		c.Height(float64(height))
		c.LayoutDrawAndPaint()
	}
}

func (c *NanoWindowComponent) initNanoContext() {
	c.initGl()
	context, err := nanovgo.NewContext(0 /* nnovgo.AntiAlias | nanovgo.StencilStrokes | nanovgo.Debug */)
	if err != nil {
		panic(err)
	}

	c.nanoContext = context
}

func (c *NanoWindowComponent) initSurface() {
	c.nanoSurface = NewNanoSurface(c.nanoContext)
}

func (c *NanoWindowComponent) onCloseWindow() {
	c.nanoContext.Delete()
	c.GlfwWindowComponent.OnClose()
}

func (c *NanoWindowComponent) Loop() {
	// Do not connect to the GPU hardware until we begin looping.
	// This allows us to set up an instance in the test environment.
	c.initGlfw()
	c.initNanoContext()
	c.OnWindowResize(c.updateSize)
	c.initSurface()
	c.LayoutDrawAndPaint()

	// Clean up GL and GLFW entities before closing
	defer c.onCloseWindow()
	for {
		t := time.Now()

		if c.getNativeWindow().ShouldClose() {
			return
		}

		c.PollEvents()
		c.LayoutDrawAndPaint()

		// Wait for whatever amount of time remains between how long we just spent,
		// and when the next frame (at fps) should be.
		waitDuration := time.Second/time.Duration(c.GetFrameRate()) - time.Since(t)
		// NOTE: Looping stops when mouse is pressed on window resizer (on macOS, but not i3wm/Ubuntu Linux)
		time.Sleep(waitDuration)
	}
}

func (c *NanoWindowComponent) LayoutDrawAndPaint() {
	// Make the component window size match the window frame buffer.
	w, h := c.getNativeWindow().GetFramebufferSize()
	winWidth, winHeight := float64(w), float64(h)
	c.Width(winWidth)
	c.Height(winHeight)

	c.nanoContext.BeginFrame(int(winWidth), int(winHeight), 1.0)
	c.Layout()
	c.GlLayout()
	c.Draw(c.nanoSurface)
	c.GlClear()
	c.nanoContext.EndFrame()
	c.SwapWindowBuffers()
}

func (c *NanoWindowComponent) GetFrameRate() int {
	return c.frameRate
}

func NewNanoWindow() Displayable {
	win := &NanoWindowComponent{}
	win.frameRate = DefaultFrameRate
	win.Title(DefaultWindowTitle)
	return win
}

// Debating whether this belongs in this file, or if they should all be
// defined in component_factory.go, or maybe someplace else?
// This is the hook that is used within the Builder context.
var NanoWindow = NewComponentFactory(NewNanoWindow, LayoutType(VerticalFlowLayoutType))
