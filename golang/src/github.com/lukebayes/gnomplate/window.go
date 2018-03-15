package gnomplate

/**
* Sample code found here:
* https://medium.com/@drgomesp/opengl-and-golang-getting-started-abcd3d96f3db
 */

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

type WindowOptions struct {
	Width  int
	Height int
	Title  string
}

type Window struct {
	target *glfw.Window
}

func (w Window) Open() Window {
	fmt.Println("Open called")
	runtime.LockOSThread()

	w.target.MakeContextCurrent()

	// TODO: Remove the following!
	context := glfw.GetCurrentContext()
	fmt.Printf("CONTEXT: %v", context)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.ClearColor(0.4, 0.4, 0.4, 0.4)

	for !w.target.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		w.target.SwapBuffers()
		glfw.PollEvents()
	}

	fmt.Println("Exiting now")
	return w
}

func CreateWindow(opts *WindowOptions) *Window {
	fmt.Println("CreateWindow called")

	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("could not initialize glfw: %v", err))
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(opts.Width, opts.Height, opts.Title, nil, nil)

	if err != nil {
		panic(fmt.Errorf("could not create opengl renderer: %v", err))
	}

	return &Window{target: win}
}
