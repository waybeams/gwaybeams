package display

import (
	"events"
	"github.com/shibukawa/nanovgo"
	"github.com/shibukawa/nanovgo/perfgraph"
)

const RobotoRegularTTF = "third_party/fonts/Roboto/Roboto-Regular.ttf"
const RobotoBoldTTF = "third_party/fonts/Roboto/Roboto-Bold.ttf"
const RobotLightTTF = "third_party/fonts/Roboto/Roboto-Light.ttf"

type NanoWindowComponent struct {
	GlfwWindowComponent

	lastHeight  int
	lastWidth   int
	nanoContext *nanovgo.Context
	nanoSurface Surface
	perfGraph   *perfgraph.PerfGraph
}

func (c *NanoWindowComponent) updateSize(width, height int) {
	if float64(width) != c.Width() || float64(height) != c.Height() {
		c.SetWidth(float64(width))
		c.SetHeight(float64(height))
		c.LayoutDrawAndPaint()
		c.Layout()
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
	robotoRegularCreated := c.nanoContext.CreateFont("Roboto", RobotoRegularTTF)
	if robotoRegularCreated == -1 {
		panic("Could not create regular font")
	}

	robotoBoldCreated := c.nanoContext.CreateFont("Roboto Bold", RobotoBoldTTF)
	if robotoBoldCreated == -1 {
		panic("Could not create Roboto-Bold font")
	}

	robotoLightCreated := c.nanoContext.CreateFont("Roboto Light", RobotLightTTF)
	if robotoLightCreated == -1 {
		panic("Could not create Roboto-Light font")
	}
}

func (c *NanoWindowComponent) initSurface() {
	c.nanoSurface = NewNanoSurface(c.nanoContext)
}

func (c *NanoWindowComponent) OnExit() {
	c.nanoContext.Delete()
	c.GlfwWindowComponent.OnClose()
}

func (c *NanoWindowComponent) ShouldExit() bool {
	return c.getNativeWindow().ShouldClose()
}

func (c *NanoWindowComponent) onEnterFrame(e Event) {
	c.LayoutDrawAndPaint()
	c.PollEvents()
	c.UpdateCursor()

	if c.ShouldExit() {
		// Stop the frame loop by destroying the Builder
		c.Builder().Destroy()
	}
}

func (c *NanoWindowComponent) Init() {
	// Do not connect to the GPU hardware until we begin looping.
	// This allows us to set up an instance in the test environment.
	c.initGlfw()
	c.initNanoContext()
	c.initNanoFonts()
	c.initSurface()
	c.perfGraph = perfgraph.NewPerfGraph("Frame Time", "Roboto")
	c.OnWindowResize(c.updateSize)

	defer c.OnExit()
	c.Builder().OnEnterFrame(c.onEnterFrame)

	// Block permanently as frame events arrive
	c.Builder().Listen()
}

func (c *NanoWindowComponent) LayoutDrawAndPaint() {
	// Currently working to remove / rework this method from the controller
	// of the frame work, to one that is simply notified when a frame happens.

	// Make the component window size match the window frame buffer.
	fbWidth, fbHeight := c.getNativeWindow().GetFramebufferSize()
	winWidth, winHeight := c.getNativeWindow().GetSize()
	// TODO(lbayes): Only set pixelRatio on init, not every frame
	pixelRatio := float32(fbWidth) / float32(winWidth)

	c.nanoContext.BeginFrame(int(fbWidth), int(winHeight), pixelRatio)

	c.Emit(NewEvent(events.EnterFrame, c, nil))

	if fbWidth != c.lastWidth || fbHeight != c.lastHeight {
		c.SetWidth(float64(fbWidth))
		c.SetHeight(float64(fbHeight))

		c.lastHeight = fbHeight
		c.lastWidth = fbWidth

		// if c.ShouldRecompose() {
		// c.RecomposeChildren()
		// }

		// c.Layout()
	}

	c.LayoutGl()
	c.ClearGl()
	c.Draw(c.nanoSurface)

	if true && c.perfGraph != nil {
		c.perfGraph.UpdateGraph()
		c.perfGraph.RenderGraph(c.nanoContext, 5, 5)
	}

	c.nanoContext.EndFrame()
	c.SwapWindowBuffers()
}

func NewNanoWindow() Displayable {
	win := &NanoWindowComponent{}
	win.SetTitle(DefaultWindowTitle)
	return win
}

var NanoWindow = NewComponentFactory("NanoWindow", NewNanoWindow,
	LayoutType(VerticalFlowLayoutType),
	Width(DefaultWindowWidth),
	Height(DefaultWindowHeight))
