package main

import (
	. "display"
	"github.com/fogleman/ease"
	"runtime"
)

func init() {
	// We need to do this so that our interactions with CGO (NanoVG/OpenGL) are
	// synchronous.
	runtime.LockOSThread()
}

func createWindow() (Displayable, error) {

	// var currentMove ComponentOption
	// moveLeft := Transition(X, 700.0, 0.0, 2000, ease.InOutCubic)
	moveRight := Transition(X, 0.0, 700.0, 2000, ease.InOutCubic)

	/*
		var currentMoveName string
			var toggleCurrentMove = func(e Event) {
				if currentMoveName == "moveLeft" {
					currentMove = moveRight
					currentMoveName = "moveRight"
				} else {
					currentMove = moveLeft
					currentMoveName = "moveLeft"
				}
			}
	*/

	return NanoWindow(
		NewBuilder(),
		ID("nano-window"),
		Width(800),
		Height(600),
		Children(func(b Builder) {
			Box(b,
				ID("moving-box"),
				ExcludeFromLayout(true),
				// OnClick(toggleCurrentMove),
				moveRight,
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
