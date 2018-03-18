package styles

import d "display"

func BackgroundColor(color uint) func(surface d.Surface, view d.Renderable) {
	return nil
}

func BorderColor(color uint) func(surface d.Surface, view d.Renderable) {
	return nil
}

func BorderSize(size int) func(surface d.Surface, view d.Renderable) {
	return nil
}

func BorderStyle(style string) func(surface d.Surface, view d.Renderable) {
	return nil
}

func Margin(size int) func(surface d.Surface, view d.Renderable) {
	return nil
}

func Padding(size int) func(surface d.Surface, view d.Renderable) {
	return nil
}
