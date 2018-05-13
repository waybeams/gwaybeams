package spec

import (
	"events"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type KeyCallback func(key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
type CharCallback func(r rune)
type MouseButtonCallback func(button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey)

type GestureSource interface {
	GetCursorPos() (xpos, ypos float64)
	SetCursorByName(name glfw.StandardCursor)
	SetCharCallback(callback CharCallback) events.Unsubscriber
	SetKeyCallback(callback KeyCallback) events.Unsubscriber
	SetMouseButtonCallback(callback MouseButtonCallback) events.Unsubscriber
}

type Window interface {
	ResizableWriter
	ResizableReader
	GestureSource

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

type InputController interface {
	Update(root ReadWriter)
}
