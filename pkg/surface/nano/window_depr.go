package nano

/*
import (
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/shibukawa/nanovgo"
	"github.com/waybeams/waybeams/pkg/surface"
	"ui"
	"ui/control"
	"uiold/opts"
)

type NanoWindowControl struct {
	GlfwWindowControl

	inputCtrl       InputController
	lastHeight      int
	lastWidth       int
	lastHoverTarget ui.Displayable
	nanoContext     *nanovgo.Context
	Surface     ui.Surface
}

func (c *NanoWindowControl) initInput() {
	c.inputCtrl = NewGlfwInput(c, c)
}

func (c *NanoWindowControl) updateSize(width, height int) {
	if float64(width) != c.Width() || float64(height) != c.Height() {
		c.SetWidth(float64(width))
		c.SetHeight(float64(height))
		c.LayoutDrawAndPaint()
		c.Layout()
	}
}

func (c *NanoWindowControl) initNanoContext() {
	c.initGl()
	// nnovgo.AntiAlias | nanovgo.StencilStrokes | nanovgo.Debug
	context, err := nanovgo.NewContext(0)
	if err != nil {
		panic(err)
	}

	c.nanoContext = context
}

func (c *NanoWindowControl) initSurface() {
	c.Surface = surface.NewNano(c.nanoContext)
}

func (c *NanoWindowControl) OnExit() {
	c.nanoContext.Delete()
	c.GlfwWindowControl.OnClose()
}

func (c *NanoWindowControl) ShouldExit() bool {
	return c.getNativeWindow().ShouldClose()
}

func (c *NanoWindowControl) enterFrameHandler(e events.Event) {
	c.inputCtrl.Update()
	c.LayoutDrawAndPaint()
	c.PollEvents()

	if c.ShouldExit() {
		// Stop the frame loop by destroying the Builder
		c.Context().Destroy()
	}
}

func (c *NanoWindowControl) Listen() {
	c.init()

	defer c.OnExit()
	// TODO(lbayes): Definitely do not like this pattern. Need to find a cleaner way to set this up.
	// Controls should generally not interact with the Builder and I do not want to require any
	// particular control TYPE to be the ROOT.
	c.Context().OnFrameEntered(c.enterFrameHandler)
	// Block permanently as frame events arrive
	c.Context().Listen()
}

func (c *NanoWindowControl) init() {
	// Do not connect to the GPU hardware until we begin looping.
	// This allows us to set up an instance in the test environment.
	c.initGlfw()
	c.initNanoContext()
	c.initSurface()
	c.initInput()
	c.OnWindowResize(c.updateSize)
}

func (c *NanoWindowControl) LayoutDrawAndPaint() {
	// Currently working to remove / rework this method from the controller
	// of the frame work, to one that is simply notified when a frame happens.

	// Make the control source size match the source frame buffer.
	fbWidth, fbHeight := c.getNativeWindow().GetFramebufferSize()
	winWidth, winHeight := c.getNativeWindow().GetSize()
	// TODO(lbayes): Only set pixelRatio on init, not every frame
	pixelRatio := float32(fbWidth) / float32(winWidth)

	c.nanoContext.BeginFrame(int(fbWidth), int(winHeight), pixelRatio)

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
	c.Context().CreateFonts(c.Surface)
	c.Draw(c.Surface)

	c.nanoContext.EndFrame()
	c.SwapWindowBuffers()
}

func NewNanoWindow() *NanoWindowControl {
	win := &NanoWindowControl{}
	win.SetTitle(ui.DefaultWindowTitle)
	return win
}

var NanoWindow = control.Define("NanoWindow",
	func() ui.Displayable { return NewNanoWindow() },
	opts.LayoutType(ui.VerticalFlowLayoutType),
	opts.Width(ui.DefaultWindowWidth),
	opts.Height(ui.DefaultWindowHeight))
*/
