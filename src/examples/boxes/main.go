package main

import (
	. "display"
	"log"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	Application(&Opts{Title: "Example"}, func(s Surface) {
		VBox(s, func() {
			log.Printf("Application main compose function called with: %v", s)
			Box(s, &Opts{FlexWidth: 1, FlexHeight: 1})
			Box(s, &Opts{FlexWidth: 1, FlexHeight: 1})
		})
	}).Loop()
}
