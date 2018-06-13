package main

import (
	"runtime"

	"github.com/waybeams/waybeams/examples/todo/ctrl"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/builder"
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/surface/webgl"
)

func init() {
	runtime.LockOSThread()
}

// CreateSurface will creates and configures a new surface.
func CreateWebglSurface() spec.Surface {
	return webgl.NewSurface(
		webgl.Font("Roboto", "./src/github.com/waybeams/waybeams/third_party/fonts/Roboto/Roboto-Regular.ttf"),
		webgl.Font("Roboto Light", "./src/github.com/waybeams/waybeams/third_party/fonts/Roboto/Roboto-Light.ttf"),
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

func CreateWebglWindow() spec.Window {
	return webgl.NewWindow()
}

func main() {
	// Create the app model and some fake data.
	m := CreateModel()

	// Create and configure the Builder.
	build := builder.New(
		builder.Factory(ctrl.AppRenderer(m)),
		builder.Surface(CreateWebglSurface()),
		builder.Window(CreateWebglWindow()),
	)

	// Loop until exit.
	build.Listen()
}
