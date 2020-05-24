package spec

import (
	"github.com/waybeams/waybeams/pkg/events"
)

type CharCallback func(r rune)

type Window interface {
	ResizableWriter
	ResizableReader

	BeginFrame()
	Close()
	EndFrame()
	FrameRate() int
    GetContentScale() (float32, float32)
	Init()
	OnResize(handler events.EventHandler) events.Unsubscriber
	PixelRatio() float32
	PollEvents()
	ShouldClose() bool
	UpdateInput(root ReadWriter)
}

type InputController interface {
	Update(root ReadWriter)
}
