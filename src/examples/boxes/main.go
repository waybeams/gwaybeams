package main

import (
	. "display"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func createWindow() (Displayable, error) {
	return GlfwWindow(NewBuilder(), Title("Test Title"), Width(640), Height(480), GlfwFrameRate(10), Children(func(b Builder) {
		Box(b, Id("header"), Height(200), FlexWidth(1))
		Box(b, Id("body"), FlexHeight(1), FlexWidth(1))
		Box(b, Id("footer"), Height(120), FlexWidth(1))
	}))
}

func main() {
	win, err := createWindow()
	if err != nil {
		panic(err)
	}
	win.(*GlfwWindowComponent).Loop()
}
