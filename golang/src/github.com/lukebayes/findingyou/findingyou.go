package main

import (
	"fmt"
	"github.com/lukebayes/gnomplate"
)

func main() {
	fmt.Println("Finding You")
	options := &gnomplate.WindowOptions{
		Title:  "Finding You",
		Height: 600,
		Width:  800,
	}
	instance := gnomplate.CreateWindow(options)
	fmt.Println("Created Application")
	instance.Open()
}
