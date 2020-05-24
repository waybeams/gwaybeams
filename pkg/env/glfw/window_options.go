package glfw

import (
	g "github.com/go-gl/glfw/v3.3/glfw"
)

type WindowOption func(*window)

func Width(width float64) WindowOption {
	return func(win *window) {
		win.SetWidth(width)
	}
}

func Height(height float64) WindowOption {
	return func(win *window) {
		win.SetHeight(height)
	}
}

func Hint(key g.Hint, value int) WindowOption {
	return func(win *window) {
		win.AddHint(WindowHint{Key: key, Value: value})
	}
}

func FrameRate(fps int) WindowOption {
	return func(win *window) {
		win.frameRate = fps
	}
}

func Title(title string) WindowOption {
	return func(win *window) {
		win.SetTitle(title)
	}
}
