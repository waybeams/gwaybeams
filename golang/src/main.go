package main

import (
	"display"
	// "fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/cairo/cairogl"
)

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
		// fmt.Printf("Width x Height: %dx%d\n", width, height)
		surface.Update(width, height)
		draw(surface)
		win.SwapBuffers()
	})

	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)
	// closer.Bind(func() {
	// close(exitC)
	// <-doneC
	// })

	fpsTicker := time.NewTicker(time.Second / 60)
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
			draw(surface)
			win.SwapBuffers()
		}
	}
}

const PI = 3.1415926

var angle = 45.0
var angleMux sync.RWMutex

func init() {
	runtime.LockOSThread()
	go func() {
		for {
			angleMux.Lock()
			angle -= 1
			if angle <= 0 {
				angle = 360.0
			}
			angleMux.Unlock()
			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func draw(surface *cairogl.Surface) {
	cr := surface.Context()
	adapter := display.NewCairoAdapter(cr)

	width, height := surface.Size()
	rectWidth := width - 20
	rectHeight := height - 20
	rectX := 10
	rectY := 10

	root := display.NewRectangle()
	opts := &display.Opts{Width: rectWidth, Height: rectHeight, X: rectX, Y: rectY}
	root.UpdateState(opts)

	display.Render(adapter, root)

	/*
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
}
