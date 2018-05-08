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

var currentMessage = func(index int) string {
	return messages[index]
}

func createWindow(opts ...context.Option) Displayable {
	grey := BgColor(0xdbd9d6ff)
	blue := BgColor(0x00acd7ff)
	pink := BgColor(0xce3262ff)

	messageIndex := 0
	inputContent := ""

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
						messageIndex = (messageIndex + 1) % len(messages)
						fmt.Println("Update Message Now to:", messages[messageIndex])
						d.Invalidate()
					}

					var clearInputContent = func(e events.Event) {
						inputContent = ""
						d.Invalidate()
					}

					var updateInputContent = func(e events.Event) {
						char := e.Payload()
						inputContent += string(char.(rune))
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
								Text(currentMessage(messageIndex)))
						}))

					VBox(c, TraitNames("control-list"), Gutter(10), Padding(10), FlexWidth(1), FlexHeight(1), Children(func() {
						TextInput(c,
							DefaultStyle,
							Width(200),
							Padding(10),
							On(events.CharEntered, updateInputContent),
							Placeholder("Full Name Here"),
							Data("TextInput.Text", inputContent),
							Text(inputContent))
						HBox(c, Gutter(10), Padding(5), FlexWidth(1), FlexHeight(1), Children(func() {
							Button(c, DefaultStyle,
								Height(50),
								OnState("active", blue),
								OnState("hovered", grey),
								OnState("pressed", pink),
								OnClick(updateMessage),
								Text("Update Label"))
							Button(c, DefaultStyle,
								Height(50),
								OnState("active", blue),
								OnState("hovered", grey),
								OnState("pressed", pink),
								OnClick(clearInputContent),
								Text("Clear Text"))

						}))
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
