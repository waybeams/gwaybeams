package main

import (
	"./todomvc"
	"clock"
	"display"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	c := clock.New()
	app, err := todomvc.Create(c)

	if err != nil {
		panic(err)
	}
	app.(display.Window).Init()
}
