package builder

import (
	"github.com/waybeams/waybeams/pkg/clock"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/layout"
	"github.com/waybeams/waybeams/pkg/spec"
)

// Maybe this should be called a Driver?
type Builder struct {
	clock            clock.Clock
	factory          spec.Factory
	isClosed         bool
	lastWindowHeight float64
	lastWindowWidth  float64
	root             spec.ReadWriter
	shouldRender     bool
	surface          spec.Surface
	window           spec.Window
}

func (b *Builder) Clock() clock.Clock {
	if b.clock == nil {
		b.clock = clock.New()
	}
	return b.clock
}

func (b *Builder) specInvalidatedHandler(e events.Event) {
	// fmt.Println("INVALIDATED!", spec.Path(e.Target().(spec.Reader)))
	b.shouldRender = true
}

func (b *Builder) renderSpecs() {
	root := b.root
	if b.shouldRender {
		win := b.window
		w, h := win.Width(), win.Height()
		b.lastWindowHeight = h
		b.lastWindowWidth = w

		// Create a new Spec tree and store it.
		root = b.factory()
		b.root = root

		root.On(events.Invalidated, b.specInvalidatedHandler)

		root.SetWidth(w)
		root.SetHeight(h)

		layout.Layout(root, b.Surface())

		b.shouldRender = false
	}

	layout.Draw(root, b.Surface())
	b.Window().UpdateInput(root)
}

func (b *Builder) Listen() {
	win := b.Window()
	win.Init()
	win.OnResize(b.windowResizedHandler)

	surface := b.Surface()
	surface.Init()

	defer b.Close()
	clock.OnFrame(b.eventPollingFrameHandler, win.FrameRate(), b.Clock())
}

func (b *Builder) windowResizedHandler(e events.Event) {
	b.shouldRender = true
	b.frameHandler(false)
}

func (b *Builder) eventPollingFrameHandler() bool {
	return b.frameHandler(true)
}

func (b *Builder) frameHandler(pollEvents bool) bool {
	// BeginFrame on the Window.
	win := b.Window()
	win.BeginFrame()

	// BeginFrame on the Surface.
	surface := b.Surface()
	surface.BeginFrame(win.Width(), win.Height())

	// Render the Specs.
	b.renderSpecs()

	// EndFrame on Surface and then Window.
	surface.EndFrame()

	win.EndFrame()

	// Return true if we should exita.
	if pollEvents {
		win.PollEvents()
	}
	return b.isClosed || win.ShouldClose()
}

func (b *Builder) Close() {
	b.Surface().Close()
	b.Window().Close()
	b.isClosed = true
}

func (b *Builder) Surface() spec.Surface {
	return b.surface
}

func (b *Builder) Root() spec.ReadWriter {
	return b.root
}

func (b *Builder) Window() spec.Window {
	return b.window
}

type Option func(b *Builder)

func Clock(c clock.Clock) Option {
	return func(b *Builder) {
		b.clock = c
	}
}

func Window(win spec.Window) Option {
	return func(b *Builder) {
		b.window = win
	}
}

func Surface(surface spec.Surface) Option {
	return func(b *Builder) {
		b.surface = surface
	}
}

func Factory(renderer spec.Factory) Option {
	return func(b *Builder) {
		b.factory = renderer
	}
}

func Root(root spec.ReadWriter) Option {
	return func(b *Builder) {
		b.root = root
	}
}

func New(options ...Option) *Builder {
	b := &Builder{}
	b.shouldRender = true
	for _, option := range options {
		option(b)
	}
	return b
}
