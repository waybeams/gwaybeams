package main

import (
	"builder"
	ctrl "controls"
	"events"
	"fmt"
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
		opts.BgColor(0xffffffff),
		opts.Padding(10),
		opts.Gutter(10),
	)

	flexStyle := opts.Bag(
		opts.FlexWidth(1),
		opts.FlexHeight(1),
	)

	headerText := opts.Bag(
		opts.FontColor(0xaf2f2f26),
		opts.FontFace("Roboto Light"),
		opts.FontSize(100),
	)

	footerText := opts.Bag(
		opts.FontColor(0xccccccff),
		opts.FontFace("Roboto"),
		opts.FontSize(18),
	)

	mainStyle := opts.Bag(
		boxStyle,
		opts.FontColor(0x111111ff),
		opts.FontFace("Roboto"),
		opts.FontSize(24),
	)

	buttonStyle := opts.Bag(
		opts.BgColor(0xf8f8f8ff),
		// opts.StrokeColor(0x333333ff),
		// opts.StrokeSize(1),
	)

	return func() spec.ReadWriter {
		return ctrl.VBox(
			opts.Key("App"),
			mainStyle,
			opts.HAlign(spec.AlignCenter),
			opts.Child(ctrl.VBox(
				boxStyle,
				flexStyle,
				opts.Key("Body"),
				opts.MaxWidth(500),
				opts.MinWidth(350),

				opts.HAlign(spec.AlignCenter),
				opts.Child(ctrl.Label(
					headerText,
					opts.Text("TODO"),
				)),
				opts.Child(ctrl.Label(
					opts.Key("TextInput"),
					opts.Height(80),
					opts.FlexWidth(1),
					opts.BgColor(0xccccccff),
				)),
				opts.Child(ctrl.Box(
					opts.Key("Items Box"),
					opts.MinHeight(100),
					opts.FlexWidth(1),
					opts.BgColor(0xeeeeeeff),
				)),
				opts.Child(ctrl.HBox(
					boxStyle,
					opts.Padding(10),
					opts.Gutter(10),
					footerText,
					opts.HAlign(spec.AlignCenter),
					opts.Key("Footer"),
					opts.FlexWidth(1),
					opts.Padding(5),
					opts.Child(ctrl.Label(
						opts.Text("2 items left"),
						buttonStyle,
					)),
					// opts.Child(ctrl.Spacer()),
					opts.Child(ctrl.Button(
						opts.Text("All"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							fmt.Println("All Clicked")
						}),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Active"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							fmt.Println("Active Clicked")
						}),
					)),
					opts.Child(ctrl.Button(
						opts.Text("Completed"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							fmt.Println("Completed Clicked")
						}),
					)),
					// opts.Child(ctrl.Spacer()),
					opts.Child(ctrl.Button(
						opts.Text("Clear Completed"),
						buttonStyle,
						opts.OnClick(func(e events.Event) {
							fmt.Println("Clear Completed Clicked")
						}),
					)),
				)),
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
		nano.Font("Roboto Light", "./third_party/fonts/Roboto/Roboto-Light.ttf"),
	)

	// Create and configure the Builder.
	build := builder.New(
		builder.Surface(surface),
		builder.Factory(renderer),
	)

	// Loop until exit.
	build.Listen()
}
