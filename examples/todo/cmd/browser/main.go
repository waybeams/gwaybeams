package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/waybeams/examples/todo/ctrl"
	"github.com/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/clock"
	"github.com/waybeams/waybeams/pkg/env/browser"
	"github.com/waybeams/waybeams/pkg/scheduler"
)

func CreateCanvas() *js.Object {
	doc := js.Global.Get("document")
	canvas := doc.Call("createElement", "canvas")

	body := doc.Get("body")
	body.Set("style", "margin:0;padding:0;")
	body.Call("appendChild", canvas)

	return canvas
}

func main() {
	canvas := browser.NewCanvasFromJsObject(CreateCanvas())

	// Create and configure the Scheduler.
	scheduler.New(
		browser.NewWindow(
			browser.BrowserWindow(js.Global.Get("window")),
			browser.Title("Todo MVC"),
		),
		browser.NewSurface(canvas),
		ctrl.AppRenderer(model.NewSample()),
		clock.New(),
	).Listen()
}
