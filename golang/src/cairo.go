package main

import (
	"fmt"
	cairo "github.com/ungerik/go-cairo"
	"time"
)

func cairoMain() {
	fmt.Println("Cairo Version:", cairo.Version())
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 800, 600)
	surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surface.SetFontSize(32)
	surface.MoveTo(10.0, 50.0)
	surface.ShowText("Hello World")

	for {
		fmt.Println("Enter Frame")
		surface.Paint()
		time.Sleep(1000 * time.Millisecond)
	}
}
