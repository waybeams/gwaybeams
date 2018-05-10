package glfw

import (
	g "github.com/go-gl/glfw/v3.2/glfw"
)

func Width(width float64) Option {
	return func(win *window) {
		win.SetWidth(width)
	}
}

func Height(height float64) Option {
	return func(win *window) {
		win.SetHeight(height)
	}
}

func Hint(key g.Hint, value int) Option {
	return func(win *window) {
		win.AddHint(WindowHint{Key: key, Value: value})
	}
}

func FrameRate(fps int) Option {
	return func(win *window) {
		win.frameRate = fps
	}
}

func Title(title string) Option {
	return func(win *window) {
		win.SetTitle(title)
	}
}
