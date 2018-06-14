package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/waybeams/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/builder"
	"github.com/waybeams/waybeams/pkg/surface/canvas"
)

func main() {
	js.Global.Get("document").Call("write", "Hello world!")

	// Create and configure the Builder.
	builder.New(
		builder.Factory(ctrl.AppRenderer(model.NewSample())),
		builder.Surface(canvas.NewSurface(
		// TODO(lbayes): Configure fonts for the canvas/window
		// canvas.Font("Roboto", "./src/github.com/waybeams/waybeams/third_party/fonts/Roboto/Roboto-Regular.ttf"),
		// canvas.Font("Roboto Light", "./src/github.com/waybeams/waybeams/third_party/fonts/Roboto/Roboto-Light.ttf"),
		)),
		builder.Window(canvas.NewWindow(
			canvas.Width(800),
			canvas.Height(600),
			canvas.Title("Todo MVC"),
		)),
	).Listen()
}
