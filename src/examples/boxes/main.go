package main

import . "display"

func CreateBoxesApp() Displayable {
	return Application(&Opts{Title: "Example"}, func(s Surface) {
		Box(s, func() {
			// Box(s, FlexWidth(1), FlexHeight(1), MaxWidth(321), MaxHeight(2423))
			Box(s, &Opts{FlexWidth: 1, FlexHeight: 1, MaxWidth: 640, MaxHeight: 480})
			Box(s, &Opts{FlexWidth: 1, FlexHeight: 1, MaxWidth: 320, MaxHeight: 280})
		})
	})
}

func main() {
	ApplicationLoop(CreateBoxesApp())
}
