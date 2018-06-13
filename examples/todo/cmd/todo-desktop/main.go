package main

import (
	"path/filepath"
	"runtime"

	"github.com/waybeams/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/builder"
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/surface/glfw"
	"github.com/waybeams/waybeams/pkg/surface/nano"
)

func init() {
	runtime.LockOSThread()
}

// CreateSurface will creates and configures a new surface.
func CreateNanoSurface() spec.Surface {
	fonts := filepath.Join(
		"src",
		"github.com",
		"waybeams",
		"waybeams",
		"third_party",
		"fonts",
		"Roboto",
	)

	return nano.NewSurface(
		nano.Font("Roboto", filepath.Join(fonts, "Roboto-Regular.ttf")),
		nano.Font("Roboto Light", filepath.Join(fonts, "Roboto-Light.ttf")),
	)
}

// CreateModel instantiates and configures a new application model.
func CreateModel() *model.App {
	m := model.New()
	m.CreateItem("Item One")
	m.CreateItem("Item Two")
	m.CreateItem("Item Three")
	m.CreateItem("Item Four")
	m.CreateItem("Item Five")
	m.CreateItem("Item Six")
	return m
}

func CreateGlfwWindow() spec.Window {
	return glfw.NewWindow(
		glfw.Width(800),
		glfw.Height(600),
		glfw.Title("Todo"),
	)
}

func main() {
	// Create the app model and some fake data.
	m := CreateModel()

	// Create and configure the Builder.
	build := builder.New(
		builder.Factory(ctrl.AppRenderer(m)),
		builder.Surface(CreateNanoSurface()),
		builder.Window(CreateGlfwWindow()),
	)

	// Loop until exit.
	build.Listen()
}
