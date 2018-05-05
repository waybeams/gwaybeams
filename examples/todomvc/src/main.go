package main

import (
	"./todomvc"
	"clock"
	"runtime"
	"ui"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	c := clock.New()
	app := todomvc.Create(c, &todomvc.TodoAppModel{})
	app.(ui.Window).Init()
}
