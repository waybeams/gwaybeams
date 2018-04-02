package main

import (
	d "display"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func createWindow() (d.Displayable, error) {
	return d.NanoWindow(d.NewBuilder(), d.Padding(10), d.Title("Test Title"), d.Width(640), d.Height(480), d.Children(func(b d.Builder) {
		d.Box(b, d.ID("header"), d.Padding(5), d.Height(100), d.FlexWidth(1))
		d.HBox(b, d.ID("body"), d.Padding(5), d.FlexHeight(3), d.FlexWidth(1), d.Children(func(b d.Builder) {
			d.Box(b, d.ID("leftNav"), d.FlexWidth(1), d.FlexHeight(1))
			d.Box(b, d.ID("content"), d.FlexWidth(3), d.FlexHeight(1))
		}))
		d.Box(b, d.ID("footer"), d.Height(80), d.FlexWidth(1))
	}))
}

func main() {
	win, err := createWindow()
	if err != nil {
		panic(err)
	}
	win.(d.Window).Loop()
}
