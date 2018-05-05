package main

import (
	. "controls"
	"ctx"
	"events"
	"fmt"
	. "opts"
	"runtime"
	. "ui"
)

func init() {
	// We need to do this so that our interactions with CGO (NanoVG/OpenGL) are
	// synchronous.
	runtime.LockOSThread()
}

var messages = []string{"ABCD", "EFGH", "IJKL", "MNOP", "QRST", "UVWX"}
var currentIndex = 0

func currentMessage() string {
	return messages[currentIndex]
}

func createWindow() Displayable {
	DefaultStyle := Bag(
		BgColor(0xdbd9d6ff),
		FontColor(0xffffffff),
		FontFace("Roboto"),
		FontSize(36),
		OnState("hovered", BgColor(0xffcc00ff)),
	)

	BlueStyle := Bag(
		BgColor(0x00acd7ff),
	)

	return NanoWindow(ctx.New(),
		ID("nano-window"),
		Padding(10),
		Title("Test Title"),
		Width(800),
		Height(610),
		Children(func(c Context) {
			/*
				Trait(c, "*",
					BgColor(0xdbd9d6ff),
					FontColor(0xffffffff),
					FontFace("Roboto"),
					FontSize(36),
				)
			*/
			// Trait(c, "Box:hovered", BgColor(0xff0000ff))
			// Trait(c, "Box:pressed", BgColor(0x00ff00ff))
			// Trait(c, "Box:disabled", BgColor(0xccccccff))

			Box(c, ID("header"), BgColor(0xce3262ff), Height(100), FlexWidth(1), Children(func() {
				Label(c,
					ID("title"),
					DefaultStyle,
					StrokeSize(1),
					FontSize(48),
					Padding(10),
					FlexWidth(1),
					Height(100),
					IsFocusable(false),
					Text("HELLO WORLD"))
			}))
			HBox(c, ID("body"), Padding(5), FlexHeight(3), FlexWidth(1), Children(func() {
				Box(c, ID("leftNav"), FlexWidth(1), FlexHeight(1), Padding(10))
				VBox(c, ID("content"), Gutter(10), FlexWidth(3), FlexHeight(1), Children(func(d Displayable) {
					var updateMessage = func(e events.Event) {
						currentIndex = (currentIndex + 1) % len(messages)
						fmt.Println("Update Message Now to:", messages[currentIndex])
						d.Invalidate()
					}

					Box(c,
						BlueStyle,
						FlexWidth(1),
						FlexHeight(1),
						Children(func() {
							Label(c,
								FlexWidth(1),
								FontSize(48),
								Height(60),
								IsFocusable(false),
								MinWidth(100),
								Padding(5),
								Text(currentMessage()))
						}))

					VBox(c, TraitNames("component-list"), Gutter(10), Padding(10), FlexWidth(1), FlexHeight(1), Children(func() {
						TextInput(c, DefaultStyle, Width(200), Height(60), Placeholder("Full Name Here"))
						Button(c, DefaultStyle, Width(200), Height(60), OnClick(updateMessage), Text("Update Label"))
					}))
				}))
			}))
			HBox(c, ID("footer"), Height(80), FlexWidth(1), Children(func() {
				// FPS(b)
			}))
		}))
}

func main() {
	win := createWindow()
	win.(Window).Init()
}
