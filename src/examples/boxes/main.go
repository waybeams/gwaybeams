package main

import (
	. "display"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	NewGlfwBuilder(WindowTitle("Test Title"), WindowSize(640, 480), BuildAndLoop(func(b Builder) {
		Sprite(b, Children(func() {
			Sprite(b, FlexWidth(1), FlexHeight(1), MaxWidth(640), MaxHeight(480))
			Sprite(b, FlexWidth(1), FlexHeight(1), MaxWidth(320), MaxHeight(240))
		}))
	}))
}
