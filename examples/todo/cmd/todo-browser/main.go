package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/waybeams/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/builder"
	"github.com/waybeams/waybeams/pkg/surface/canvas"
)

func createPageContext() *js.Object {
	doc := js.Global.Get("document")
	pageContext := doc.Call("createElement", "canvas")

	body := doc.Get("body")
	body.Set("style", "margin:0;padding:0;")
	body.Call("appendChild", pageContext)

	return pageContext
}

func main() {
	pageContext := createPageContext()

	// Create and configure the Builder.
	builder.New(
		builder.Factory(ctrl.AppRenderer(model.NewSample())),
		builder.Surface(canvas.NewSurface(
			canvas.PageContext(pageContext),
		// TODO(lbayes): Configure fonts for the canvas/window
		// canvas.Font("Roboto", "./src/github.com/waybeams/waybeams/third_party/fonts/Roboto/Roboto-Regular.ttf"),
		// canvas.Font("Roboto Light", "./src/github.com/waybeams/waybeams/third_party/fonts/Roboto/Roboto-Light.ttf"),
		)),
		builder.Window(canvas.NewWindow(
			canvas.BrowserWindow(js.Global.Get("window")),
			canvas.Title("Todo MVC"),
		)),
	).Listen()
}
