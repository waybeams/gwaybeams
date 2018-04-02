package display

import (
	"github.com/golang-ui/cairo/cairogl"
	"time"
)

type CairoWindowComponent struct {
	GlfwWindowComponent

	cairoGlSurface      *cairogl.Surface
	cairoSurfaceAdapter Surface
}

func (c *CairoWindowComponent) updateSize(width, height int) {
	if float64(width) != c.GetWidth() || float64(height) != c.GetHeight() {
		c.cairoGlSurface.Update(width, height)
		// enqueue a render request
		c.LayoutDrawAndPaint()
	}
}

func (c *CairoWindowComponent) initCairoGl() {
	c.initGl()
	c.cairoGlSurface = cairogl.NewSurface(int(c.GetWidth()), int(c.GetHeight()))
}

func (c *CairoWindowComponent) initSurface() {
	// Create the Epiphyte OffsetSurface (manages offset) ->
	// Cairo Surface Adapter (indirection for Cairo context w/API calls) ->
	// Native CGO library Cairo surface wrapper
	c.cairoSurfaceAdapter = NewCairoSurfaceAdapter(c.cairoGlSurface.Context())
}

func (c *CairoWindowComponent) OnCloseWindow() {
	c.cairoGlSurface.Destroy()
	c.GlfwWindowComponent.OnClose()
}

func (c *CairoWindowComponent) Loop() {
	c.initGlfw()
	c.initCairoGl()
	c.OnWindowResize(c.updateSize)
	c.initSurface()

	c.LayoutDrawAndPaint()

	// Clean up GL and GLFW entities before closing
	defer c.OnCloseWindow()
	for {
		t := time.Now()

		if c.getNativeWindow().ShouldClose() {
			c.OnCloseWindow()
			return
		}

		c.PollEvents()
		// Don't want to force layouts on every render.
		// Need a layout engine to determine when/what to Layout()
		c.LayoutDrawAndPaint()

		// Wait for whatever amount of time remains between how long we just spent,
		// and when the next frame (at fps) should be.
		waitDuration := time.Second/time.Duration(c.GetFrameRate()) - time.Since(t)
		// NOTE: Looping stops when mouse is pressed on window resizer (on macOS, but not i3wm/Ubuntu Linux)
		time.Sleep(waitDuration)
	}
}

func (c *CairoWindowComponent) LayoutCairo() {
	c.cairoGlSurface.Update(int(c.GetWidth()), int(c.GetHeight()))
	c.Layout()
	c.LayoutGl()
}

func (c *CairoWindowComponent) DrawCairo() {
	c.Draw(c.cairoSurfaceAdapter)
	c.ClearGl()
	c.cairoGlSurface.Draw()
	c.SwapWindowBuffers()
}

func (c *CairoWindowComponent) LayoutDrawAndPaint() {
	// Make the component window size match the window frame buffer.
	w, h := c.getNativeWindow().GetFramebufferSize()
	winWidth, winHeight := float64(w), float64(h)
	c.Width(winWidth)
	c.Height(winHeight)

	c.LayoutCairo()
	c.DrawCairo()
}

func (c *CairoWindowComponent) GetFrameRate() int {
	return c.frameRate
}

func NewCairoWindow() Displayable {
	win := &CairoWindowComponent{}
	win.frameRate = DefaultFrameRate
	win.Title(DefaultWindowTitle)
	return win
}

// CairoWindow is available to use with Builders.
var CairoWindow = NewComponentFactory(NewCairoWindow, LayoutType(VerticalFlowLayoutType))
