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

func createWindow(opts ...ctx.Option) Displayable {
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

	return NanoWindow(ctx.New(opts...),
		ID("nano-window"),
		Padding(10),
		Title("Test Title"),
		Width(800),
		Height(610),
		Children(func(c Context) {
			Box(c, ID("header"), BgColor(0xce3262ff), Height(100), FlexWidth(1), Children(func() {
				Label(c,
					ID("title"),
					DefaultStyle,
					BgColor(0xffcc00ff),
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
								DefaultStyle,
								FlexWidth(1),
								FontSize(48),
								Height(60),
								IsFocusable(false),
								MinWidth(100),
								Padding(5),
								Text(currentMessage()))
						}))

					VBox(c, TraitNames("component-list"), Gutter(10), Padding(10), FlexWidth(1), FlexHeight(1), Children(func() {
						TextInput(c, DefaultStyle, BgColor(0xffcc00ff), Width(200), Height(60), Placeholder("Full Name Here"))
						Button(c, DefaultStyle, OnClick(updateMessage), Text("Update Label"))
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
		ctx.Font("Roboto", "./third_party/fonts/Roboto/Roboto-Regular.ttf"),
	)
	win.(*NanoWindowComponent).Listen()
}
