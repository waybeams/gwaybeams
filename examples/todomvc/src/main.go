package main

import (
	"./todomvc"
	"clock"
	"controls"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	c := clock.New()
	win := todomvc.Create(c, &todomvc.TodoAppModel{})
	win.(*controls.NanoWindowComponent).Listen()
}
