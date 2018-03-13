package main

import (
	"fmt"
	"github.com/lukebayes/gnomplate"
)

func main() {
	fmt.Println("Finding You")
	instance := gnomplate.CreateWindow("Finding You", 80, 600)
	fmt.Println("Created Application")
	instance.Open()
}
