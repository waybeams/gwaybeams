package display

import "log"

func DrawRectangle(surface Surface, disp Displayable) {
	log.Println("-----------------------------")
	surface.MoveTo(disp.GetX(), disp.GetY())
	surface.SetRgba(1, 1, 0, 1)

	log.Printf("DrawRectangle with %vX%v", disp.GetWidth(), disp.GetHeight())
	surface.DrawRectangle(disp.GetX(), disp.GetY(), disp.GetWidth(), disp.GetHeight())
	surface.FillPreserve()

	surface.SetLineWidth(2)
	surface.SetRgba(0, 0, 0, 1)
	surface.Stroke()
}
