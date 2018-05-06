package controls

import (
	"component"
	"events"
	"github.com/shibukawa/nanovgo"
	"github.com/shibukawa/nanovgo/perfgraph"
	"log"
	"opts"
	"surface"
	"ui"
)

const RobotoRegularTTF = "third_party/fonts/Roboto/Roboto-Regular.ttf"
const RobotoBoldTTF = "third_party/fonts/Roboto/Roboto-Bold.ttf"
const RobotLightTTF = "third_party/fonts/Roboto/Roboto-Light.ttf"

type NanoWindowComponent struct {
	GlfwWindowComponent

	inputCtrl       InputController
	lastHeight      int
	lastWidth       int
	lastHoverTarget ui.Displayable
	nanoContext     *nanovgo.Context
	nanoSurface     ui.Surface
	perfGraph       *perfgraph.PerfGraph
}

func (c *NanoWindowComponent) initInput() {
	c.inputCtrl = NewGlfwInput(c, c)
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
		log.Print("Could not create regular font")
	}

	robotoBoldCreated := c.nanoContext.CreateFont("Roboto Bold", RobotoBoldTTF)
	if robotoBoldCreated == -1 {
		log.Print("Could not create Roboto-Bold font")
	}

	robotoLightCreated := c.nanoContext.CreateFont("Roboto Light", RobotLightTTF)
	if robotoLightCreated == -1 {
		log.Print("Could not create Roboto-Light font")
	}
}

func (c *NanoWindowComponent) initSurface() {
	c.nanoSurface = surface.NewNano(c.nanoContext)
}

func (c *NanoWindowComponent) OnExit() {
	c.nanoContext.Delete()
	c.GlfwWindowComponent.OnClose()
}

func (c *NanoWindowComponent) ShouldExit() bool {
	return c.getNativeWindow().ShouldClose()
}

func (c *NanoWindowComponent) enterFrameHandler(e events.Event) {
	c.inputCtrl.Update()
	c.LayoutDrawAndPaint()
	c.PollEvents()

	if c.ShouldExit() {
		// Stop the frame loop by destroying the Builder
		c.Context().Destroy()
	}
}

func (c *NanoWindowComponent) Listen() {
	c.init()

	defer c.OnExit()
	// TODO(lbayes): Definitely do not like this pattern. Need to find a cleaner way to set this up.
	// Components should generally not interact with the Builder and I do not want to require any
	// particular component TYPE to be the ROOT.
	c.Context().OnFrameEntered(c.enterFrameHandler)
	// Block permanently as frame events arrive
	c.Context().Listen()
}

func (c *NanoWindowComponent) init() {
	// Do not connect to the GPU hardware until we begin looping.
	// This allows us to set up an instance in the test environment.
	c.initGlfw()
	c.initNanoContext()
	c.initNanoFonts()
	c.initSurface()
	c.initInput()
	c.perfGraph = perfgraph.NewPerfGraph("Frame Time", "Roboto")
	c.OnWindowResize(c.updateSize)
}

func (c *NanoWindowComponent) LayoutDrawAndPaint() {
	// Currently working to remove / rework this method from the controller
	// of the frame work, to one that is simply notified when a frame happens.

	// Make the component source size match the source frame buffer.
	fbWidth, fbHeight := c.getNativeWindow().GetFramebufferSize()
	winWidth, winHeight := c.getNativeWindow().GetSize()
	// TODO(lbayes): Only set pixelRatio on init, not every frame
	pixelRatio := float32(fbWidth) / float32(winWidth)

	c.nanoContext.BeginFrame(int(fbWidth), int(winHeight), pixelRatio)

	// c.Emit(events.New(events.FrameEntered, c, nil))

	if fbWidth != c.lastWidth || fbHeight != c.lastHeight {
		c.SetWidth(float64(fbWidth))
		c.SetHeight(float64(fbHeight))

		c.lastHeight = fbHeight
		c.lastWidth = fbWidth
	}

	if c.ShouldRecompose() {
		c.RecomposeChildren()
	}

	c.LayoutGl()
	c.ClearGl()
	c.Draw(c.nanoSurface)

	if false && c.perfGraph != nil {
		c.perfGraph.UpdateGraph()
		c.perfGraph.RenderGraph(c.nanoContext, 5, 5)
	}

	c.nanoContext.EndFrame()
	c.SwapWindowBuffers()
}

func NewNanoWindow() ui.Displayable {
	win := &NanoWindowComponent{}
	win.SetTitle(ui.DefaultWindowTitle)
	return win
}

var NanoWindow = component.Define("NanoWindow", NewNanoWindow,
	opts.LayoutType(ui.VerticalFlowLayoutType),
	opts.Width(ui.DefaultWindowWidth),
	opts.Height(ui.DefaultWindowHeight))
