package main

import . "display"

func CreateBoxesApp() Displayable {
	return Application(&Opts{Title: "Example"}, func(s Surface) {
		VBox(s, func() {
			Box(s, &Opts{FlexWidth: 1, FlexHeight: 1})
			Box(s, &Opts{FlexWidth: 1, FlexHeight: 1})
		})
	})
}

func main() {
	ApplicationLoop(CreateBoxesApp())
}
