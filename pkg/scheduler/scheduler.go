package scheduler

import (
    "fmt"
	"github.com/waybeams/waybeams/pkg/clock"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/layout"
	"github.com/waybeams/waybeams/pkg/spec"
)

const shouldPollEvents = true

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
	shouldLayout     bool
	shouldDraw       bool
	surface          spec.Surface
	window           spec.Window
}

func (s *Scheduler) specInvalidatedHandler(e events.Event) {
	s.shouldRender = true
}

func (s *Scheduler) renderSpecs() {
	if s.shouldRender {
		root := s.root
		w, h := s.window.Width(), s.window.Height()
		s.lastWindowHeight = h
		s.lastWindowWidth = w

		// Create a new Spec tree and store it.
		root = s.factory()
		s.root = root
		root.On(events.Invalidated, s.specInvalidatedHandler)
	}
}

func (s *Scheduler) layoutSpecs() {
	if s.shouldRender || s.shouldLayout {
		s.root.SetWidth(s.window.Width())
		s.root.SetHeight(s.window.Height())

		layout.Layout(s.root, s.surface)
	}
}

func (s *Scheduler) drawSpecs() {
	if s.shouldRender || s.shouldLayout {
		layout.Draw(s.root, s.surface)
	}
}

func (s *Scheduler) Listen() {
	s.window.Init()
	s.window.OnResize(s.windowResizedHandler)
	s.surface.Init()

	defer s.Close()

	s.clock.OnFrame(func() bool {
		return s.frameHandler(shouldPollEvents)
	}, s.window.FrameRate())
}

func (s *Scheduler) windowResizedHandler(e events.Event) {
	s.shouldLayout = true
}

func (s *Scheduler) frameHandler(pollEvents bool) bool {
	if s.shouldRender || s.shouldLayout {
        w := s.window.Width()
        h := s.window.Height()

        fmt.Println("frameHandler with w: and h:", w, h)

		// BeginFrame on the Window.
		s.window.BeginFrame()

        sX, _ := s.window.GetContentScale()
        s.surface.SetPixelRatio(sX)

        // Configure the Surface demensions.
		s.surface.SetWidth(w)
		s.surface.SetHeight(h)

		// BeginFrame on the Surface.
		s.surface.BeginFrame()

        // Set content scaling for high DPI screens.
        // s.surface.SetScale(sX, sY)

		// Render the Specs.
		s.renderSpecs()
		s.layoutSpecs()
		s.drawSpecs()

		// EndFrame on Surface and then Window.
		s.surface.EndFrame()
		s.window.EndFrame()

		s.shouldRender = false
		s.shouldLayout = false
	}

	s.window.UpdateInput(s.root)

	// Return true if we should exit.
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
		shouldLayout: true,
		window:       w,
		surface:      s,
		factory:      f,
		clock:        c,
	}
}
