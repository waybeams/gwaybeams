package scheduler

import (
	"github.com/waybeams/waybeams/pkg/clock"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/layout"
	"github.com/waybeams/waybeams/pkg/spec"
)

// Scheduler manages Specification lifecycle and rendering interactions with
// the host environment.
type Scheduler struct {
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

func (s *Scheduler) specInvalidatedHandler(e events.Event) {
	s.shouldRender = true
}

func (s *Scheduler) renderSpecs() {
	root := s.root
	if s.shouldRender {
		w, h := s.window.Width(), s.window.Height()
		s.lastWindowHeight = h
		s.lastWindowWidth = w

		// Create a new Spec tree and store it.
		root = s.factory()
		s.root = root

		root.On(events.Invalidated, s.specInvalidatedHandler)

		root.SetWidth(w)
		root.SetHeight(h)

		layout.Layout(root, s.surface)

		s.shouldRender = false
	}

	layout.Draw(root, s.surface)
	s.window.UpdateInput(root)
}

func (s *Scheduler) Listen() {
	s.window.Init()
	s.window.OnResize(s.windowResizedHandler)

	s.surface.Init()

	defer s.Close()

	s.clock.OnFrame(func() bool {
		return s.frameHandler(true)
	}, s.window.FrameRate())
}

func (s *Scheduler) windowResizedHandler(e events.Event) {
	s.shouldRender = true
}

func (s *Scheduler) frameHandler(pollEvents bool) bool {
	// BeginFrame on the Window.
	s.window.BeginFrame()

	// BeginFrame on the Surface.
	s.surface.SetWidth(s.window.Width())
	s.surface.SetHeight(s.window.Height())
	s.surface.BeginFrame()

	// Render the Specs.
	s.renderSpecs()

	// EndFrame on Surface and then Window.
	s.surface.EndFrame()

	s.window.EndFrame()

	// Return true if we should exita.
	if pollEvents {
		s.window.PollEvents()
	}
	return s.isClosed || s.window.ShouldClose()
}

func (s *Scheduler) Close() {
	s.surface.Close()
	s.Window().Close()
	s.isClosed = true
}

func (s *Scheduler) Root() spec.ReadWriter {
	return s.root
}

func (s *Scheduler) Window() spec.Window {
	return s.window
}

func New(w spec.Window, s spec.Surface, f spec.Factory, c clock.Clock) *Scheduler {
	return &Scheduler{
		shouldRender: true,
		window:       w,
		surface:      s,
		factory:      f,
		clock:        c,
	}
}
