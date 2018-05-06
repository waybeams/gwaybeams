package controls

import (
	"component"
	"events"
	"github.com/shibukawa/nanovgo"
	"opts"
	"surface"
	"ui"
)

type NanoWindowComponent struct {
	GlfwWindowComponent

	inputCtrl       InputController
	lastHeight      int
	lastWidth       int
	lastHoverTarget ui.Displayable
	nanoContext     *nanovgo.Context
	nanoSurface     ui.Surface
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
	c.initSurface()
	c.initInput()
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
	c.Context().CreateFonts(c.nanoSurface)
	c.Draw(c.nanoSurface)

	c.nanoContext.EndFrame()
	c.SwapWindowBuffers()
}

func NewNanoWindow() *NanoWindowComponent {
	win := &NanoWindowComponent{}
	win.SetTitle(ui.DefaultWindowTitle)
	return win
}

var NanoWindow = component.Define("NanoWindow",
	func() ui.Displayable { return NewNanoWindow() },
	opts.LayoutType(ui.VerticalFlowLayoutType),
	opts.Width(ui.DefaultWindowWidth),
	opts.Height(ui.DefaultWindowHeight))
