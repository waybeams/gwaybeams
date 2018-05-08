package main

import (
	"events"
	"fmt"
	"runtime"
	. "ui"
	"ui/context"
	. "ui/controls"
	. "ui/opts"
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

func createWindow(opts ...context.Option) Displayable {
	grey := BgColor(0xdbd9d6ff)
	blue := BgColor(0x00acd7ff)
	pink := BgColor(0xce3262ff)

	DefaultStyle := Bag(
		grey,
		FontColor(0xffffffff),
		FontFace("Roboto"),
		FontSize(36),
		OnState("hovered", blue),
	)

	return NanoWindow(context.New(opts...),
		ID("nano-window"),
		Padding(10),
		Title("Test Title"),
		Width(800),
		Height(610),
		Children(func(c Context) {
			Box(c, ID("header"), pink, FlexWidth(1), Children(func() {
				Label(c,
					ID("title"),
					DefaultStyle,
					StrokeSize(1),
					FontSize(48),
					Padding(30),
					FlexWidth(1),
					IsFocusable(false),
					Text("HELLO WORLD"))
			}))
			HBox(c, ID("body"), Padding(5), FlexHeight(1), FlexWidth(1), Children(func() {
				Box(c, ID("leftNav"), FlexWidth(1), FlexHeight(1), Padding(10))
				VBox(c, ID("content"), Gutter(10), FlexWidth(3), FlexHeight(1), Children(func(d Displayable) {
					var updateMessage = func(e events.Event) {
						currentIndex = (currentIndex + 1) % len(messages)
						fmt.Println("Update Message Now to:", messages[currentIndex])
						d.Invalidate()
					}

					Box(c,
						FlexWidth(1),
						FlexHeight(1),
						Children(func() {
							Label(c,
								DefaultStyle,
								FlexWidth(1),
								Padding(20),
								FontSize(48),
								IsFocusable(false),
								MinWidth(100),
								Text(currentMessage()))
						}))

					VBox(c, TraitNames("control-list"), Gutter(10), Padding(10), FlexWidth(1), FlexHeight(1), Children(func() {
						TextInput(c,
							DefaultStyle,
							Width(200),
							Padding(20),
							Placeholder("Full Name Here"))
						Button(c, DefaultStyle,
							OnState("active", blue),
							OnState("hovered", grey),
							OnState("pressed", pink),
							OnClick(updateMessage),
							Text("Update Label"))
					}))
				}))
			}))
			HBox(c, ID("footer"), Height(80), FlexWidth(1), Children(func() {
				// FPS(b)
			}))
		}))
}

func main() {
	win := createWindow(
		// NOTE(lbayes): Font refs are relative to root for binaries, but relative to
		// files for tests. Not sure why :-(
		context.Font("Roboto", "./third_party/fonts/Roboto/Roboto-Regular.ttf"),
	)
	win.(*NanoWindowControl).Listen()
}
