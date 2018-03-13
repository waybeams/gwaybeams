package gnomplate

import (
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

type Window struct {
	target *glfw.Window
}

func (w Window) Open() {
	w.target.MakeContextCurrent()

	for !w.target.ShouldClose() {
		w.target.SwapBuffers()
		glfw.PollEvents()
	}

	fmt.Println("Exiting now")
}

func CreateWindow(title string, height int, width int) *Window {
	fmt.Println("Start called")
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("could not initialize glfw: %v", err))
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(800, 600, "Hello world", nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create opengl renderer: %v", err))
	}

	return &Window{target: win}
}
