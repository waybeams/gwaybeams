package main

import (
	. "display"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func createWindow() (Displayable, error) {
	return NanoWindow(NewBuilder(), Padding(10), Title("Test Title"), Children(func(b Builder) {
		Box(b, ID("header"), Height(100), FlexWidth(1), Children(func(b Builder) {
			Label(b, ID("title"), BgColor(0x33ff33ff), FontSize(48), Padding(10), FlexWidth(1), Height(100), Text("HELLO WORLD"))
		}))
		HBox(b, ID("body"), Padding(5), FlexHeight(3), FlexWidth(1), Children(func(b Builder) {
			Box(b, ID("leftNav"), FlexWidth(1), FlexHeight(1))
			Box(b, ID("content"), FlexWidth(3), FlexHeight(1))
		}))
		Box(b, ID("footer"), Height(80), FlexWidth(1))
	}))
}

func main() {
	win, err := createWindow()
	if err != nil {
		panic(err)
	}
	win.(Window).Loop()
}
