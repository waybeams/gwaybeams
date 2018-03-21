package main

import (
	. "display"
)

func main() {
	Application(&Opts{Title: "Example"}, func(s Surface) {
		VBox(s, func() {
			Box(s, &Opts{FlexWidth: 1, FlexHeight: 1})
			Box(s, &Opts{FlexWidth: 1, FlexHeight: 1})
		})
	}).Loop()
}
