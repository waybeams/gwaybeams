package display

func Rectangle(disp Displayable, surface Surface) {
	surface.MakeRectangle(disp.GetX(), disp.GetY(), disp.GetWidth(), disp.GetHeight())

	surface.SetRgba(1, 1, 0, 1)
	surface.FillPreserve()

	surface.SetLineWidth(2)
	surface.SetRgba(0, 0, 0, 1)
	surface.Stroke()
}
