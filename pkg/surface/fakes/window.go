package fakes

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/spec"
	g "github.com/waybeams/waybeams/pkg/surface/glfw"
)

type FakeWindow struct {
	width      float64
	height     float64
	pixelRatio float64
	frameRate  int
}

func (f *FakeWindow) FrameRate() int {
	return f.frameRate
}

func (f *FakeWindow) SetWidth(width float64) {
	f.width = width
}

func (f *FakeWindow) Width() float64 {
	return f.width
}

func (f *FakeWindow) SetHeight(height float64) {
	f.height = height
}

func (f *FakeWindow) Height() float64 {
	return f.height
}

func (f *FakeWindow) BeginFrame() {
}

func (f *FakeWindow) EndFrame() {
}

func (f *FakeWindow) Close() {
}

func (f *FakeWindow) PixelRatio() float64 {
	return f.pixelRatio
}

func (f *FakeWindow) ShouldClose() bool {
	return false
}

func (f *FakeWindow) UpdateInput(root spec.ReadWriter) {
	panic("FakeWIndow.UpdateInput not implemented")
}

func NewFakeWindow() *FakeWindow {
	return &FakeWindow{}
}

// FakeGestureSource is a minimal struct that is used for testing Gestures.
type FakeGestureSource struct {
	xpos          float64
	ypos          float64
	CursorName    glfw.StandardCursor
	CharCallback  spec.CharCallback
	KeyCallback   g.KeyCallback
	MouseCallback g.MouseButtonCallback
}

func (f *FakeGestureSource) SetCursorPos(xpos, ypos float64) {
	f.xpos = xpos
	f.ypos = ypos
}

func (f *FakeGestureSource) GetCursorPos() (xpos, ypos float64) {
	return f.xpos, f.ypos
}

func (f *FakeGestureSource) SetCursorByName(name glfw.StandardCursor) {
	f.CursorName = name
}

func (f *FakeGestureSource) SetKeyCallback(callback g.KeyCallback) events.Unsubscriber {
	f.KeyCallback = callback
	return func() bool {
		f.KeyCallback = nil
		return true
	}
}

func (f *FakeGestureSource) SetCharCallback(callback spec.CharCallback) events.Unsubscriber {
	f.CharCallback = callback
	return func() bool {
		f.CharCallback = nil
		return true
	}
}

func (f *FakeGestureSource) SetMouseButtonCallback(callback g.MouseButtonCallback) events.Unsubscriber {
	f.MouseCallback = callback
	return func() bool {
		f.MouseCallback = nil
		return true
	}
}

func NewFakeGestureSource() *FakeGestureSource {
	return &FakeGestureSource{}
}
