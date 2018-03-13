package gnomplate

import (
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

const defaultHeight = 600
const defaultTitle = ""
const defaultWidth = 800

type WindowOptions struct {
	Width  int
	Height int
	Title  string
}

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

func CreateWindow(opts *WindowOptions) *Window {
	if opts.Height == 0 {
		opts.Height = defaultHeight
	}
	if opts.Width == 0 {
		opts.Width = defaultWidth
	}
	if opts.Title == "" {
		opts.Title = defaultTitle
	}

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

	win, err := glfw.CreateWindow(opts.Width, opts.Height, opts.Title, nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create opengl renderer: %v", err))
	}

	return &Window{target: win}
}
