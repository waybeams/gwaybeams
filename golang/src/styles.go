package styles

import "display"

func BackgroundColor(color uint) func(canvas Canvas, view Displayable) {
	return nil
}

func BorderColor(color uint) func(canvas Canvas, view Displayable) {
	return nil
}

func BorderSize(size int) func(canvas Canvas, view Displayable) {
	return nil
}

func BorderStyle(style string) func(canvas Canvas, view Displayable) {
	return nil
}
