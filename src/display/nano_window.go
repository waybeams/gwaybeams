package display

import (
	"github.com/shibukawa/nanovgo"
	"github.com/shibukawa/nanovgo/perfgraph"
	"time"
)

const RobotoRegularTTF = "third_party/fonts/Roboto/Roboto-Regular.ttf"
const RobotoBoldTTF = "third_party/fonts/Roboto/Roboto-Bold.ttf"

type NanoWindowComponent struct {
	GlfwWindowComponent

	lastHeight  int
	lastWidth   int
	nanoContext *nanovgo.Context
	nanoSurface Surface
	perfGraph   *perfgraph.PerfGraph
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

func (c *NanoWindowComponent) initNanoFonts() {
	regularCreated := c.nanoContext.CreateFont("sans", RobotoRegularTTF)
	if regularCreated == -1 {
		panic("Could not create regular font")
	}

	boldCreated := c.nanoContext.CreateFont("sans-bold", RobotoBoldTTF)
	if boldCreated == -1 {
		panic("Could not create bold font")
	}
}

func (c *NanoWindowComponent) initSurface() {
	c.nanoSurface = NewNanoSurface(c.nanoContext)
}

func (c *NanoWindowComponent) onCloseWindow() {
	c.nanoContext.Delete()
	c.GlfwWindowComponent.OnClose()
}

func (c *NanoWindowComponent) Init() {
	// Do not connect to the GPU hardware until we begin looping.
	// This allows us to set up an instance in the test environment.
	c.initGlfw()
	c.initNanoContext()
	c.initNanoFonts()
	c.initSurface()
	c.perfGraph = perfgraph.NewPerfGraph("Frame Time", "sans")
	c.OnWindowResize(c.updateSize)
	// Clean up GL and GLFW entities before closing
	defer c.onCloseWindow()
	for {
		t := time.Now()

		if c.getNativeWindow().ShouldClose() {
			return
		}

		c.LayoutDrawAndPaint()
		c.PollEvents()
		c.UpdateCursor()

		// Wait for whatever amount of time remains between how long we just spent,
		// and when the next frame (at fps) should be.
		waitDuration := time.Second/time.Duration(c.GetFrameRate()) - time.Since(t)
		// NOTE: Looping stops when mouse is pressed on window resizer (on macOS, but not i3wm/Ubuntu Linux)
		time.Sleep(waitDuration)
	}
}

func (c *NanoWindowComponent) LayoutDrawAndPaint() {
	// Make the component window size match the window frame buffer.
	fbWidth, fbHeight := c.getNativeWindow().GetFramebufferSize()
	winWidth, winHeight := c.getNativeWindow().GetSize()
	pixelRatio := float32(fbWidth) / float32(winWidth)

	c.Width(float64(fbWidth))
	c.Height(float64(fbHeight))

	if fbWidth != c.lastWidth || fbHeight != c.lastHeight {
		c.lastHeight = fbHeight
		c.lastWidth = fbWidth
		c.Layout()

		// IF SOMETHING ELSE HAPPENED
		c.LayoutGl()
		c.ClearGl()
		c.nanoContext.BeginFrame(int(fbWidth), int(winHeight), pixelRatio)

		c.Layout()
		c.Draw(c.nanoSurface)

		if false && c.perfGraph != nil {
			c.perfGraph.UpdateGraph()
			c.perfGraph.RenderGraph(c.nanoContext, 5, 5)
		}

		c.nanoContext.EndFrame()
		// c.EnableGlDepthTest()
		c.SwapWindowBuffers()
	}
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

var NanoWindow = NewComponentFactory("NanoWindow", NewNanoWindow,
	LayoutType(VerticalFlowLayoutType),
	Width(DefaultWindowWidth),
	Height(DefaultWindowHeight))
