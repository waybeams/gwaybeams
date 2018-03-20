package main

import (
	. "display"
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/cairo/cairogl"
)

func init() {
	runtime.LockOSThread()
}

func newMain() {
	Application(&Opts{Width: 800, Height: 600, Title: "Epiphyte Demo"}, func(s Surface) {
		VBox(s, func() {
			HBox(s, func() {
				Label(s, &Opts{Width: 200, HAlign: LeftAlign, Styles: []string{"logo"}}, "Epiphyte")
			})
		})
	})
}

// FrameRate is a temporary setting for render loop
const FrameRate = 12

// Main entry point to help explore this nascent GUI toolkit
// Most of the code in this file was copied from examples provided by the
// graphical libraries that we're using.
func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	win, err := glfw.CreateWindow(420, 420, "Cairo Demo", nil, nil)
	if err != nil {
		panic(err)
	}
	win.MakeContextCurrent()

	ww, wh := win.GetSize()
	width, height := win.GetFramebufferSize()
	log.Printf("glfw: created window %dx%d (framebuffer: %dx%d)", ww, wh, width, height)

	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))
	surface := cairogl.NewSurface(width, height)
	win.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		surface.Update(width, height)
		// NOTE(lbayes): This is needed on macOS b/c the OS will stop the render timer
		// while the mouse button is pressed on a window resizer. This causes the window
		// simply go black until the button is released. This fix has the unfortunate
		// side effect of making resizes laggy and slow feeling, especially on Linux.
		// Instead of forcing the draw here at a rate that can be much higher than our
		// framerate, we should simply enqueue a draw request that will take place in
		// the future.
		draw(win, surface)
	})

	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)
	// closer.Bind(func() {
	// close(exitC)
	// <-doneC
	// })

	fpsTicker := time.NewTicker(time.Second / FrameRate)
	for {
		select {
		case <-exitC:
			surface.Destroy()
			glfw.Terminate()
			fpsTicker.Stop()
			close(doneC)
			return
		case <-fpsTicker.C:
			if win.ShouldClose() {
				close(exitC)
				continue
			}
			glfw.PollEvents()
			draw(win, surface)
		}
	}
}

var lastWidth = 0
var lastHeight = 0

func draw(win *glfw.Window, surface *cairogl.Surface) {
	width, height := surface.Size()

	if lastWidth != width || lastHeight != height {
		lastWidth = width
		lastHeight = height
	} else {
		return
	}

	cr := surface.Context()
	cairoSurface := NewCairoSurface(cr)

	CreateRenderer(cairoSurface, func(s Surface) {
		Window(s, &Opts{X: 0, Y: 0, Width: float64(width), Height: float64(height)}, func() {
			Box(s, &Opts{FlexWidth: 1, FlexHeight: 1})
			Box(s, &Opts{FlexWidth: 1, FlexHeight: 1})
		})
	})

	/*
		// From example code, draws a rotating line inside a circle and box with some text that I added
			cairo.SetSourceRgba(cr, 0.1, 0.1, 0.1, 1)
			cairo.Paint(cr)

			cairo.SetSourceRgba(cr, 0.9, 0.9, 0.9, 1)
			cairo.SelectFontFace(cr, "serif", cairo.FontSlantNormal, cairo.FontWeightBold)
			cairo.SetFontSize(cr, 32)
			cairo.MoveTo(cr, 60.0, 50.0)
			cairo.ShowText(cr, "Hello World")

			offset := 50.0
			cairo.SetSourceRgba(cr, 1, 1, 1, 1)
			cairo.SetLineWidth(cr, 5)
			cairo.MakeRectangle(cr, 10+offset, 10+offset, 300, 300)
			cairo.Stroke(cr)

			xc := 160.0 + offset
			yc := 150.0 + offset
			radius := 100.0
			angleMux.RLock()
			angle1 := angle * (PI / 180.0)
			angleMux.RUnlock()
			angle2 := 180.0 * (PI / 180.0)

			cairo.SetLineWidth(cr, 3.0)
			cairo.SetSourceRgba(cr, 1, 1, 1, 1)
			cairo.Arc(cr, xc, yc, radius, angle1, angle2)
			cairo.Stroke(cr)

			cairo.SetSourceRgba(cr, 1, 0.2, 0.2, 0.6)
			cairo.SetLineWidth(cr, 6.0)

			cairo.Arc(cr, xc, yc, 10.0, 0, 2*PI)
			cairo.Fill(cr)

			cairo.Arc(cr, xc, yc, radius, angle1, angle2)
			cairo.LineTo(cr, xc, yc)
			cairo.Arc(cr, xc, yc, radius, angle1, angle2)
			cairo.LineTo(cr, xc, yc)
			cairo.Stroke(cr)
	*/

	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(1, 1, 1, 1)
	surface.Draw()
	win.SwapBuffers()
}
