package main

import (
	. "display"
	"fmt"
	"runtime"
)

func init() {
	// We need to do this so that our interactions with CGO (NanoVG/OpenGL) are
	// synchronous.
	runtime.LockOSThread()
}

var messages = []string{"ABCD", "EFGH", "IJKL", "MNOP", "QRST", "UVWX"}
var currentIndex = 0
var currentMessage = messages[currentIndex]

func createWindow() (Displayable, error) {
	return NanoWindow(NewBuilder(), ID("nano-window"), Padding(10), Title("Test Title"), Children(func(b Builder) {
		Trait(b, "*",
			BgColor(0xccccccff),
			FontColor(0x222222ff),
			FontFace("Roboto"),
			FontSize(36),
		)
		// Trait(b, "Box:hovered", BgColor(0xff0000ff))
		// Trait(b, "Box:pressed", BgColor(0x00ff00ff))
		// Trait(b, "Box:disabled", BgColor(0xccccccff))

		Box(b, ID("header"), Height(100), FlexWidth(1), Children(func() {
			Label(b,
				ID("title"),
				BgColor(0x33ff33ff),
				FontSize(48),
				Padding(10),
				FlexWidth(1),
				Height(100),
				IsFocusable(false),
				Text("HELLO WORLD"))
		}))
		HBox(b, ID("body"), Padding(5), FlexHeight(3), FlexWidth(1), Children(func() {
			Box(b, ID("leftNav"), FlexWidth(1), FlexHeight(1), Padding(10))
			VBox(b, ID("content"), Gutter(10), FlexWidth(3), FlexHeight(1), Children(func(d Displayable) {
				var updateMessage = func(e Event) {
					fmt.Println("Update Message!")
					currentIndex = (currentIndex + 1) % len(messages)
					d.Invalidate()
				}

				Label(b,
					BgColor(0xff0000ff),
					FlexWidth(1),
					FontSize(48),
					Height(60),
					IsFocusable(false),
					MinWidth(100),
					Padding(5),
					Text(currentMessage))

				VBox(b, TraitNames("component-list"), Gutter(10), Padding(10), FlexWidth(1), FlexHeight(1), Children(func() {
					TextInput(b, Width(200), Height(40), Placeholder("Full Name Here"))
					Button(b, Width(200), Height(60), OnClick(updateMessage), Text("Update Label"))
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
