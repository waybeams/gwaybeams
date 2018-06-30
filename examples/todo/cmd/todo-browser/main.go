package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/waybeams/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/builder"
	"github.com/waybeams/waybeams/pkg/env/browser"
)

func createCanvas() *js.Object {
	doc := js.Global.Get("document")
	canvas := doc.Call("createElement", "canvas")

	body := doc.Get("body")
	body.Set("style", "margin:0;padding:0;")
	body.Call("appendChild", canvas)

	return canvas
}

func main() {
	canvas := browser.NewCanvasFromJsObject(createCanvas())

	// Create and configure the Builder.
	builder.New(
		builder.Factory(ctrl.AppRenderer(model.NewSample())),
		builder.Surface(browser.NewSurface(canvas)),
		builder.Window(browser.NewWindow(
			browser.BrowserWindow(js.Global.Get("window")),
			browser.Title("Todo MVC"),
		)),
	).Listen()
}
