package main

import (
	"github.com/waybeams/waybeams/examples/todo/controls"
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/builder"
	"github.com/waybeams/waybeams/pkg/spec"
	"github.com/waybeams/waybeams/pkg/surface/nano"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func CreateSurface() spec.Surface {
	return nano.New(
		nano.Font("Roboto", "./src/github.com/waybeams/waybeams/third_party/fonts/Roboto/Roboto-Regular.ttf"),
		nano.Font("Roboto Light", "./src/github.com/waybeams/waybeams/third_party/fonts/Roboto/Roboto-Light.ttf"),
	)
}

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

func main() {
	// Create the app model and some fake data.
	model := CreateModel()

	// Create the Application specification.
	renderer := controls.AppRenderer(model)

	// Create and configure the NanoSurface.
	surface := CreateSurface()

	// Create and configure the Builder.
	build := builder.New(
		builder.Surface(surface),
		builder.Factory(renderer),
	)

	// Loop until exit.
	build.Listen()
}
