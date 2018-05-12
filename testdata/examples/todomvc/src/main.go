package main

import (
	"./todomvc"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	win := todomvc.App(&todomvc.TodoAppModel{})
	// context.Font("Roboto", "./third_party/fonts/Roboto/Roboto-Regular.ttf"),
	// context.Font("Roboto-Thin", "./third_party/fonts/Roboto/Roboto-Thin.ttf"),
	// context.Font("Roboto-Light", "./third_party/fonts/Roboto/Roboto-Light.ttf"),
	// context.Font("Roboto-Bold", "./third_party/fonts/Roboto/Roboto-Bold.ttf"),
	win.Listen()
}
