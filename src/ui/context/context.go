package context

import (
	"clock"
	"events"
	"font"
	"ui"
)

type Option func(c *baseContext)

//-----------------------------------------------------------------------------
// Context Implementation

// baseContext is a concrete implementation of the Context interface that
// is required in order to instantiate any component.
type baseContext struct {
	builder     ui.Builder
	clock       clock.Clock
	root        ui.Displayable
	isDestroyed bool
	emitter     events.Emitter
	fonts       map[string]*font.Font
}

func (b *baseContext) getEmitter() events.Emitter {
	if b.emitter == nil {
		b.emitter = events.NewEmitter()
	}
	return b.emitter
}

func (b *baseContext) getFonts() map[string]*font.Font {
	if b.fonts == nil {
		b.fonts = make(map[string]*font.Font)
	}
	return b.fonts
}

func (b *baseContext) AddFont(name string, path string) {
	fonts := b.getFonts()
	if fonts[name] == nil {
		fonts[name] = font.New(name, path)
	}
}

func (b *baseContext) Font(name string) *font.Font {
	return b.getFonts()[name]
}

func (b *baseContext) OnFrameEntered(handler events.EventHandler) events.Unsubscriber {
	return b.getEmitter().On(events.FrameEntered, handler)
}

func (b *baseContext) Root() ui.Displayable {
	return b.root
}

// Builder returns the configured (or default) Builder instance.
func (b *baseContext) Builder() ui.Builder {
	if b.builder == nil {
		b.builder = ui.NewBuilder()
	}
	return b.builder
}

// Clock returns the configured (or default) Clock instance.
func (b *baseContext) Clock() clock.Clock {
	if b.clock == nil {
		// Go get the clock from the provided Root component's builder
		b.clock = clock.New()
	}
	return b.clock
}

// Destroy will clean up all composed entity state.
func (b *baseContext) Destroy() {
	b.builder.Destroy()
	b.isDestroyed = true
}

func (b *baseContext) CreateFonts(s ui.Surface) {
	for _, font := range b.getFonts() {
		if !font.Created {
			s.CreateFont(font.Name, font.Path)
			font.Created = true
		}
	}
}

func (b *baseContext) Listen() {
	var frameHandler = func() bool {
		b.getEmitter().Emit(events.New(events.FrameEntered, nil, nil))
		return b.isDestroyed
	}
	// TODO(lbayes): Move the DefaultFrameRate constant
	// TODO(lbayes): FIX ONFRAME interface... so gross!
	clock.OnFrame(frameHandler, ui.DefaultFrameRate, b.Clock())
}

//-----------------------------------------------------------------------------
// Context Constructor

// New creates and returns a configured Context instance.
func New(options ...Option) *baseContext {
	instance := &baseContext{}
	for _, option := range options {
		option(instance)
	}
	return instance
}

//-----------------------------------------------------------------------------
// Context Options Below

// Builder Option
func Builder(b ui.Builder) Option {
	return func(c *baseContext) {
		c.builder = b
	}
}

// Clock Option
func Clock(c clock.Clock) Option {
	return func(context *baseContext) {
		context.clock = c
	}
}

func Font(name string, path string) Option {
	return func(context *baseContext) {
		context.AddFont(name, path)
	}
}

//-----------------------------------------------------------------------------
// Test Helper

func NewTestContext() ui.Context {
	return New(Font("Roboto", "../../../testdata/Roboto-Regular.ttf"))
}
