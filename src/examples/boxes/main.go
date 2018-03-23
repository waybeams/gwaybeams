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
		Component(b, Children(func() {
			Component(b, FlexWidth(1), FlexHeight(1), MaxWidth(640), MaxHeight(480))
			Component(b, FlexWidth(1), FlexHeight(1), MaxWidth(320), MaxHeight(240))
		}))
	}))
}
