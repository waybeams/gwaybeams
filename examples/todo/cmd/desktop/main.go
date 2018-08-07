package main

import (
	"path/filepath"
	"runtime"

	"github.com/waybeams/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/clock"
	"github.com/waybeams/waybeams/pkg/env/glfw"
	"github.com/waybeams/waybeams/pkg/env/nano"
	"github.com/waybeams/waybeams/pkg/scheduler"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// Create and configure the Scheduler.
	scheduler.New(
		glfw.NewWindow(
			glfw.Width(800),
			glfw.Height(600),
			glfw.Title("Todo"),
		),
		nano.NewSurface(
			nano.AddFont("Roboto", filepath.Join("third_party", "fonts", "Roboto", "Roboto-Regular.ttf")),
			nano.AddFont("Roboto Light", filepath.Join("third_party", "fonts", "Roboto", "Roboto-Light.ttf")),
		),
		ctrl.AppRenderer(model.NewSample()),
		clock.New(),
	).Listen()
}
