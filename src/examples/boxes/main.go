package main

import (
	. "display"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func createWindow() (Displayable, error) {
	return GlfwWindow(NewBuilder(), Title("Test Title"), Width(640), Height(480), Children(func(b Builder) {
		Box(b, FlexWidth(1), FlexHeight(1), MaxWidth(640), MaxHeight(480))
		Box(b, FlexWidth(1), FlexHeight(1), MaxWidth(320), MaxHeight(240))
	}))
}

func main() {
	win, err := createWindow()
	if err != nil {
		panic(err)
	}
	win.(*GlfwWindowComponent).Loop()
}
