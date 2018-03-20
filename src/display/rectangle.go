package display

func DrawRectangle(surface Surface, disp Displayable) {
	// fmt.Println(">> DrawRect for:", disp.GetX(), disp.GetY(), disp.GetWidth(), disp.GetHeight())
	surface.MoveTo(disp.GetX(), disp.GetY())
	surface.SetRgba(1, 1, 0, 1)

	surface.DrawRectangle(disp.GetX(), disp.GetY(), disp.GetWidth(), disp.GetHeight())
	surface.FillPreserve()

	surface.SetLineWidth(2)
	surface.SetRgba(0, 0, 0, 1)
	surface.Stroke()
}
