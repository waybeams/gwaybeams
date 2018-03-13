package main

import (
	"fmt"
	"github.com/lukebayes/gnomplate"
)

func main() {
	fmt.Println("Finding You")
	instance := &gnomplate.Application{600, 800}
	fmt.Printf("Created Applicatin with Width %d and Height %d.\n", instance.Width, instance.Height)
}
