package main

import (
	. "display"
	"runtime"
	"time"
)

func init() {
	runtime.LockOSThread()
}

var messages = []string{"ABCD", "EFGH", "IJKL", "MNOP", "QRST", "UVWX"}
var currentIndex = 0
var currentMessage = messages[currentIndex]

func updateMessage(callback func()) {
	go func() {
		time.Sleep(time.Second * 1)
		currentIndex = (currentIndex + 1) % len(messages)
		currentMessage = messages[currentIndex]
		callback()
	}()
}

func createWindow() (Displayable, error) {
	return NanoWindow(NewBuilder(), ID("nano-window"), Padding(10), Title("Test Title"), Children(func(b Builder) {
		Trait(b, "*",
			BgColor(0xccccccff),
			FontFace("sans"),
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
		}))
		HBox(b, ID("body"), Padding(5), FlexHeight(3), FlexWidth(1), Children(func() {
			Box(b, ID("leftNav"), FlexWidth(1), FlexHeight(1), Padding(10))
			Box(b, ID("content"), FlexWidth(3), FlexHeight(1), Children(func(d Displayable) {
				updateMessage(func() {
					d.InvalidateChildren()
				})
				Label(b,
					BgColor(0xff0000ff),
					MinWidth(100),
					FlexWidth(1),
					Height(60),
					FontSize(48),
					Padding(5),
					Text(currentMessage))
			}))
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
