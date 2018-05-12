package builder

import (
	"clock"
	"events"
	"layout"
	"spec"
	"surface/nano"
	"win/glfw"
)

const TitleOffset = 30

// Maybe this should be called a Driver?
type Builder struct {
	clock            clock.Clock
	window           spec.Window
	surface          spec.Surface
	renderer         func() spec.ReadWriter
	root             spec.ReadWriter
	isClosed         bool
	lastWindowWidth  float64
	lastWindowHeight float64
}

func (b *Builder) Clock() clock.Clock {
	if b.clock == nil {
		b.clock = clock.New()
	}
	return b.clock
}

func (b *Builder) renderSpecs() {
	win := b.window
	w, h := win.Width(), win.Height()
	b.lastWindowHeight = h
	b.lastWindowWidth = w

	// Create a new Spec tree and store it.
	root := b.renderer()
	b.root = root

	// NO IDEA WHY, but my surface is drawing with these offsets.
	// Need to figure this out and remove the magic number.
	root.SetWidth(w)
	root.SetHeight(h - TitleOffset)

	surface := b.Surface()
	layout.Layout(root, surface)
	layout.Draw(root, surface)
}

func (b *Builder) Listen() {
	win := b.Window()
	win.Init()
	win.OnResize(b.windowResizedHandler)

	surface := b.Surface()
	surface.Init()

	defer b.Close()
	clock.OnFrame(b.createFrameHandler(), win.FrameRate(), b.Clock())
}

func (b *Builder) windowResizedHandler(e events.Event) {
	b.frameHandler(false)
}

func (b *Builder) createFrameHandler() func() bool {
	return func() bool {
		return b.frameHandler(true)
	}
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
	if b.surface == nil {
		b.surface = nano.New()
	}
	return b.surface
}

func (b *Builder) Root() spec.ReadWriter {
	return b.root
}

func (b *Builder) Window() spec.Window {
	if b.window == nil {
		b.window = glfw.New()
	}
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

func Renderer(renderer func() spec.ReadWriter) Option {
	return func(b *Builder) {
		b.renderer = renderer
	}
}

func Root(root spec.ReadWriter) Option {
	return func(b *Builder) {
		b.root = root
	}
}

func New(options ...Option) *Builder {
	b := &Builder{}
	for _, option := range options {
		option(b)
	}
	return b
}
