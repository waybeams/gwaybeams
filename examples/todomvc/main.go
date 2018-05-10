package main

import (
	"builder"
	ctrl "controls"
	"opts"
	"runtime"
	"spec"
	"surface/nano"
)

func init() {
	runtime.LockOSThread()
}

func CreateAppRenderer() func() spec.ReadWriter {
	boxStyle := opts.Bag(
		opts.BgColor(0xffcc00ff),
		opts.StrokeSize(1),
		opts.StrokeColor(0x666666ff),
		opts.Padding(10),
		opts.Gutter(10),
	)

	mainStyle := opts.Bag(
		boxStyle,
		opts.FontColor(0x111111ff),
		opts.FontFace("Roboto"),
		opts.FontSize(24),
	)

	headerStyle := opts.Bag(
		boxStyle,
		opts.FontColor(0x000000ff),
		opts.BgColor(0xffffffff),
		opts.FontFace("Roboto"),
		opts.FontSize(36),
	)

	return func() spec.ReadWriter {
		return ctrl.VBox(
			mainStyle,
			opts.Key("App"),
			opts.Child(ctrl.HBox(
				boxStyle,
				opts.Key("Header"),
				opts.FlexWidth(1),
				opts.Height(80),
				opts.Child(
					ctrl.Label(
						headerStyle,
						opts.Text("Hello World"))),
			)),
			opts.Child(ctrl.VBox(
				boxStyle,
				opts.Key("Body"),
				opts.FlexWidth(1),
				opts.FlexHeight(1),
			)),
			opts.Child(ctrl.HBox(
				boxStyle,
				opts.Key("Footer"),
				opts.FlexWidth(1),
				opts.Height(100),
			)),
		)
	}
}

func main() {
	// Create the Application specification.
	renderer := CreateAppRenderer()

	// Create and configure the NanoSurface.
	surface := nano.New(
		nano.Font("Roboto", "./third_party/fonts/Roboto/Roboto-Regular.ttf"),
	)

	// Create and configure the Builder.
	build := builder.New(
		builder.Surface(surface),
		builder.Renderer(renderer),
	)

	// Loop until exit.
	build.Listen()
}
