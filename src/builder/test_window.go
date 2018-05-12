package builder

import (
	"spec"
	"surface/nano"
)

func TestWindow(options ...spec.Option) spec.Option {
	var renderer func() spec.ReadWriter

	surface := nano.New(
		nano.Font("Roboto", "../../third_party/fonts/Roboto/Roboto-Regular.ttf"),
		nano.Font("Roboto Light", "../../third_party/fonts/Roboto/Roboto-Light.ttf"),
	)

	return func(r spec.ReadWriter) {
		renderer = func() spec.ReadWriter {
			return r
		}

		build := New(
			Surface(surface),
			Renderer(renderer),
		)
		build.Listen()
	}
}
