package main

import (
	"./todomvc"
	"controls"
	"ctx"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	win := todomvc.Create(&todomvc.TodoAppModel{},
		ctx.Font("Roboto", "./third_party/fonts/Roboto/Roboto-Regular.ttf"),
		ctx.Font("Roboto-Thin", "./third_party/fonts/Roboto/Roboto-Thin.ttf"),
		ctx.Font("Roboto-Light", "./third_party/fonts/Roboto/Roboto-Light.ttf"),
		ctx.Font("Roboto-Bold", "./third_party/fonts/Roboto/Roboto-Bold.ttf"),
	)
	win.(*controls.NanoWindowComponent).Listen()
}
