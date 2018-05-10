package spec

import "font"

// Surface is an interface that should hide concrete drawing implementations
// from controls. Using this interface should allow us to reasonably easily
// swap rendering backends (e.g., NanoVg, Cairo, Skia, HTML Canvas, etc.)
type Surface interface {
	Init()

	Font(name string) *font.Font

	// Arc draws an arc from the x,y point along angle 1 and 2 at the provided radius.
	Arc(xc, yc, radius, angle1, angle2 float64)

	// Begin a path to stroke or fill.
	BeginPath()

	BeginFrame(w, h float64)

	EndFrame()

	Close()

	CreateFont(name, path string)

	// DebugDumpPathCache will print the current Path cache to log.
	DebugDumpPathCache()

	// Fill will fill the previously drawn shape.
	Fill()

	// Rect draws a rectangle from x and y to width and height.
	Rect(x, y, width, height float64)

	// Rect draws a rectangle with rounded corners from x and y to width and height.
	RoundedRect(x, y, width, height, radius float64)

	// SetStrokeWidth configures the width in pixels of the next shape.
	SetStrokeWidth(width float64)

	// SetFillColor configures the fill color as an RGBA hex value (0xffcc00ff)
	SetFillColor(color uint)

	// SetStrokeColor configures the stroke color as an RGBA hex value (0xffcc00ff)
	SetStrokeColor(color uint)

	// Stroke draws a stroke around the previous shape.
	Stroke()

	// GetOffsetSurfaceFor provides offset surface for nested controls so that
	// they can use local coordinates for positioning.
	// GetOffsetSurfaceFor(d Reader) Surface

	SetFontSize(size float64)

	SetFontFace(face string)

	Text(x float64, y float64, text string)
}
