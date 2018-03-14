package main

/**
* Sample code found here:
* https://medium.com/@drgomesp/opengl-and-golang-getting-started-abcd3d96f3db
 */

import (
	"github.com/lukebayes/gnomplate"
)

func main() {
	opts := &gnomplate.WindowOptions{Title: "Gnomplate"}
	window := gnomplate.CreateWindow(opts)
	window.Open()
}
