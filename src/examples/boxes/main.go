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
		Box(b, Children(func() {
			Box(b, FlexWidth(1), FlexHeight(1), MaxWidth(640), MaxHeight(480))
			Box(b, FlexWidth(1), FlexHeight(1), MaxWidth(320), MaxHeight(240))
		}))
	}))
}
