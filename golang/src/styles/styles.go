package styles

import d "display"

func BackgroundColor(color uint) func(canvas d.Canvas, view d.Displayable) {
	return nil
}

func BorderColor(color uint) func(canvas d.Canvas, view d.Displayable) {
	return nil
}

func BorderSize(size int) func(canvas d.Canvas, view d.Displayable) {
	return nil
}

func BorderStyle(style string) func(canvas d.Canvas, view d.Displayable) {
	return nil
}
