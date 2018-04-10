package main

import (
	. "display"
	"github.com/fogleman/ease"
	"runtime"
	"time"
)

func init() {
	// We need to do this so that our interactions with CGO (NanoVG/OpenGL) are
	// synchronous.
	runtime.LockOSThread()
}

func createWindow() (Displayable, error) {

	return NanoWindow(
		NewBuilder(),
		ID("nano-window"),
		Width(800),
		Height(600),
		Children(func(b Builder) {
			Box(b,
				ID("moving-box"),
				ExcludeFromLayout(true),
				Transition(X, 200.0, 400.0, time.Millisecond*200, ease.OutCubic),
				Y(200),
				BgColor(0xffcc00ff),
				Width(100),
				Height(100))
		}))
}

func main() {
	win, err := createWindow()
	if err != nil {
		panic(err)
	}

	win.(Window).Init()
}
