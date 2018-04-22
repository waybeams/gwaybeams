package main

import (
	"clock"
	. "display"
	"runtime"
	"time"
)

func init() {
	// We need to do this so that our interactions with CGO (NanoVG/OpenGL) are
	// synchronous.
	runtime.LockOSThread()
}

var messages = []string{"ABCD", "EFGH", "IJKL", "MNOP", "QRST", "UVWX"}
var currentIndex = 0
var currentMessage = messages[currentIndex]

func updateMessage(clock clock.Clock, callback func()) {
	go func() {
		clock.Sleep(time.Second * 1)
		currentIndex = (currentIndex + 1) % len(messages)
		currentMessage = messages[currentIndex]
		callback()
	}()
}

func createWindow() (Displayable, error) {
	return NanoWindow(NewBuilder(), ID("nano-window"), Padding(10), Title("Test Title"), Children(func(b Builder) {
		Trait(b, "*",
			BgColor(0xccccccff),
			FontFace("Roboto"),
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
			VBox(b, ID("content"), Gutter(10), FlexWidth(3), FlexHeight(1), Children(func(d Displayable) {
				updateMessage(b.Clock(), func() {
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

				VBox(b, TraitNames("component-list"), Gutter(10), Padding(10), FlexWidth(1), FlexHeight(1), Children(func() {
					Label(b, Width(200), Height(40), Text("Full Name:"))
					TextInput(b, Width(200), Height(60), Placeholder("Name Here"))
				}))
			}))
		}))
		HBox(b, ID("footer"), Height(80), FlexWidth(1), Children(func() {
			// FPS(b)
		}))
	}))
}

func main() {
	win, err := createWindow()
	if err != nil {
		panic(err)
	}

	win.(Window).Init()
}
