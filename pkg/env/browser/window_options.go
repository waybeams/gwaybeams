package browser

import (
	"github.com/gopherjs/gopherjs/js"
)

type WindowOption func(w *window)

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

func BrowserWindow(bwin *js.Object) WindowOption {
	return func(win *window) {
		win.browserWindow = bwin
	}
}
