package main

import (
	"path/filepath"
	"runtime"

	"github.com/waybeams/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/builder"
	"github.com/waybeams/waybeams/pkg/env/glfw"
	"github.com/waybeams/waybeams/pkg/env/nano"
)

func init() {
	runtime.LockOSThread()
}

// Get the relative font path for the provided font name.
func fontPathFor(fontFileName string) string {
	return filepath.Join(
		"src",
		"github.com",
		"waybeams",
		"waybeams",
		"third_party",
		"fonts",
		"Roboto",
		fontFileName,
	)
}

func main() {
	// Create and configure the Builder.
	builder.New(
		glfw.NewWindow(
			glfw.Width(800),
			glfw.Height(600),
			glfw.Title("Todo"),
		),
		nano.NewSurface(
			nano.AddFont("Roboto", fontPathFor("Roboto-Regular.ttf")),
			nano.AddFont("Roboto Light", fontPathFor("Roboto-Light.ttf")),
		),
		ctrl.AppRenderer(model.NewSample()),
	).Listen()
}
