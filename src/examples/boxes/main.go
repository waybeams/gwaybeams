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
		Trait(b, "*",
			BgColor(0xccccccff),
			FontFace("sans"),
			FontSize(12),
			FontColor(0x222222ff),
		)
		// Trait(b, "Box:mouseover", BgColor(0xff0000ff))
		// Trait(b, "Box:mousedown", BgColor(0x00ff00ff))

		Box(b, ID("header"), Height(100), FlexWidth(1), Children(func() {
			Label(b,
				ID("title"),
				BgColor(0x33ff33ff),
				FontSize(48),
				Padding(10),
				FlexWidth(1),
				Height(100),
				Text("HELLO WORLD"))
			Button(b, ID("One"))
			Button(b, ID("Two"))
			Button(b, ID("Three"))
		}))
		HBox(b, ID("body"), Padding(5), FlexHeight(3), FlexWidth(1), Children(func() {
			VBox(b, ID("leftNav"), FlexWidth(1), FlexHeight(1), Padding(10))
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
	win.(Window).Init()
}
