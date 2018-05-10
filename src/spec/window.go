package spec

import "events"

type Window interface {
	ResizableWriter
	ResizableReader

	BeginFrame()
	Close()
	EndFrame()
	FrameRate() int
	Init()
	OnResize(handler events.EventHandler) events.Unsubscriber
	PixelRatio() float64
	PollEvents()
	ShouldClose() bool
}
