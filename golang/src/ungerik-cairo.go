package main

// Trying to get a 2d drawing surface that reasonably renders text into a cross
// platform Window.

// FWIW: Here's an example of Cairo in an SDL window  (in C++)
// https://github.com/cubicool/cairo-gl-sdl2/blob/master/src/sdl-example.cpp

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/ungerik/go-cairo"
	"log"
	"runtime"
	"time"
	"unsafe"
)

// Making debugging a little slower
const FRAME_RATE = 12

// Auto-called by Go runtime when this module is ready for execution.
func init() {
	fmt.Println("Main.init called")
	runtime.LockOSThread()
}

func initOpenGl(window *glfw.Window) *cairo.Surface {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	width, height := window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))
	log.Printf("GL WIDTH %d x HEIGHT %d", int32(width), int32(height))

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, width, height)
	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		log.Println("Window FrameBufferSize updated!")
		onFrameBufferResized(w, width, height)
		// surface.Update(width, height)
	})

	return surface
}

func onFrameBufferResized(window *glfw.Window, width, height int) {
	log.Printf("glfw: framebuffer size: %dx%d", width, height)
}

func initGlfw() *glfw.Window {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(800, 600, "Finding You", nil, nil)

	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	return window
}

func draw(window *glfw.Window, surface *cairo.Surface) {
	fmt.Println("DRAWING NOW!")
	surface.SetSourceRGBA(0, 1, 0, 1)
	surface.Paint()

	surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surface.SetFontSize(32.0)
	surface.LineTo(100.0, 100.0)
	surface.ShowText("Hello World")

	surface.MoveTo(100, 200)
	surface.SetLineWidth(5)
	surface.SetSourceRGB(1, 0, 1)
	surface.Rectangle(100, 200, 300, 140)
	surface.Stroke()
	surface.Fill()

	surface.Flush()
}

func render(window *glfw.Window, surface *cairo.Surface) {
	fmt.Println("LOOP NOW")
	glfw.PollEvents()

	width, height := window.GetFramebufferSize()
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.ONE, gl.ONE_MINUS_SRC_ALPHA)

	gl.Viewport(0, 0, int32(width), int32(height))
	// gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.5, 0.5, 0.5, 1)

	// textureValue := uint32(0)
	// texturePtr := &textureValue
	//
	// gl.GenTextures(1, texturePtr)
	// gl.BindTexture(gl.TEXTURE_2D, textureValue)
	// gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(width), int32(height), 0, gl.BGRA, gl.UNSIGNED_BYTE, nil)

	data := surface.GetData()
	gl.TexImage2D(gl.TEXTURE_2D, 0, 4, int32(width), int32(height), 0, gl.BGRA, gl.UNSIGNED_BYTE, unsafe.Pointer(&data))

	window.SwapBuffers()
}

func main() {
	window := initGlfw()
	defer glfw.Terminate()

	surface := initOpenGl(window)

	fpsTicker := time.NewTicker(time.Second / FRAME_RATE)

	draw(window, surface)

	for !window.ShouldClose() {
		select {
		case <-fpsTicker.C:
			render(window, surface)
			draw(window, surface)
		}
	}

	fmt.Println("EXITING NOW")
	surface.Destroy()
}
